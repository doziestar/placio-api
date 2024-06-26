// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/customblock"
	"placio-app/ent/media"
	"placio-app/ent/website"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WebsiteCreate is the builder for creating a Website entity.
type WebsiteCreate struct {
	config
	mutation *WebsiteMutation
	hooks    []Hook
}

// SetDomainName sets the "domainName" field.
func (wc *WebsiteCreate) SetDomainName(s string) *WebsiteCreate {
	wc.mutation.SetDomainName(s)
	return wc
}

// SetHeadingText sets the "heading_text" field.
func (wc *WebsiteCreate) SetHeadingText(s string) *WebsiteCreate {
	wc.mutation.SetHeadingText(s)
	return wc
}

// SetNillableHeadingText sets the "heading_text" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableHeadingText(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetHeadingText(*s)
	}
	return wc
}

// SetBusinessLogo sets the "business_logo" field.
func (wc *WebsiteCreate) SetBusinessLogo(s string) *WebsiteCreate {
	wc.mutation.SetBusinessLogo(s)
	return wc
}

// SetNillableBusinessLogo sets the "business_logo" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableBusinessLogo(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetBusinessLogo(*s)
	}
	return wc
}

// SetBusinessName sets the "business_name" field.
func (wc *WebsiteCreate) SetBusinessName(s string) *WebsiteCreate {
	wc.mutation.SetBusinessName(s)
	return wc
}

// SetNillableBusinessName sets the "business_name" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableBusinessName(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetBusinessName(*s)
	}
	return wc
}

// SetBannerSectionBackgroundImage sets the "banner_section_background_image" field.
func (wc *WebsiteCreate) SetBannerSectionBackgroundImage(s string) *WebsiteCreate {
	wc.mutation.SetBannerSectionBackgroundImage(s)
	return wc
}

// SetNillableBannerSectionBackgroundImage sets the "banner_section_background_image" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableBannerSectionBackgroundImage(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetBannerSectionBackgroundImage(*s)
	}
	return wc
}

// SetBannerSectionBackgroundColor sets the "banner_section_background_color" field.
func (wc *WebsiteCreate) SetBannerSectionBackgroundColor(s string) *WebsiteCreate {
	wc.mutation.SetBannerSectionBackgroundColor(s)
	return wc
}

// SetNillableBannerSectionBackgroundColor sets the "banner_section_background_color" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableBannerSectionBackgroundColor(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetBannerSectionBackgroundColor(*s)
	}
	return wc
}

// SetBannerSectionText sets the "banner_section_text" field.
func (wc *WebsiteCreate) SetBannerSectionText(s string) *WebsiteCreate {
	wc.mutation.SetBannerSectionText(s)
	return wc
}

// SetNillableBannerSectionText sets the "banner_section_text" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableBannerSectionText(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetBannerSectionText(*s)
	}
	return wc
}

// SetThreeItemsSectionHeadingText sets the "three_items_section_heading_text" field.
func (wc *WebsiteCreate) SetThreeItemsSectionHeadingText(s string) *WebsiteCreate {
	wc.mutation.SetThreeItemsSectionHeadingText(s)
	return wc
}

// SetNillableThreeItemsSectionHeadingText sets the "three_items_section_heading_text" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableThreeItemsSectionHeadingText(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetThreeItemsSectionHeadingText(*s)
	}
	return wc
}

// SetThreeItemsSectionDetailsText sets the "three_items_section_details_text" field.
func (wc *WebsiteCreate) SetThreeItemsSectionDetailsText(s string) *WebsiteCreate {
	wc.mutation.SetThreeItemsSectionDetailsText(s)
	return wc
}

// SetNillableThreeItemsSectionDetailsText sets the "three_items_section_details_text" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableThreeItemsSectionDetailsText(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetThreeItemsSectionDetailsText(*s)
	}
	return wc
}

// SetThreeItemsSectionItemOneText sets the "three_items_section_item_one_text" field.
func (wc *WebsiteCreate) SetThreeItemsSectionItemOneText(s string) *WebsiteCreate {
	wc.mutation.SetThreeItemsSectionItemOneText(s)
	return wc
}

// SetNillableThreeItemsSectionItemOneText sets the "three_items_section_item_one_text" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableThreeItemsSectionItemOneText(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetThreeItemsSectionItemOneText(*s)
	}
	return wc
}

