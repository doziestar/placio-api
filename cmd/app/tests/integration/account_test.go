package integration

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type AccountIntegrationTest struct {
	suite.Suite
}

func TestAccount(t *testing.T) {
	suite.Run(t, new(AccountIntegrationTest))
}

func (suite *AccountIntegrationTest) TestAccountCreation() {
	resp, err := client.Post("http://localhost:7070/api/v1/accounts/create-account", "application/json", bytes.NewBuffer([]byte(fmt.Sprintf(`{"name": "%s", "email": "%s", "password": "%s", "confirm_password": "%s", "account_type": "%s"}`, userData.Name, userData.Email, userData.Password, userData.ConfirmPassword, userData.AccountType))))
	suite.NoError(err)

	suite.Equal(http.StatusCreated, resp.StatusCode)

	//expected := {
	//	"status": "ok",
	//	"message": "Account created successfully",
	//	"data": {
	//
	//	}
	//}

	//b, err := io.ReadAll(resp.Body)
	//suite.NoError(err)
	//
	//actual := string(b)
	//
	//suite.JSONEq(expected, actual)
}
