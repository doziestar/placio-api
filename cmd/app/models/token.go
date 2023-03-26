package models

import (
	"context"
	"errors"
	"fmt"
	"os"
	"placio-app/Dto"
	"placio-pkg/hash"
	"placio-pkg/logger"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Token represents the token data in the database
type Token struct {
	ID                string `gorm:"primaryKey"`
	Provider          string `gorm:"not null"`
	Jwt               string
	TokenID           string
	Access            string
	AccessTokenExpiry time.Time
	Refresh           string
	UserID            string `gorm:"not null"`
	CodeCreateAt      time.Time
	CodeExpiresIn     time.Duration
	AccessCreateAt    time.Time
	AccessExpiresIn   time.Duration
	RefreshCreateAt   time.Time
	RefreshExpiresIn  time.Duration
	ProviderID        string
	Email             string
}

// BeforeCreate is a hook that is called before creating a token
func (t *Token) BeforeCreate(tx *gorm.DB) error {
	t.ID = GenerateID()
	//t.TokenID = GenerateID()
	return nil
}

// VerifyRefreshToken verifies the provided refresh token and returns the user ID associated with it.
func (t *Token) VerifyRefreshToken(refreshToken string, db *gorm.DB) (string, error) {
	var tokenData Token
	//hashedRefreshToken, err := hash.DecryptString(refreshToken, "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
	//if err != nil {
	//	logger.Error(context.Background(), err.Error())
	//	return "", errors.New("invalid refresh token")
	//}
	if err := db.Where("refresh = ?", refreshToken).First(&tokenData).Error; err != nil {
		return "", errors.New("invalid refresh token")
	}

	return tokenData.UserID, nil
}

// GenerateTokens generates a new access token and refresh token for the given user ID.
func GenerateTokens(userID string) (string, string, string, error) {
	tokenId := GenerateID()
	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"jti": tokenId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(1 * time.Hour).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", "", "", err
	}

	// Generate refresh token
	refreshToken := uuid.NewString()

	return accessTokenString, refreshToken, tokenId, nil
}

// RefreshTokens refreshes the access token and refresh token using the provided refresh token.
func (t *Token) RefreshTokens(refreshToken string, db *gorm.DB) (*Dto.TokenResponse, error) {
	// Verify refresh token and get user ID
	userID, err := t.VerifyRefreshToken(refreshToken, db)
	if err != nil {
		return nil, err
	}

	logger.Info(context.Background(), fmt.Sprintf("User ID: %s", userID))
	// Generate new access and refresh tokens
	newAccessToken, newRefreshToken, tokenID, err := GenerateTokens(userID)
	logger.Info(context.Background(), fmt.Sprintf("New Access Token: %s", newAccessToken))
	logger.Info(context.Background(), fmt.Sprintf("New Refresh Token: %s", newRefreshToken))
	if err != nil {
		return nil, err
	}

	// Update the token record in the database with the new access and refresh tokens
	var tokenData Token
	if err := db.Where("refresh = ?", refreshToken).First(&tokenData).Error; err != nil {
		return nil, errors.New("could not find token to update")
	}

	tokenData.Access = newAccessToken
	tokenData.TokenID = tokenID
	tokenData.UserID = userID
	tokenData.Refresh = newRefreshToken
	tokenData.AccessCreateAt = time.Now()
	tokenData.AccessExpiresIn = 1 * time.Hour * 24 * 7 // 7 days
	tokenData.RefreshCreateAt = time.Now()
	tokenData.RefreshExpiresIn = 1 * time.Hour * 24 * 30 // 30 days

	err = tokenData.Save(db)
	if err != nil {
		logger.Error(context.Background(), err.Error())
		return nil, errors.New("could not update token")
	}

	return &Dto.TokenResponse{
		AccessToken:      tokenData.Access,
		ExpiresIn:        1 * time.Hour * 24 * 7, // 7 days
		RefreshToken:     tokenData.Refresh,
		RefreshExpiresIn: 1 * time.Hour * 24 * 30, // 30 days
		UserId:           userID,
		TokenID:          tokenID,
	}, nil
}

