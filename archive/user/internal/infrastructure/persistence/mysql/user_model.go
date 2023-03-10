/*
Package mysql holds view model repositories
*/
package mysql

import (
	"placio-api/cmd/user/internal/application/access"
	"placio-api/pkg/mysql"
)

// User model
type User struct {
	ID         string           `json:"id"`
	Email      string           `json:"email"`
	FacebookID mysql.NullString `json:"facebook_id"`
	GoogleID   mysql.NullString `json:"google_id"`
	Role       uint8            `json:"role"`
}

// GetID the id
func (u User) GetID() string {
	return u.ID
}

// GetEmail the email
func (u User) GetEmail() string {
	return u.Email
}

// GetFacebookID facebook id
func (u User) GetFacebookID() string {
	return u.FacebookID.String
}

// GetGoogleID google id
func (u User) GetGoogleID() string {
	return u.GoogleID.String
}

// GetRole returns user role
func (u User) GetRole() access.Role {
	return access.Role(u.Role)
}