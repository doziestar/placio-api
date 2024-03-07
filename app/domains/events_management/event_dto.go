package events_management

import "time"

type EventDTO struct {
	ID                          string                 `json:"id"`
	Name                        string                 `json:"name"`
	EventType                   string                 `json:"eventType"`
	Status                      string                 `json:"status"`
	Location                    string                 `json:"location"`
	URL                         string                 `json:"url"`
	Title                       string                 `json:"title"`
	TimeZone                    string                 `json:"timeZone"`
	StartTime                   time.Time              `json:"startTime"`
	EndTime                     time.Time              `json:"endTime"`
	StartDate                   string                 `json:"startDate"`
	EndDate                     string                 `json:"endDate"`
	Frequency                   string                 `json:"frequency"`
	FrequencyInterval           string                 `json:"frequencyInterval"`
	FrequencyDayOfWeek          string                 `json:"frequencyDayOfWeek"`
	FrequencyDayOfMonth         string                 `json:"frequencyDayOfMonth"`
	FrequencyMonthOfYear        string                 `json:"frequencyMonthOfYear"`
	VenueType                   string                 `json:"venueType"`
	VenueName                   string                 `json:"venueName"`
	VenueAddress                string                 `json:"venueAddress"`
	VenueCity                   string                 `json:"venueCity"`
	VenueState                  string                 `json:"venueState"`
	VenueCountry                string                 `json:"venueCountry"`
	VenueZip                    string                 `json:"venueZip"`
	VenueLat                    string                 `json:"venueLat"`
	VenueLon                    string                 `json:"venueLon"`
	VenueURL                    string                 `json:"venueUrl"`
	VenuePhone                  string                 `json:"venuePhone"`
	VenueEmail                  string                 `json:"venueEmail"`
	Tags                        []string               `json:"tags"`
	Description                 string                 `json:"description"`
	CoverImage                  string                 `json:"coverImage"`
	CreatedAt                   time.Time              `json:"createdAt"`
	UpdatedAt                   time.Time              `json:"updatedAt"`
	MapCoordinates              map[string]interface{} `json:"mapCoordinates"`
	Longitude                   string                 `json:"longitude"`
	Latitude                    string                 `json:"latitude"`
	SearchText                  string                 `json:"searchText"`
	RelevanceScore              float64                `json:"relevanceScore"`
	FollowerCount               int                    `json:"followerCount"`
	FollowingCount              int                    `json:"followingCount"`
	EventSettings               map[string]interface{} `json:"eventSettings"`
	OrganizerInfo               []OrganizerInfo        `json:"organizerInfo"`
	IsPremium                   bool                   `json:"isPremium"`
	IsPublished                 bool                   `json:"isPublished"`
	IsOnline                    bool                   `json:"isOnline"`
	IsFree                      bool                   `json:"isFree"`
	IsPaid                      bool                   `json:"isPaid"`
	IsPublic                    bool                   `json:"isPublic"`
	IsOnlineOnly                bool                   `json:"isOnlineOnly"`
	IsInPersonOnly              bool                   `json:"isInPersonOnly"`
	IsHybrid                    bool                   `json:"isHybrid"`
	IsOnlineAndInPerson         bool                   `json:"isOnlineAndInPerson"`
	IsOnlineAndInPersonOnly     bool                   `json:"isOnlineAndInPersonOnly"`
	IsOnlineAndInPersonOrHybrid bool                   `json:"isOnlineAndInPersonOrHybrid"`
	LikedByCurrentUser          bool                   `json:"likedByCurrentUser"`
	FollowedByCurrentUser       bool                   `json:"followedByCurrentUser"`
	RegistrationType            string                 `json:"registrationType"`
	RegistrationURL             string                 `json:"registrationUrl"`
	IsPhysicallyAccessible      bool                   `json:"isPhysicallyAccessible"`
	AccessibilityInfo           string                 `json:"accessibilityInfo"`
	IsVirtuallyAccessible       bool                   `json:"isVirtuallyAccessible"`
}

type EventFilter struct {
	EventType string `form:"eventType"`
	Status    string `form:"status"`
	Location  string `form:"location"`
	Title     string `form:"title"`
	TimeZone  string `form:"timeZone"`
	StartDate struct {
		From string `form:"startDateFrom"`
		To   string `form:"startDateTo"`
	} `form:"startDate"`
	EndDate struct {
		From string `form:"endDateFrom"`
		To   string `form:"endDateTo"`
	} `form:"endDate"`
	StartTime struct {
		From time.Time `form:"startTimeFrom"`
		To   time.Time `form:"startTimeTo"`
	} `form:"startTime"`
	EndTime struct {
		From time.Time `form:"endTimeFrom"`
		To   time.Time `form:"endTimeTo"`
	} `form:"endTime"`
}

