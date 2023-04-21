package integration

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"log"
	"os"
	"placio-app/models"
	"placio-app/service"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	fixtures *testfixtures.Loader
	userSvc  *service.UserService
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	// Connect to the test database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		"postgres-db",
		"dozie",
		"918273645dozie",
		"test_db",
		"5432",
		"disable",
		"UTC",
	)
	var err error
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to test database: %v", err)
	}

	// Create tables in the test database
	err = db.AutoMigrate(&models.User{}, &models.Account{})
	if err != nil {
		log.Fatalf("failed to create tables in test database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get *sql.DB from *gorm.DB: %v", err)
	}

	// Load test fixtures
	fixtures, err = testfixtures.New(
		testfixtures.Database(sqlDB),
		testfixtures.Dialect("sqlite"),
		testfixtures.Directory("fixtures"),
	)

	if err != nil {
		log.Fatalf("failed to set up test fixtures: %v", err)
	}

	// Initialize UserService with the test database
	userSvc = &service.UserService{
		User:    &models.User{},
		DB:      db,
		Account: &models.Account{},
	}
}

func tearDown() {
	// You can add code to clean up the test environment if necessary
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		log.Fatalf("failed to load test fixtures: %v", err)
	}
}

func TestCheckIfUserNameOrEmailExists(t *testing.T) {
	prepareTestDatabase()

	testCases := []struct {
		userName string
		email    string
		exists   bool
	}{
		{"john_doe", "", true},
		{"", "john.doe@example.com", true},
		{"nonexistent", "", false},
		{"", "nonexistent@example.com", false},
	}

	for _, tc := range testCases {
		exists, err := userSvc.CheckIfUserNameOrEmailExists(tc.userName, tc.email)
		if err != nil {
			t.Errorf("CheckIfUserNameOrEmailExists(%q, %q) returned error: %v", tc.userName, tc.email, err)
		}
		if exists != tc.exists {
			t.Errorf("CheckIfUserNameOrEmailExists(%q, %q) returned %v, expected %v", tc.userName, tc.email, exists, tc.exists)
		}
	}
}

func TestGetUserByUserName(t *testing.T) {
	prepareTestDatabase()

	user, err := userSvc.User.GetUserByUserName("john_doe", userSvc.DB)
	if err != nil {
		t.Errorf("GetUserByUserName(john_doe) returned error: %v", err)
	}
	if user == nil || user.Username != "john_doe" {
		t.Errorf("GetUserByUserName(john_doe) returned invalid user: %v", user)
	}

	user, err = userSvc.User.GetUserByUserName("nonexistent", userSvc.DB)
	if err == nil || err.Error() != "user not found" {
		t.Errorf("GetUserByUserName(nonexistent) returned error: %v, expected 'user not found'", err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	prepareTestDatabase()

	user, err := userSvc.User.GetUserByEmail("john.doe@example.com", userSvc.DB)
	if err != nil {
		t.Errorf("GetUserByEmail(john.doe@example.com) returned error: %v", err)
	}
	if user == nil || user.Email != "john.doe@example.com" {
		t.Errorf("GetUserByEmail(john.doe@example.com) returned invalid user: %v", user)
	}

	user, err = userSvc.User.GetUserByEmail("nonexistent@example.com", userSvc.DB)
	if err == nil || err.Error() != "user not found" {
		t.Errorf("GetUserByEmail(nonexistent@example.com) returned error: %v, expected 'user not found'", err)
	}
}
