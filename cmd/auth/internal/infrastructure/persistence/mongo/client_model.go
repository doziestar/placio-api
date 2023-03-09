package mongo

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
)

// Client model
type Client struct {
	ID          string   `json:"id" bson:"client_id"`
	UserID      string   `json:"user_id" bson:"user_id"`
	Secret      string   `json:"secret" bson:"secret"`
	Domain      string   `json:"domain" bson:"domain"`
	RedirectURL string   `json:"redirect_url" bson:"redirect_url"`
	Scopes      []string `json:"scopes" bson:"scopes"`
}

func (c *Client) GetID() string {
	return c.ID
}

func (c *Client) GetUserID() string {
	return c.UserID
}

func (c *Client) GetSecret() string {
	return c.Secret
}

func (c *Client) GetDomain() string {
	return c.Domain
}

func (c *Client) GetRedirectURL() string {
	return c.RedirectURL
}

func (c *Client) GetScopes() []string {
	return c.Scopes
}

func (c *Client) ClientInfo() (oauth2.ClientInfo, error) {
	return &models.Client{
		ID:     c.ID,
		Secret: c.Secret,
		Domain: c.Domain,
		UserID: c.UserID,
	}, nil
}
