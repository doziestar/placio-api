package settings

type SettingsController struct {
	generalSettingsService       GeneralSettingsService
	notificationsSettingsService NotificationsSettingsService
	accountSettingsService       AccountSettingsService
	connectedAccountService      ConnectedAccountService
	contentSettingsService       ContentSettingsService
	// Add other services if needed
}
