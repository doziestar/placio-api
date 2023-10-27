// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/post"
	"placio-app/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Post is the model entity for the Post schema.
type Post struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Content holds the value of the "Content" field.
	Content string `json:"Content,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// Privacy holds the value of the "Privacy" field.
	Privacy post.Privacy `json:"Privacy,omitempty"`
	// LikedByMe holds the value of the "LikedByMe" field.
	LikedByMe bool `json:"LikedByMe,omitempty"`
	// LikeCount holds the value of the "LikeCount" field.
	LikeCount int `json:"LikeCount,omitempty"`
	// CommentCount holds the value of the "CommentCount" field.
	CommentCount int `json:"CommentCount,omitempty"`
	// ShareCount holds the value of the "ShareCount" field.
	ShareCount int `json:"ShareCount,omitempty"`
	// ViewCount holds the value of the "ViewCount" field.
	ViewCount int `json:"ViewCount,omitempty"`
	// IsSponsored holds the value of the "IsSponsored" field.
	IsSponsored bool `json:"IsSponsored,omitempty"`
	// IsPromoted holds the value of the "IsPromoted" field.
	IsPromoted bool `json:"IsPromoted,omitempty"`
	// IsBoosted holds the value of the "IsBoosted" field.
	IsBoosted bool `json:"IsBoosted,omitempty"`
	// IsPinned holds the value of the "IsPinned" field.
	IsPinned bool `json:"IsPinned,omitempty"`
	// IsHidden holds the value of the "IsHidden" field.
	IsHidden bool `json:"IsHidden,omitempty"`
	// RepostCount holds the value of the "RepostCount" field.
	RepostCount int `json:"RepostCount,omitempty"`
	// IsRepost holds the value of the "IsRepost" field.
	IsRepost bool `json:"IsRepost,omitempty"`
	// RelevanceScore holds the value of the "RelevanceScore" field.
	RelevanceScore int `json:"RelevanceScore,omitempty"`
	// SearchText holds the value of the "SearchText" field.
	SearchText string `json:"SearchText,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PostQuery when eager-loading is set.
	Edges              PostEdges `json:"edges"`
	business_posts     *string
	post_original_post *string
	user_posts         *string
	selectValues       sql.SelectValues
}

