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
	GetLoggedInUser(ctx *fiber.Ctx) (*Dto.UserAccountResponse, error)
}

type UserService struct {
	user    *models.User
	db      *gorm.DB
	account *models.Account
}

func NewUserService(db *gorm.DB, user *models.User, account *models.Account) *UserService {
	return &UserService{user, db, account}
}

// Get Logged in user
func (u *UserService) GetLoggedInUser(ctx *fiber.Ctx) (*Dto.UserAccountResponse, error) {
	// get user from context
	user := ctx.Locals("user").(string)
	logger.Info(ctx.Context(), "user: "+user)
	// get user from db
	userData, err := u.user.GetUserById(user, u.db)
	if err != nil {
		return nil, err
	}

	// get current active account
	account, err := u.account.GetAccount(userData.CurrentActiveAccount, u.db)

	response := &Dto.UserAccountResponse{
		Name:           userData.Name,
		Email:          userData.Email,
		Disabled:       userData.Disabled,
		SupportEnabled: userData.SupportEnabled,
		UserId:         userData.ID,
		Account: &Dto.Account{
			ID:          account.ID,
			Permission:  account.Permission,
			AccountType: account.AccountType,
			AccountID:   account.AccountID,
			AccountSetting: Dto.AccountSetting{
				ID:                      account.AccountSetting.ID,
				AccountID:               account.AccountSetting.AccountID,
				TwoFactorAuthentication: account.AccountSetting.TwoFactorAuthentication,
				BlockedUsers:            account.AccountSetting.BlockedUsers,
				MutedUsers:              account.AccountSetting.MutedUsers,
			},
			Onboarded: account.Onboarded,
			//Interests:      account.Interests,
			UserID:   account.UserID,
			Plan:     account.Plan,
			Active:   account.Active,
			Status:   account.Status,
			Disabled: account.Disabled,
		},
	}

	return response, nil

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
