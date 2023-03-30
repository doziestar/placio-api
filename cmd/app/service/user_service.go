package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"placio-app/Dto"
	"placio-app/models"
	"time"
)

type IUser interface {
	CreateUser(userData Dto.SignUpDto, ctx *fiber.Ctx, db *gorm.DB) (*models.User, error)
	Get(id, email, account, permission string, social *models.Social, db *gorm.DB) (*[]models.User, error)
	GetUserByID(userID string, db *gorm.DB) (*models.User, error)
	GetByEmail(email string, db *gorm.DB) (*models.User, error)
	AddToAccount(userID, accountId, secondAccountId string)
	UpdateUser(userId string, accountId string, lastActive time.Time, disabled bool)
	GetUserAccounts(userID string, db *gorm.DB) (*[]models.Account, error)
	AddInterest(userID, accountID, interest string) error
	UpdateInterest(userID, accountID, oldInterest, newInterest string) error
	AddAccount(accountID, permission string, db *gorm.DB) error
	DeleteAccount(userId, accountId string, db *gorm.DB) error
	UpdateUserProfile(userId, accountId string, data Dto.UpdateProfileDto, db *gorm.DB) error
	DeleteUser(userId, accountID string, db *gorm.DB) error
}

type UserService struct {
	*models.User
	db *gorm.DB
}

func NewUserService(db *gorm.DB, user *models.User) *UserService {
	return &UserService{user, db}
}

func (u *UserService) CreateUser(userData Dto.SignUpDto, ctx *fiber.Ctx, db *gorm.DB) (*models.User, error) {
	return nil, nil
}

func (u *UserService) Get(id, email, account, permission string, social *models.Social, db *gorm.DB) (*[]models.User, error) {
	return nil, nil
}

func (u *UserService) GetUserByID(userID string, db *gorm.DB) (*models.User, error) {
	return nil, nil
}

func (u *UserService) GetByEmail(email string, db *gorm.DB) (*models.User, error) {
	return nil, nil
}

func (u *UserService) AddToAccount(userID, accountId, secondAccountId string) {

}

func (u *UserService) UpdateUser(userId string, accountId string, lastActive time.Time, disabled bool) {

}

func (u *UserService) GetUserAccounts(userID string, db *gorm.DB) (*[]models.Account, error) {
	return nil, nil
}

func (u *UserService) AddInterest(userID, accountID, interest string) error {
	return nil
}

func (u *UserService) UpdateInterest(userID, accountID, oldInterest, newInterest string) error {
	return nil
}

func (u *UserService) AddAccount(accountID, permission string, db *gorm.DB) error {
	return nil
}

func (u *UserService) DeleteAccount(userId, accountId string, db *gorm.DB) error {
	return nil
}

func (u *UserService) UpdateUserProfile(userId, accountId string, data Dto.UpdateProfileDto, db *gorm.DB) error {
	return nil
}

func (u *UserService) DeleteUser(userId, accountID string, db *gorm.DB) error {
	return nil
}
