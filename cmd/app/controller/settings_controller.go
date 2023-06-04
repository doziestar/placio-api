package controller

import "placio-app/service"

type SettingsController struct {
	generalSettingsService service.GeneralSettingsService
	notificationsSettingsService service.NotificationsSettingsService
	accountSettingsService service.AccountSettingsService
	connectedAccountService service.ConnectedAccountService
	contentSettingsService service.ContentSettingsService
	// Add other services if needed
}