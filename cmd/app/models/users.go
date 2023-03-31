package models

import (
	"context"
	"errors"
	"fmt"
	"os"
	"placio-app/Dto"
	"placio-app/database"
	errs "placio-app/errors"
	"placio-pkg/logger"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db = database.DB

type User struct {
	ID                   string     `gorm:"primaryKey,unique,column:id"`
	CreatedAt            time.Time  `gorm:"column:created_at"`
	UpdatedAt            time.Time  `gorm:"column:updated_at"`
	DeletedAt            *time.Time `gorm:"column:deleted_at"`
	Fingerprint          string     `gorm:"column:fingerprint"`
	Name                 string     `gorm:"column:name"`
	Email                string     `gorm:"unique,column:email"`
	Password             string     `gorm:"column:password"`
	DateCreated          time.Time  `gorm:"column:date_created"`
	LastActive           time.Time  `gorm:"column:last_active"`
	Disabled             bool       `gorm:"column:disabled"`
	SupportEnabled       bool       `gorm:"column:support_enabled"`
	TwoFactorAuthEnabled bool       `gorm:"column:2fa_enabled"`
	TwoFASecret          string     `gorm:"column:2fa_secret"`
	TwoFABackupCode      string     `gorm:"column:2fa_backup_code"`
	DefaultAccount       string     `gorm:"column:default_account"`
	FacebookID           string     `gorm:"column:facebook_id"`
	TwitterID            string     `gorm:"column:twitter_id"`
	DefaultAccountID     string     `gorm:"column:default_account_id"`
	Accounts             []Account  `gorm:"foreignKey:UserID"`
	IP                   string     `gorm:"column:ip"`
	UserAgent            string     `gorm:"column:user_agent"`
	Twitter              *TwitterAccount
	Facebook             *FacebookAccount
	Google               *GoogleAccount
	HasPassword          bool   `gorm:"column:has_password"`
	Onboarded            bool   `gorm:"column:onboarded"`
	AccountID            string `gorm:"column:account_id"`
	Permission           string `gorm:"column:permission"`
	GeneralSettingsID    string
	GeneralSettings      GeneralSettings
}

type TwitterAccount struct {
	AccessToken  string
	RefreshToken string
	UserID       string `gorm:"column:user_id"`
	UserName     string `gorm:"column:user_name"`
	CodeVerifier string `gorm:"column:code_verifier"`
	State        string
	Name         string
	DateCreated  time.Time `gorm:"column:date_created"`
	ExpiresIn    time.Time `gorm:"column:expires_in"`
}

type FacebookAccount struct {
	AccessToken  string
	RefreshToken string
	UserID       string `gorm:"column:user_id"`
	UserName     string `gorm:"column:user_name"`
	CodeVerifier string `gorm:"column:code_verifier"`
	State        string
	Name         string
	DateCreated  time.Time `gorm:"column:date_created"`
	ExpiresIn    time.Time `gorm:"column:expires_in"`
}

type GoogleAccount struct {
	AccessToken  string
	RefreshToken string
	UserID       string `gorm:"column:user_id"`
	Email        string
	DateCreated  time.Time
}

type Social struct {
	Provider string
	ID       string
}

// func DecryptFingerprint(fingerprint string) (string, error) {
//     return crypto.Decrypt(fingerprint)
// }

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// u.ID = uuid.New().String()
	//err = u.HashPassword()
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (u *User) GenerateUserFields(userData Dto.SignUpDto, c *fiber.Ctx) {
	// Generate and set fields for the user
	logger.Info(context.Background(), "Generating user fields")
	u.ID = GenerateID()

	// Set the creation and last active dates
	now := time.Now()
	u.DateCreated = now
	u.LastActive = now

	// Set default values for some fields
	u.SupportEnabled = false
	u.TwoFactorAuthEnabled = false
	u.HasPassword = false
	u.Onboarded = false
	u.Disabled = false
	u.Fingerprint = c.Get("fingerprint")
	u.TwoFABackupCode = ""
	u.TwoFASecret = ""
	u.Permission = "user"
	//user.AccountID = ""
	u.UserAgent = c.Get("user-agents")
	u.IP = c.IP()
	u.Email = userData.Email
	u.Name = userData.Name
	u.Password = userData.Password

}

func (u *User) EncryptPassword() error {
	// Encrypt the password if present
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
		u.HasPassword = true
	}
	return nil
}

