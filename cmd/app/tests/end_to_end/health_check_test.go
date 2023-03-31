package end_to_end

import (
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"testing"
)

type HealthCheckTest struct {
	suite.Suite
}

func TestHealthCheck(t *testing.T) {
	suite.Run(t, new(HealthCheckTest))
}

func (suite *HealthCheckTest) TestHealthCheck() {
	c := http.Client{}
	resp, err := c.Get("http://localhost:7070/api/v1/health")
	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)

	expected := `{"status":"ok"}`
	b, err := io.ReadAll(resp.Body)
	suite.NoError(err)

	actual := string(b)

	suite.JSONEq(expected, actual)
}
