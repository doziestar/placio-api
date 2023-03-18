package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = user.ID
	claims["name"] = user.Name
	claims["exp"] = now.Add(69).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	accessToken, err := tokenByte.SignedString([]byte("secret"))
	refreshToken, err := tokenByte.SignedString([]byte("secret"))
	if err != nil {
		return Dto.Token{}, errors.ErrForbidden
	}
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
		AccessCreateAt:      time.Time{},
		AccessExpiresIn:     0,
		Refresh:             refreshToken,
		RefreshCreateAt:     time.Time{},
		RefreshExpiresIn:    0,
	}, nil
}
