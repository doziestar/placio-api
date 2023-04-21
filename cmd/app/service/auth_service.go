package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"placio-app/Dto"
	"placio-app/database"
	errs "placio-app/errors"
	"placio-app/models"
	"placio-pkg/logger"
)

type IAuth interface {
	LogIn(data Dto.LoginDto) (Dto.UserResponseDto, error)
	LogOut()
	SignUp(data Dto.SignUpDto) (Dto.UserResponseDto, error)
	ResetPassword()
	ChangePassword()
	VerifyEmail()
	VerifyPhone()
	// VerifyID()
	// VerifyAddress()
	// VerifyDocument()
	// VerifyFace()
	// VerifyFingerPrint()
	// VerifyVoice()
	// VerifySignature()
	// VerifyOTP()
	// VerifyPIN()
	// VerifyBiometric()
	// VerifySecurityQuestion()
	// VerifySecurityAnswer()
	// VerifySecurityCode()
	// VerifySecurityPhrase()
}

type Auth struct {
	db *gorm.DB
	models.User
}

func NewAuthService(db *gorm.DB, user *models.User) IAuth {
	return &Auth{db, *user}
}

func (a *Auth) LogIn(data Dto.LoginDto) (Dto.UserResponseDto, error) {
	// check data is valid
	email, password := data.IsValid()
	if email == nil || password == nil {
		return Dto.UserResponseDto{}, errs.ErrInvalid
	}

	// check User exists
	user := a.db.Where("email = ?", email).First(&a.User)
	if user.Error != nil {
		return Dto.UserResponseDto{}, errs.ErrNotFound
	}

	// check password
	if !a.User.ComparePassword(password.(string)) {
		return Dto.UserResponseDto{}, errs.ErrInvalid
	}

	// get User details
	//userDetails := a.DB.Model(&a.User).Association("UserDetails").Find(&a.User.UserDetails)
	userData := models.User{
		ID:    a.User.ID,
		Name:  a.User.Name,
		Email: a.User.Email,
	}
	// generate token
	token, err := a.User.GenerateToken(userData)
	if err != nil {
		return Dto.UserResponseDto{}, errs.ErrInternal
	}

	// return User details
	return Dto.UserResponseDto{
		ID:                    a.User.ID,
		Email:                 a.User.Email,
		Name:                  a.User.Name,
		AccessToken:           token.Access,
		RefreshToken:          token.Refresh,
		AccessTokenExpiresIn:  token.AccessExpiresIn,
		RefreshTokenExpiresIn: token.RefreshExpiresIn,
	}, nil
}

func (a *Auth) LogOut() {

}

func (a *Auth) SignUp(data Dto.SignUpDto) (userRespDto Dto.UserResponseDto, err error) {
	logger.Info(context.Background(), data.Email)

	// Check if User already exists
	var user models.User
	err = database.DB.Where("email = ?", data.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User does not exist, so create a new User
			newUser := &models.User{
				Name:     data.Name,
				Email:    data.Email,
				Password: data.Password,
			}
			err = database.DB.Create(newUser).Error
			if err != nil {
				err = errs.ErrInternal
				return
			}

			// generate token
			token, tokenErr := newUser.GenerateToken(*newUser)
			if tokenErr != nil {
				err = errs.ErrInternal
				return
			}

			userRespDto = Dto.UserResponseDto{
				ID:                    newUser.ID,
				Email:                 newUser.Email,
				Name:                  newUser.Name,
				AccessToken:           token.Access,
				RefreshToken:          token.Refresh,
				AccessTokenExpiresIn:  token.AccessExpiresIn,
				RefreshTokenExpiresIn: token.RefreshExpiresIn,
			}

			return
		} else {
			err = errs.ErrInternal
			return
		}
	}

	// User already exists, return error
	err = errs.ErrAlreadyExists
	return userRespDto, err
}

func (a *Auth) ResetPassword() {

}

func (a *Auth) ChangePassword() {

}

func (a *Auth) VerifyEmail() {

}

func (a *Auth) VerifyPhone() {

}
