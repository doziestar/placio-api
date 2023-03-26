package models

import (
	"context"
	"errors"
	"fmt"
	"os"
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

// Save creates new or updates an existing token
func (t *Token) Save(db *gorm.DB) error {
	if t.Access != "" {
		// data["access"] = crypto.Encrypt(data["access"])
	}
	if t.Refresh != "" {
		// data["refresh"] = crypto.Encrypt(data["refresh"])
	}

	// is there already a token for this provider?
	var tokenData Token
	if err := db.Where("provider = ? AND user_id = ?", t.Provider, t.UserID).First(&tokenData).Error; err == nil {
		// update existing token
		tokenData.Jwt = t.Jwt
		tokenData.Access = t.Access
		tokenData.Refresh = t.Refresh
		return db.Save(&tokenData).Error
	}

	// create a new token
	newToken := Token{
		ID:               uuid.NewString(),
		Provider:         t.Provider,
		TokenID:          t.TokenID,
		Jwt:              t.Jwt,
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
