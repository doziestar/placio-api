// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/amenity"
	"placio-app/ent/booking"
	"placio-app/ent/business"
	"placio-app/ent/category"
	"placio-app/ent/categoryassignment"
	"placio-app/ent/event"
	"placio-app/ent/faq"
	"placio-app/ent/media"
	"placio-app/ent/menu"
	"placio-app/ent/place"
	"placio-app/ent/rating"
	"placio-app/ent/reservation"
	"placio-app/ent/review"
	"placio-app/ent/room"
	"placio-app/ent/user"
	"placio-app/ent/userfollowplace"
	"placio-app/ent/userlikeplace"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlaceCreate is the builder for creating a Place entity.
type PlaceCreate struct {
	config
	mutation *PlaceMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PlaceCreate) SetName(s string) *PlaceCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetType sets the "type" field.
func (pc *PlaceCreate) SetType(s string) *PlaceCreate {
	pc.mutation.SetType(s)
	return pc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableType(s *string) *PlaceCreate {
	if s != nil {
		pc.SetType(*s)
	}
	return pc
}

// SetDescription sets the "description" field.
func (pc *PlaceCreate) SetDescription(s string) *PlaceCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableDescription(s *string) *PlaceCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetLocation sets the "location" field.
func (pc *PlaceCreate) SetLocation(s string) *PlaceCreate {
	pc.mutation.SetLocation(s)
	return pc
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableLocation(s *string) *PlaceCreate {
	if s != nil {
		pc.SetLocation(*s)
	}
	return pc
}

// SetEmail sets the "email" field.
func (pc *PlaceCreate) SetEmail(s string) *PlaceCreate {
	pc.mutation.SetEmail(s)
	return pc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableEmail(s *string) *PlaceCreate {
	if s != nil {
		pc.SetEmail(*s)
	}
	return pc
}

// SetPhone sets the "phone" field.
func (pc *PlaceCreate) SetPhone(s string) *PlaceCreate {
	pc.mutation.SetPhone(s)
	return pc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (pc *PlaceCreate) SetNillablePhone(s *string) *PlaceCreate {
	if s != nil {
		pc.SetPhone(*s)
	}
	return pc
}

// SetWebsite sets the "website" field.
func (pc *PlaceCreate) SetWebsite(s string) *PlaceCreate {
	pc.mutation.SetWebsite(s)
	return pc
}

// SetNillableWebsite sets the "website" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableWebsite(s *string) *PlaceCreate {
	if s != nil {
		pc.SetWebsite(*s)
	}
	return pc
}

// SetCoverImage sets the "cover_image" field.
func (pc *PlaceCreate) SetCoverImage(s string) *PlaceCreate {
	pc.mutation.SetCoverImage(s)
	return pc
}

// SetNillableCoverImage sets the "cover_image" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableCoverImage(s *string) *PlaceCreate {
	if s != nil {
		pc.SetCoverImage(*s)
	}
	return pc
}

// SetPicture sets the "picture" field.
func (pc *PlaceCreate) SetPicture(s string) *PlaceCreate {
	pc.mutation.SetPicture(s)
	return pc
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (pc *PlaceCreate) SetNillablePicture(s *string) *PlaceCreate {
	if s != nil {
		pc.SetPicture(*s)
	}
	return pc
}

// SetCountry sets the "country" field.
func (pc *PlaceCreate) SetCountry(s string) *PlaceCreate {
	pc.mutation.SetCountry(s)
	return pc
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableCountry(s *string) *PlaceCreate {
	if s != nil {
		pc.SetCountry(*s)
	}
	return pc
}

// SetCity sets the "city" field.
func (pc *PlaceCreate) SetCity(s string) *PlaceCreate {
	pc.mutation.SetCity(s)
	return pc
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableCity(s *string) *PlaceCreate {
	if s != nil {
		pc.SetCity(*s)
	}
	return pc
}

// SetState sets the "state" field.
func (pc *PlaceCreate) SetState(s string) *PlaceCreate {
	pc.mutation.SetState(s)
	return pc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableState(s *string) *PlaceCreate {
	if s != nil {
		pc.SetState(*s)
	}
	return pc
}

// SetPlaceSettings sets the "place_settings" field.
func (pc *PlaceCreate) SetPlaceSettings(m map[string]interface{}) *PlaceCreate {
	pc.mutation.SetPlaceSettings(m)
	return pc
}

// SetOpeningHours sets the "opening_hours" field.
func (pc *PlaceCreate) SetOpeningHours(m map[string]interface{}) *PlaceCreate {
	pc.mutation.SetOpeningHours(m)
	return pc
}

// SetSocialMedia sets the "social_media" field.
func (pc *PlaceCreate) SetSocialMedia(m map[string]interface{}) *PlaceCreate {
	pc.mutation.SetSocialMedia(m)
	return pc
}

// SetPaymentOptions sets the "payment_options" field.
func (pc *PlaceCreate) SetPaymentOptions(m map[string]interface{}) *PlaceCreate {
	pc.mutation.SetPaymentOptions(m)
	return pc
}

// SetTags sets the "tags" field.
func (pc *PlaceCreate) SetTags(s []string) *PlaceCreate {
	pc.mutation.SetTags(s)
	return pc
}

// SetFeatures sets the "features" field.
func (pc *PlaceCreate) SetFeatures(s []string) *PlaceCreate {
	pc.mutation.SetFeatures(s)
	return pc
}

// SetAdditionalInfo sets the "additional_info" field.
func (pc *PlaceCreate) SetAdditionalInfo(m map[string]interface{}) *PlaceCreate {
	pc.mutation.SetAdditionalInfo(m)
	return pc
}

// SetImages sets the "images" field.
func (pc *PlaceCreate) SetImages(s []string) *PlaceCreate {
	pc.mutation.SetImages(s)
	return pc
}

// SetAvailability sets the "availability" field.
func (pc *PlaceCreate) SetAvailability(m map[string]interface{}) *PlaceCreate {
	pc.mutation.SetAvailability(m)
	return pc
}

// SetSpecialOffers sets the "special_offers" field.
func (pc *PlaceCreate) SetSpecialOffers(s string) *PlaceCreate {
	pc.mutation.SetSpecialOffers(s)
	return pc
}

// SetNillableSpecialOffers sets the "special_offers" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableSpecialOffers(s *string) *PlaceCreate {
	if s != nil {
		pc.SetSpecialOffers(*s)
	}
	return pc
}

// SetSustainabilityScore sets the "sustainability_score" field.
func (pc *PlaceCreate) SetSustainabilityScore(f float64) *PlaceCreate {
	pc.mutation.SetSustainabilityScore(f)
	return pc
}

// SetNillableSustainabilityScore sets the "sustainability_score" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableSustainabilityScore(f *float64) *PlaceCreate {
	if f != nil {
		pc.SetSustainabilityScore(*f)
	}
	return pc
}

// SetMapCoordinates sets the "map_coordinates" field.
func (pc *PlaceCreate) SetMapCoordinates(m map[string]interface{}) *PlaceCreate {
	pc.mutation.SetMapCoordinates(m)
	return pc
}

// SetLongitude sets the "longitude" field.
func (pc *PlaceCreate) SetLongitude(s string) *PlaceCreate {
	pc.mutation.SetLongitude(s)
	return pc
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableLongitude(s *string) *PlaceCreate {
	if s != nil {
		pc.SetLongitude(*s)
	}
	return pc
}

// SetLatitude sets the "latitude" field.
func (pc *PlaceCreate) SetLatitude(s string) *PlaceCreate {
	pc.mutation.SetLatitude(s)
	return pc
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableLatitude(s *string) *PlaceCreate {
	if s != nil {
		pc.SetLatitude(*s)
	}
	return pc
}

// SetSearchText sets the "search_text" field.
func (pc *PlaceCreate) SetSearchText(s string) *PlaceCreate {
	pc.mutation.SetSearchText(s)
	return pc
}

// SetNillableSearchText sets the "search_text" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableSearchText(s *string) *PlaceCreate {
	if s != nil {
		pc.SetSearchText(*s)
	}
	return pc
}

// SetRelevanceScore sets the "relevance_score" field.
func (pc *PlaceCreate) SetRelevanceScore(f float64) *PlaceCreate {
	pc.mutation.SetRelevanceScore(f)
	return pc
}

// SetNillableRelevanceScore sets the "relevance_score" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableRelevanceScore(f *float64) *PlaceCreate {
	if f != nil {
		pc.SetRelevanceScore(*f)
	}
	return pc
}

// SetFollowerCount sets the "follower_count" field.
func (pc *PlaceCreate) SetFollowerCount(i int) *PlaceCreate {
	pc.mutation.SetFollowerCount(i)
	return pc
}

// SetNillableFollowerCount sets the "follower_count" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableFollowerCount(i *int) *PlaceCreate {
	if i != nil {
		pc.SetFollowerCount(*i)
	}
	return pc
}

// SetFollowingCount sets the "following_count" field.
func (pc *PlaceCreate) SetFollowingCount(i int) *PlaceCreate {
	pc.mutation.SetFollowingCount(i)
	return pc
}

// SetNillableFollowingCount sets the "following_count" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableFollowingCount(i *int) *PlaceCreate {
	if i != nil {
		pc.SetFollowingCount(*i)
	}
	return pc
}

// SetIsPremium sets the "is_Premium" field.
func (pc *PlaceCreate) SetIsPremium(b bool) *PlaceCreate {
	pc.mutation.SetIsPremium(b)
	return pc
}

// SetNillableIsPremium sets the "is_Premium" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableIsPremium(b *bool) *PlaceCreate {
	if b != nil {
		pc.SetIsPremium(*b)
	}
	return pc
}

// SetIsPublished sets the "is_published" field.
func (pc *PlaceCreate) SetIsPublished(b bool) *PlaceCreate {
	pc.mutation.SetIsPublished(b)
	return pc
}

// SetNillableIsPublished sets the "is_published" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableIsPublished(b *bool) *PlaceCreate {
	if b != nil {
		pc.SetIsPublished(*b)
	}
	return pc
}

