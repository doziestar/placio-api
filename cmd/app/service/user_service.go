package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"placio-app/Dto"
	"placio-app/models"
	"placio-pkg/logger"
	"time"
)

type IUser interface {
	CreateUser(userData Dto.SignUpDto, ctx *fiber.Ctx) (*models.User, error)
	Get(id, email, account, permission string, social *models.Social) (*[]models.User, error)
	GetUserByID(userID string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	AddToAccount(userID, accountId, secondAccountId string)
	UpdateUser(userId string, accountId string, lastActive time.Time, disabled bool)
	GetUserAccounts(userID string) (*[]models.Account, error)
	AddInterest(userID, accountID, interest string) error
	UpdateInterest(userID, accountID, oldInterest, newInterest string) error
	AddAccount(accountID, permission string) error
	DeleteAccount(userId, accountId string) error
	UpdateUserProfile(userId, accountId string, data Dto.UpdateProfileDto) error
	DeleteUser(userId, accountID string) error
	GetLoggedInUser(userId string) (*models.User, error)
	CheckIfUserNameOrEmailExists(userName, email string) (bool, error)
}

type UserService struct {
	user    *models.User
	db      *gorm.DB
	account *models.Account
}

func NewUserService(db *gorm.DB, user *models.User, account *models.Account) *UserService {
	return &UserService{user, db, account}
}

// GetLoggedInUser Logged in user
func (u *UserService) GetLoggedInUser(userId string) (*models.User, error) {
	// get user from db
	userData, err := u.user.GetUserById(userId, u.db)
	if err != nil {
		return nil, err
	}

	return userData, nil

}

func (u *UserService) CreateUser(userData Dto.SignUpDto, ctx *fiber.Ctx) (*models.User, error) {
	return nil, nil
}

func (u *UserService) Get(id, email, account, permission string, social *models.Social) (*[]models.User, error) {
	return nil, nil
}

func (u *UserService) GetUserByID(userID string) (*models.User, error) {
	logger.Info(u.db.Statement.Context, "user id: "+userID)
	user, err := u.user.GetUserById(userID, u.db)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) GetByEmail(email string) (*models.User, error) {
	return nil, nil
}

func (u *UserService) AddToAccount(userID, accountId, secondAccountId string) {

}

func (u *UserService) UpdateUser(userId string, accountId string, lastActive time.Time, disabled bool) {

}

func (u *UserService) GetUserAccounts(userID string) (*[]models.Account, error) {
	return nil, nil
}

func (u *UserService) AddInterest(userID, accountID, interest string) error {
	return nil
}

func (u *UserService) UpdateInterest(userID, accountID, oldInterest, newInterest string) error {
	return nil
}

func (u *UserService) AddAccount(accountID, permission string) error {
	return nil
}

func (u *UserService) DeleteAccount(userId, accountId string) error {
	return nil
}

func (u *UserService) UpdateUserProfile(userId, accountId string, data Dto.UpdateProfileDto) error {
	return nil
}

func (u *UserService) DeleteUser(userId, accountID string) error {
	return nil
}

// CheckIfUserNameOrEmailExists checks if username or email exists
func (u *UserService) CheckIfUserNameOrEmailExists(userName, email string) (bool, error) {
	if userName != "" {
		user, err := u.user.GetUserByUserName(userName, u.db)
		if err != nil {
			return false, err
		}
		if user != nil {
			return true, nil
		}
	} else if email != "" {
		user, err := u.user.GetUserByEmail(email, u.db)
		if err != nil {
			return false, err
		}
		if user != nil {
			return true, nil
		}
	}
	return false, nil
}