// PostEdges holds the relations/edges for other nodes in the graph.
type PostEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// BusinessAccount holds the value of the business_account edge.
	BusinessAccount *Business `json:"business_account,omitempty"`
	// Medias holds the value of the medias edge.
	Medias []*Media `json:"medias,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// Likes holds the value of the likes edge.
	Likes []*Like `json:"likes,omitempty"`
	// Categories holds the value of the categories edge.
	Categories []*Category `json:"categories,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// Reposts holds the value of the reposts edge.
	Reposts *Post `json:"reposts,omitempty"`
	// OriginalPost holds the value of the original_post edge.
	OriginalPost []*Post `json:"original_post,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// BusinessAccountOrErr returns the BusinessAccount value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) BusinessAccountOrErr() (*Business, error) {
	if e.loadedTypes[1] {
		if e.BusinessAccount == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: business.Label}
		}
		return e.BusinessAccount, nil
	}
	return nil, &NotLoadedError{edge: "business_account"}
}

// MediasOrErr returns the Medias value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) MediasOrErr() ([]*Media, error) {
	if e.loadedTypes[2] {
		return e.Medias, nil
	}
	return nil, &NotLoadedError{edge: "medias"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[3] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// LikesOrErr returns the Likes value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) LikesOrErr() ([]*Like, error) {
	if e.loadedTypes[4] {
		return e.Likes, nil
	}
	return nil, &NotLoadedError{edge: "likes"}
}

// CategoriesOrErr returns the Categories value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) CategoriesOrErr() ([]*Category, error) {
	if e.loadedTypes[5] {
		return e.Categories, nil
	}
	return nil, &NotLoadedError{edge: "categories"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[6] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// RepostsOrErr returns the Reposts value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) RepostsOrErr() (*Post, error) {
	if e.loadedTypes[7] {
		if e.Reposts == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: post.Label}
		}
		return e.Reposts, nil
	}
	return nil, &NotLoadedError{edge: "reposts"}
}

// OriginalPostOrErr returns the OriginalPost value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) OriginalPostOrErr() ([]*Post, error) {
	if e.loadedTypes[8] {
		return e.OriginalPost, nil
	}
	return nil, &NotLoadedError{edge: "original_post"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Post) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case post.FieldLikedByMe, post.FieldIsSponsored, post.FieldIsPromoted, post.FieldIsBoosted, post.FieldIsPinned, post.FieldIsHidden, post.FieldIsRepost:
			values[i] = new(sql.NullBool)
		case post.FieldLikeCount, post.FieldCommentCount, post.FieldShareCount, post.FieldViewCount, post.FieldRepostCount, post.FieldRelevanceScore:
			values[i] = new(sql.NullInt64)
		case post.FieldID, post.FieldContent, post.FieldPrivacy, post.FieldSearchText:
			values[i] = new(sql.NullString)
		case post.FieldCreatedAt, post.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case post.ForeignKeys[0]: // business_posts
			values[i] = new(sql.NullString)
		case post.ForeignKeys[1]: // post_original_post
			values[i] = new(sql.NullString)
		case post.ForeignKeys[2]: // user_posts
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Post fields.
func (po *Post) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				po.ID = value.String
			}
		case post.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Content", values[i])
			} else if value.Valid {
				po.Content = value.String
			}
		case post.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case post.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		case post.FieldPrivacy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Privacy", values[i])
			} else if value.Valid {
				po.Privacy = post.Privacy(value.String)
			}
		case post.FieldLikedByMe:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field LikedByMe", values[i])
			} else if value.Valid {
				po.LikedByMe = value.Bool
			}
		case post.FieldLikeCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field LikeCount", values[i])
			} else if value.Valid {
				po.LikeCount = int(value.Int64)
			}
		case post.FieldCommentCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CommentCount", values[i])
			} else if value.Valid {
				po.CommentCount = int(value.Int64)
			}
		case post.FieldShareCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ShareCount", values[i])
			} else if value.Valid {
				po.ShareCount = int(value.Int64)
			}
		case post.FieldViewCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ViewCount", values[i])
			} else if value.Valid {
				po.ViewCount = int(value.Int64)
			}
		case post.FieldIsSponsored:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field IsSponsored", values[i])
			} else if value.Valid {
				po.IsSponsored = value.Bool
			}
		case post.FieldIsPromoted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field IsPromoted", values[i])
			} else if value.Valid {
				po.IsPromoted = value.Bool
			}
		case post.FieldIsBoosted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field IsBoosted", values[i])
			} else if value.Valid {
				po.IsBoosted = value.Bool
			}
		case post.FieldIsPinned:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field IsPinned", values[i])
			} else if value.Valid {
				po.IsPinned = value.Bool
			}
		case post.FieldIsHidden:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field IsHidden", values[i])
			} else if value.Valid {
				po.IsHidden = value.Bool
			}
		case post.FieldRepostCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field RepostCount", values[i])
			} else if value.Valid {
				po.RepostCount = int(value.Int64)
			}
		case post.FieldIsRepost:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field IsRepost", values[i])
			} else if value.Valid {
				po.IsRepost = value.Bool
			}
		case post.FieldRelevanceScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field RelevanceScore", values[i])
			} else if value.Valid {
				po.RelevanceScore = int(value.Int64)
			}
		case post.FieldSearchText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field SearchText", values[i])
			} else if value.Valid {
				po.SearchText = value.String
			}
		case post.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_posts", values[i])
			} else if value.Valid {
				po.business_posts = new(string)
				*po.business_posts = value.String
			}
		case post.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field post_original_post", values[i])
			} else if value.Valid {
				po.post_original_post = new(string)
				*po.post_original_post = value.String
			}
		case post.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_posts", values[i])
			} else if value.Valid {
				po.user_posts = new(string)
				*po.user_posts = value.String
			}
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Post.
// This includes values selected through modifiers, order, etc.
func (po *Post) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Post entity.
func (po *Post) QueryUser() *UserQuery {
	return NewPostClient(po.config).QueryUser(po)
}

// QueryBusinessAccount queries the "business_account" edge of the Post entity.
func (po *Post) QueryBusinessAccount() *BusinessQuery {
	return NewPostClient(po.config).QueryBusinessAccount(po)
}

// QueryMedias queries the "medias" edge of the Post entity.
func (po *Post) QueryMedias() *MediaQuery {
	return NewPostClient(po.config).QueryMedias(po)
}

// QueryComments queries the "comments" edge of the Post entity.
func (po *Post) QueryComments() *CommentQuery {
	return NewPostClient(po.config).QueryComments(po)
}

// QueryLikes queries the "likes" edge of the Post entity.
func (po *Post) QueryLikes() *LikeQuery {
	return NewPostClient(po.config).QueryLikes(po)
}

// QueryCategories queries the "categories" edge of the Post entity.
func (po *Post) QueryCategories() *CategoryQuery {
	return NewPostClient(po.config).QueryCategories(po)
}

// QueryNotifications queries the "notifications" edge of the Post entity.
func (po *Post) QueryNotifications() *NotificationQuery {
	return NewPostClient(po.config).QueryNotifications(po)
}

// QueryReposts queries the "reposts" edge of the Post entity.
func (po *Post) QueryReposts() *PostQuery {
	return NewPostClient(po.config).QueryReposts(po)
}

// QueryOriginalPost queries the "original_post" edge of the Post entity.
func (po *Post) QueryOriginalPost() *PostQuery {
	return NewPostClient(po.config).QueryOriginalPost(po)
}

// Update returns a builder for updating this Post.
// Note that you need to call Post.Unwrap() before calling this method if this Post
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Post) Update() *PostUpdateOne {
	return NewPostClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Post entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Post) Unwrap() *Post {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Post is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Post) String() string {
	var builder strings.Builder
	builder.WriteString("Post(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("Content=")
	builder.WriteString(po.Content)
	builder.WriteString(", ")
	builder.WriteString("CreatedAt=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("Privacy=")
	builder.WriteString(fmt.Sprintf("%v", po.Privacy))
	builder.WriteString(", ")
	builder.WriteString("LikedByMe=")
	builder.WriteString(fmt.Sprintf("%v", po.LikedByMe))
	builder.WriteString(", ")
	builder.WriteString("LikeCount=")
	builder.WriteString(fmt.Sprintf("%v", po.LikeCount))
	builder.WriteString(", ")
	builder.WriteString("CommentCount=")
	builder.WriteString(fmt.Sprintf("%v", po.CommentCount))
	builder.WriteString(", ")
	builder.WriteString("ShareCount=")
	builder.WriteString(fmt.Sprintf("%v", po.ShareCount))
	builder.WriteString(", ")
	builder.WriteString("ViewCount=")
	builder.WriteString(fmt.Sprintf("%v", po.ViewCount))
	builder.WriteString(", ")
	builder.WriteString("IsSponsored=")
	builder.WriteString(fmt.Sprintf("%v", po.IsSponsored))
	builder.WriteString(", ")
	builder.WriteString("IsPromoted=")
	builder.WriteString(fmt.Sprintf("%v", po.IsPromoted))
	builder.WriteString(", ")
	builder.WriteString("IsBoosted=")
	builder.WriteString(fmt.Sprintf("%v", po.IsBoosted))
	builder.WriteString(", ")
	builder.WriteString("IsPinned=")
	builder.WriteString(fmt.Sprintf("%v", po.IsPinned))
	builder.WriteString(", ")
	builder.WriteString("IsHidden=")
	builder.WriteString(fmt.Sprintf("%v", po.IsHidden))
	builder.WriteString(", ")
	builder.WriteString("RepostCount=")
	builder.WriteString(fmt.Sprintf("%v", po.RepostCount))
	builder.WriteString(", ")
	builder.WriteString("IsRepost=")
	builder.WriteString(fmt.Sprintf("%v", po.IsRepost))
	builder.WriteString(", ")
	builder.WriteString("RelevanceScore=")
	builder.WriteString(fmt.Sprintf("%v", po.RelevanceScore))
	builder.WriteString(", ")
	builder.WriteString("SearchText=")
	builder.WriteString(po.SearchText)
	builder.WriteByte(')')
	return builder.String()
}

// Posts is a parsable slice of Post.
type Posts []*Post
