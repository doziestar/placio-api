package feature_releases

type FeatureReleaseDTO struct {
	FeatureName       string                 `json:"feature_name"`
	Description       string                 `json:"description"`
	State             string                 `json:"state"`
	EligibilityRules  map[string]interface{} `json:"eligibility_rules"`
	ReleaseDate       string                 `json:"release_date"`
	DocumentationLink string                 `json:"documentation_link"`
	Metadata          map[string]interface{} `json:"metadata"`
}
