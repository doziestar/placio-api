package service

import (
	"context"
	"placio-app/ent"
)

type HelpService interface {
	// ...
	CreateHelp(ctx context.Context, userID, category, subject, body, media string) (*ent.Help, error)
	ResolveHelp(ctx context.Context, helpID string) (*ent.Help, error)
	DeleteHelp(ctx context.Context, helpID string) error
}

type HelpServiceImpl struct {
	client *ent.Client
}

func NewHelpService(client *ent.Client) HelpService {
	return &HelpServiceImpl{client: client}
}

func (s *HelpServiceImpl) CreateHelp(ctx context.Context, userID, category, subject, body, media string) (*ent.Help, error) {
	return s.client.Help.
		Create().
		SetUserID(userID).
		SetCategory(category).
		SetSubject(subject).
		SetBody(body).
		SetMedia(media).
		Save(ctx)
}

func (s *HelpServiceImpl) ResolveHelp(ctx context.Context, helpID string) (*ent.Help, error) {
	return s.client.Help.
		UpdateOneID(helpID).
		SetStatus("resolved").
		Save(ctx)
}

func (s *HelpServiceImpl) DeleteHelp(ctx context.Context, helpID string) error {
	return s.client.Help.
		DeleteOneID(helpID).
		Exec(ctx)
}
