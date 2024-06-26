// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/website"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Website is the model entity for the Website schema.
type Website struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// DomainName holds the value of the "domainName" field.
	DomainName string `json:"domainName,omitempty"`
	// HeadingText holds the value of the "heading_text" field.
	HeadingText string `json:"heading_text,omitempty"`
	// BusinessLogo holds the value of the "business_logo" field.
	BusinessLogo string `json:"business_logo,omitempty"`
	// BusinessName holds the value of the "business_name" field.
	BusinessName string `json:"business_name,omitempty"`
	// BannerSectionBackgroundImage holds the value of the "banner_section_background_image" field.
	BannerSectionBackgroundImage string `json:"banner_section_background_image,omitempty"`
	// BannerSectionBackgroundColor holds the value of the "banner_section_background_color" field.
	BannerSectionBackgroundColor string `json:"banner_section_background_color,omitempty"`
	// BannerSectionText holds the value of the "banner_section_text" field.
	BannerSectionText string `json:"banner_section_text,omitempty"`
	// ThreeItemsSectionHeadingText holds the value of the "three_items_section_heading_text" field.
	ThreeItemsSectionHeadingText string `json:"three_items_section_heading_text,omitempty"`
	// ThreeItemsSectionDetailsText holds the value of the "three_items_section_details_text" field.
	ThreeItemsSectionDetailsText string `json:"three_items_section_details_text,omitempty"`
	// ThreeItemsSectionItemOneText holds the value of the "three_items_section_item_one_text" field.
	ThreeItemsSectionItemOneText string `json:"three_items_section_item_one_text,omitempty"`
	// ThreeItemsSectionItemTwoText holds the value of the "three_items_section_item_two_text" field.
	ThreeItemsSectionItemTwoText string `json:"three_items_section_item_two_text,omitempty"`
	// ThreeItemsSectionItemThreeText holds the value of the "three_items_section_item_three_text" field.
	ThreeItemsSectionItemThreeText string `json:"three_items_section_item_three_text,omitempty"`
	// BannerTwoSectionBackgroundImage holds the value of the "banner_two_section_background_image" field.
	BannerTwoSectionBackgroundImage string `json:"banner_two_section_background_image,omitempty"`
	// BannerTwoSectionBackgroundColor holds the value of the "banner_two_section_background_color" field.
	BannerTwoSectionBackgroundColor string `json:"banner_two_section_background_color,omitempty"`
	// BannerTwoLeftSectionHeadingText holds the value of the "banner_two_left_section_heading_text" field.
	BannerTwoLeftSectionHeadingText string `json:"banner_two_left_section_heading_text,omitempty"`
	// BannerTwoLeftSectionDetailsText holds the value of the "banner_two_left_section_details_text" field.
	BannerTwoLeftSectionDetailsText string `json:"banner_two_left_section_details_text,omitempty"`
	// BannerTwoLeftSectionButtonText holds the value of the "banner_two_left_section_button_text" field.
	BannerTwoLeftSectionButtonText string `json:"banner_two_left_section_button_text,omitempty"`
	// BannerTwoLeftSectionButtonLink holds the value of the "banner_two_left_section_button_link" field.
	BannerTwoLeftSectionButtonLink string `json:"banner_two_left_section_button_link,omitempty"`
	// BannerTwoRightSideImage holds the value of the "banner_two_right_side_image" field.
	BannerTwoRightSideImage string `json:"banner_two_right_side_image,omitempty"`
	// AchievementsSection holds the value of the "achievements_section" field.
	AchievementsSection map[string]interface{} `json:"achievements_section,omitempty"`
	// InventorySectionHeadingText holds the value of the "Inventory_section_heading_text" field.
	InventorySectionHeadingText string `json:"Inventory_section_heading_text,omitempty"`
	// CreationDate holds the value of the "creationDate" field.
	CreationDate time.Time `json:"creationDate,omitempty"`
	// LastUpdated holds the value of the "lastUpdated" field.
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Keywords holds the value of the "keywords" field.
	Keywords string `json:"keywords,omitempty"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty"`
	// Logo holds the value of the "logo" field.
	Logo string `json:"logo,omitempty"`
	// Favicon holds the value of the "favicon" field.
	Favicon string `json:"favicon,omitempty"`
	// Facebook holds the value of the "facebook" field.
	Facebook string `json:"facebook,omitempty"`
	// Twitter holds the value of the "twitter" field.
	Twitter string `json:"twitter,omitempty"`
	// Instagram holds the value of the "instagram" field.
	Instagram string `json:"instagram,omitempty"`
	// Youtube holds the value of the "youtube" field.
	Youtube string `json:"youtube,omitempty"`
	// Linkedin holds the value of the "linkedin" field.
	Linkedin string `json:"linkedin,omitempty"`
	// Pinterest holds the value of the "pinterest" field.
	Pinterest string `json:"pinterest,omitempty"`
	// MapCoordinates holds the value of the "mapCoordinates" field.
	MapCoordinates map[string]interface{} `json:"mapCoordinates,omitempty"`
	// Longitude holds the value of the "longitude" field.
	Longitude string `json:"longitude,omitempty"`
	// Latitude holds the value of the "latitude" field.
	Latitude string `json:"latitude,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// ZipCode holds the value of the "zipCode" field.
	ZipCode string `json:"zipCode,omitempty"`
	// PhoneNumber holds the value of the "phoneNumber" field.
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// MetaTags holds the value of the "metaTags" field.
	MetaTags map[string]interface{} `json:"metaTags,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WebsiteQuery when eager-loading is set.
	Edges             WebsiteEdges `json:"edges"`
	business_websites *string
	template_websites *string
	selectValues      sql.SelectValues
}

