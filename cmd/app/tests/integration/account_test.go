package integration

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccountEndToEndTest struct {
	suite.Suite
}

func TestAccountEndToEnd(t *testing.T) {
	suite.Run(t, new(AccountEndToEndTest))
}