// SetLikedByCurrentUser sets the "likedByCurrentUser" field.
func (pc *PlaceCreate) SetLikedByCurrentUser(b bool) *PlaceCreate {
	pc.mutation.SetLikedByCurrentUser(b)
	return pc
}

// SetNillableLikedByCurrentUser sets the "likedByCurrentUser" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableLikedByCurrentUser(b *bool) *PlaceCreate {
	if b != nil {
		pc.SetLikedByCurrentUser(*b)
	}
	return pc
}

// SetFollowedByCurrentUser sets the "followedByCurrentUser" field.
func (pc *PlaceCreate) SetFollowedByCurrentUser(b bool) *PlaceCreate {
	pc.mutation.SetFollowedByCurrentUser(b)
	return pc
}

// SetNillableFollowedByCurrentUser sets the "followedByCurrentUser" field if the given value is not nil.
func (pc *PlaceCreate) SetNillableFollowedByCurrentUser(b *bool) *PlaceCreate {
	if b != nil {
		pc.SetFollowedByCurrentUser(*b)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PlaceCreate) SetID(s string) *PlaceCreate {
	pc.mutation.SetID(s)
	return pc
}

// SetBusinessID sets the "business" edge to the Business entity by ID.
func (pc *PlaceCreate) SetBusinessID(id string) *PlaceCreate {
	pc.mutation.SetBusinessID(id)
	return pc
}

// SetNillableBusinessID sets the "business" edge to the Business entity by ID if the given value is not nil.
func (pc *PlaceCreate) SetNillableBusinessID(id *string) *PlaceCreate {
	if id != nil {
		pc = pc.SetBusinessID(*id)
	}
	return pc
}

// SetBusiness sets the "business" edge to the Business entity.
func (pc *PlaceCreate) SetBusiness(b *Business) *PlaceCreate {
	return pc.SetBusinessID(b.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (pc *PlaceCreate) AddUserIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddUserIDs(ids...)
	return pc
}

// AddUsers adds the "users" edges to the User entity.
func (pc *PlaceCreate) AddUsers(u ...*User) *PlaceCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddUserIDs(ids...)
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (pc *PlaceCreate) AddReviewIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddReviewIDs(ids...)
	return pc
}

// AddReviews adds the "reviews" edges to the Review entity.
func (pc *PlaceCreate) AddReviews(r ...*Review) *PlaceCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddReviewIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (pc *PlaceCreate) AddEventIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddEventIDs(ids...)
	return pc
}

