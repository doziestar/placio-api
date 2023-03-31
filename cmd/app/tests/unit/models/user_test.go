package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"placio-app/Dto"
	"placio-app/models"
	"testing"
)

var (
	user     models.User
	userData = Dto.SignUpDto{
		Name:        "Test User",
		Email:       "test@example.com",
		Password:    "test_password",
		AccountType: "user",
	}

	app = fiber.New(fiber.Config{})
	ctx = app.AcquireCtx(&fasthttp.RequestCtx{})
)

func TestAccountCreationProcess(t *testing.T) {
	// Setup
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	require.NoError(t, err)

	db.AutoMigrate(&models.User{}, &models.GeneralSettings{}, &models.NotificationsSettings{}, &models.AccountSettings{}, &models.ContentSettings{}, &models.Account{})

	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)

	// Test GenerateUserFields

	user.GenerateUserFields(userData, c)
	assert.Equal(t, userData.Email, user.Email)
	assert.Equal(t, userData.Name, user.Name)
	assert.Equal(t, userData.Password, user.Password)

	// Test EncryptPassword
	err = user.EncryptPassword()
	assert.NoError(t, err)
	assert.NotEqual(t, userData.Password, user.Password)

	// Test CreateGeneralSettings
	generalSettings := &models.GeneralSettings{}
	createdGeneralSettings, err := generalSettings.CreateGeneralSettings(user.ID, db)
	assert.NoError(t, err)
	assert.NotNil(t, createdGeneralSettings)

	// Test CreateAccount
	account := &models.Account{}
	createdAccount, err := account.CreateAccount(user.ID, "owner", userData.AccountType, db)
	assert.NoError(t, err)
	assert.NotNil(t, createdAccount)

	// Test CreateUser
	log.Println("Creating user", user.ID, user.Email, userData.Email)
	createdUser, err := user.CreateUser(userData, c, db)
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)

	// Check if the createdUser has the same email as the test user data
	var foundUser models.User
	db.First(&foundUser, "id = ?", createdUser.ID)
	assert.Equal(t, userData.Email, foundUser.Email)

	// Tear down
	db.Migrator().DropTable(&models.User{}, &models.GeneralSettings{}, &models.NotificationsSettings{}, &models.AccountSettings{}, &models.ContentSettings{}, &models.Account{})
}

func TestGenerateUserFields(t *testing.T) {

	user.GenerateUserFields(userData, ctx)

	assert.Equal(t, user.Email, "test@example.com")
	assert.False(t, user.HasPassword)
	assert.Equal(t, user.Permission, "user")
	assert.Equal(t, user.Name, "Test User")
	assert.Equal(t, user.IP, "0.0.0.0")
}

func TestEncryptPassword(t *testing.T) {
	user.GenerateUserFields(userData, ctx)

	assert.NotNil(t, user.Password)
	assert.Equal(t, user.Password, userData.Password)

	err := user.EncryptPassword()
	assert.NoError(t, err)

	assert.NotEqual(t, userData.Password, user.Password)
}