// SetThreeItemsSectionItemTwoText sets the "three_items_section_item_two_text" field.
func (wc *WebsiteCreate) SetThreeItemsSectionItemTwoText(s string) *WebsiteCreate {
	wc.mutation.SetThreeItemsSectionItemTwoText(s)
	return wc
}

// SetNillableThreeItemsSectionItemTwoText sets the "three_items_section_item_two_text" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableThreeItemsSectionItemTwoText(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetThreeItemsSectionItemTwoText(*s)
	}
	return wc
}

// SetThreeItemsSectionItemThreeText sets the "three_items_section_item_three_text" field.
func (wc *WebsiteCreate) SetThreeItemsSectionItemThreeText(s string) *WebsiteCreate {
	wc.mutation.SetThreeItemsSectionItemThreeText(s)
	return wc
}

// SetNillableThreeItemsSectionItemThreeText sets the "three_items_section_item_three_text" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableThreeItemsSectionItemThreeText(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetThreeItemsSectionItemThreeText(*s)
	}
	return wc
}

// SetBannerTwoSectionBackgroundImage sets the "banner_two_section_background_image" field.
func (wc *WebsiteCreate) SetBannerTwoSectionBackgroundImage(s string) *WebsiteCreate {
	wc.mutation.SetBannerTwoSectionBackgroundImage(s)
	return wc
}

// SetNillableBannerTwoSectionBackgroundImage sets the "banner_two_section_background_image" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableBannerTwoSectionBackgroundImage(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetBannerTwoSectionBackgroundImage(*s)
	}
	return wc
}

// SetBannerTwoSectionBackgroundColor sets the "banner_two_section_background_color" field.
func (wc *WebsiteCreate) SetBannerTwoSectionBackgroundColor(s string) *WebsiteCreate {
	wc.mutation.SetBannerTwoSectionBackgroundColor(s)
	return wc
}

// SetNillableBannerTwoSectionBackgroundColor sets the "banner_two_section_background_color" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableBannerTwoSectionBackgroundColor(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetBannerTwoSectionBackgroundColor(*s)
	}
	return wc
}

// SetBannerTwoLeftSectionHeadingText sets the "banner_two_left_section_heading_text" field.
func (wc *WebsiteCreate) SetBannerTwoLeftSectionHeadingText(s string) *WebsiteCreate {
	wc.mutation.SetBannerTwoLeftSectionHeadingText(s)
	return wc
}

// SetNillableBannerTwoLeftSectionHeadingText sets the "banner_two_left_section_heading_text" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableBannerTwoLeftSectionHeadingText(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetBannerTwoLeftSectionHeadingText(*s)
	}
	return wc
}

// SetBannerTwoLeftSectionDetailsText sets the "banner_two_left_section_details_text" field.
func (wc *WebsiteCreate) SetBannerTwoLeftSectionDetailsText(s string) *WebsiteCreate {
	wc.mutation.SetBannerTwoLeftSectionDetailsText(s)
	return wc
}

// SetNillableBannerTwoLeftSectionDetailsText sets the "banner_two_left_section_details_text" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableBannerTwoLeftSectionDetailsText(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetBannerTwoLeftSectionDetailsText(*s)
	}
	return wc
}

// SetBannerTwoLeftSectionButtonText sets the "banner_two_left_section_button_text" field.
func (wc *WebsiteCreate) SetBannerTwoLeftSectionButtonText(s string) *WebsiteCreate {
	wc.mutation.SetBannerTwoLeftSectionButtonText(s)
	return wc
}

// SetNillableBannerTwoLeftSectionButtonText sets the "banner_two_left_section_button_text" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableBannerTwoLeftSectionButtonText(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetBannerTwoLeftSectionButtonText(*s)
	}
	return wc
}

// SetBannerTwoLeftSectionButtonLink sets the "banner_two_left_section_button_link" field.
func (wc *WebsiteCreate) SetBannerTwoLeftSectionButtonLink(s string) *WebsiteCreate {
	wc.mutation.SetBannerTwoLeftSectionButtonLink(s)
	return wc
}