func (u *User) CreateUser(userData Dto.SignUpDto, c *fiber.Ctx, db *gorm.DB) (*User, error) {
	// Generate and set fields for the user
	u.GenerateUserFields(userData, c)

	// Encrypt the password if present
	err := u.EncryptPassword()

	// Create a new general settings record in the database
	var settings GeneralSettings
	userSettings, err := settings.CreateGeneralSettings(u.ID, db)
	if err != nil {
		return &User{}, err
	}

	u.GeneralSettingsID = userSettings.ID

	// Create a new user record in the database
	err = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	// Create a new account record in the database
	var accountRecord Account
	account, err := accountRecord.CreateAccount(u.ID, "owner", userData.AccountType, db)
	if err != nil {
		return &User{}, err
	}

	//// Update the account record with the account ID
	err = db.Model(&accountRecord).Update("account_id", account.ID).Update("default_account_id", account.ID).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func Get(id uuid.UUID, email string, account string, social *Social, permission string) ([]User, error) {
	var users []User
	var cond = make(map[string]interface{})

	if account != "" {
		cond["accounts.id"] = account
	}
	if permission != "" {
		cond["accounts.permission"] = permission
	}

	if social != nil {
		providerIDField := fmt.Sprintf("%s_id", social.Provider)
		cond[providerIDField] = social.ID

		err := database.DB.Joins("JOIN accounts ON users.id = accounts.user_id").
			Where(cond).Find(&users).Error
		if err != nil {
			return nil, err
		}
	} else {
		if id != uuid.Nil {
			cond["id"] = id
		}
		if email != "" {
			cond["email"] = email
		}

		err := database.DB.Joins("JOIN accounts ON users.id = accounts.user_id").
			Where(cond).Find(&users).Error
		if err != nil {
			return nil, err
		}
	}

	for i, user := range users {
		user.Accounts = nil
		user.HasPassword = user.Password != ""
		user.Password = ""
		if account != "" {
			user.AccountID = account
		} else {
			user.AccountID = user.DefaultAccount
		}
		for _, account := range user.Accounts {
			if account.ID == user.AccountID {
				user.Permission = account.Permission
				user.Onboarded = account.Onboarded
				break
			}
		}
		users[i] = user
	}

	return users, nil
}

func (u *User) GetUserById(id string, db *gorm.DB) (*User, error) {
	err := db.Model(&User{}).Where("id = ?", id).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	return u, nil
}

// GetAccounts returns a list of accounts that the user with the given ID is attached to
func GetAccounts(db *gorm.DB, userID string) ([]Account, error) {
	var userAccounts []Account

	// Join the User and Account tables to get the account data
	err := db.Table("users").
		Select("users.id, users.email, account.id AS account_id, account.permission, account.name").
		Joins("LEFT JOIN account ON users.account_id = account.id").
		Where("users.id = ?", userID).
		Find(&userAccounts).Error

	if err != nil {
		return nil, err
	}

	// Group the accounts by ID and format the results
	accountMap := make(map[string][]Account)
	for _, account := range userAccounts {
		accountMap[account.ID] = append(accountMap[account.ID], account)
	}

	var results []Account
	for _, accounts := range accountMap {
		result := Account{
			ID:         accounts[0].ID,
			UserID:     accounts[0].UserID,
			Permission: accounts[0].Permission,
		}
		results = append(results, result)
	}

	return results, nil
}

// addInterest adds the specified interest to the user's account
func AddInterest(userID string, accountID string, interest string) error {
	// get user by ID
	var user User
	result := database.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return result.Error
	}

	// find account by ID
	var account Account
	result = database.DB.Where("id = ?", accountID).First(&account)
	if result.Error != nil {
		return result.Error
	}

	// add interest to account
	// account.Interests = append(account.Interests, interest)

	// save account
	result = database.DB.Save(&account)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// updateInterest updates the specified interest in the user's account
func UpdateInterest(userID string, accountID string, oldInterest string, newInterest string) error {
	// get user by ID
	var user User
	result := database.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return result.Error
	}

	// find account by ID
	var account Account
	result = database.DB.Where("id = ?", accountID).First(&account)
	if result.Error != nil {
		return result.Error
	}

	// find index of old interest in account interests
	var index = -1
	// for i, v := range account.Interests {
	// 	if v == oldInterest {
	// 		index = i
	// 		break
	// 	}
	// }

	// if old interest is not found, return an error
	if index == -1 {
		return errors.New("old interest not found")
	}

	// replace old interest with new interest
	// account.Interests[index] = newInterest

	// save account
	result = database.DB.Save(&account)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// AddAccount assigns a user to an account with a specified permission
func (u *User) AddAccount(db *gorm.DB, accountID string, permission string) error {
	var user User
	result := db.Where("id = ?", u.ID).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("No user with that ID")
	}
	if result.Error != nil {
		return result.Error
	}

	account := Account{
		ID:         accountID,
		UserID:     u.ID,
		Permission: permission,
		Onboarded:  false,
	}

	user.Accounts = append(user.Accounts, account)

	result = db.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteAccount removes a user from an account
func DeleteAccount(userID string, accountID string) error {
	var user User
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return errors.New("No user with that ID")
	}

	// Find the index of the account with the given ID
	var accountIndex int = -1
	for i, a := range user.Accounts {
		if a.ID == accountID {
			accountIndex = i
			break
		}
	}

	if accountIndex == -1 {
		return errors.New("User is not attached to that account")
	}

	user.Accounts = append(user.Accounts[:accountIndex], user.Accounts[accountIndex+1:]...)

	return db.Save(&user).Error
}

