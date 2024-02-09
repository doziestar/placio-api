package websites

import (
	"context"
	"log"
	businessService "placio-app/domains/business"
	"placio-app/domains/media"
	"placio-app/domains/users"
	"placio-app/ent"
	"placio-app/ent/business"
	"placio-app/ent/website"
	"time"

	"github.com/google/uuid"
)

type IWebsite interface {
	GetBusinessWebsite(ctx context.Context, businessID, domainName string) (*ent.Website, error)
	CreateBusinessWebsite(ctx context.Context, businessID string, websiteData *ent.Website) (*ent.Website, error)
	UpdateBusinessWebsite(ctx context.Context, businessID string, websiteData *ent.Website) (*ent.Website, error)
	VerifyDomainName(ctx context.Context, domainName string) (bool, error)
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

func (w *WebsiteService) VerifyDomainName(ctx context.Context, domainName string) (bool, error) {
	_, err := w.client.Website.Query().Where(website.DomainName(domainName)).First(ctx)
	if ent.IsNotFound(err) {
		return false, nil
	}
	return true, nil
}

func (w *WebsiteService) GetBusinessWebsite(ctx context.Context, businessID, domainName string) (*ent.Website, error) {
	var websiteData *ent.Website
	var err error

	if domainName != "" {
		websiteData, err = w.client.Website.
			Query().
			Where(website.DomainName(domainName)).
			WithCustomBlocks().
			WithAssets().
			WithBusiness(func(bq *ent.BusinessQuery) {
				bq.WithPlaces(func(pq *ent.PlaceQuery) {
					pq.WithMenus(func(mq *ent.MenuQuery) {
						mq.WithMenuItems(func(miq *ent.MenuItemQuery) {
							miq.WithMedia()
						})

						mq.WithMedia()
					})
					pq.WithRoomCategories(func(rcq *ent.RoomCategoryQuery) {
						rcq.WithRooms(func(rq *ent.RoomQuery) {
							rq.WithMedia()
						})
						rcq.WithMedia()
					})
					pq.WithEvents(func(eq *ent.EventQuery) {
					})
					pq.WithMedias()
				})

				bq.WithEvents(func(eq *ent.EventQuery) {
					// eq.WithEventItems()
				})
			}).
			First(ctx)
		if err != nil {
			return nil, err
		}
		return websiteData, nil
	}
	websiteData, err = w.client.Website.Query().
		Where(website.HasBusinessWith(business.ID(businessID))).
		WithCustomBlocks().
		WithAssets().
		WithBusiness().
		First(ctx)
	if err != nil {
		return nil, err
	}
	return websiteData, nil
}

func (w *WebsiteService) CreateBusinessWebsite(ctx context.Context, businessID string, websiteData *ent.Website) (*ent.Website, error) {
	log.Println("websiteData", websiteData)
	websiteData, err := w.client.Website.Create().
		SetID(uuid.New().String()).
		SetBusinessID(businessID).
		SetDomainName(websiteData.DomainName).
		SetBannerSectionBackgroundColor(websiteData.BannerSectionBackgroundColor).
		SetThreeItemsSectionHeadingText(websiteData.ThreeItemsSectionHeadingText).
		SetThreeItemsSectionItemOneText(websiteData.ThreeItemsSectionItemOneText).
		SetThreeItemsSectionItemTwoText(websiteData.ThreeItemsSectionItemTwoText).
		SetThreeItemsSectionItemThreeText(websiteData.ThreeItemsSectionItemThreeText).
		SetThreeItemsSectionDetailsText(websiteData.ThreeItemsSectionDetailsText).
		SetBannerTwoSectionBackgroundColor(websiteData.BannerTwoSectionBackgroundColor).
		SetBannerTwoLeftSectionHeadingText(websiteData.BannerTwoLeftSectionHeadingText).
		SetBannerTwoLeftSectionDetailsText(websiteData.BannerTwoLeftSectionDetailsText).
		SetBannerTwoRightSideImage(websiteData.BannerTwoRightSideImage).
		SetInventorySectionHeadingText(websiteData.InventorySectionHeadingText).
		SetHeadingText(websiteData.HeadingText).
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
		SetLastUpdated(time.Now()).
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
		SetDomainName(websiteData.DomainName).
		SetBannerSectionBackgroundColor(websiteData.BannerSectionBackgroundColor).
		SetThreeItemsSectionHeadingText(websiteData.ThreeItemsSectionHeadingText).
		SetThreeItemsSectionItemOneText(websiteData.ThreeItemsSectionItemOneText).
		SetThreeItemsSectionItemTwoText(websiteData.ThreeItemsSectionItemTwoText).
		SetThreeItemsSectionItemThreeText(websiteData.ThreeItemsSectionItemThreeText).
		SetThreeItemsSectionDetailsText(websiteData.ThreeItemsSectionDetailsText).
		SetBannerTwoSectionBackgroundColor(websiteData.BannerTwoSectionBackgroundColor).
		SetBannerTwoLeftSectionHeadingText(websiteData.BannerTwoLeftSectionHeadingText).
		SetBannerTwoLeftSectionDetailsText(websiteData.BannerTwoLeftSectionDetailsText).
		SetBannerTwoRightSideImage(websiteData.BannerTwoRightSideImage).
		SetInventorySectionHeadingText(websiteData.InventorySectionHeadingText).
		SetHeadingText(websiteData.HeadingText).
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
		SetLastUpdated(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return website, nil
}