// Save creates new or updates an existing token
func (t *Token) Save(db *gorm.DB) error {
	var err error
	if t.Access != "" {
		logger.Info(context.Background(), "encrypting access token")
		logger.Info(context.Background(), t.Access)
		t.Access, err = hash.EncryptString(t.Access, "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
		if err != nil {
			return errors.New("could not encrypt access token")
		}
	}
	if t.Refresh != "" {
		logger.Info(context.Background(), "encrypting refresh token")
		logger.Info(context.Background(), t.Refresh)
		t.Refresh, err = hash.EncryptString(t.Refresh, "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
		if err != nil {
			return errors.New("could not encrypt refresh token")
		}
	}

	// is there already a token for this provider?
	var tokenData Token
	logger.Info(context.Background(), "checking for existing token")
	logger.Info(context.Background(), t.Provider)
	logger.Info(context.Background(), t.UserID)
	if err := db.Where("provider = ? AND user_id = ?", t.Provider, t.UserID).First(&tokenData).Error; err == nil {
		// update existing token
		fmt.Println("updating token")
		tokenData.Jwt = t.Jwt
		tokenData.Access = t.Access
		tokenData.Refresh = t.Refresh
		return db.Save(&t).Error
	}

	//access, err := hash.EncryptString(t.Access, "access")
	//if err != nil {
	//	return err
	//}
	//
	//refresh, err := hash.EncryptString(t.Refresh, "refresh")
	//if err != nil {
	//	return err
	//}

	// create a new token
	newToken := Token{
		ID:               uuid.NewString(),
		Provider:         t.Provider,
		TokenID:          t.TokenID,
		Jwt:              t.Access,
		Access:           t.Access,
		Refresh:          t.Refresh,
		UserID:           t.UserID,
		CodeCreateAt:     time.Time{},
		CodeExpiresIn:    0,
		AccessCreateAt:   t.AccessCreateAt,
		AccessExpiresIn:  t.AccessExpiresIn,
		RefreshCreateAt:  t.RefreshCreateAt,
		RefreshExpiresIn: t.RefreshExpiresIn,
	}
	logger.Info(context.Background(), "creating new token")
	return db.Create(&newToken).Error
}

// Get returns the token for the given user and provider
func (t *Token) Get(id, provider, user string, skipDecryption bool, db *gorm.DB) ([]Token, error) {
	var data []Token
	query := db.Where("user_id = ?", user)
	if id != "" {
		query = query.Where("id = ?", id)
	}
	if provider != "" {
		query = query.Where("provider = ?", provider)
	}
	if err := query.Find(&data).Error; err != nil {
		return nil, err
	}

	if !skipDecryption {
		for i := range data {
			if data[i].Access != "" {
				// data[i].Access = crypto.Decrypt(data[i].Access)
			}
			if data[i].Refresh != "" {
				// data[i].Refresh = crypto.Decrypt(data[i].Refresh)
			}
		}
	}

	return data, nil
}

// Verify checks if a token is present for the given user and provider
func (t *Token) Verify(provider, user string) bool {
	var count int64
	db.Model(&Token{}).Where("provider = ? AND user_id = ?", provider, user).Count(&count)
	return count > 0
}

// VerifyToken checks if a token is present for the given user and provider
func (t *Token) VerifyToken(provider string, token string) (*Token, error) {
	// verify token using jwt
	tokenData, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	// get token data
	var data Token

	// check if token is valid
	if claims, ok := tokenData.Claims.(jwt.MapClaims); ok && tokenData.Valid {
		// get token data
		db.Where("provider = ? AND user_id = ?", provider, claims["sub"]).First(&data)
	}

	return &data, nil
}

// Delete deletes the token with the given ID, provider, and user
func (t *Token) Delete(id, provider, user string, db *gorm.DB) error {
	query := db.Model(&Token{}).Where("user_id = ?", user)
	if id != "" {
		logger.Info(context.Background(), fmt.Sprintf("Delete id: %s", id))
		query = query.Where("token_id = ?", id)
	}
	if provider != "" {
		logger.Info(context.Background(), fmt.Sprintf("Delete provider: %s", provider))
		query = query.Where("provider = ?", provider)
	}
	logger.Info(context.Background(), fmt.Sprintf("Delete query: %s", query.Statement.SQL.String()))
	result := query.Delete(&Token{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no token found")
	}
	return nil
	//return query.Delete(&Token{}).Error
}

// Generate GenerateJwt generates a new JWT token for the given user
func (t *Token) Generate(userId, accountId, provider, email string) (*Token, error) {
	// generate a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"aud": accountId,
		"iss": provider,
		// "jti":   providerID,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
		"iat":   time.Now().Unix(),
		"email": email,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	accessToken, err := token.SignedString([]byte(os.Getenv("ASSESS_TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}
	refreshToken, err := token.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}

	codeCreateAt := time.Now().UTC()

	return &Token{
		Jwt:              tokenString,
		Access:           accessToken,
		Refresh:          refreshToken,
		Provider:         provider,
		UserID:           userId,
		CodeCreateAt:     codeCreateAt,
		CodeExpiresIn:    codeCreateAt.Add(time.Hour * 24 * 7).Sub(codeCreateAt),
		AccessCreateAt:   codeCreateAt,
		AccessExpiresIn:  time.Hour * 24 * 7,
		RefreshCreateAt:  codeCreateAt,
		RefreshExpiresIn: time.Hour * 24 * 30,
	}, nil
}
