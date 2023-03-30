package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"placio-app/models"
	"testing"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	// Setup
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	db = database

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.GeneralSettings{}, models.NotificationsSettings{}, models.AccountSettings{}, models.ContentSettings{})

	// Run tests
	exitCode := m.Run()

	// Tear down
	db.Migrator().DropTable(&models.User{}, &models.Account{}, &models.GeneralSettings{})
	os.Exit(exitCode)
}