// SetNillableBannerTwoLeftSectionButtonLink sets the "banner_two_left_section_button_link" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableBannerTwoLeftSectionButtonLink(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetBannerTwoLeftSectionButtonLink(*s)
	}
	return wc
}

// SetBannerTwoRightSideImage sets the "banner_two_right_side_image" field.
func (wc *WebsiteCreate) SetBannerTwoRightSideImage(s string) *WebsiteCreate {
	wc.mutation.SetBannerTwoRightSideImage(s)
	return wc
}

// SetNillableBannerTwoRightSideImage sets the "banner_two_right_side_image" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableBannerTwoRightSideImage(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetBannerTwoRightSideImage(*s)
	}
	return wc
}

// SetAchievementsSection sets the "achievements_section" field.
func (wc *WebsiteCreate) SetAchievementsSection(m map[string]interface{}) *WebsiteCreate {
	wc.mutation.SetAchievementsSection(m)
	return wc
}

// SetInventorySectionHeadingText sets the "Inventory_section_heading_text" field.
func (wc *WebsiteCreate) SetInventorySectionHeadingText(s string) *WebsiteCreate {
	wc.mutation.SetInventorySectionHeadingText(s)
	return wc
}

// SetNillableInventorySectionHeadingText sets the "Inventory_section_heading_text" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableInventorySectionHeadingText(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetInventorySectionHeadingText(*s)
	}
	return wc
}

// SetCreationDate sets the "creationDate" field.
func (wc *WebsiteCreate) SetCreationDate(t time.Time) *WebsiteCreate {
	wc.mutation.SetCreationDate(t)
	return wc
}

// SetNillableCreationDate sets the "creationDate" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableCreationDate(t *time.Time) *WebsiteCreate {
	if t != nil {
		wc.SetCreationDate(*t)
	}
	return wc
}

// SetLastUpdated sets the "lastUpdated" field.
func (wc *WebsiteCreate) SetLastUpdated(t time.Time) *WebsiteCreate {
	wc.mutation.SetLastUpdated(t)
	return wc
}

// SetTitle sets the "title" field.
func (wc *WebsiteCreate) SetTitle(s string) *WebsiteCreate {
	wc.mutation.SetTitle(s)
	return wc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableTitle(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetTitle(*s)
	}
	return wc
}

// SetDescription sets the "description" field.
func (wc *WebsiteCreate) SetDescription(s string) *WebsiteCreate {
	wc.mutation.SetDescription(s)
	return wc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableDescription(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetDescription(*s)
	}
	return wc
}

// SetKeywords sets the "keywords" field.
func (wc *WebsiteCreate) SetKeywords(s string) *WebsiteCreate {
	wc.mutation.SetKeywords(s)
	return wc
}

// SetNillableKeywords sets the "keywords" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableKeywords(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetKeywords(*s)
	}
	return wc
}

// SetLanguage sets the "language" field.
func (wc *WebsiteCreate) SetLanguage(s string) *WebsiteCreate {
	wc.mutation.SetLanguage(s)
	return wc
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableLanguage(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetLanguage(*s)
	}
	return wc
}

// SetLogo sets the "logo" field.
func (wc *WebsiteCreate) SetLogo(s string) *WebsiteCreate {
	wc.mutation.SetLogo(s)
	return wc
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableLogo(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetLogo(*s)
	}
	return wc
}

// SetFavicon sets the "favicon" field.
func (wc *WebsiteCreate) SetFavicon(s string) *WebsiteCreate {
	wc.mutation.SetFavicon(s)
	return wc
}

// SetNillableFavicon sets the "favicon" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableFavicon(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetFavicon(*s)
	}
	return wc
}

// SetFacebook sets the "facebook" field.
func (wc *WebsiteCreate) SetFacebook(s string) *WebsiteCreate {
	wc.mutation.SetFacebook(s)
	return wc
}

// SetNillableFacebook sets the "facebook" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableFacebook(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetFacebook(*s)
	}
	return wc
}

// SetTwitter sets the "twitter" field.
func (wc *WebsiteCreate) SetTwitter(s string) *WebsiteCreate {
	wc.mutation.SetTwitter(s)
	return wc
}

// SetNillableTwitter sets the "twitter" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableTwitter(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetTwitter(*s)
	}
	return wc
}

