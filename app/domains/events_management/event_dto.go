package events_management

import "time"

type EventDTO struct {
	ID                   string
	Name                 string
	EventType            string
	Status               string
	Location             string
	URL                  string
	Title                string
	TimeZone             string
	StartTime            time.Time
	EndTime              time.Time
	StartDate            string
	EndDate              string
	Frequency            string
	FrequencyInterval    string
	FrequencyDayOfWeek   string
	FrequencyDayOfMonth  string
	FrequencyMonthOfYear string
	VenueType            string
	VenueName            string
	VenueAddress         string
	VenueCity            string
	VenueState           string
	VenueCountry         string
	VenueZIP             string
	VenueLat             string
	VenueLon             string
	VenueURL             string
	VenuePhone           string
	VenueEmail           string
	Tags                 []string
	Description          string
	EventSettings        map[string]interface{}
	CoverImage           string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type EventFilter struct {
	EventType string
	Status    string
	Location  string
	Title     string
	TimeZone  string
	StartDate struct {
		From string
		To   string
	}
	EndDate struct {
		From string
		To   string
	}
	StartTime struct {
		From time.Time
		To   time.Time
	}
	EndTime struct {
		From time.Time
		To   time.Time
	}
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