// Password returns the hashed password for the specified user and account
func Password(db *gorm.DB, userID uint, accountID uint) (string, error) {
	var user User
	result := db.Joins("Account").First(&user, "users.id = ? AND accounts.id = ?", userID, accountID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("user with id %d and account id %d not found", userID, accountID)
		}
		return "", result.Error
	}

	return user.Password, nil
}

func (u *User) VerifyPassword(plainPassword string) (bool, error) {
	//result := db.Joins("JOIN accounts ON users.account_id = accounts.id").Where("users.id = ? AND accounts.id = ?", userID, accountID).First(u)
	//
	//if result.Error != nil {
	//	return false, result.Error
	//}
	logger.Info(context.Background(), fmt.Sprintf("password: %s", plainPassword))
	logger.Info(context.Background(), fmt.Sprintf("u.Password: %s", u.Password))
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword))
	if err != nil {
		logger.Error(context.Background(), fmt.Sprintf("error comparing password hashes: %s", err))
		return false, err
	}

	u.Password = ""
	return true, nil
}

// SavePassword saves a new password for the user with the given ID.
// If not executed via a password reset request, the user is notified
// by email that their password has been changed.
// passwordReset determines if password update is part of reset.
//func (u *User) SavePassword(id string, password string, reset bool) error {
//	// Encrypt the password.
//	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
//	if err != nil {
//		return err
//	}
//
//	// Save the password hash in the database.
//	user := &User{}
//	result := db.Where("id = ?", id).First(user)
//	if result.Error != nil {
//		return result.Error
//	}
//
//	result = db.Model(user).Update("password", hash)
//	if result.Error != nil {
//		return result.Error
//	}
//
//	return nil
//}

// func (u *User) SaveBackupCode(db *pg.DB, code string) error {
//     u.BackupCode = code
//     return u.Update(db)
// }

// func (u *User) VerifyBackupCode(db *pg.DB, code string) bool {
//     return u.BackupCode == code
// }

// func (u *User) SaveTwoFactorToken(db *pg.DB, token string) error {
//     u.TwoFactorToken = token
//     return u.Update(db)
// }

// func (u *User) VerifyTwoFactorToken(db *pg.DB, token string) bool {
//     return u.TwoFactorToken == token
// }

// func (u *User) SaveTwoFactorAuth(db *pg.DB, enabled bool) error {
//     u.TwoFactorAuth = enabled
//     return u.Update(db)
// }

// func (u *User) VerifyTwoFactorAuth(db *pg.DB) bool {
//     return u.TwoFactorAuth
// }

// // SaveBackupCode saves the hashed backup code for a user
// func SaveBackupCode(db *pg.DB, id string, code string) error {
//     // Hash the backup code
//     hash, err := bcrypt.GenerateFromPassword([]byte(code), 10)
//     if err != nil {
//         return err
//     }

//     // Update the user with the hashed backup code
//     user := &User{ID: id}
//     _, err = db.Model(user).
//         Set("two_factor_auth.backup_code = ?", hash).
//         Update()
//     if err != nil {
//         return err
//     }

//     return nil
// }