// SetInstagram sets the "instagram" field.
func (wc *WebsiteCreate) SetInstagram(s string) *WebsiteCreate {
	wc.mutation.SetInstagram(s)
	return wc
}

// SetNillableInstagram sets the "instagram" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableInstagram(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetInstagram(*s)
	}
	return wc
}

// SetYoutube sets the "youtube" field.
func (wc *WebsiteCreate) SetYoutube(s string) *WebsiteCreate {
	wc.mutation.SetYoutube(s)
	return wc
}

// SetNillableYoutube sets the "youtube" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableYoutube(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetYoutube(*s)
	}
	return wc
}

// SetLinkedin sets the "linkedin" field.
func (wc *WebsiteCreate) SetLinkedin(s string) *WebsiteCreate {
	wc.mutation.SetLinkedin(s)
	return wc
}

// SetNillableLinkedin sets the "linkedin" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableLinkedin(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetLinkedin(*s)
	}
	return wc
}

// SetPinterest sets the "pinterest" field.
func (wc *WebsiteCreate) SetPinterest(s string) *WebsiteCreate {
	wc.mutation.SetPinterest(s)
	return wc
}

// SetNillablePinterest sets the "pinterest" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillablePinterest(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetPinterest(*s)
	}
	return wc
}

// SetMapCoordinates sets the "mapCoordinates" field.
func (wc *WebsiteCreate) SetMapCoordinates(m map[string]interface{}) *WebsiteCreate {
	wc.mutation.SetMapCoordinates(m)
	return wc
}

// SetLongitude sets the "longitude" field.
func (wc *WebsiteCreate) SetLongitude(s string) *WebsiteCreate {
	wc.mutation.SetLongitude(s)
	return wc
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableLongitude(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetLongitude(*s)
	}
	return wc
}

// SetLatitude sets the "latitude" field.
func (wc *WebsiteCreate) SetLatitude(s string) *WebsiteCreate {
	wc.mutation.SetLatitude(s)
	return wc
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableLatitude(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetLatitude(*s)
	}
	return wc
}

// SetAddress sets the "address" field.
func (wc *WebsiteCreate) SetAddress(s string) *WebsiteCreate {
	wc.mutation.SetAddress(s)
	return wc
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableAddress(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetAddress(*s)
	}
	return wc
}

// SetCity sets the "city" field.
func (wc *WebsiteCreate) SetCity(s string) *WebsiteCreate {
	wc.mutation.SetCity(s)
	return wc
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableCity(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetCity(*s)
	}
	return wc
}

// SetState sets the "state" field.
func (wc *WebsiteCreate) SetState(s string) *WebsiteCreate {
	wc.mutation.SetState(s)
	return wc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableState(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetState(*s)
	}
	return wc
}

// SetCountry sets the "country" field.
func (wc *WebsiteCreate) SetCountry(s string) *WebsiteCreate {
	wc.mutation.SetCountry(s)
	return wc
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableCountry(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetCountry(*s)
	}
	return wc
}

// SetZipCode sets the "zipCode" field.
func (wc *WebsiteCreate) SetZipCode(s string) *WebsiteCreate {
	wc.mutation.SetZipCode(s)
	return wc
}

// SetNillableZipCode sets the "zipCode" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableZipCode(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetZipCode(*s)
	}
	return wc
}

// SetPhoneNumber sets the "phoneNumber" field.
func (wc *WebsiteCreate) SetPhoneNumber(s string) *WebsiteCreate {
	wc.mutation.SetPhoneNumber(s)
	return wc
}

// SetNillablePhoneNumber sets the "phoneNumber" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillablePhoneNumber(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetPhoneNumber(*s)
	}
	return wc
}