// AddEvents adds the "events" edges to the Event entity.
func (pc *PlaceCreate) AddEvents(e ...*Event) *PlaceCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pc.AddEventIDs(ids...)
}

// AddAmenityIDs adds the "amenities" edge to the Amenity entity by IDs.
func (pc *PlaceCreate) AddAmenityIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddAmenityIDs(ids...)
	return pc
}

// AddAmenities adds the "amenities" edges to the Amenity entity.
func (pc *PlaceCreate) AddAmenities(a ...*Amenity) *PlaceCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pc.AddAmenityIDs(ids...)
}

// AddMenuIDs adds the "menus" edge to the Menu entity by IDs.
func (pc *PlaceCreate) AddMenuIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddMenuIDs(ids...)
	return pc
}

// AddMenus adds the "menus" edges to the Menu entity.
func (pc *PlaceCreate) AddMenus(m ...*Menu) *PlaceCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pc.AddMenuIDs(ids...)
}

// AddMediaIDs adds the "medias" edge to the Media entity by IDs.
func (pc *PlaceCreate) AddMediaIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddMediaIDs(ids...)
	return pc
}

// AddMedias adds the "medias" edges to the Media entity.
func (pc *PlaceCreate) AddMedias(m ...*Media) *PlaceCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pc.AddMediaIDs(ids...)
}

