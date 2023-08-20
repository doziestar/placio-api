package inventory

type IndustryType string

const (
	IndustryHotel      IndustryType = "hotel"
	IndustryRestaurant IndustryType = "restaurant"
	IndustryBar        IndustryType = "bar"
	IndustryClub       IndustryType = "club"
	IndustryGym        IndustryType = "gym"
	IndustryEvents     IndustryType = "events"
	IndustryRetail     IndustryType = "retail"
	IndustryOther      IndustryType = "other"
)

func IsValidIndustryType(it IndustryType) bool {
	switch it {
	case IndustryHotel, IndustryRestaurant, IndustryBar, IndustryClub, IndustryGym, IndustryEvents, IndustryRetail, IndustryOther:
		return true
	default:
		return false
	}
}
