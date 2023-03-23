package controller

import (
	"context"
	"errors"
	Dto "placio-app/Dto"
	"placio-app/database"
	errs "placio-app/errors"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"
	"placio-pkg/logger"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	userData *models.User
	useEmail bool
	account  *models.Account
	login    *models.LoginModel
	token    models.Token
)

// Login godoc
// @Summary Login and get token
// @Description Login to the server
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} Dto.UserResponseDto
// @Failure 400 {object} map[string]interface{}
// @QueryParam email query string true "email"
// @QueryParam password query string true "password"
// @Body 200 {object} Dto.LoginDto
// @Router /api/v1/auth/login [get]
func Login(c *fiber.Ctx) error {
	// get data from request body
	data := new(Dto.LoginDto)
	if err := c.BodyParser(data); err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": "invalid data", "data": err})
	}

	// validate data
	email, password := data.IsValid()

	// call service
	auth := service.NewAuth(&database.Database{}, &models.User{
		Email:    email.(string),
		Password: password.(string),
	})

	user, err := auth.LogIn(*data)
	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": "invalid data", "data": err})
	}

	return c.JSON(fiber.Map{"status": "ok", "data": user})

}

// SignUp godoc
// @Summary Sign up
// @Description Sign up to the server
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} Dto.UserResponseDto
// @Failure 400 {object} map[string]interface{}
// @Param user body Dto.SignUpDto true "user"
// @Body 200 {object} Dto.SignUpDto
// @Router /api/v1/auth/signup [post]
func SignUp(c *fiber.Ctx) error {
	var data Dto.SignUpDto
	logger.Info(context.Background(), string(c.Body()))
	if err := c.BodyParser(&data); err != nil {
		logger.Info(context.Background(), data.Email)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "invalid data", "data": err})
	}

	// validate data
	userData, err := data.IsValid()
	if err != nil {
		logger.Info(context.Background(), err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "invalid data", "data": err})
	}
	logger.Info(context.Background(), userData.Email)

	// call service
	auth := service.NewAuth(&database.Database{}, &models.User{
		Email:    userData.Email,
		Password: userData.Password,
		Name:     userData.Name,
		//Role:     userData.Role,
	})

	user, err := auth.SignUp(data)
	if errors.Is(err, errs.ErrAlreadyExists) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "user already exists", "data": err})
	}

	return c.JSON(fiber.Map{"status": "ok", "data": user})
}

