package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"placio-app/Dto"
	"placio-app/database"
	"placio-app/models"
	"placio-pkg/logger"
	"time"
)

type IAccountService interface {
	CreateUserAccount(data *Dto.SignUpDto, ctx *gin.Context) (*fiber.Map, error)
	SwitchUserAccount(accountId, userId string) (*models.User, error)
	CreateBusinessAccount(data *Dto.AddAccountDto, ctx *gin.Context) (*Dto.UserAccountResponse, error)
	GetAccount(ctx *gin.Context) (*Dto.UserAccountResponse, error)
	GetAccounts(ctx *gin.Context) (*[]models.Account, error)
	MakeAccountDefault(accountId string, ctx *gin.Context) (*Dto.UserAccountResponse, error)
	Follow(followerID, followingID string) error
	Unfollow(followerID, followingID string) error
	ListFollowers(accountID string) ([]models.Account, error)
	ListFollowing(accountID string) ([]models.Account, error)
}

type AccountService struct {
	db      *gorm.DB
	account models.Account
	user    models.User
}

func NewAccountService(db *gorm.DB, account models.Account, user models.User) *AccountService {
	return &AccountService{db: db, account: account, user: user}
}

func (s *AccountService) CreateUserAccount(data *Dto.SignUpDto, ctx *gin.Context) (*fiber.Map, error) {

	// check if User has already registered an Account
	userData, err := s.user.GetByEmail(data.Email, database.DB)

	//if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	//	// continue if User doesn't exist
	//} else {
	//	return &fiber.Map{
	//		"error": "Internal Server Error",
	//	}, err
	//}

	// check if User has already registered an Account
	if userData != nil {
		// User already owns an Account
		if userData.Permission == "owner" {
			return &fiber.Map{
				"inputError": "email",
				"message":    "You have already registered an Account",
			}, nil
		}

		// flag for authController to notify onboarding ui
		// that the User's existing Account was used
		duplicateUser := true
		hasPassword := userData.HasPassword

		// save the new password if it exists and User doesn't have one
		if !hasPassword && data.Password != "" {
			if err := s.user.SavePassword(userData.ID, data.Password); err != nil {
				return &fiber.Map{
					"error": "Internal Server Error",
				}, err
			}
		}

		ctx.Set("duplicate_user", duplicateUser)
		ctx.Set("has_password", hasPassword)
	}
	//permission := func() string {
	//	if userData != nil {
	//		return userData.Permission
	//	}
	//	return "owner"
	//}()
	logger.Info(context.Background(), "CreateAccount")

	//Account := new(models.Account)
	// create the User and assign to Account
	newUser, err := s.user.CreateUser(*data, ctx, database.DB)
	if err != nil {
		return &fiber.Map{
			"error": "Internal Server Error",
		}, err
	}

	//var token *models.Token

	tokenData, err := s.user.GenerateToken(*newUser)
	if err != nil {
		return &fiber.Map{
			"error": "Internal Server Error",
		}, err
	}

	//c.Locals("token", tokenData)
	var newData = &models.Token{
		Provider:         "app",
		Access:           tokenData.Access,
		TokenID:          tokenData.TokenID,
		Refresh:          tokenData.Refresh,
		UserID:           tokenData.UserID,
		CodeCreateAt:     time.Time{},
		CodeExpiresIn:    tokenData.CodeExpiresIn,
		AccessCreateAt:   time.Time{},
		AccessExpiresIn:  tokenData.AccessExpiresIn,
		RefreshCreateAt:  time.Time{},
		RefreshExpiresIn: tokenData.RefreshExpiresIn,
	}

	logger.Info(context.Background(), fmt.Sprintf("newData: %v", newData))
	err = newData.Save(database.DB)
	if err != nil {
		return &fiber.Map{
			"error": err.Error(),
		}, err
	}

	err = s.user.Login(ctx, database.DB)
	if err != nil {
		return &fiber.Map{
			"error": err.Error(),
		}, err
	}

	mail := new(models.EmailContent)
	// send welcome email
	//if err := mail.Send(newUser.Email, "new-Account", userData.ToJson()); err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": "Internal Server Error",
	//	})
	//}
	if err := mail.SendEmailToTerminal(newUser.Email); err != nil {
		return &fiber.Map{
			"error": "Internal Server Error",
		}, err
	}

	responseData := s.user.GenerateUserResponse(newData)

	return &fiber.Map{
		"message": "Account created successfully",
		"data":    responseData,
	}, nil
}

