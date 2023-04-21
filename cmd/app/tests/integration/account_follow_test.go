package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"placio-app/models"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountServiceMock struct {
	mock.Mock
}

func (m *AccountServiceMock) Follow(followerID, followingID string) error {
	args := m.Called(followerID, followingID)
	return args.Error(0)
}

func (m *AccountServiceMock) Unfollow(followerID, followingID string) error {
	args := m.Called(followerID, followingID)
	return args.Error(0)
}

func (m *AccountServiceMock) ListFollowers(accountID string) ([]models.Account, error) {
	args := m.Called(accountID)
	return args.Get(0).([]models.Account), args.Error(1)
}

func (m *AccountServiceMock) ListFollowing(accountID string) ([]models.Account, error) {
	args := m.Called(accountID)
	return args.Get(0).([]models.Account), args.Error(1)
}

type AccountControllerTestSuite struct {
	app     *fiber.App
	service *AccountServiceMock
}

func (suite *AccountControllerTestSuite) Setup(t *testing.T) {
	suite.app = fiber.New()
	suite.service = new(AccountServiceMock)
	//TODO: Fix this
	//con := controller.NewAccountController(suite.service, nil)
	//con.RegisterRoutes(suite.app.Group("/api/v1/accounts"))
}

func TestFollowAccount(t *testing.T) {
	suite := new(AccountControllerTestSuite)
	suite.Setup(t)

	suite.service.On("Follow", "1", "2").Return(nil)

	body := bytes.NewBufferString("following_id=2")
	req := httptest.NewRequest("POST", "/api/v1/accounts/1/follow", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := suite.app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	suite.service.AssertExpectations(t)
}

func TestUnfollowAccount(t *testing.T) {
	suite := new(AccountControllerTestSuite)
	suite.Setup(t)

	suite.service.On("Unfollow", "1", "2").Return(nil)

	body := bytes.NewBufferString("following_id=2")
	req := httptest.NewRequest("POST", "/api/v1/accounts/1/unfollow", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := suite.app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	suite.service.AssertExpectations(t)
}

func TestGetFollowers(t *testing.T) {
	suite := new(AccountControllerTestSuite)
	suite.Setup(t)

	followers := []models.Account{
		{ID: "2", Name: "Account 2"},
		{ID: "3", Name: "Account 3"},
	}

	suite.service.On("ListFollowers", "1").Return(followers, nil)

	req := httptest.NewRequest("GET", "/api/v1/accounts/1/followers", nil)
	resp, err := suite.app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result []models.Account
	json.NewDecoder(resp.Body).Decode(&result)

	assert.Equal(t, followers, result)
	suite.service.AssertExpectations(t)
}

func TestGetFollowing(t *testing.T) {
	suite := new(AccountControllerTestSuite)
	suite.Setup(t)

	following := []models.Account{
		{ID: "2", Name: "Account 2"},
		{ID: "3", Name: "Account 3"},
	}

	suite.service.On("ListFollowing", "1").Return(following, nil)

	req := httptest.NewRequest("GET", "/api/v1/accounts/1/following", nil)
	resp, err := suite.app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result []models.Account
	json.NewDecoder(resp.Body).Decode(&result)

	assert.Equal(t, following, result)
	suite.service.AssertExpectations(t)
}