// WebsiteEdges holds the relations/edges for other nodes in the graph.
type WebsiteEdges struct {
	// Business holds the value of the business edge.
	Business *Business `json:"business,omitempty"`
	// CustomBlocks holds the value of the customBlocks edge.
	CustomBlocks []*CustomBlock `json:"customBlocks,omitempty"`
	// Assets holds the value of the assets edge.
	Assets []*Media `json:"assets,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// BusinessOrErr returns the Business value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WebsiteEdges) BusinessOrErr() (*Business, error) {
	if e.loadedTypes[0] {
		if e.Business == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: business.Label}
		}
		return e.Business, nil
	}
	return nil, &NotLoadedError{edge: "business"}
}

// CustomBlocksOrErr returns the CustomBlocks value or an error if the edge
// was not loaded in eager-loading.
func (e WebsiteEdges) CustomBlocksOrErr() ([]*CustomBlock, error) {
	if e.loadedTypes[1] {
		return e.CustomBlocks, nil
	}
	return nil, &NotLoadedError{edge: "customBlocks"}
}

// AssetsOrErr returns the Assets value or an error if the edge
// was not loaded in eager-loading.
func (e WebsiteEdges) AssetsOrErr() ([]*Media, error) {
	if e.loadedTypes[2] {
		return e.Assets, nil
	}
	return nil, &NotLoadedError{edge: "assets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Website) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case website.FieldAchievementsSection, website.FieldMapCoordinates, website.FieldMetaTags:
			values[i] = new([]byte)
		case website.FieldID, website.FieldDomainName, website.FieldHeadingText, website.FieldBusinessLogo, website.FieldBusinessName, website.FieldBannerSectionBackgroundImage, website.FieldBannerSectionBackgroundColor, website.FieldBannerSectionText, website.FieldThreeItemsSectionHeadingText, website.FieldThreeItemsSectionDetailsText, website.FieldThreeItemsSectionItemOneText, website.FieldThreeItemsSectionItemTwoText, website.FieldThreeItemsSectionItemThreeText, website.FieldBannerTwoSectionBackgroundImage, website.FieldBannerTwoSectionBackgroundColor, website.FieldBannerTwoLeftSectionHeadingText, website.FieldBannerTwoLeftSectionDetailsText, website.FieldBannerTwoLeftSectionButtonText, website.FieldBannerTwoLeftSectionButtonLink, website.FieldBannerTwoRightSideImage, website.FieldInventorySectionHeadingText, website.FieldTitle, website.FieldDescription, website.FieldKeywords, website.FieldLanguage, website.FieldLogo, website.FieldFavicon, website.FieldFacebook, website.FieldTwitter, website.FieldInstagram, website.FieldYoutube, website.FieldLinkedin, website.FieldPinterest, website.FieldLongitude, website.FieldLatitude, website.FieldAddress, website.FieldCity, website.FieldState, website.FieldCountry, website.FieldZipCode, website.FieldPhoneNumber, website.FieldEmail:
			values[i] = new(sql.NullString)
		case website.FieldCreationDate, website.FieldLastUpdated:
			values[i] = new(sql.NullTime)
		case website.ForeignKeys[0]: // business_websites
			values[i] = new(sql.NullString)
		case website.ForeignKeys[1]: // template_websites
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Website fields.
func (w *Website) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case website.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				w.ID = value.String
			}
		case website.FieldDomainName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domainName", values[i])
			} else if value.Valid {
				w.DomainName = value.String
			}
		case website.FieldHeadingText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field heading_text", values[i])
			} else if value.Valid {
				w.HeadingText = value.String
			}
		case website.FieldBusinessLogo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_logo", values[i])
			} else if value.Valid {
				w.BusinessLogo = value.String
			}
		case website.FieldBusinessName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_name", values[i])
			} else if value.Valid {
				w.BusinessName = value.String
			}
		case website.FieldBannerSectionBackgroundImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_section_background_image", values[i])
			} else if value.Valid {
				w.BannerSectionBackgroundImage = value.String
			}
		case website.FieldBannerSectionBackgroundColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_section_background_color", values[i])
			} else if value.Valid {
				w.BannerSectionBackgroundColor = value.String
			}
		case website.FieldBannerSectionText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_section_text", values[i])
			} else if value.Valid {
				w.BannerSectionText = value.String
			}
		case website.FieldThreeItemsSectionHeadingText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field three_items_section_heading_text", values[i])
			} else if value.Valid {
				w.ThreeItemsSectionHeadingText = value.String
			}
		case website.FieldThreeItemsSectionDetailsText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field three_items_section_details_text", values[i])
			} else if value.Valid {
				w.ThreeItemsSectionDetailsText = value.String
			}
		case website.FieldThreeItemsSectionItemOneText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field three_items_section_item_one_text", values[i])
			} else if value.Valid {
				w.ThreeItemsSectionItemOneText = value.String
			}
		case website.FieldThreeItemsSectionItemTwoText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field three_items_section_item_two_text", values[i])
			} else if value.Valid {
				w.ThreeItemsSectionItemTwoText = value.String
			}
		case website.FieldThreeItemsSectionItemThreeText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field three_items_section_item_three_text", values[i])
			} else if value.Valid {
				w.ThreeItemsSectionItemThreeText = value.String
			}
		case website.FieldBannerTwoSectionBackgroundImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_two_section_background_image", values[i])
			} else if value.Valid {
				w.BannerTwoSectionBackgroundImage = value.String
			}
		case website.FieldBannerTwoSectionBackgroundColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_two_section_background_color", values[i])
			} else if value.Valid {
				w.BannerTwoSectionBackgroundColor = value.String
			}
		case website.FieldBannerTwoLeftSectionHeadingText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_two_left_section_heading_text", values[i])
			} else if value.Valid {
				w.BannerTwoLeftSectionHeadingText = value.String
			}
		case website.FieldBannerTwoLeftSectionDetailsText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_two_left_section_details_text", values[i])
			} else if value.Valid {
				w.BannerTwoLeftSectionDetailsText = value.String
			}
		case website.FieldBannerTwoLeftSectionButtonText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_two_left_section_button_text", values[i])
			} else if value.Valid {
				w.BannerTwoLeftSectionButtonText = value.String
			}
		case website.FieldBannerTwoLeftSectionButtonLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_two_left_section_button_link", values[i])
			} else if value.Valid {
				w.BannerTwoLeftSectionButtonLink = value.String
			}
		case website.FieldBannerTwoRightSideImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_two_right_side_image", values[i])
			} else if value.Valid {
				w.BannerTwoRightSideImage = value.String
			}
		case website.FieldAchievementsSection:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field achievements_section", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &w.AchievementsSection); err != nil {
					return fmt.Errorf("unmarshal field achievements_section: %w", err)
				}
			}
		case website.FieldInventorySectionHeadingText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Inventory_section_heading_text", values[i])
			} else if value.Valid {
				w.InventorySectionHeadingText = value.String
			}
		case website.FieldCreationDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field creationDate", values[i])
			} else if value.Valid {
				w.CreationDate = value.Time
			}
		case website.FieldLastUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lastUpdated", values[i])
			} else if value.Valid {
				w.LastUpdated = value.Time
			}
		case website.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				w.Title = value.String
			}
		case website.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				w.Description = value.String
			}
		case website.FieldKeywords:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field keywords", values[i])
			} else if value.Valid {
				w.Keywords = value.String
			}
		case website.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				w.Language = value.String
			}
		case website.FieldLogo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo", values[i])
			} else if value.Valid {
				w.Logo = value.String
			}
		case website.FieldFavicon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field favicon", values[i])
			} else if value.Valid {
				w.Favicon = value.String
			}
		case website.FieldFacebook:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field facebook", values[i])
			} else if value.Valid {
				w.Facebook = value.String
			}
		case website.FieldTwitter:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field twitter", values[i])
			} else if value.Valid {
				w.Twitter = value.String
			}
		case website.FieldInstagram:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field instagram", values[i])
			} else if value.Valid {
				w.Instagram = value.String
			}
		case website.FieldYoutube:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field youtube", values[i])
			} else if value.Valid {
				w.Youtube = value.String
			}
		case website.FieldLinkedin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field linkedin", values[i])
			} else if value.Valid {
				w.Linkedin = value.String
			}
		case website.FieldPinterest:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pinterest", values[i])
			} else if value.Valid {
				w.Pinterest = value.String
			}
		case website.FieldMapCoordinates:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field mapCoordinates", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &w.MapCoordinates); err != nil {
					return fmt.Errorf("unmarshal field mapCoordinates: %w", err)
				}
			}
		case website.FieldLongitude:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				w.Longitude = value.String
			}
		case website.FieldLatitude:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				w.Latitude = value.String
			}
		case website.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				w.Address = value.String
			}
		case website.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				w.City = value.String
			}
		case website.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				w.State = value.String
			}
		case website.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				w.Country = value.String
			}
		case website.FieldZipCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field zipCode", values[i])
			} else if value.Valid {
				w.ZipCode = value.String
			}
		case website.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phoneNumber", values[i])
			} else if value.Valid {
				w.PhoneNumber = value.String
			}
		case website.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				w.Email = value.String
			}
		case website.FieldMetaTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metaTags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &w.MetaTags); err != nil {
					return fmt.Errorf("unmarshal field metaTags: %w", err)
				}
			}
		case website.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_websites", values[i])
			} else if value.Valid {
				w.business_websites = new(string)
				*w.business_websites = value.String
			}
		case website.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field template_websites", value)
			} else if value.Valid {
				w.template_websites = new(string)
				*w.template_websites = string(value.Int64)
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Website.
// This includes values selected through modifiers, order, etc.
func (w *Website) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// QueryBusiness queries the "business" edge of the Website entity.
func (w *Website) QueryBusiness() *BusinessQuery {
	return NewWebsiteClient(w.config).QueryBusiness(w)
}

// QueryCustomBlocks queries the "customBlocks" edge of the Website entity.
func (w *Website) QueryCustomBlocks() *CustomBlockQuery {
	return NewWebsiteClient(w.config).QueryCustomBlocks(w)
}

// QueryAssets queries the "assets" edge of the Website entity.
func (w *Website) QueryAssets() *MediaQuery {
	return NewWebsiteClient(w.config).QueryAssets(w)
}

// Update returns a builder for updating this Website.
// Note that you need to call Website.Unwrap() before calling this method if this Website
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Website) Update() *WebsiteUpdateOne {
	return NewWebsiteClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Website entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Website) Unwrap() *Website {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Website is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Website) String() string {
	var builder strings.Builder
	builder.WriteString("Website(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("domainName=")
	builder.WriteString(w.DomainName)
	builder.WriteString(", ")
	builder.WriteString("heading_text=")
	builder.WriteString(w.HeadingText)
	builder.WriteString(", ")
	builder.WriteString("business_logo=")
	builder.WriteString(w.BusinessLogo)
	builder.WriteString(", ")
	builder.WriteString("business_name=")
	builder.WriteString(w.BusinessName)
	builder.WriteString(", ")
	builder.WriteString("banner_section_background_image=")
	builder.WriteString(w.BannerSectionBackgroundImage)
	builder.WriteString(", ")
	builder.WriteString("banner_section_background_color=")
	builder.WriteString(w.BannerSectionBackgroundColor)
	builder.WriteString(", ")
	builder.WriteString("banner_section_text=")
	builder.WriteString(w.BannerSectionText)
	builder.WriteString(", ")
	builder.WriteString("three_items_section_heading_text=")
	builder.WriteString(w.ThreeItemsSectionHeadingText)
	builder.WriteString(", ")
	builder.WriteString("three_items_section_details_text=")
	builder.WriteString(w.ThreeItemsSectionDetailsText)
	builder.WriteString(", ")
	builder.WriteString("three_items_section_item_one_text=")
	builder.WriteString(w.ThreeItemsSectionItemOneText)
	builder.WriteString(", ")
	builder.WriteString("three_items_section_item_two_text=")
	builder.WriteString(w.ThreeItemsSectionItemTwoText)
	builder.WriteString(", ")
	builder.WriteString("three_items_section_item_three_text=")
	builder.WriteString(w.ThreeItemsSectionItemThreeText)
	builder.WriteString(", ")
	builder.WriteString("banner_two_section_background_image=")
	builder.WriteString(w.BannerTwoSectionBackgroundImage)
	builder.WriteString(", ")
	builder.WriteString("banner_two_section_background_color=")
	builder.WriteString(w.BannerTwoSectionBackgroundColor)
	builder.WriteString(", ")
	builder.WriteString("banner_two_left_section_heading_text=")
	builder.WriteString(w.BannerTwoLeftSectionHeadingText)
	builder.WriteString(", ")
	builder.WriteString("banner_two_left_section_details_text=")
	builder.WriteString(w.BannerTwoLeftSectionDetailsText)
	builder.WriteString(", ")
	builder.WriteString("banner_two_left_section_button_text=")
	builder.WriteString(w.BannerTwoLeftSectionButtonText)
	builder.WriteString(", ")
	builder.WriteString("banner_two_left_section_button_link=")
	builder.WriteString(w.BannerTwoLeftSectionButtonLink)
	builder.WriteString(", ")
	builder.WriteString("banner_two_right_side_image=")
	builder.WriteString(w.BannerTwoRightSideImage)
	builder.WriteString(", ")
	builder.WriteString("achievements_section=")
	builder.WriteString(fmt.Sprintf("%v", w.AchievementsSection))
	builder.WriteString(", ")
	builder.WriteString("Inventory_section_heading_text=")
	builder.WriteString(w.InventorySectionHeadingText)
	builder.WriteString(", ")
	builder.WriteString("creationDate=")
	builder.WriteString(w.CreationDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("lastUpdated=")
	builder.WriteString(w.LastUpdated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(w.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(w.Description)
	builder.WriteString(", ")
	builder.WriteString("keywords=")
	builder.WriteString(w.Keywords)
	builder.WriteString(", ")
	builder.WriteString("language=")
	builder.WriteString(w.Language)
	builder.WriteString(", ")
	builder.WriteString("logo=")
	builder.WriteString(w.Logo)
	builder.WriteString(", ")
	builder.WriteString("favicon=")
	builder.WriteString(w.Favicon)
	builder.WriteString(", ")
	builder.WriteString("facebook=")
	builder.WriteString(w.Facebook)
	builder.WriteString(", ")
	builder.WriteString("twitter=")
	builder.WriteString(w.Twitter)
	builder.WriteString(", ")
	builder.WriteString("instagram=")
	builder.WriteString(w.Instagram)
	builder.WriteString(", ")
	builder.WriteString("youtube=")
	builder.WriteString(w.Youtube)
	builder.WriteString(", ")
	builder.WriteString("linkedin=")
	builder.WriteString(w.Linkedin)
	builder.WriteString(", ")
	builder.WriteString("pinterest=")
	builder.WriteString(w.Pinterest)
	builder.WriteString(", ")
	builder.WriteString("mapCoordinates=")
	builder.WriteString(fmt.Sprintf("%v", w.MapCoordinates))
	builder.WriteString(", ")
	builder.WriteString("longitude=")
	builder.WriteString(w.Longitude)
	builder.WriteString(", ")
	builder.WriteString("latitude=")
	builder.WriteString(w.Latitude)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(w.Address)
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(w.City)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(w.State)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(w.Country)
	builder.WriteString(", ")
	builder.WriteString("zipCode=")
	builder.WriteString(w.ZipCode)
	builder.WriteString(", ")
	builder.WriteString("phoneNumber=")
	builder.WriteString(w.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(w.Email)
	builder.WriteString(", ")
	builder.WriteString("metaTags=")
	builder.WriteString(fmt.Sprintf("%v", w.MetaTags))
	builder.WriteByte(')')
	return builder.String()
}

// Websites is a parsable slice of Website.
type Websites []*Website
