package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"placio-app/Dto"
	"placio-app/errors"
	"time"
)

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Profile  []Profile `gorm:"foreignKey:UserID"`
	Role     string    `gorm:"type:varchar(50);default:'user';not null"`
	Provider string    `gorm:"type:varchar(50);default:'local';not null"`
	Photo    string    `gorm:"not null;default:'default.png'"`
	Verified bool      `gorm:"not null;default:false"`
	// MessagesSent       []Message      `gorm:"foreignKey:SenderID"`
	// MessagesReceived   []Message      `gorm:"foreignKey:RecipientID"`
	// Conversations      []Conversation `gorm:"many2many:conversation_participant"`
	// Groups             []Group        `gorm:"many2many:group_membership"`
	// VoiceNotesSent     []VoiceNote    `gorm:"foreignKey:SenderID"`
	// VoiceNotesReceived []VoiceNote    `gorm:"foreignKey:RecipientID"`
	// Notifications      []Notification
	// Bookings           []Booking
	// Payments           []Payment
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	err = u.HashPassword()
	if err != nil {
		return err
	}
	return nil
}

// HashPassword hashes the password
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.ErrForbidden
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	return u.Password == password
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) GetID() string {
	return u.ID
}

// GenerateToken generates a new token for the user
func (u *User) GenerateToken(user User) (Dto.Token, error) {
	// Create token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["name"] = user.Name
	claims["iat"] = time.Now().UTC().Unix()
	claims["exp"] = time.Now().UTC().Add(time.Minute * 15).Unix()

	// Generate access and refresh tokens
	accessToken, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return Dto.Token{}, errors.ErrForbidden
	}

	refreshToken := uuid.NewString()

	// Set token expiration times
	accessCreateAt := time.Now().UTC()
	accessExpiresIn := accessCreateAt.Add(time.Hour * 24 * 7)
	refreshCreateAt := accessCreateAt
	refreshExpiresIn := refreshCreateAt.Add(time.Hour * 24 * 30)

	return Dto.Token{
		ClientID:            "",
		UserID:              user.ID,
		RedirectURI:         "",
		Scope:               "",
		Code:                "",
		CodeChallenge:       "",
		CodeChallengeMethod: "",
		CodeCreateAt:        time.Time{},
		CodeExpiresIn:       0,
		Access:              accessToken,
		AccessCreateAt:      accessCreateAt,
		AccessExpiresIn:     time.Duration(accessExpiresIn.Unix()),
		Refresh:             refreshToken,
		RefreshCreateAt:     refreshCreateAt,
		RefreshExpiresIn:    time.Duration(refreshExpiresIn.Unix()),
	}, nil
}
