package Dto

type CreatePlaceDTO struct {
	ID                  string                  `json:"id"`
	Name                string                  `json:"name"`
	Type                string                  `json:"type"`
	Description         *string                 `json:"description,omitempty"`
	Location            string                  `json:"location"`
	Email               *string                 `json:"email,omitempty"`
	Phone               *string                 `json:"phone,omitempty"`
	Website             *string                 `json:"website,omitempty"`
	CoverImage          *string                 `json:"cover_image,omitempty"`
	Picture             *string                 `json:"picture,omitempty"`
	PlaceSettings       *map[string]interface{} `json:"place_settings,omitempty"`
	OpeningHours        *map[string]interface{} `json:"opening_hours,omitempty"`
	SocialMedia         *map[string]interface{} `json:"social_media,omitempty"`
	Tags                *[]string               `json:"tags,omitempty"`
	Features            *[]string               `json:"features,omitempty"`
	AdditionalInfo      *map[string]interface{} `json:"additional_info,omitempty"`
	Images              *[]string               `json:"images,omitempty"`
	Availability        *map[string]interface{} `json:"availability,omitempty"`
	SpecialOffers       *string                 `json:"special_offers,omitempty"`
	SustainabilityScore *float64                `json:"sustainability_score,omitempty"`
	MapCoordinates      map[string]interface{}  `json:"map_coordinates"`
	SearchText          *string                 `json:"search_text,omitempty"`
	RelevanceScore      *float64                `json:"relevance_score,omitempty"`
	BusinessID          string                  `json:"business_id"`
}

type UpdatePlaceDTO struct {
	Name                *string                 `json:"name,omitempty"`
	Type                *string                 `json:"type,omitempty"`
	Description         *string                 `json:"description,omitempty"`
	Location            *string                 `json:"location,omitempty"`
	Email               *string                 `json:"email,omitempty"`
	Phone               *string                 `json:"phone,omitempty"`
	Website             *string                 `json:"website,omitempty"`
	CoverImage          *string                 `json:"cover_image,omitempty"`
	Picture             *string                 `json:"picture,omitempty"`
	PlaceSettings       *map[string]interface{} `json:"place_settings,omitempty"`
	OpeningHours        *map[string]interface{} `json:"opening_hours,omitempty"`
	SocialMedia         *map[string]interface{} `json:"social_media,omitempty"`
	Tags                *[]string               `json:"tags,omitempty"`
	Features            *[]string               `json:"features,omitempty"`
	AdditionalInfo      *map[string]interface{} `json:"additional_info,omitempty"`
	Images              *[]string               `json:"images,omitempty"`
	Availability        *map[string]interface{} `json:"availability,omitempty"`
	SpecialOffers       *string                 `json:"special_offers,omitempty"`
	SustainabilityScore *float64                `json:"sustainability_score,omitempty"`
	MapCoordinates      *map[string]interface{} `json:"map_coordinates,omitempty"`
	SearchText          *string                 `json:"search_text,omitempty"`
	RelevanceScore      *float64                `json:"relevance_score,omitempty"`
}