func (s *AccountService) CreateBusinessAccount(data *Dto.AddAccountDto, ctx *gin.Context) (*Dto.UserAccountResponse, error) {
	// get the User from the context
	userID, ok := ctx.Get("User")
	if !ok {
		return nil, errors.New("User not found")
	}
	logger.Info(context.Background(), fmt.Sprintf("User: %v", userID))

	// get the User's Account
	account := new(models.Account)
	user := new(models.User)

	userData, err := user.GetUserById(userID.(string), database.DB)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	// create the business Account
	businessAccount, err := account.CreateBusinessAccount(userData, *data, database.DB)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}
	return businessAccount.GenerateUserAccountResponse(database.DB), nil
}

func (s *AccountService) SwitchUserAccount(accountId, userId string) (*models.User, error) {

	// switch the User's Account
	if err := s.user.SwitchAccount(accountId, userId, database.DB); err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	// get the User's Account
	userData, err := s.user.GetUserById(userId, database.DB)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return userData, nil
}

func (s *AccountService) GetAccount(ctx *gin.Context) (*Dto.UserAccountResponse, error) {
	// get the User from the context
	userID, ok := ctx.Get("User")
	if !ok {
		return nil, errors.New("User not found")
	}
	accountId := ctx.Param("accountId")
	logger.Info(context.Background(), fmt.Sprintf("User: %v", userID))
	logger.Info(context.Background(), fmt.Sprintf("accountId: %v", accountId))
	logger.Info(context.Background(), fmt.Sprintf("User: %v", userID))

	// get the User's Account
	userModel := new(models.User)

	userData, err := userModel.GetUserById(userID.(string), database.DB)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	if accountId != "" {
		for _, account := range userData.Accounts {
			if account.ID == accountId {
				return account.GenerateUserAccountResponse(database.DB), nil
			}
		}
	}

	for _, account := range userData.Accounts {
		if account.ID == userData.ActiveAccountID {
			return account.GenerateUserAccountResponse(database.DB), nil
		}
	}

	return nil, errors.New("no Account found")
}

func (s *AccountService) GetAccounts(ctx *gin.Context) (*[]models.Account, error) {
	// get the User from the context
	userID, ok := ctx.Get("User")
	if !ok {
		return nil, errors.New("User not found")
	}
	logger.Info(context.Background(), fmt.Sprintf("User: %v", userID))

	// get the User's Account
	userModel := new(models.User)

	userData, err := userModel.GetUserById(userID.(string), database.DB)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return &userData.Accounts, nil
}

func (s *AccountService) MakeAccountDefault(accountId string, ctx *gin.Context) (*Dto.UserAccountResponse, error) {
	// get the User from the context
	userID, ok := ctx.Get("User")
	if !ok {
		return nil, errors.New("User not found")
	}
	logger.Info(context.Background(), fmt.Sprintf("User: %v", userID))

	// get the User's Account
	var user models.User
	userData := s.db.Preload("Accounts").Where("id = ?", userID.(string)).First(&user)
	if userData.Error != nil {
		sentry.CaptureException(userData.Error)
		return nil, userData.Error
	}

	for _, account := range user.Accounts {
		if account.ID == accountId {
			//User.DefaultAccountID = Account.ID
			if err := s.db.Save(&user).Error; err != nil {
				sentry.CaptureException(err)
				return nil, err
			}
			return account.GenerateUserAccountData(&user), nil
		}
	}

	return nil, errors.New("no Account found")
}

// Follow an account
func (s *AccountService) Follow(followerID, followingID string) error {
	// Add logic to prevent an account from following itself
	if followerID == followingID {
		return errors.New("cannot follow oneself")
	}

	// Check if the relationship already exists
	var follow models.Follow
	if err := s.db.Where("follower_id = ? AND following_id = ?", followerID, followingID).First(&follow).Error; err == nil {
		return errors.New("already following")
	}

	// Create the relationship
	newFollow := models.Follow{
		ID:          uuid.New().String(),
		FollowerID:  followerID,
		FollowingID: followingID,
		CreatedAt:   time.Now(),
	}
	return s.db.Create(&newFollow).Error
}

// Unfollow an account
func (s *AccountService) Unfollow(followerID, followingID string) error {
	return s.db.Where("follower_id = ? AND following_id = ?", followerID, followingID).Delete(models.Follow{}).Error
}

// ListFollowers List followers of an account
func (s *AccountService) ListFollowers(accountID string) ([]models.Account, error) {
	var followers []models.Account
	err := s.db.Joins("JOIN follows ON follows.follower_id = accounts.id").
		Where("follows.following_id = ?", accountID).
		Find(&followers).Error
	return followers, err
}

// ListFollowing List accounts being followed by an account
func (s *AccountService) ListFollowing(accountID string) ([]models.Account, error) {
	var following []models.Account
	err := s.db.Joins("JOIN follows ON follows.following_id = accounts.id").
		Where("follows.follower_id = ?", accountID).
		Find(&following).Error
	return following, err
}