// // VerifyBackupCode verifies the provided backup code for a user
// func VerifyBackupCode(db *pg.DB, id string, email string, account string, code string) bool {
//     // Get the user by ID or email and account
//     var user User
//     query := db.Model(&user).
//         Where("id = ?", id).
//         Where("email = ?", email).
//         Where("account_id = ?", account).
//         Select()

//     if query.Err() != nil {
//         return false
//     }

//     // Compare the provided backup code with the user's hashed backup code
//     err := bcrypt.CompareHashAndPassword(user.TwoFactorAuth.BackupCode, []byte(code))
//     if err != nil {
//         return false
//     }

//     return true
// }

// // GetSecret returns the 2FA secret for a user
// func GetSecret(db *pg.DB, id string, email string) (string, error) {
//     // Get the user by ID or email
//     var user User
//     query := db.Model(&user).
//         Where("id = ?", id).
//         Where("email = ?", email).
//         Select()

//     if query.Err() != nil {
//         return "", query.Err()
//     }

//     // Decrypt and return the 2FA secret
//     secret, err := Decrypt(user.TwoFactorAuth.Secret)
//     if err != nil {
//         return "", err
//     }

//     return secret, nil
// }

// // Decrypt decrypts a 2FA secret
// func Decrypt(ciphertext string) (string, error) {
//     // TODO: Implement decryption
//     return "", nil
// }

// UpdateUserProfile updates the user profile
func UpdateUserProfile(id string, accountID string, data map[string]interface{}) error {
	// Update nested objects
	if onboarded, ok := data["onboarded"].(bool); ok {
		var user User
		result := db.Where("id = ? AND account.id = ?", id, accountID).First(&user)
		if result.Error != nil {
			return result.Error
		}
		if user.ID == "" {
			return errors.New("No user with that ID")
		}
		user.Accounts[0].Onboarded = onboarded
		result = db.Save(&user)
		if result.Error != nil {
			return result.Error
		}
	} else if permission, ok := data["permission"].(string); ok {
		var user User
		result := db.Where("id = ? AND account.id = ?", id, accountID).First(&user)
		if result.Error != nil {
			return result.Error
		}
		if user.ID == "" {
			return errors.New("No user with that ID")
		}
		user.Accounts[0].Permission = permission
		result = db.Save(&user)
		if result.Error != nil {
			return result.Error
		}
	} else {
		result := db.Model(&User{}).Where("id = ? AND account.id = ?", id, accountID).Updates(data)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

// DeleteUser deletes the user with the given ID and account ID
func DeleteUser(id string, account string) error {
	return db.Where("id = ? AND account.id = ?", id, account).Delete(&User{}).Error
}

// UpdateTwitter updates the Twitter data for the user with the given ID
func UpdateTwitter(id string, data interface{}) error {
	return db.Model(&User{}).Where("id = ?", id).Update("twitter", data).Error
}

// HashPassword hashes the password
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errs.ErrForbidden
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
	tokenID := GenerateID()
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	//claims["name"] = user.Name
	//claims["email"] = user.Email
	////claims["account"] = user.Accounts[0].ID
	//claims["permissions"] = user.Accounts[0].Permission
	//claims["onboarded"] = user.Accounts[0].Onboarded
	//claims["provider"] = "app"
	//claims["accountID"] = user.Accounts[0].ID
	claims["jti"] = tokenID
	claims["iat"] = time.Now().UTC().Unix()
	claims["exp"] = time.Now().UTC().Add(time.Hour * 24 * 7).Unix()

	// Generate access and refresh tokens
	accessToken, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return Dto.Token{}, errs.ErrForbidden
	}

	refreshToken := uuid.NewString()

	// Set token expiration times
	accessCreateAt := time.Now().UTC()
	accessExpiresIn := accessCreateAt.Add(time.Hour * 24 * 7)
	refreshCreateAt := accessCreateAt
	refreshExpiresIn := refreshCreateAt.Add(time.Hour * 24 * 30)

	return Dto.Token{
		UserID:           user.ID,
		TokenID:          tokenID,
		CodeCreateAt:     time.Time{},
		CodeExpiresIn:    time.Duration(accessExpiresIn.Unix()),
		Access:           accessToken,
		AccessCreateAt:   accessCreateAt,
		AccessExpiresIn:  time.Duration(accessExpiresIn.Unix()),
		Refresh:          refreshToken,
		RefreshCreateAt:  refreshCreateAt,
		RefreshExpiresIn: time.Duration(refreshExpiresIn.Unix()),
	}, nil
}

