package models

import "encoding/json"

func UnmarshalAuth0UserData(data []byte) (Auth0UserData, error) {
	var r Auth0UserData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Auth0UserData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Auth0UserData struct {
	Blocked           bool        `json:"blocked"`
	EmailVerified     bool        `json:"email_verified"`
	Email             string      `json:"email"`
	PhoneNumber       string      `json:"phone_number"`
	PhoneVerified     bool        `json:"phone_verified"`
	UserMetadata      Metadata    `json:"user_metadata"`
	AppMetadata       AppMetadata `json:"app_metadata"`
	GivenName         string      `json:"given_name"`
	FamilyName        string      `json:"family_name"`
	Name              string      `json:"name"`
	Nickname          string      `json:"nickname"`
	Picture           string      `json:"picture"`
	VerifyEmail       bool        `json:"verify_email"`
	VerifyPhoneNumber bool        `json:"verify_phone_number"`
	Password          string      `json:"password"`
	Connection        string      `json:"connection"`
	ClientID          string      `json:"client_id"`
	Username          string      `json:"username"`
}

func UnmarshalMetadata(data []byte) (Metadata, error) {
	var r Metadata
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Metadata) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Metadata struct {
	CreatedAt string      `json:"CreatedAt"`
	UpdatedAt string      `json:"UpdatedAt"`
	DeletedAt interface{} `json:"DeletedAt"`
	Language  string      `json:"Language"`
	Theme     string      `json:"Theme"`
	Privacy   string      `json:"Privacy"`
	Content   Content     `json:"Content"`
}

type Content struct {
	MediaVisibility       string `json:"MediaVisibility"`
	ExplicitContentFilter string `json:"ExplicitContentFilter"`
	DefaultPostPrivacy    string `json:"DefaultPostPrivacy"`
	AutoplayVideos        bool   `json:"AutoplayVideos"`
	DisplaySensitiveMedia bool   `json:"DisplaySensitiveMedia"`
}

type AppMetadata struct {
	EmailNotifications         bool `json:"EmailNotifications"`
	PushNotifications          bool `json:"PushNotifications"`
	DirectMessageNotifications bool `json:"DirectMessageNotifications"`
	LikeNotifications          bool `json:"LikeNotifications"`
	CommentNotifications       bool `json:"CommentNotifications"`
	MentionNotifications       bool `json:"MentionNotifications"`
	FollowNotifications        bool `json:"FollowNotifications"`
}

func UnmarshalAppMetadata(data []byte) (AppMetadata, error) {
	var r AppMetadata
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AppMetadata) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
