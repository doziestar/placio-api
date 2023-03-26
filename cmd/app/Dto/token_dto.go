package Dto

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/go-oauth2/oauth2/v4"
	"strings"
	"time"
)

type TokenDto struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// MarshalJSON is a custom JSON marshaler for TokenDto.
func (m *TokenDto) MarshalJSON() ([]byte, error) {
	type Alias TokenDto
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	})
}

// UnmarshalJSON is a custom JSON unmarshaler for TokenDto.
func (m *TokenDto) UnmarshalJSON(data []byte) error {
	type Alias TokenDto
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

// NewToken create to token model instance
func NewTokenInit() *Token {
	return &Token{}
}

// Token token model
type Token struct {
	TokenID             string        `json:"TokenID"`
	ClientID            string        `json:"ClientID"`
	UserID              string        `json:"UserID"`
	Scope               string        `json:"Scope"`
	Code                string        `json:"Code"`
	CodeCreateAt        time.Time     `json:"CodeCreateAt"`
	CodeExpiresIn       time.Duration `json:"CodeExpiresIn"`
	Access              string        `json:"Access"`
	AccessCreateAt      time.Time     `json:"AccessCreateAt"`
	AccessExpiresIn     time.Duration `json:"AccessExpiresIn"`
	Refresh             string        `json:"Refresh"`
	RefreshCreateAt     time.Time     `json:"RefreshCreateAt"`
	RefreshExpiresIn    time.Duration `json:"RefreshExpiresIn"`
	RedirectURI         string        `json:"RedirectURI"`
	CodeChallengeMethod string        `json:"CodeChallengeMethod"`
	CodeChallenge       string        `json:"CodeChallenge"`
}

// New create to token model instance
func (t *Token) New() oauth2.TokenInfo {
	return NewTokenInit()
}

// GetClientID the client id
func (t *Token) GetClientID() string {
	return t.ClientID
}

// SetClientID the client id
func (t *Token) SetClientID(clientID string) {
	t.ClientID = clientID
}

// GetUserID the user id
func (t *Token) GetUserID() string {
	return t.UserID
}

// SetUserID the user id
func (t *Token) SetUserID(userID string) {
	t.UserID = userID
}

// GetRedirectURI redirect URI
func (t *Token) GetRedirectURI() string {
	return t.RedirectURI
}

// SetRedirectURI redirect URI
func (t *Token) SetRedirectURI(redirectURI string) {
	t.RedirectURI = redirectURI
}

// GetScope get scope of authorization
func (t *Token) GetScope() string {
	return t.Scope
}

// SetScope get scope of authorization
func (t *Token) SetScope(scope string) {
	t.Scope = scope
}

// GetCode authorization code
func (t *Token) GetCode() string {
	return t.Code
}

// SetCode authorization code
func (t *Token) SetCode(code string) {
	t.Code = code
}

// GetCodeCreateAt create Time
func (t *Token) GetCodeCreateAt() time.Time {
	return t.CodeCreateAt
}

// SetCodeCreateAt create Time
func (t *Token) SetCodeCreateAt(createAt time.Time) {
	t.CodeCreateAt = createAt
}

// GetCodeExpiresIn the lifetime in seconds of the authorization code
func (t *Token) GetCodeExpiresIn() time.Duration {
	return t.CodeExpiresIn
}

// SetCodeExpiresIn the lifetime in seconds of the authorization code
func (t *Token) SetCodeExpiresIn(exp time.Duration) {
	t.CodeExpiresIn = exp
}

// GetCodeChallenge challenge code
func (t *Token) GetCodeChallenge() string {
	return t.CodeChallenge
}

// SetCodeChallenge challenge code
func (t *Token) SetCodeChallenge(code string) {
	t.CodeChallenge = code
}

// GetCodeChallengeMethod challenge method
func (t *Token) GetCodeChallengeMethod() oauth2.CodeChallengeMethod {
	return oauth2.CodeChallengeMethod(t.CodeChallengeMethod)
}

// SetCodeChallengeMethod challenge method
func (t *Token) SetCodeChallengeMethod(method oauth2.CodeChallengeMethod) {
	t.CodeChallengeMethod = string(method)
}

// GetAccess access Token
func (t *Token) GetAccess() string {
	return t.Access
}

// SetAccess access Token
func (t *Token) SetAccess(access string) {
	t.Access = access
}

// GetAccessCreateAt create Time
func (t *Token) GetAccessCreateAt() time.Time {
	return t.AccessCreateAt
}

// SetAccessCreateAt create Time
func (t *Token) SetAccessCreateAt(createAt time.Time) {
	t.AccessCreateAt = createAt
}

// GetAccessExpiresIn the lifetime in seconds of the access token
func (t *Token) GetAccessExpiresIn() time.Duration {
	return t.AccessExpiresIn
}

// SetAccessExpiresIn the lifetime in seconds of the access token
func (t *Token) SetAccessExpiresIn(exp time.Duration) {
	t.AccessExpiresIn = exp
}

// GetRefresh refresh Token
func (t *Token) GetRefresh() string {
	return t.Refresh
}

// SetRefresh refresh Token
func (t *Token) SetRefresh(refresh string) {
	t.Refresh = refresh
}

// GetRefreshCreateAt create Time
func (t *Token) GetRefreshCreateAt() time.Time {
	return t.RefreshCreateAt
}

// SetRefreshCreateAt create Time
func (t *Token) SetRefreshCreateAt(createAt time.Time) {
	t.RefreshCreateAt = createAt
}

// GetRefreshExpiresIn the lifetime in seconds of the refresh token
func (t *Token) GetRefreshExpiresIn() time.Duration {
	return t.RefreshExpiresIn
}

// SetRefreshExpiresIn the lifetime in seconds of the refresh token
func (t *Token) SetRefreshExpiresIn(exp time.Duration) {
	t.RefreshExpiresIn = exp
}

// ResponseType the type of authorization request
type ResponseType string

// define the type of authorization request
const (
	Code     ResponseType = "code"
	NewToken ResponseType = "new_token"
)

func (rt ResponseType) String() string {
	return string(rt)
}

// GrantType authorization model
type GrantType string

// define authorization model
const (
	AuthorizationCode   GrantType = "authorization_code"
	PasswordCredentials GrantType = "password"
	ClientCredentials   GrantType = "client_credentials"
	Refreshing          GrantType = "refresh_token"
	Implicit            GrantType = "__implicit"
)

func (gt GrantType) String() string {
	if gt == AuthorizationCode ||
		gt == PasswordCredentials ||
		gt == ClientCredentials ||
		gt == Refreshing {
		return string(gt)
	}
	return ""
}

// CodeChallengeMethod PCKE method
type CodeChallengeMethod string

const (
	// CodeChallengePlain PCKE Method
	CodeChallengePlain CodeChallengeMethod = "plain"
	// CodeChallengeS256 PCKE Method
	CodeChallengeS256 CodeChallengeMethod = "S256"
)

func (ccm CodeChallengeMethod) String() string {
	if ccm == CodeChallengePlain ||
		ccm == CodeChallengeS256 {
		return string(ccm)
	}
	return ""
}

// Validate code challenge
func (ccm CodeChallengeMethod) Validate(cc, ver string) bool {
	switch ccm {
	case CodeChallengePlain:
		return cc == ver
	case CodeChallengeS256:
		s256 := sha256.Sum256([]byte(ver))
		// trim padding
		a := strings.TrimRight(base64.URLEncoding.EncodeToString(s256[:]), "=")
		b := strings.TrimRight(cc, "=")
		return a == b
	default:
		return false
	}
}