func (u *User) GenerateUserResponse(token *Token) Dto.UserResponse {
	return Dto.UserResponse{
		User: &Dto.User{
			ID:          u.ID,
			Name:        u.Name,
			Email:       u.Email,
			Disabled:    false,
			HasPassword: false,
			Onboarded:   false,
			Account: func(db *gorm.DB) []Dto.Account {
				var account []Account
				// find accounts with user id
				err := db.Where("user_id = ?", u.ID).First(&account).Error
				if err != nil {
					return []Dto.Account{}
				}
				// convert to dto
				var dto []Dto.Account
				for _, a := range account {
					dto = append(dto, Dto.Account{
						ID:          a.ID,
						Permission:  a.Permission,
						AccountType: a.AccountType,
						AccountID:   a.AccountID,
						Onboarded:   a.Onboarded,
						//Interests:   a.Interests,
						UserID:   a.UserID,
						Plan:     a.Plan,
						Active:   a.Active,
						Status:   a.Status,
						Disabled: a.Disabled,
					})
				}
				return dto
			}(database.DB),
			Permission: "",
			GeneralSettings: func(db *gorm.DB) Dto.GeneralSettings {
				var settings GeneralSettings
				// find settings with user id
				err := db.Where("ID = ?", u.GeneralSettingsID).First(&settings).Error
				if err != nil {
					return Dto.GeneralSettings{}
				}
				// convert to dto
				return Dto.GeneralSettings{
					ID:       settings.ID,
					Language: settings.Language,
					Theme:    settings.Theme,
				}
			}(database.DB),
		},
		Token: &Dto.UserToken{
			UserID:           u.ID,
			Access:           token.Access,
			AccessExpiresIn:  int64(token.AccessExpiresIn),
			Refresh:          token.Refresh,
			RefreshExpiresIn: int64(token.RefreshExpiresIn),
		},
	}
}

func (u *User) Login(c *fiber.Ctx, db *gorm.DB) error {
	var login *Login

	err := db.Where("email = ?", u.Email).First(&login).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			login = &Login{
				IP:      c.IP(),
				Browser: c.Get("User-Agent"),
				Time:    time.Now(),
				Device:  c.Get("Device"),
				UserID:  u.ID,
			}
			err := db.Create(&login).Error
			if err != nil {
				return err
			}
		}
	}

	login.IP = c.IP()
	login.Browser = c.Get("User-Agent")
	login.Time = time.Now()
	login.Device = c.Get("Device")
	login.UserID = u.ID

	err = db.Save(&login).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetByEmail(email string, db *gorm.DB) (*User, error) {
	var user *User
	err := db.Preload("Accounts").Preload("GeneralSettings").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	logger.Info(context.Background(), fmt.Sprintf("User with email %s found", user))
	return user, err
}

func (u *User) SavePassword(id string, password string) error {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	user.Password = password
	err = db.Save(&user).Error
	return err
}

func (u *User) Create(userData *Dto.SignUpDto, ctx *fiber.Ctx, db *gorm.DB) (*User, error) {
	user := &User{
		ID:        uuid.NewString(),
		Email:     userData.Email,
		Name:      userData.Name,
		IP:        ctx.IP(),
		UserAgent: ctx.Get("User-Agent"),

		Password: userData.Password,
	}
	err := db.Create(&user).Error
	return user, err

}

func (u *User) AddToAccount(id string, id2 string, s string) interface{} {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	user.Accounts = append(user.Accounts, Account{
		ID:        id2,
		AccountID: s,
	})
	err = db.Save(&user).Error
	return err

}

func (u *User) ToJson() map[string]string {
	return map[string]string{
		"UserId": u.ID,
		"email":  u.Email,
		"name":   u.Name,
	}

}

func (u *User) UpdateUser(userId string, accountId string, lastActive time.Time, disabled bool) {
	var user User
	db.Where("id = ? AND account.id = ?", userId, accountId).First(&user)
	if user.ID == "" {
		return
	}
	user.Accounts[0].LastActive = lastActive
	user.Accounts[0].Disabled = disabled
	db.Save(&user)

}

//
//func (u *User) Get(id string, s string, id2 string) (interface{}, interface{}) {
//
//}