// LogOut godoc
// @Summary Log out
// @Description Log out from the server
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/logout [get]
func LogOut(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

// RefreshToken godoc
// @Summary Refresh token
// @Description Refresh token
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/refresh [get]
func RefreshToken(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

// ResetPassword godoc
// @Summary Reset password
// @Description Reset password
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/reset [get]
func ResetPassword(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

// ChangePassword godoc
// @Summary Change password
// @Description Change password
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/change [get]
func ChangePassword(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

// VerifyEmail godoc
// @Summary Verify email
// @Description Verify email
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/verify [get]
func VerifyEmail(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

// VerifyPhone godoc
// @Summary Verify phone
// @Description Verify phone
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/verify [get]
func VerifyPhone(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

// SigninRequest defines the request body for the signin route
type SigninRequest struct {
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	Token        string `json:"token,omitempty"`
	Provider     string `json:"provider,omitempty"`
	ProviderID   string `json:"provider_id,omitempty"`
	MagicViewURL string `json:"magic_view_url,omitempty"`
}

// SigninResponse defines the response body for the signin route
type SigninResponse struct {
	Message       string       `json:"message,omitempty"`
	TwoFARequired bool         `json:"2fa_required,omitempty"`
	Token         models.Token `json:"token,omitempty"`
}

func (s *SigninRequest) IsValid() error {
	if s.Email == "" && s.Token == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Email or token is required")
	}
	if s.Email != "" && s.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Password is required")
	}
	return nil
}

// covert SigninRequest to json
func (s *SigninRequest) ToJson() map[string]interface{} {
	mySigninRequestMap := map[string]interface{}{
		"email":          s.Email,
		"password":       s.Password,
		"token":          s.Token,
		"provider":       s.Provider,
		"provider_id":    s.ProviderID,
		"magic_view_url": s.MagicViewURL,
	}
	return mySigninRequestMap
}

// Signin authenticates a user via email/password or social network
// Signin authenticates a user and returns an access token.
// The function performs the following steps:
// 1. Parses the incoming request body into a SigninRequest.
// 2. Validates the input data.
// 3. Checks if the user exists.
// 4. Verifies the user's password.
// 5. Retrieves the user's account data.
// 6. Logs the sign-in attempt and checks for suspicious activity.
// 7. Generates an access token for the authenticated user.
// 8. Returns the access token in a SigninResponse.
//
// @Summary Authenticate a user
// @Description Authenticate a user and return an access token
// @Tags authentication
// @Accept json
// @Produce json
// @Param SigninRequest body SigninRequest true "Sign In Data"
// @Success 200 {object} SigninResponse "Successfully signed in"
// @Failure 400 {object} fiber.Error "Bad Request"
// @Failure 401 {object} fiber.Error "Unauthorized"
// @Router /api/v1/auth/signin [post]
func Signin(c *fiber.Ctx) error {
	data := new(SigninRequest)
	if err := c.BodyParser(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	var token models.Token
	var user models.User

	if data.Email != "" {
		useEmail = true
		data.Provider = "app"
		if err := utility.Validate(data.ToJson(), []string{"email", "password"}); err != nil {
			return err
		}
	} else {
		if err := utility.Validate(data.ToJson(), []string{"token"}); err != nil {
			return err
		}
		decode, err := token.VerifyToken("app", data.Token)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}
		data.Provider = decode.Provider
		data.ProviderID = decode.ProviderID
		data.Email = decode.Email
	}

	// check user exists
	var err error
	if useEmail {
		userData, err = user.GetByEmail(data.Email)
	} else {
		//userData, err = user.Get(nil, "", nil, map[string]string{
		//	"provider": data.Provider,
		//	"id":       data.ProviderID,
		//})
	}
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Please enter the correct login details")
	}

	// verify password
	if useEmail {
		verified, err := user.VerifyPassword(userData.ID, userData.AccountID, data.Password)
		if err != nil {
			return err
		}
		if !verified {
			return fiber.NewError(fiber.StatusUnauthorized, "Please enter the correct login details")
		}
	}

	// get the account
	accountData, err := account.GetAccount(userData.AccountID)
	if err != nil {
		return err
	}
	if !accountData.Active {
		return fiber.NewError(fiber.StatusUnauthorized, "Your account has been deactivated. Please contact support.")
	}

	// log the sign in and check if it's suspicious
	log, err := login.Create(userData.ID, c.IP(), c.Get("User-Agent"), c.Get("Device"))
	if err != nil {
		return err
	}
	loginVerification, err := login.Verify(userData.ID, log)
	if err != nil {
		return err
	}
	if loginVerification.Suspicious {
		return fiber.NewError(fiber.StatusUnauthorized, "Your account has been flagged for suspicious activity. Please contact support.")
	}

	// generate the token
	userToken, err := token.Generate(userData.ID, userData.AccountID, data.Provider, data.ProviderID, data.Email)
	if err != nil {
		return err
	}

	// return the token
	return c.JSON(SigninResponse{
		Message:       "You have successfully signed in",
		TwoFARequired: userData.TwoFactorAuthEnabled,
		Token:         *userToken,
	})
}

func Authenticate(c *fiber.Ctx, userData models.User) error {
	accountData, err := account.GetAccount(userData.AccountID)
	if err != nil {
		return err
	}

	subscription, err := account.GetSubscription(userData.AccountID)
	if err != nil {
		return err
	}

	userAccounts, err := account.GetUserAccount(userData.ID)
	if err != nil {
		return err
	}

	// create & store the token
	jwt, err := token.Generate(userData.ID, userData.AccountID, "app", userData.AccountID, userData.Email)
	if err != nil {
		return err
	}

	err = token.Save("app", map[string]string{
		"provider_id":  userData.AccountID,
		"email":        userData.Email,
		"accessToken":  jwt.Access,
		"refreshToken": jwt.Refresh,
		//"expires":     jwt.AccessExpiresIn,

	}, userData.ID)
	if err != nil {
		return err
	}

	userData.UpdateUser(userData.ID, userData.AccountID, time.Now(), false)

	// return user to client
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"token":        jwt,
		"subscription": subscription.Status,
		"plan":         accountData.Plan,
		"permission":   userData.Permission,
		"name":         userData.Name,
		"accounts":     userAccounts,
		"account_id":   userData.AccountID,
		"has_password": userData.HasPassword,
		"onboarded":    userData.Onboarded,
	})
}

//func Signout(c *fiber.Ctx) error {
//	// destroy social tokens
//	err := token.Delete(nil, c.Locals("provider").(string), c.Locals("user").(int))
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": "Failed to sign out",
//		})
//	}
//	return c.SendStatus(fiber.StatusOK)
//}

// auth.switch()
// let the user switch account
//func Switch(c *fiber.Ctx) error {
//	// Get the user and account ID from the request
//	userID := c.Locals("user").(string)
//	accountID := c.Params("account")
//
//	// Check if the user belongs to this account
//	userData, err := userData.Get(userID, "", accountID)
//	if err != nil {
//		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
//			"message": "You don't belong to this account.",
//		})
//	}
//
//	return Authenticate(c, userData)
//}

// Impersonate impersonates a user without a password (accessible via master account only)
//func Impersonate(c *fiber.Ctx) error {
//	// Get the token from the request body
//	token := new(struct {
//		Token string `json:"token"`
//	})
//	if err := c.BodyParser(token); err != nil {
//		return errors.Wrap(err, "failed to parse request body")
//	}
//
//	// Verify the authorization token
//	data, err := auth.VerifyToken(token.Token)
//	if err != nil {
//		return errors.Wrap(err, "failed to verify token")
//	}
//
//	// Check if the token is valid and has master permission
//	if data.UserID == "" || data.Permission != "master" {
//		return utility.NewError(fiber.StatusUnauthorized, "Invalid token")
//	}
//
//	// Get the user data
//	userData, err := user.Get(data.UserID)
//	if err != nil {
//		return errors.Wrap(err, "failed to get user data")
//	}
//
//	// Authenticate the user and generate a new token
//	jwt, err := auth.GenerateToken(userData.AccountID, userData.ID, userData.Permission, "app")
//	if err != nil {
//		return errors.Wrap(err, "failed to generate token")
//	}
//
//	// Save the token to the database
//	if err := auth.SaveToken("app", jwt.AccessToken, userData.ID); err != nil {
//		return errors.Wrap(err, "failed to save token to database")
//	}
//
//	// Update user's last active timestamp and disabled status
//	if err := user.Update(userData.ID, userData.AccountID, user.UpdateParams{
//		LastActive: time.Now(),
//		Disabled:   false,
//	}); err != nil {
//		return errors.Wrap(err, "failed to update user")
//	}
//
//	// Return the response
//	return c.JSON(fiber.Map{
//		"token":        jwt.AccessToken,
//		"subscription": "",
//		"plan":         "",
//		"permission":   userData.Permission,
//		"name":         userData.Name,
//		"accounts":     nil,
//		"account_id":   userData.AccountID,
//		"has_password": false,
//		"onboarded":    userData.Onboarded,
//	})
//}

// GetAuthStatus retrieves the auth status of a user
//func GetAuthStatus(c *fiber.Ctx) error {
//	// Check if there's a valid account/user
//	var hasJWT, hasSocialToken, usingSocialSignin bool
//
//	// Check if there's an active JWT
//	if c.Provider() == "app" {
//		usingSocialSignin = false
//		hasJWT = token.Verify("app", c.UserID())
//	}
//
//	// Check if there's an active access token if the user is
//	// signed in via social network or was their account de-authed
//	if c.Provider() != "app" {
//		usingSocialSignin = true
//		hasSocialToken = token.Verify(c.Provider(), c.UserID())
//	}
//
//	// Check if the user has an active subscription
//	subscription := account.Subscription(c.AccountID())
//	userAccounts := user.Accounts(c.UserID())
//	user.UpdateLastActive(c.UserID(), c.AccountID(), time.Now())
//
//	// Return the auth status
//	return c.Status(fiber.StatusOK).JSON(fiber.Map{
//		"data": fiber.Map{
//			"jwt_token":     hasJWT,
//			"social_token":  hasSocialToken,
//			"subscription":  subscription.Status,
//			"accounts":      userAccounts,
//			"account_id":    c.AccountID(),
//			"authenticated": usingSocialSignin && hasSocialToken || !usingSocialSignin && hasJWT,
//		},
//	})
//}

// / MagicVerify verifies a magic token
//func MagicVerify(c *fiber.Ctx) error {
//	data := new(struct {
//		Token string `json:"token"`
//	})
//	if err := c.BodyParser(data); err != nil {
//		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
//	}
//
//	magicToken, err := Verify(data.Token)
//	if err != nil {
//		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
//	}
//
//	userData, err := user.Get(magicToken.UserID)
//	if err != nil {
//		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
//	}
//
//	// log the sign in and check if it's suspicious
//	log, err := login.Create(userData.ID, c)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create login record")
//	}
//	risk, err := login.Verify(userData.ID, log)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).SendString("Failed to verify login risk")
//	}
//
//	// notify the user of suspicious logins
//	if risk.Level > 0 {
//		err = mail.Send(mail.Options{
//			To:        userData.Email,
//			Subject:   "New sign-in on your account",
//			Template:  "new_signin",
//			Variables: mail.Variables{"ip": risk.Flag.IP, "time": risk.Time, "device": risk.Flag.Device, "browser": risk.Flag.Browser},
//		})
//		if err != nil {
//			return c.Status(fiber.StatusInternalServerError).SendString("Failed to send email")
//		}
//	}
//
//	// 2fa is required
//	if userData.TwoFactorEnabled {
//		// notify the client and use email to identify the user when sending otp
//		// send a token so the otp password screen can't be accessed directly without a password
//		jwt, err := Token(TokenData{Email: userData.Email, Provider: "app"}, nil, time.Minute*5)
//		if err != nil {
//			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate JWT token")
//		}
//		return c.JSON(fiber.Map{
//			"2fa_required": true,
//			"token":        jwt,
//		})
//	}
//
//	// authenticate the user
//	userAccounts, err := user.Account(userData.ID)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve user accounts")
//	}
//
//	err = TokenSave("app", userData.ID, TokenData{Access: Token(TokenData{AccountID: userData.AccountID, UserID: userData.ID, Permission: userData.Permission})})
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save token")
//	}
//
//	err = user.Update(userData.ID, userData.AccountID, user.UpdateData{LastActive: time.Now(), Disabled: false})
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update user data")
//	}
//
//	return c.JSON(fiber.Map{
//		"token":        Token(TokenData{AccountID: userData.AccountID, UserID: userData.ID, Permission: userData.Permission}),
//		"subscription": subscription.Status,
//		"plan":         accountData.Plan,
//		"permission":   userData.Permission,
//		"name":         userData.Name,
//		"accounts":     userAccounts,
//		"account_id":   userData.AccountID,
//		"has_password": userData.Password != "",
//		"onboarded":    userData.Onboarded,
//	})
//}

//func SocialAuthHandler(provider string, signInURL string, socialURL string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// Use the passport-go middleware to authenticate the user
//		if err := passport.Authenticate(c, provider, passport.Options{
//			FailureRedirect: signInURL,
//		}); err != nil {
//			c.Redirect(http.StatusFound, fmt.Sprintf("%s?error=%s", signInURL, url.QueryEscape(err.Error())))
//			return
//		}
//
//		// Get the user's profile from the authenticated session
//		profile, ok := passport.ProfileFromContext(c)
//		if !ok {
//			c.Redirect(http.StatusFound, fmt.Sprintf("%s?error=%s", signInURL, url.QueryEscape("Unable to get user profile")))
//			return
//		}
//
//		// Authenticate the user
//		email := ""
//		if len(profile.Emails) > 0 {
//			email = profile.Emails[0].Value
//		}
//		userData, err := user.Get(nil, email, nil, map[string]interface{}{
//			"provider": provider,
//			"id":       profile.ID,
//		})
//		if err != nil {
//			c.Redirect(http.StatusFound, fmt.Sprintf("%s?error=%s", signInURL, url.QueryEscape("Unable to get user data")))
//			return
//		}
//
//		if userData != nil {
//			// Generate a JWT token and redirect the user to the social URL
//			jwt, err := auth.Token(map[string]interface{}{
//				"provider":    provider,
//				"provider_id": profile.ID,
//				"email":       email,
//			}, nil, 300)
//			if err != nil {
//				c.Redirect(http.StatusFound, fmt.Sprintf("%s?error=%s", signInURL, url.QueryEscape("Unable to generate JWT token")))
//				return
//			}
//			c.Redirect(http.StatusFound, fmt.Sprintf("%s?provider=%s&token=%s", socialURL, provider, jwt))
//		} else {
//			c.Redirect(http.StatusFound, fmt.Sprintf("%s?error=%s", signInURL, url.QueryEscape("You're not registered")))
//		}
//	}
//}
