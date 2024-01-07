// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// AccountSettings is the client for interacting with the AccountSettings builders.
	AccountSettings *AccountSettingsClient
	// AccountWallet is the client for interacting with the AccountWallet builders.
	AccountWallet *AccountWalletClient
	// Amenity is the client for interacting with the Amenity builders.
	Amenity *AmenityClient
	// Booking is the client for interacting with the Booking builders.
	Booking *BookingClient
	// Business is the client for interacting with the Business builders.
	Business *BusinessClient
	// BusinessFollowBusiness is the client for interacting with the BusinessFollowBusiness builders.
	BusinessFollowBusiness *BusinessFollowBusinessClient
	// BusinessFollowEvent is the client for interacting with the BusinessFollowEvent builders.
	BusinessFollowEvent *BusinessFollowEventClient
	// BusinessFollowUser is the client for interacting with the BusinessFollowUser builders.
	BusinessFollowUser *BusinessFollowUserClient
	// Category is the client for interacting with the Category builders.
	Category *CategoryClient
	// CategoryAssignment is the client for interacting with the CategoryAssignment builders.
	CategoryAssignment *CategoryAssignmentClient
	// Chat is the client for interacting with the Chat builders.
	Chat *ChatClient
	// Comment is the client for interacting with the Comment builders.
	Comment *CommentClient
	// CustomBlock is the client for interacting with the CustomBlock builders.
	CustomBlock *CustomBlockClient
	// Event is the client for interacting with the Event builders.
	Event *EventClient
	// FAQ is the client for interacting with the FAQ builders.
	FAQ *FAQClient
	// FeatureRelease is the client for interacting with the FeatureRelease builders.
	FeatureRelease *FeatureReleaseClient
	// Fitness is the client for interacting with the Fitness builders.
	Fitness *FitnessClient
	// Help is the client for interacting with the Help builders.
	Help *HelpClient
	// InventoryAttribute is the client for interacting with the InventoryAttribute builders.
	InventoryAttribute *InventoryAttributeClient
	// InventoryType is the client for interacting with the InventoryType builders.
	InventoryType *InventoryTypeClient
	// Like is the client for interacting with the Like builders.
	Like *LikeClient
	// Media is the client for interacting with the Media builders.
	Media *MediaClient
	// Menu is the client for interacting with the Menu builders.
	Menu *MenuClient
	// MenuItem is the client for interacting with the MenuItem builders.
	MenuItem *MenuItemClient
	// Notification is the client for interacting with the Notification builders.
	Notification *NotificationClient
	// Order is the client for interacting with the Order builders.
	Order *OrderClient
	// OrderItem is the client for interacting with the OrderItem builders.
	OrderItem *OrderItemClient
	// Payment is the client for interacting with the Payment builders.
	Payment *PaymentClient
	// Permission is the client for interacting with the Permission builders.
	Permission *PermissionClient
	// Place is the client for interacting with the Place builders.
	Place *PlaceClient
	// PlaceInventory is the client for interacting with the PlaceInventory builders.
	PlaceInventory *PlaceInventoryClient
	// PlaceInventoryAttribute is the client for interacting with the PlaceInventoryAttribute builders.
	PlaceInventoryAttribute *PlaceInventoryAttributeClient
	// PlaceTable is the client for interacting with the PlaceTable builders.
	PlaceTable *PlaceTableClient
	// Plan is the client for interacting with the Plan builders.
	Plan *PlanClient
	// Post is the client for interacting with the Post builders.
	Post *PostClient
	// Price is the client for interacting with the Price builders.
	Price *PriceClient
	// Rating is the client for interacting with the Rating builders.
	Rating *RatingClient
	// Reaction is the client for interacting with the Reaction builders.
	Reaction *ReactionClient
	// Reservation is the client for interacting with the Reservation builders.
	Reservation *ReservationClient
	// ReservationBlock is the client for interacting with the ReservationBlock builders.
	ReservationBlock *ReservationBlockClient
	// Resourse is the client for interacting with the Resourse builders.
	Resourse *ResourseClient
	// Review is the client for interacting with the Review builders.
	Review *ReviewClient
	// Room is the client for interacting with the Room builders.
	Room *RoomClient
	// RoomCategory is the client for interacting with the RoomCategory builders.
	RoomCategory *RoomCategoryClient
	// Staff is the client for interacting with the Staff builders.
	Staff *StaffClient
	// Subscription is the client for interacting with the Subscription builders.
	Subscription *SubscriptionClient
	// Template is the client for interacting with the Template builders.
	Template *TemplateClient
	// Ticket is the client for interacting with the Ticket builders.
	Ticket *TicketClient
	// TicketOption is the client for interacting with the TicketOption builders.
	TicketOption *TicketOptionClient
	// TransactionHistory is the client for interacting with the TransactionHistory builders.
	TransactionHistory *TransactionHistoryClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserBusiness is the client for interacting with the UserBusiness builders.
	UserBusiness *UserBusinessClient
	// UserFollowBusiness is the client for interacting with the UserFollowBusiness builders.
	UserFollowBusiness *UserFollowBusinessClient
	// UserFollowEvent is the client for interacting with the UserFollowEvent builders.
	UserFollowEvent *UserFollowEventClient
	// UserFollowPlace is the client for interacting with the UserFollowPlace builders.
	UserFollowPlace *UserFollowPlaceClient
	// UserFollowUser is the client for interacting with the UserFollowUser builders.
	UserFollowUser *UserFollowUserClient
	// UserLikePlace is the client for interacting with the UserLikePlace builders.
	UserLikePlace *UserLikePlaceClient
	// Website is the client for interacting with the Website builders.
	Website *WebsiteClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once
	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Commit method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	txDriver.mu.Lock()
	hooks := append([]CommitHook(nil), txDriver.onCommit...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onCommit = append(txDriver.onCommit, f)
	txDriver.mu.Unlock()
}

