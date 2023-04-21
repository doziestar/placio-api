package utility

import (
	"placio-app/utility"
	"testing"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		email    string
		name     string
		password string
		username string
		errMsg   string
	}{
		{"john.doe@example.com", "John Doe", "P@ssw0rd", "john_doe", ""},
		{"", "John Doe", "P@ssw0rd", "john_doe", "invalid email address"},
		{"john.doe@example", "John Doe", "P@ssw0rd", "john_doe", "invalid email address"},
		{"john.doe@example.com", "", "P@ssw0rd", "john_doe", "invalid name: name must be at least 3 characters, and contain only letters, spaces, and hyphens"},
		{"john.doe@example.com", "J0hn", "P@ssw0rd", "john_doe", "invalid name: name must be at least 3 characters, and contain only letters, spaces, and hyphens"},
		{"john.doe@example.com", "John Doe", "", "john_doe", "invalid password: password must be at least 8 characters and contain at least one uppercase, one lowercase, and one number"},
		{"john.doe@example.com", "John Doe", "password", "john_doe", "invalid password: password must be at least 8 characters and contain at least one uppercase, one lowercase, and one number"},
		{"john.doe@example.com", "John Doe", "P@ssw0rd", "", "invalid username: username must be at least 4 characters, and contain only letters, numbers, and underscores"},
		{"john.doe@example.com", "John Doe", "P@ssw0rd", "john", ""},
	}

	for _, tc := range testCases {
		err := utility.Validate(tc.email, tc.name, tc.password, tc.username)
		if err != nil && err.Error() != tc.errMsg {
			t.Errorf("Validate(%q, %q, %q, %q) returned error %q, expected %q", tc.email, tc.name, tc.password, tc.username, err.Error(), tc.errMsg)
		}
	}
}
