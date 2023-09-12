package notifications

import (
	"context"
	"placio-app/ent"
)

const (
	OrderNotification = iota
	MessageNotification
	FollowerNotification
	EventNotification
	CustomNotification
)

type INotification interface {
	SendNotification(ctx context.Context, notificationData *Notification) error
	CreateNotification(ctx context.Context, notification *Notification) (*ent.Notification, error)
	GetNotification(ctx context.Context, id string) (*ent.Notification, error)
	GetNotifications(ctx context.Context, userID string, limit int, offset int) ([]*ent.Notification, error)
	GetUnreadNotifications(ctx context.Context, userID string, limit int, offset int) ([]*ent.Notification, error)
	GetUnreadNotificationsCount(ctx context.Context, userID string) (int, error)
	GetUnreadNotificationsCountByType(ctx context.Context, userID string, notificationType int) (int, error)
	DeleteNotification(ctx context.Context, id string) error
}

type NotificationService struct {
	client *ent.Client
}

func NewNotificationService(client *ent.Client) *NotificationService {
	return &NotificationService{
		client: client,
	}
}

func (n *NotificationService) SendNotification(ctx context.Context, notificationData *Notification) error {
	return nil
}

func (n *NotificationService) CreateNotification(ctx context.Context, notification *Notification) (*ent.Notification, error) {
	return nil, nil
}

func (n *NotificationService) GetNotification(ctx context.Context, id string) (*ent.Notification, error) {
	return nil, nil
}

func (n *NotificationService) GetNotifications(ctx context.Context, userID string, limit int, offset int) ([]*ent.Notification, error) {
	return nil, nil
}

func (n *NotificationService) GetUnreadNotifications(ctx context.Context, userID string, limit int, offset int) ([]*ent.Notification, error) {
	return nil, nil
}

func (n *NotificationService) GetUnreadNotificationsCount(ctx context.Context, userID string) (int, error) {
	return 0, nil
}

func (n *NotificationService) GetUnreadNotificationsCountByType(ctx context.Context, userID string, notificationType int) (int, error) {
	return 0, nil
}

func (n *NotificationService) DeleteNotification(ctx context.Context, id string) error {
	return nil
}