// SetEmail sets the "email" field.
func (wc *WebsiteCreate) SetEmail(s string) *WebsiteCreate {
	wc.mutation.SetEmail(s)
	return wc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (wc *WebsiteCreate) SetNillableEmail(s *string) *WebsiteCreate {
	if s != nil {
		wc.SetEmail(*s)
	}
	return wc
}

// SetMetaTags sets the "metaTags" field.
func (wc *WebsiteCreate) SetMetaTags(m map[string]interface{}) *WebsiteCreate {
	wc.mutation.SetMetaTags(m)
	return wc
}

// SetID sets the "id" field.
func (wc *WebsiteCreate) SetID(s string) *WebsiteCreate {
	wc.mutation.SetID(s)
	return wc
}

// SetBusinessID sets the "business" edge to the Business entity by ID.
func (wc *WebsiteCreate) SetBusinessID(id string) *WebsiteCreate {
	wc.mutation.SetBusinessID(id)
	return wc
}

// SetBusiness sets the "business" edge to the Business entity.
func (wc *WebsiteCreate) SetBusiness(b *Business) *WebsiteCreate {
	return wc.SetBusinessID(b.ID)
}

// AddCustomBlockIDs adds the "customBlocks" edge to the CustomBlock entity by IDs.
func (wc *WebsiteCreate) AddCustomBlockIDs(ids ...string) *WebsiteCreate {
	wc.mutation.AddCustomBlockIDs(ids...)
	return wc
}

// AddCustomBlocks adds the "customBlocks" edges to the CustomBlock entity.
func (wc *WebsiteCreate) AddCustomBlocks(c ...*CustomBlock) *WebsiteCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wc.AddCustomBlockIDs(ids...)
}

// AddAssetIDs adds the "assets" edge to the Media entity by IDs.
func (wc *WebsiteCreate) AddAssetIDs(ids ...string) *WebsiteCreate {
	wc.mutation.AddAssetIDs(ids...)
	return wc
}

// AddAssets adds the "assets" edges to the Media entity.
func (wc *WebsiteCreate) AddAssets(m ...*Media) *WebsiteCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return wc.AddAssetIDs(ids...)
}

// Mutation returns the WebsiteMutation object of the builder.
func (wc *WebsiteCreate) Mutation() *WebsiteMutation {
	return wc.mutation
}