type (
	// Rollbacker is the interface that wraps the Rollback method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	txDriver.mu.Lock()
	hooks := append([]RollbackHook(nil), txDriver.onRollback...)
	txDriver.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	txDriver := tx.config.driver.(*txDriver)
	txDriver.mu.Lock()
	txDriver.onRollback = append(txDriver.onRollback, f)
	txDriver.mu.Unlock()
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.AccountSettings = NewAccountSettingsClient(tx.config)
	tx.AccountWallet = NewAccountWalletClient(tx.config)
	tx.Amenity = NewAmenityClient(tx.config)
	tx.Booking = NewBookingClient(tx.config)
	tx.Business = NewBusinessClient(tx.config)
	tx.BusinessFollowBusiness = NewBusinessFollowBusinessClient(tx.config)
	tx.BusinessFollowEvent = NewBusinessFollowEventClient(tx.config)
	tx.BusinessFollowUser = NewBusinessFollowUserClient(tx.config)
	tx.Category = NewCategoryClient(tx.config)
	tx.CategoryAssignment = NewCategoryAssignmentClient(tx.config)
	tx.Chat = NewChatClient(tx.config)
	tx.Comment = NewCommentClient(tx.config)
	tx.CustomBlock = NewCustomBlockClient(tx.config)
	tx.Event = NewEventClient(tx.config)
	tx.FAQ = NewFAQClient(tx.config)
	tx.FeatureRelease = NewFeatureReleaseClient(tx.config)
	tx.Fitness = NewFitnessClient(tx.config)
	tx.Help = NewHelpClient(tx.config)
	tx.InventoryAttribute = NewInventoryAttributeClient(tx.config)
	tx.InventoryType = NewInventoryTypeClient(tx.config)
	tx.Like = NewLikeClient(tx.config)
	tx.Media = NewMediaClient(tx.config)
	tx.Menu = NewMenuClient(tx.config)
	tx.MenuItem = NewMenuItemClient(tx.config)
	tx.Notification = NewNotificationClient(tx.config)
	tx.Order = NewOrderClient(tx.config)
	tx.OrderItem = NewOrderItemClient(tx.config)
	tx.Payment = NewPaymentClient(tx.config)
	tx.Permission = NewPermissionClient(tx.config)
	tx.Place = NewPlaceClient(tx.config)
	tx.PlaceInventory = NewPlaceInventoryClient(tx.config)
	tx.PlaceInventoryAttribute = NewPlaceInventoryAttributeClient(tx.config)
	tx.PlaceTable = NewPlaceTableClient(tx.config)
	tx.Plan = NewPlanClient(tx.config)
	tx.Post = NewPostClient(tx.config)
	tx.Price = NewPriceClient(tx.config)
	tx.Rating = NewRatingClient(tx.config)
	tx.Reaction = NewReactionClient(tx.config)
	tx.Reservation = NewReservationClient(tx.config)
	tx.ReservationBlock = NewReservationBlockClient(tx.config)
	tx.Resourse = NewResourseClient(tx.config)
	tx.Review = NewReviewClient(tx.config)
	tx.Room = NewRoomClient(tx.config)
	tx.RoomCategory = NewRoomCategoryClient(tx.config)
	tx.Staff = NewStaffClient(tx.config)
	tx.Subscription = NewSubscriptionClient(tx.config)
	tx.Template = NewTemplateClient(tx.config)
	tx.Ticket = NewTicketClient(tx.config)
	tx.TicketOption = NewTicketOptionClient(tx.config)
	tx.TransactionHistory = NewTransactionHistoryClient(tx.config)
	tx.User = NewUserClient(tx.config)
	tx.UserBusiness = NewUserBusinessClient(tx.config)
	tx.UserFollowBusiness = NewUserFollowBusinessClient(tx.config)
	tx.UserFollowEvent = NewUserFollowEventClient(tx.config)
	tx.UserFollowPlace = NewUserFollowPlaceClient(tx.config)
	tx.UserFollowUser = NewUserFollowUserClient(tx.config)
	tx.UserLikePlace = NewUserLikePlaceClient(tx.config)
	tx.Website = NewWebsiteClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: AccountSettings.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
	// completion hooks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v any) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v any) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
