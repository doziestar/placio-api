package notifications

import (
	"context"
	"github.com/google/uuid"
	"placio-app/ent"
	"placio-app/ent/notification"
	"placio-app/ent/user"
	"time"
)

const (
	OrderNotification = iota
	MessageNotification
	FollowerNotification
	EventNotification
	CustomNotification
)

type INotification interface {
	SendNotification(ctx context.Context, notificationData *ent.Notification) error
	CreateNotification(ctx context.Context, notification *ent.Notification) (*ent.Notification, error)
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

func (n *NotificationService) SendNotification(ctx context.Context, notificationData *ent.Notification) error {
	// Logic to send a notification
	// This could be sending a push notification, email, or any other form of notification.
	// You may want to use external services like Firebase, SendGrid, etc.
	return nil
}

func (n *NotificationService) CreateNotification(ctx context.Context, notification *ent.Notification) (*ent.Notification, error) {
	create := n.client.Notification.Create().
		SetID(uuid.New().String()).
		SetTitle(notification.Title).
		SetMessage(notification.Message).
		SetType(notification.Type).
		SetTriggeredBy(notification.TriggeredBy).
		SetTriggeredTo(notification.TriggeredTo).
		SetNotifiableType(notification.NotifiableType).
		SetNotifiableID(notification.NotifiableID).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now())

	if notification.Edges.User != nil {
		create.AddUser(notification.Edges.User...)
	}
	if notification.Edges.BusinessAccount != nil {
		create.AddBusinessAccount(notification.Edges.BusinessAccount...)
	}
	if notification.Edges.Place != nil {
		create.AddPlace(notification.Edges.Place...)
	}
	if notification.Edges.Post != nil {
		create.AddPost(notification.Edges.Post...)
	}
	if notification.Edges.Comment != nil {
		create.AddComment(notification.Edges.Comment...)
	}

	return create.Save(ctx)
}

func (n *NotificationService) GetNotification(ctx context.Context, id string) (*ent.Notification, error) {
	// we need set it the is_read to true
	notification, err := n.client.Notification.Query().
		Where(notification.ID(id)).
		WithUser().
		WithBusinessAccount().
		WithPlace().
		WithPost().
		WithComment().
		First(ctx)

	if err != nil {
		return nil, err
	}

	if notification.IsRead == false {
		notification.Update().
			SetIsRead(true).
			Save(ctx)
	}

	return notification, nil
}

func (n *NotificationService) GetNotifications(ctx context.Context, userID string, limit int, offset int) ([]*ent.Notification, error) {
	return n.client.Notification.Query().
		Where(notification.HasUserWith(user.ID(userID))).
		Limit(limit).
		Offset(offset).
		All(ctx)
}

func (n *NotificationService) GetUnreadNotifications(ctx context.Context, userID string, limit int, offset int) ([]*ent.Notification, error) {
	return n.client.Notification.Query().
		Where(notification.HasUserWith(user.ID(userID)), notification.IsRead(false)).
		Limit(limit).
		Offset(offset).
		All(ctx)
}

func (n *NotificationService) GetUnreadNotificationsCount(ctx context.Context, userID string) (int, error) {
	return n.client.Notification.Query().
		Where(notification.HasUserWith(user.ID(userID)), notification.IsRead(false)).
		Count(ctx)
}

func (n *NotificationService) GetUnreadNotificationsCountByType(ctx context.Context, userID string, notificationType int) (int, error) {
	return n.client.Notification.Query().
		Where(
			notification.HasUserWith(user.ID(userID)),
			notification.IsRead(false),
			notification.TypeEQ(notificationType),
		).
		Count(ctx)
}

func (n *NotificationService) DeleteNotification(ctx context.Context, id string) error {
	return n.client.Notification.DeleteOneID(id).Exec(ctx)
}
