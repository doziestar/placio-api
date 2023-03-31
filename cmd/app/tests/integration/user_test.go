package integration

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserIntegrationTest struct {
	suite.Suite
}

func TestUserIntegration(t *testing.T) {
	suite.Run(t, new(UserIntegrationTest))
}

func (suite *UserIntegrationTest) TestUserIntegration() {
	suite.Equal(1, 1)
}
