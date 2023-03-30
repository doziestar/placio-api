package service

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"placio-app/Dto"
	"placio-app/database"
	"placio-app/models"
	"placio-pkg/logger"
	"time"
)

type IAccountService interface {
	CreateUserAccount(data *Dto.SignUpDto, ctx *fiber.Ctx) (*fiber.Map, error)
}

type AccountService struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{db: db}
}

func (s *AccountService) CreateUserAccount(data *Dto.SignUpDto, ctx *fiber.Ctx) (*fiber.Map, error) {
	user := new(models.User)

	// check if user has already registered an account
	userData, _ := user.GetByEmail(data.Email, database.DB)

	//if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	//	// continue if user doesn't exist
	//} else {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": "Internal Server Error",
	//	})
	//}

	logger.Info(context.Background(), fmt.Sprintf("userData: %v", userData))

	if userData != nil {
		// user already owns an account
		if userData.Permission == "owner" {
			return &fiber.Map{
				"inputError": "email",
				"message":    "You have already registered an account",
			}, nil
		}

		// flag for authController to notify onboarding ui
		// that the user's existing account was used
		duplicateUser := true
		hasPassword := userData.HasPassword

		// save the new password if it exists and user doesn't have one
		if !hasPassword && data.Password != "" {
			if err := user.SavePassword(userData.ID, data.Password); err != nil {
				return &fiber.Map{
					"error": "Internal Server Error",
				}, err
			}
		}

		ctx.Locals("duplicate_user", duplicateUser)
		ctx.Locals("has_password", hasPassword)
	}
	//permission := func() string {
	//	if userData != nil {
	//		return userData.Permission
	//	}
	//	return "owner"
	//}()
	logger.Info(context.Background(), "CreateAccount")

	//account := new(models.Account)

	// create the user and assign to account
	newUser, err := user.CreateUser(*data, ctx, database.DB)
	if err != nil {
		return &fiber.Map{
			"error": "Internal Server Error",
		}, err
	}

	//var token *models.Token

	tokenData, err := user.GenerateToken(*newUser)
	if err != nil {
		return &fiber.Map{
			"error": "Internal Server Error",
		}, err
	}

	//c.Locals("token", tokenData)
	var newData = &models.Token{
		Provider: "app",
		Jwt:      tokenData.Access,
		Access:   tokenData.Access,
		TokenID:  tokenData.TokenID,
		//AccessTokenExpiry: tokenData.AccessExpiresIn,
		Refresh:          tokenData.Refresh,
		UserID:           tokenData.UserID,
		CodeCreateAt:     time.Time{},
		CodeExpiresIn:    tokenData.CodeExpiresIn,
		AccessCreateAt:   time.Time{},
		AccessExpiresIn:  tokenData.AccessExpiresIn,
		RefreshCreateAt:  time.Time{},
		RefreshExpiresIn: tokenData.RefreshExpiresIn,
		ProviderID:       "",
	}

	logger.Info(context.Background(), fmt.Sprintf("newData: %v", newData))
	err = newData.Save(database.DB)
	if err != nil {
		return &fiber.Map{
			"error": err.Error(),
		}, err
	}

	err = user.Login(ctx, database.DB)
	if err != nil {
		return &fiber.Map{
			"error": err.Error(),
		}, err
	}

	mail := new(models.EmailContent)
	// send welcome email
	//if err := mail.Send(newUser.Email, "new-account", userData.ToJson()); err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": "Internal Server Error",
	//	})
	//}
	if err := mail.SendEmailToTerminal(newUser.Email); err != nil {
		return &fiber.Map{
			"error": "Internal Server Error",
		}, err
	}

	responseData := user.GenerateUserResponse(newData)

	return &fiber.Map{
		"message": "Account created successfully",
		"data":    responseData,
	}, nil
}
