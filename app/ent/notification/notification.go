// Code generated by ent, DO NOT EDIT.

package notification

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the notification type in the database.
	Label = "notification"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldIsRead holds the string denoting the is_read field in the database.
	FieldIsRead = "is_read"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldUnreadCount holds the string denoting the unread_count field in the database.
	FieldUnreadCount = "unread_count"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldNotifiableType holds the string denoting the notifiable_type field in the database.
	FieldNotifiableType = "notifiable_type"
	// FieldNotifiableID holds the string denoting the notifiable_id field in the database.
	FieldNotifiableID = "notifiable_id"
	// FieldTriggeredBy holds the string denoting the triggered_by field in the database.
	FieldTriggeredBy = "triggered_by"
	// FieldTriggeredTo holds the string denoting the triggered_to field in the database.
	FieldTriggeredTo = "triggered_to"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeBusinessAccount holds the string denoting the business_account edge name in mutations.
	EdgeBusinessAccount = "business_account"
	// EdgePlace holds the string denoting the place edge name in mutations.
	EdgePlace = "place"
	// EdgePost holds the string denoting the post edge name in mutations.
	EdgePost = "post"
	// EdgeComment holds the string denoting the comment edge name in mutations.
	EdgeComment = "comment"
	// Table holds the table name of the notification in the database.
	Table = "notifications"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "user_notifications"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// BusinessAccountTable is the table that holds the business_account relation/edge. The primary key declared below.
	BusinessAccountTable = "business_notifications"
	// BusinessAccountInverseTable is the table name for the Business entity.
	// It exists in this package in order to avoid circular dependency with the "business" package.
	BusinessAccountInverseTable = "businesses"
	// PlaceTable is the table that holds the place relation/edge. The primary key declared below.
	PlaceTable = "place_notifications"
	// PlaceInverseTable is the table name for the Place entity.
	// It exists in this package in order to avoid circular dependency with the "place" package.
	PlaceInverseTable = "places"
	// PostTable is the table that holds the post relation/edge. The primary key declared below.
	PostTable = "post_notifications"
	// PostInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostInverseTable = "posts"
	// CommentTable is the table that holds the comment relation/edge. The primary key declared below.
	CommentTable = "comment_notifications"
	// CommentInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentInverseTable = "comments"
)

// Columns holds all SQL columns for notification fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldMessage,
	FieldLink,
	FieldIsRead,
	FieldType,
	FieldUnreadCount,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldNotifiableType,
	FieldNotifiableID,
	FieldTriggeredBy,
	FieldTriggeredTo,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "notification_id"}
	// BusinessAccountPrimaryKey and BusinessAccountColumn2 are the table columns denoting the
	// primary key for the business_account relation (M2M).
	BusinessAccountPrimaryKey = []string{"business_id", "notification_id"}
	// PlacePrimaryKey and PlaceColumn2 are the table columns denoting the
	// primary key for the place relation (M2M).
	PlacePrimaryKey = []string{"place_id", "notification_id"}
	// PostPrimaryKey and PostColumn2 are the table columns denoting the
	// primary key for the post relation (M2M).
	PostPrimaryKey = []string{"post_id", "notification_id"}
	// CommentPrimaryKey and CommentColumn2 are the table columns denoting the
	// primary key for the comment relation (M2M).
	CommentPrimaryKey = []string{"comment_id", "notification_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// MessageValidator is a validator for the "message" field. It is called by the builders before save.
	MessageValidator func(string) error
	// LinkValidator is a validator for the "link" field. It is called by the builders before save.
	LinkValidator func(string) error
	// DefaultIsRead holds the default value on creation for the "is_read" field.
	DefaultIsRead bool
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType int
	// DefaultUnreadCount holds the default value on creation for the "unread_count" field.
	DefaultUnreadCount int
	// NotifiableTypeValidator is a validator for the "notifiable_type" field. It is called by the builders before save.
	NotifiableTypeValidator func(string) error
	// NotifiableIDValidator is a validator for the "notifiable_id" field. It is called by the builders before save.
	NotifiableIDValidator func(string) error
	// TriggeredByValidator is a validator for the "triggered_by" field. It is called by the builders before save.
	TriggeredByValidator func(string) error
	// TriggeredToValidator is a validator for the "triggered_to" field. It is called by the builders before save.
	TriggeredToValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Notification queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByMessage orders the results by the message field.
func ByMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessage, opts...).ToFunc()
}

// ByLink orders the results by the link field.
func ByLink(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLink, opts...).ToFunc()
}

// ByIsRead orders the results by the is_read field.
func ByIsRead(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsRead, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByUnreadCount orders the results by the unread_count field.
func ByUnreadCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnreadCount, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByNotifiableType orders the results by the notifiable_type field.
func ByNotifiableType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotifiableType, opts...).ToFunc()
}

// ByNotifiableID orders the results by the notifiable_id field.
func ByNotifiableID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotifiableID, opts...).ToFunc()
}

// ByTriggeredBy orders the results by the triggered_by field.
func ByTriggeredBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTriggeredBy, opts...).ToFunc()
}

// ByTriggeredTo orders the results by the triggered_to field.
func ByTriggeredTo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTriggeredTo, opts...).ToFunc()
}

// ByUserCount orders the results by user count.
func ByUserCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserStep(), opts...)
	}
}

// ByUser orders the results by user terms.
func ByUser(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBusinessAccountCount orders the results by business_account count.
func ByBusinessAccountCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBusinessAccountStep(), opts...)
	}
}

// ByBusinessAccount orders the results by business_account terms.
func ByBusinessAccount(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessAccountStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPlaceCount orders the results by place count.
func ByPlaceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlaceStep(), opts...)
	}
}

// ByPlace orders the results by place terms.
func ByPlace(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlaceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPostCount orders the results by post count.
func ByPostCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostStep(), opts...)
	}
}

// ByPost orders the results by post terms.
func ByPost(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCommentCount orders the results by comment count.
func ByCommentCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommentStep(), opts...)
	}
}

// ByComment orders the results by comment terms.
func ByComment(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
	)
}
func newBusinessAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessAccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, BusinessAccountTable, BusinessAccountPrimaryKey...),
	)
}
func newPlaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PlaceTable, PlacePrimaryKey...),
	)
}
func newPostStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PostTable, PostPrimaryKey...),
	)
}
func newCommentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, CommentTable, CommentPrimaryKey...),
	)
}
