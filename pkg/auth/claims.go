package auth

import (
	"fmt"

	apperrors "placio-api/pkg/errors"
	"placio-api/pkg/identity"
)

type Claims struct {
	jwt.StandardClaims
	Identity *identity.Identity `json:"identity,omitempty"`
}

func (c *Claims) Valid() error {
	if c.Identity == nil {
		return apperrors.Wrap(fmt.Errorf("Identity must be set"))
	}

	return c.StandardClaims.Valid()
}
