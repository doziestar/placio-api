package websites

import (
	"context"
	"github.com/google/uuid"
	"log"
	businessService "placio-app/domains/business"
	"placio-app/domains/media"
	"placio-app/domains/users"
	"placio-app/ent"
	"placio-app/ent/business"
	"placio-app/ent/website"
)

type IWebsite interface {
	GetBusinessWebsite(ctx context.Context, businessID string) (*ent.Website, error)
	CreateBusinessWebsite(ctx context.Context, businessID string, websiteData *ent.Website) (*ent.Website, error)
	UpdateBusinessWebsite(ctx context.Context, businessID string, websiteData *ent.Website) (*ent.Website, error)
}

type WebsiteService struct {
	client          *ent.Client
	businessService businessService.BusinessAccountService
	userService     users.UserService
	mediaService    media.MediaService
}

func NewWebsiteService(client *ent.Client, businessService businessService.BusinessAccountService, userService users.UserService, mediaService media.MediaService) *WebsiteService {
	return &WebsiteService{
		client:          client,
		businessService: businessService,
		userService:     userService,
		mediaService:    mediaService,
	}
}

func (w *WebsiteService) GetBusinessWebsite(ctx context.Context, businessID string) (*ent.Website, error) {
	website, err := w.client.Website.Query().Where(website.HasBusinessWith(business.ID(businessID))).First(ctx)
	if err != nil {
		return nil, err
	}
	return website, nil
}

func (w *WebsiteService) CreateBusinessWebsite(ctx context.Context, businessID string, websiteData *ent.Website) (*ent.Website, error) {
	log.Println("websiteData", websiteData)
	websiteData, err := w.client.Website.Create().
		SetID(uuid.New().String()).
		SetBusinessID(businessID).
		SetDomainName(websiteData.DomainName).
		SetAchievementsSection(websiteData.AchievementsSection).
		SetAddress(websiteData.Address).
		SetBannerSectionBackgroundColor(websiteData.BannerSectionBackgroundColor).
		SetBannerSectionBackgroundImage(websiteData.BannerSectionBackgroundImage).
		SetTitle(websiteData.Title).
		SetDescription(websiteData.Description).
		SetEmail(websiteData.Email).
		SetState(websiteData.State).
		SetCountry(websiteData.Country).
		SetPinterest(websiteData.Pinterest).
		SetBannerSectionText(websiteData.BannerSectionText).
		SetMapCoordinates(websiteData.MapCoordinates).
		SetFacebook(websiteData.Facebook).
		SetInstagram(websiteData.Instagram).
		SetTwitter(websiteData.Twitter).
		SetYoutube(websiteData.Youtube).
		SetLinkedin(websiteData.Linkedin).
		SetBannerTwoLeftSectionButtonLink(websiteData.BannerTwoLeftSectionButtonLink).
		SetBannerTwoLeftSectionButtonText(websiteData.BannerTwoLeftSectionButtonText).
		SetBannerTwoLeftSectionHeadingText(websiteData.BannerTwoLeftSectionHeadingText).
		SetBannerTwoSectionBackgroundImage(websiteData.BannerTwoSectionBackgroundImage).
		SetBusinessLogo(websiteData.BusinessLogo).
		SetBusinessName(websiteData.BusinessName).
		SetKeywords(websiteData.Keywords).
		SetLogo(websiteData.Logo).
		SetLanguage(websiteData.Language).
		SetLatitude(websiteData.Latitude).
		SetLongitude(websiteData.Longitude).
		Save(ctx)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}

	return websiteData, nil
}

func (w *WebsiteService) UpdateBusinessWebsite(ctx context.Context, businessID string, websiteData *ent.Website) (*ent.Website, error) {
	// update the website
	website, err := w.client.Website.UpdateOneID(websiteData.ID).
		SetBusinessID(businessID).
		SetAchievementsSection(websiteData.AchievementsSection).
		SetAddress(websiteData.Address).
		SetBannerSectionBackgroundColor(websiteData.BannerSectionBackgroundColor).
		SetBannerSectionBackgroundImage(websiteData.BannerSectionBackgroundImage).
		SetTitle(websiteData.Title).
		SetDescription(websiteData.Description).
		SetEmail(websiteData.Email).
		SetState(websiteData.State).
		SetCountry(websiteData.Country).
		SetPinterest(websiteData.Pinterest).
		SetBannerSectionText(websiteData.BannerSectionText).
		SetMapCoordinates(websiteData.MapCoordinates).
		SetFacebook(websiteData.Facebook).
		SetInstagram(websiteData.Instagram).
		SetTwitter(websiteData.Twitter).
		SetYoutube(websiteData.Youtube).
		SetLinkedin(websiteData.Linkedin).
		SetBannerTwoLeftSectionButtonLink(websiteData.BannerTwoLeftSectionButtonLink).
		SetBannerTwoLeftSectionButtonText(websiteData.BannerTwoLeftSectionButtonText).
		SetBannerTwoLeftSectionHeadingText(websiteData.BannerTwoLeftSectionHeadingText).
		SetBannerTwoSectionBackgroundImage(websiteData.BannerTwoSectionBackgroundImage).
		SetBusinessLogo(websiteData.BusinessLogo).
		SetBusinessName(websiteData.BusinessName).
		SetKeywords(websiteData.Keywords).
		SetLogo(websiteData.Logo).
		SetLanguage(websiteData.Language).
		SetLatitude(websiteData.Latitude).
		SetLongitude(websiteData.Longitude).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return website, nil
}