// AddRoomIDs adds the "rooms" edge to the Room entity by IDs.
func (pc *PlaceCreate) AddRoomIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddRoomIDs(ids...)
	return pc
}

// AddRooms adds the "rooms" edges to the Room entity.
func (pc *PlaceCreate) AddRooms(r ...*Room) *PlaceCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddRoomIDs(ids...)
}

// AddReservationIDs adds the "reservations" edge to the Reservation entity by IDs.
func (pc *PlaceCreate) AddReservationIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddReservationIDs(ids...)
	return pc
}

// AddReservations adds the "reservations" edges to the Reservation entity.
func (pc *PlaceCreate) AddReservations(r ...*Reservation) *PlaceCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddReservationIDs(ids...)
}

// AddBookingIDs adds the "bookings" edge to the Booking entity by IDs.
func (pc *PlaceCreate) AddBookingIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddBookingIDs(ids...)
	return pc
}

// AddBookings adds the "bookings" edges to the Booking entity.
func (pc *PlaceCreate) AddBookings(b ...*Booking) *PlaceCreate {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pc.AddBookingIDs(ids...)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (pc *PlaceCreate) AddCategoryIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddCategoryIDs(ids...)
	return pc
}

// AddCategories adds the "categories" edges to the Category entity.
func (pc *PlaceCreate) AddCategories(c ...*Category) *PlaceCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCategoryIDs(ids...)
}

// AddCategoryAssignmentIDs adds the "categoryAssignments" edge to the CategoryAssignment entity by IDs.
func (pc *PlaceCreate) AddCategoryAssignmentIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddCategoryAssignmentIDs(ids...)
	return pc
}

// AddCategoryAssignments adds the "categoryAssignments" edges to the CategoryAssignment entity.
func (pc *PlaceCreate) AddCategoryAssignments(c ...*CategoryAssignment) *PlaceCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCategoryAssignmentIDs(ids...)
}

