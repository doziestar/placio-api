package websites

import (
	"context"
	"placio-app/domains/business"
	"placio-app/domains/media"
	"placio-app/domains/users"
	"placio-app/ent"
)

type IWebsite interface {
	GetBusinessWebsite(ctx context.Context, businessID string) (*ent.Website, error)
	CreateBusinessWebsite(ctx context.Context, businessID string) (*ent.Website, error)
	UpdateBusinessWebsite(ctx context.Context, businessID string) (*ent.Website, error)
}

type WebsiteService struct {
	client          *ent.Client
	businessService business.BusinessAccountService
	userService     users.UserService
	mediaService    media.MediaService
}

func NewWebsiteService(client *ent.Client, businessService business.BusinessAccountService, userService users.UserService, mediaService media.MediaService) *WebsiteService {
	return &WebsiteService{
		client:          client,
		businessService: businessService,
		userService:     userService,
		mediaService:    mediaService,
	}
}

func (w *WebsiteService) GetBusinessWebsite(ctx context.Context, businessID string) (*ent.Website, error) {
	return nil, nil
}

func (w *WebsiteService) CreateBusinessWebsite(ctx context.Context, businessID string) (*ent.Website, error) {
	return nil, nil
}

func (w *WebsiteService) UpdateBusinessWebsite(ctx context.Context, businessID string) (*ent.Website, error) {
	return nil, nil
}