// Save creates the Website in the database.
func (wc *WebsiteCreate) Save(ctx context.Context) (*Website, error) {
	wc.defaults()
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WebsiteCreate) SaveX(ctx context.Context) *Website {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WebsiteCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WebsiteCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WebsiteCreate) defaults() {
	if _, ok := wc.mutation.CreationDate(); !ok {
		v := website.DefaultCreationDate()
		wc.mutation.SetCreationDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WebsiteCreate) check() error {
	if _, ok := wc.mutation.DomainName(); !ok {
		return &ValidationError{Name: "domainName", err: errors.New(`ent: missing required field "Website.domainName"`)}
	}
	if _, ok := wc.mutation.CreationDate(); !ok {
		return &ValidationError{Name: "creationDate", err: errors.New(`ent: missing required field "Website.creationDate"`)}
	}
	if _, ok := wc.mutation.LastUpdated(); !ok {
		return &ValidationError{Name: "lastUpdated", err: errors.New(`ent: missing required field "Website.lastUpdated"`)}
	}
	if _, ok := wc.mutation.BusinessID(); !ok {
		return &ValidationError{Name: "business", err: errors.New(`ent: missing required edge "Website.business"`)}
	}
	return nil
}

func (wc *WebsiteCreate) sqlSave(ctx context.Context) (*Website, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Website.ID type: %T", _spec.ID.Value)
		}
	}
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WebsiteCreate) createSpec() (*Website, *sqlgraph.CreateSpec) {
	var (
		_node = &Website{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(website.Table, sqlgraph.NewFieldSpec(website.FieldID, field.TypeString))
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wc.mutation.DomainName(); ok {
		_spec.SetField(website.FieldDomainName, field.TypeString, value)
		_node.DomainName = value
	}
	if value, ok := wc.mutation.HeadingText(); ok {
		_spec.SetField(website.FieldHeadingText, field.TypeString, value)
		_node.HeadingText = value
	}
	if value, ok := wc.mutation.BusinessLogo(); ok {
		_spec.SetField(website.FieldBusinessLogo, field.TypeString, value)
		_node.BusinessLogo = value
	}
	if value, ok := wc.mutation.BusinessName(); ok {
		_spec.SetField(website.FieldBusinessName, field.TypeString, value)
		_node.BusinessName = value
	}
	if value, ok := wc.mutation.BannerSectionBackgroundImage(); ok {
		_spec.SetField(website.FieldBannerSectionBackgroundImage, field.TypeString, value)
		_node.BannerSectionBackgroundImage = value
	}
	if value, ok := wc.mutation.BannerSectionBackgroundColor(); ok {
		_spec.SetField(website.FieldBannerSectionBackgroundColor, field.TypeString, value)
		_node.BannerSectionBackgroundColor = value
	}
	if value, ok := wc.mutation.BannerSectionText(); ok {
		_spec.SetField(website.FieldBannerSectionText, field.TypeString, value)
		_node.BannerSectionText = value
	}
	if value, ok := wc.mutation.ThreeItemsSectionHeadingText(); ok {
		_spec.SetField(website.FieldThreeItemsSectionHeadingText, field.TypeString, value)
		_node.ThreeItemsSectionHeadingText = value
	}
	if value, ok := wc.mutation.ThreeItemsSectionDetailsText(); ok {
		_spec.SetField(website.FieldThreeItemsSectionDetailsText, field.TypeString, value)
		_node.ThreeItemsSectionDetailsText = value
	}
	if value, ok := wc.mutation.ThreeItemsSectionItemOneText(); ok {
		_spec.SetField(website.FieldThreeItemsSectionItemOneText, field.TypeString, value)
		_node.ThreeItemsSectionItemOneText = value
	}
	if value, ok := wc.mutation.ThreeItemsSectionItemTwoText(); ok {
		_spec.SetField(website.FieldThreeItemsSectionItemTwoText, field.TypeString, value)
		_node.ThreeItemsSectionItemTwoText = value
	}
	if value, ok := wc.mutation.ThreeItemsSectionItemThreeText(); ok {
		_spec.SetField(website.FieldThreeItemsSectionItemThreeText, field.TypeString, value)
		_node.ThreeItemsSectionItemThreeText = value
	}
	if value, ok := wc.mutation.BannerTwoSectionBackgroundImage(); ok {
		_spec.SetField(website.FieldBannerTwoSectionBackgroundImage, field.TypeString, value)
		_node.BannerTwoSectionBackgroundImage = value
	}
	if value, ok := wc.mutation.BannerTwoSectionBackgroundColor(); ok {
		_spec.SetField(website.FieldBannerTwoSectionBackgroundColor, field.TypeString, value)
		_node.BannerTwoSectionBackgroundColor = value
	}
	if value, ok := wc.mutation.BannerTwoLeftSectionHeadingText(); ok {
		_spec.SetField(website.FieldBannerTwoLeftSectionHeadingText, field.TypeString, value)
		_node.BannerTwoLeftSectionHeadingText = value
	}
	if value, ok := wc.mutation.BannerTwoLeftSectionDetailsText(); ok {
		_spec.SetField(website.FieldBannerTwoLeftSectionDetailsText, field.TypeString, value)
		_node.BannerTwoLeftSectionDetailsText = value
	}
	if value, ok := wc.mutation.BannerTwoLeftSectionButtonText(); ok {
		_spec.SetField(website.FieldBannerTwoLeftSectionButtonText, field.TypeString, value)
		_node.BannerTwoLeftSectionButtonText = value
	}
	if value, ok := wc.mutation.BannerTwoLeftSectionButtonLink(); ok {
		_spec.SetField(website.FieldBannerTwoLeftSectionButtonLink, field.TypeString, value)
		_node.BannerTwoLeftSectionButtonLink = value
	}
	if value, ok := wc.mutation.BannerTwoRightSideImage(); ok {
		_spec.SetField(website.FieldBannerTwoRightSideImage, field.TypeString, value)
		_node.BannerTwoRightSideImage = value
	}
	if value, ok := wc.mutation.AchievementsSection(); ok {
		_spec.SetField(website.FieldAchievementsSection, field.TypeJSON, value)
		_node.AchievementsSection = value
	}
	if value, ok := wc.mutation.InventorySectionHeadingText(); ok {
		_spec.SetField(website.FieldInventorySectionHeadingText, field.TypeString, value)
		_node.InventorySectionHeadingText = value
	}
	if value, ok := wc.mutation.CreationDate(); ok {
		_spec.SetField(website.FieldCreationDate, field.TypeTime, value)
		_node.CreationDate = value
	}
	if value, ok := wc.mutation.LastUpdated(); ok {
		_spec.SetField(website.FieldLastUpdated, field.TypeTime, value)
		_node.LastUpdated = value
	}
	if value, ok := wc.mutation.Title(); ok {
		_spec.SetField(website.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.SetField(website.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := wc.mutation.Keywords(); ok {
		_spec.SetField(website.FieldKeywords, field.TypeString, value)
		_node.Keywords = value
	}
	if value, ok := wc.mutation.Language(); ok {
		_spec.SetField(website.FieldLanguage, field.TypeString, value)
		_node.Language = value
	}
	if value, ok := wc.mutation.Logo(); ok {
		_spec.SetField(website.FieldLogo, field.TypeString, value)
		_node.Logo = value
	}
	if value, ok := wc.mutation.Favicon(); ok {
		_spec.SetField(website.FieldFavicon, field.TypeString, value)
		_node.Favicon = value
	}
	if value, ok := wc.mutation.Facebook(); ok {
		_spec.SetField(website.FieldFacebook, field.TypeString, value)
		_node.Facebook = value
	}
	if value, ok := wc.mutation.Twitter(); ok {
		_spec.SetField(website.FieldTwitter, field.TypeString, value)
		_node.Twitter = value
	}
	if value, ok := wc.mutation.Instagram(); ok {
		_spec.SetField(website.FieldInstagram, field.TypeString, value)
		_node.Instagram = value
	}
	if value, ok := wc.mutation.Youtube(); ok {
		_spec.SetField(website.FieldYoutube, field.TypeString, value)
		_node.Youtube = value
	}
	if value, ok := wc.mutation.Linkedin(); ok {
		_spec.SetField(website.FieldLinkedin, field.TypeString, value)
		_node.Linkedin = value
	}
	if value, ok := wc.mutation.Pinterest(); ok {
		_spec.SetField(website.FieldPinterest, field.TypeString, value)
		_node.Pinterest = value
	}
	if value, ok := wc.mutation.MapCoordinates(); ok {
		_spec.SetField(website.FieldMapCoordinates, field.TypeJSON, value)
		_node.MapCoordinates = value
	}
	if value, ok := wc.mutation.Longitude(); ok {
		_spec.SetField(website.FieldLongitude, field.TypeString, value)
		_node.Longitude = value
	}
	if value, ok := wc.mutation.Latitude(); ok {
		_spec.SetField(website.FieldLatitude, field.TypeString, value)
		_node.Latitude = value
	}
	if value, ok := wc.mutation.Address(); ok {
		_spec.SetField(website.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := wc.mutation.City(); ok {
		_spec.SetField(website.FieldCity, field.TypeString, value)
		_node.City = value
	}
	if value, ok := wc.mutation.State(); ok {
		_spec.SetField(website.FieldState, field.TypeString, value)
		_node.State = value
	}
	if value, ok := wc.mutation.Country(); ok {
		_spec.SetField(website.FieldCountry, field.TypeString, value)
		_node.Country = value
	}
	if value, ok := wc.mutation.ZipCode(); ok {
		_spec.SetField(website.FieldZipCode, field.TypeString, value)
		_node.ZipCode = value
	}
	if value, ok := wc.mutation.PhoneNumber(); ok {
		_spec.SetField(website.FieldPhoneNumber, field.TypeString, value)
		_node.PhoneNumber = value
	}
	if value, ok := wc.mutation.Email(); ok {
		_spec.SetField(website.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := wc.mutation.MetaTags(); ok {
		_spec.SetField(website.FieldMetaTags, field.TypeJSON, value)
		_node.MetaTags = value
	}
	if nodes := wc.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   website.BusinessTable,
			Columns: []string{website.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.business_websites = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.CustomBlocksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   website.CustomBlocksTable,
			Columns: []string{website.CustomBlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customblock.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.AssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   website.AssetsTable,
			Columns: []string{website.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WebsiteCreateBulk is the builder for creating many Website entities in bulk.
type WebsiteCreateBulk struct {
	config
	err      error
	builders []*WebsiteCreate
}

// Save creates the Website entities in the database.
func (wcb *WebsiteCreateBulk) Save(ctx context.Context) ([]*Website, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Website, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WebsiteMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WebsiteCreateBulk) SaveX(ctx context.Context) []*Website {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WebsiteCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WebsiteCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