// AddFaqIDs adds the "faqs" edge to the FAQ entity by IDs.
func (pc *PlaceCreate) AddFaqIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddFaqIDs(ids...)
	return pc
}

// AddFaqs adds the "faqs" edges to the FAQ entity.
func (pc *PlaceCreate) AddFaqs(f ...*FAQ) *PlaceCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pc.AddFaqIDs(ids...)
}

// AddLikedByUserIDs adds the "likedByUsers" edge to the UserLikePlace entity by IDs.
func (pc *PlaceCreate) AddLikedByUserIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddLikedByUserIDs(ids...)
	return pc
}

// AddLikedByUsers adds the "likedByUsers" edges to the UserLikePlace entity.
func (pc *PlaceCreate) AddLikedByUsers(u ...*UserLikePlace) *PlaceCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddLikedByUserIDs(ids...)
}

// AddFollowerUserIDs adds the "followerUsers" edge to the UserFollowPlace entity by IDs.
func (pc *PlaceCreate) AddFollowerUserIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddFollowerUserIDs(ids...)
	return pc
}

// AddFollowerUsers adds the "followerUsers" edges to the UserFollowPlace entity.
func (pc *PlaceCreate) AddFollowerUsers(u ...*UserFollowPlace) *PlaceCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddFollowerUserIDs(ids...)
}

// AddRatingIDs adds the "ratings" edge to the Rating entity by IDs.
func (pc *PlaceCreate) AddRatingIDs(ids ...string) *PlaceCreate {
	pc.mutation.AddRatingIDs(ids...)
	return pc
}

// AddRatings adds the "ratings" edges to the Rating entity.
func (pc *PlaceCreate) AddRatings(r ...*Rating) *PlaceCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddRatingIDs(ids...)
}

// Mutation returns the PlaceMutation object of the builder.
func (pc *PlaceCreate) Mutation() *PlaceMutation {
	return pc.mutation
}

// Save creates the Place in the database.
func (pc *PlaceCreate) Save(ctx context.Context) (*Place, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlaceCreate) SaveX(ctx context.Context) *Place {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PlaceCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PlaceCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PlaceCreate) defaults() error {
	if _, ok := pc.mutation.CoverImage(); !ok {
		v := place.DefaultCoverImage
		pc.mutation.SetCoverImage(v)
	}
	if _, ok := pc.mutation.FollowerCount(); !ok {
		v := place.DefaultFollowerCount
		pc.mutation.SetFollowerCount(v)
	}
	if _, ok := pc.mutation.FollowingCount(); !ok {
		v := place.DefaultFollowingCount
		pc.mutation.SetFollowingCount(v)
	}
	if _, ok := pc.mutation.IsPremium(); !ok {
		v := place.DefaultIsPremium
		pc.mutation.SetIsPremium(v)
	}
	if _, ok := pc.mutation.IsPublished(); !ok {
		v := place.DefaultIsPublished
		pc.mutation.SetIsPublished(v)
	}
	if _, ok := pc.mutation.LikedByCurrentUser(); !ok {
		v := place.DefaultLikedByCurrentUser
		pc.mutation.SetLikedByCurrentUser(v)
	}
	if _, ok := pc.mutation.FollowedByCurrentUser(); !ok {
		v := place.DefaultFollowedByCurrentUser
		pc.mutation.SetFollowedByCurrentUser(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PlaceCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Place.name"`)}
	}
	if _, ok := pc.mutation.FollowerCount(); !ok {
		return &ValidationError{Name: "follower_count", err: errors.New(`ent: missing required field "Place.follower_count"`)}
	}
	if _, ok := pc.mutation.FollowingCount(); !ok {
		return &ValidationError{Name: "following_count", err: errors.New(`ent: missing required field "Place.following_count"`)}
	}
	if _, ok := pc.mutation.IsPremium(); !ok {
		return &ValidationError{Name: "is_Premium", err: errors.New(`ent: missing required field "Place.is_Premium"`)}
	}
	if _, ok := pc.mutation.IsPublished(); !ok {
		return &ValidationError{Name: "is_published", err: errors.New(`ent: missing required field "Place.is_published"`)}
	}
	if _, ok := pc.mutation.LikedByCurrentUser(); !ok {
		return &ValidationError{Name: "likedByCurrentUser", err: errors.New(`ent: missing required field "Place.likedByCurrentUser"`)}
	}
	if _, ok := pc.mutation.FollowedByCurrentUser(); !ok {
		return &ValidationError{Name: "followedByCurrentUser", err: errors.New(`ent: missing required field "Place.followedByCurrentUser"`)}
	}
	if v, ok := pc.mutation.ID(); ok {
		if err := place.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Place.id": %w`, err)}
		}
	}
	return nil
}