// CheckInMethod describes the method of check-in, be it QR codes, facial recognition, or carrier pigeons.
type CheckInMethod string

const (
	CheckInMethodQRCode            CheckInMethod = "QR_CODE"
	CheckInMethodFacialRecognition CheckInMethod = "FACIAL_RECOGNITION"
	CheckInMethodEvehtIDAndUser    CheckInMethod = "EVENT_ID_AND_USER"
)

// SocialPlatform outlines the social media venues where event can shine.
type SocialPlatform string

const (
	SocialPlatformFacebook  SocialPlatform = "FACEBOOK"
	SocialPlatformTwitter   SocialPlatform = "TWITTER"
	SocialPlatformInstagram SocialPlatform = "INSTAGRAM"
	SocialPlatformLinkedIn  SocialPlatform = "LINKEDIN"
	SocialPlatformYouTube   SocialPlatform = "YOUTUBE"
	SocialPlatformTikTok    SocialPlatform = "TIKTOK"
)

// EventAnalyticsDTO holds the crystal ball of event's success, full of graphs and numbers.
type EventAnalyticsDTO struct {
	// Analytics data like AttendeeCount, EngagementRate, RevenueGenerated, etc.
}

// EventNotificationDTO is your event's herald, announcing news to all and sundry.
type EventNotificationDTO struct {
	// Details for crafting notifications like Message, TargetAudience, SendTime, etc.
}

// OnSiteToolsDTO are the gadgets and gizmos that make the event run without a hitch on the day.
type OnSiteToolsDTO struct {
	// On-site tools info like CheckInStations, BadgePrinters, PaymentKiosks, etc.
}

// ComplianceRulesDTO are the rules of the realm, keeping the event in line with the law.
type ComplianceRulesDTO struct {
	// Details related to compliance like DataRetentionPolicies, AccessibilityStandards, etc.
}

// AdDTO is event's billboard, bright and attention-grabbing.
type AdDTO struct {
	// Fields for AdContent, TargetAudience, Duration, Analytics, etc.
}

// AdvancedTicketOptions are the advanced ticketing options for the event.
type AdvancedTicketOptions struct {
	// Options for advanced ticketing like SeatReservations, SpecialPricing, etc.
}

// PersonalizationPreferences are the preferences of the attendees, making the event a personalized experience.
type PersonalizationPreferences struct {
	// Preferences like DietaryRestrictions, PreferredSessions, etc.
}

// VenueMapDetails are the details of the venue map, showing the way to the event's treasure.
type VenueMapDetails struct {
	// Details like VenueLayout, RealTimeUpdates, AttendeeTracking, etc.
}

// GamificationOptions are the options for gamifying the event experience.
type GamificationOptions struct {
	// Options like Rewards, Leaderboards, Challenges, etc.
}

// NetworkingOptions are the options for facilitating networking among the attendees.
type NetworkingOptions struct {
	// Options like RealTimeChat, NetworkingFacilitation, etc.
}

// AdvancedAnalyticsDTO are the advanced analytics for the event, predicting the future and providing actionable insights.
type AdvancedAnalyticsDTO struct {
	// Analytics like PredictiveModelling, ActionableInsights, etc.
}

// VendorDTO are the details of the vendors and sponsors, making the event a grand affair.
type VendorDTO struct {
	// Details like Booths, SponsoredSessions, etc.
}

// IncidentReportDTO are the details of the incidents reported and the response system for event management.
type IncidentReportDTO struct {
	// Details like RealTimeReporting, ResponseSystem, etc.
}

// CustomAppFeatures are the features of the custom event app, making it unique and special.
type CustomAppFeatures struct {
	// Features like CustomUI, SpecialFunctionalities, etc.
}

// IntegrationDetailsDTO are the details of the integration with external services and APIs for additional functionalities.
type IntegrationDetailsDTO struct {
	// Details like ExternalServices, APIs, AdditionalFunctionalities, etc.
}

// LoyaltyOptionsDTO are the options for the loyalty and rewards program for frequent attendees.
type LoyaltyOptionsDTO struct {
	// Options like Rewards, Points, Leaderboards, etc.
}

// AccessibilityOptionsDTO are the options for providing real-time accessibility services for attendees with disabilities.
type AccessibilityOptionsDTO struct {
	// Options like RealTimeServices, AccessibilityStandards, etc.
}

// LanguageSupportDTO are the options for multi-language support for international attendees.
type LanguageSupportDTO struct {
	// Options like LanguageOptions, TranslationServices, etc.
}

type CustomAppDetails struct {
}

// OrganizerInput represents the necessary input to identify an organizer.
type OrganizerInput struct {
	OrganizerID   string
	OrganizerType string
}

type OrganizerInfo struct {
	ID   string
	Type string
}

// ContextKey is a value for use with context.WithValue
type ContextKey string

const OrganizerContextKey ContextKey = "organizers"