func (pc *PlaceCreate) sqlSave(ctx context.Context) (*Place, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Place.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PlaceCreate) createSpec() (*Place, *sqlgraph.CreateSpec) {
	var (
		_node = &Place{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(place.Table, sqlgraph.NewFieldSpec(place.FieldID, field.TypeString))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(place.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.GetType(); ok {
		_spec.SetField(place.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(place.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Location(); ok {
		_spec.SetField(place.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := pc.mutation.Email(); ok {
		_spec.SetField(place.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := pc.mutation.Phone(); ok {
		_spec.SetField(place.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := pc.mutation.Website(); ok {
		_spec.SetField(place.FieldWebsite, field.TypeString, value)
		_node.Website = value
	}
	if value, ok := pc.mutation.CoverImage(); ok {
		_spec.SetField(place.FieldCoverImage, field.TypeString, value)
		_node.CoverImage = value
	}
	if value, ok := pc.mutation.Picture(); ok {
		_spec.SetField(place.FieldPicture, field.TypeString, value)
		_node.Picture = value
	}
	if value, ok := pc.mutation.Country(); ok {
		_spec.SetField(place.FieldCountry, field.TypeString, value)
		_node.Country = value
	}
	if value, ok := pc.mutation.City(); ok {
		_spec.SetField(place.FieldCity, field.TypeString, value)
		_node.City = value
	}
	if value, ok := pc.mutation.State(); ok {
		_spec.SetField(place.FieldState, field.TypeString, value)
		_node.State = value
	}
	if value, ok := pc.mutation.PlaceSettings(); ok {
		_spec.SetField(place.FieldPlaceSettings, field.TypeJSON, value)
		_node.PlaceSettings = value
	}
	if value, ok := pc.mutation.OpeningHours(); ok {
		_spec.SetField(place.FieldOpeningHours, field.TypeJSON, value)
		_node.OpeningHours = value
	}
	if value, ok := pc.mutation.SocialMedia(); ok {
		_spec.SetField(place.FieldSocialMedia, field.TypeJSON, value)
		_node.SocialMedia = value
	}
	if value, ok := pc.mutation.PaymentOptions(); ok {
		_spec.SetField(place.FieldPaymentOptions, field.TypeJSON, value)
		_node.PaymentOptions = value
	}
	if value, ok := pc.mutation.Tags(); ok {
		_spec.SetField(place.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := pc.mutation.Features(); ok {
		_spec.SetField(place.FieldFeatures, field.TypeJSON, value)
		_node.Features = value
	}
	if value, ok := pc.mutation.AdditionalInfo(); ok {
		_spec.SetField(place.FieldAdditionalInfo, field.TypeJSON, value)
		_node.AdditionalInfo = value
	}
	if value, ok := pc.mutation.Images(); ok {
		_spec.SetField(place.FieldImages, field.TypeJSON, value)
		_node.Images = value
	}
	if value, ok := pc.mutation.Availability(); ok {
		_spec.SetField(place.FieldAvailability, field.TypeJSON, value)
		_node.Availability = value
	}
	if value, ok := pc.mutation.SpecialOffers(); ok {
		_spec.SetField(place.FieldSpecialOffers, field.TypeString, value)
		_node.SpecialOffers = value
	}
	if value, ok := pc.mutation.SustainabilityScore(); ok {
		_spec.SetField(place.FieldSustainabilityScore, field.TypeFloat64, value)
		_node.SustainabilityScore = value
	}
	if value, ok := pc.mutation.MapCoordinates(); ok {
		_spec.SetField(place.FieldMapCoordinates, field.TypeJSON, value)
		_node.MapCoordinates = value
	}
	if value, ok := pc.mutation.Longitude(); ok {
		_spec.SetField(place.FieldLongitude, field.TypeString, value)
		_node.Longitude = value
	}
	if value, ok := pc.mutation.Latitude(); ok {
		_spec.SetField(place.FieldLatitude, field.TypeString, value)
		_node.Latitude = value
	}
	if value, ok := pc.mutation.SearchText(); ok {
		_spec.SetField(place.FieldSearchText, field.TypeString, value)
		_node.SearchText = value
	}
	if value, ok := pc.mutation.RelevanceScore(); ok {
		_spec.SetField(place.FieldRelevanceScore, field.TypeFloat64, value)
		_node.RelevanceScore = value
	}
	if value, ok := pc.mutation.FollowerCount(); ok {
		_spec.SetField(place.FieldFollowerCount, field.TypeInt, value)
		_node.FollowerCount = value
	}
	if value, ok := pc.mutation.FollowingCount(); ok {
		_spec.SetField(place.FieldFollowingCount, field.TypeInt, value)
		_node.FollowingCount = value
	}
	if value, ok := pc.mutation.IsPremium(); ok {
		_spec.SetField(place.FieldIsPremium, field.TypeBool, value)
		_node.IsPremium = value
	}
	if value, ok := pc.mutation.IsPublished(); ok {
		_spec.SetField(place.FieldIsPublished, field.TypeBool, value)
		_node.IsPublished = value
	}
	if value, ok := pc.mutation.LikedByCurrentUser(); ok {
		_spec.SetField(place.FieldLikedByCurrentUser, field.TypeBool, value)
		_node.LikedByCurrentUser = value
	}
	if value, ok := pc.mutation.FollowedByCurrentUser(); ok {
		_spec.SetField(place.FieldFollowedByCurrentUser, field.TypeBool, value)
		_node.FollowedByCurrentUser = value
	}
	if nodes := pc.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   place.BusinessTable,
			Columns: []string{place.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.business_places = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   place.UsersTable,
			Columns: place.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   place.ReviewsTable,
			Columns: []string{place.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   place.EventsTable,
			Columns: []string{place.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.AmenitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   place.AmenitiesTable,
			Columns: place.AmenitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amenity.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   place.MenusTable,
			Columns: []string{place.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.MediasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   place.MediasTable,
			Columns: place.MediasPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.RoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   place.RoomsTable,
			Columns: []string{place.RoomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ReservationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   place.ReservationsTable,
			Columns: []string{place.ReservationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reservation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.BookingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   place.BookingsTable,
			Columns: []string{place.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   place.CategoriesTable,
			Columns: []string{place.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CategoryAssignmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   place.CategoryAssignmentsTable,
			Columns: []string{place.CategoryAssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(categoryassignment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.FaqsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   place.FaqsTable,
			Columns: place.FaqsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(faq.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.LikedByUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   place.LikedByUsersTable,
			Columns: []string{place.LikedByUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userlikeplace.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.FollowerUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   place.FollowerUsersTable,
			Columns: []string{place.FollowerUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userfollowplace.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.RatingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   place.RatingsTable,
			Columns: []string{place.RatingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rating.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlaceCreateBulk is the builder for creating many Place entities in bulk.
type PlaceCreateBulk struct {
	config
	builders []*PlaceCreate
}

// Save creates the Place entities in the database.
func (pcb *PlaceCreateBulk) Save(ctx context.Context) ([]*Place, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Place, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlaceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PlaceCreateBulk) SaveX(ctx context.Context) []*Place {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PlaceCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PlaceCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
