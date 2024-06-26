// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/category"
	"placio-app/ent/media"
	"placio-app/ent/menu"
	"placio-app/ent/place"
	"placio-app/ent/placeinventory"
	"placio-app/ent/plan"
	"placio-app/ent/post"
	"placio-app/ent/review"
	"placio-app/ent/room"
	"placio-app/ent/roomcategory"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MediaCreate is the builder for creating a Media entity.
type MediaCreate struct {
	config
	mutation *MediaMutation
	hooks    []Hook
}

// SetURL sets the "URL" field.
func (mc *MediaCreate) SetURL(s string) *MediaCreate {
	mc.mutation.SetURL(s)
	return mc
}

// SetMediaType sets the "MediaType" field.
func (mc *MediaCreate) SetMediaType(s string) *MediaCreate {
	mc.mutation.SetMediaType(s)
	return mc
}

// SetCreatedAt sets the "CreatedAt" field.
func (mc *MediaCreate) SetCreatedAt(t time.Time) *MediaCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (mc *MediaCreate) SetNillableCreatedAt(t *time.Time) *MediaCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (mc *MediaCreate) SetUpdatedAt(t time.Time) *MediaCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (mc *MediaCreate) SetNillableUpdatedAt(t *time.Time) *MediaCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetLikeCount sets the "likeCount" field.
func (mc *MediaCreate) SetLikeCount(i int) *MediaCreate {
	mc.mutation.SetLikeCount(i)
	return mc
}

// SetNillableLikeCount sets the "likeCount" field if the given value is not nil.
func (mc *MediaCreate) SetNillableLikeCount(i *int) *MediaCreate {
	if i != nil {
		mc.SetLikeCount(*i)
	}
	return mc
}

// SetDislikeCount sets the "dislikeCount" field.
func (mc *MediaCreate) SetDislikeCount(i int) *MediaCreate {
	mc.mutation.SetDislikeCount(i)
	return mc
}

// SetNillableDislikeCount sets the "dislikeCount" field if the given value is not nil.
func (mc *MediaCreate) SetNillableDislikeCount(i *int) *MediaCreate {
	if i != nil {
		mc.SetDislikeCount(*i)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MediaCreate) SetID(s string) *MediaCreate {
	mc.mutation.SetID(s)
	return mc
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (mc *MediaCreate) SetPostID(id string) *MediaCreate {
	mc.mutation.SetPostID(id)
	return mc
}

// SetNillablePostID sets the "post" edge to the Post entity by ID if the given value is not nil.
func (mc *MediaCreate) SetNillablePostID(id *string) *MediaCreate {
	if id != nil {
		mc = mc.SetPostID(*id)
	}
	return mc
}

// SetPost sets the "post" edge to the Post entity.
func (mc *MediaCreate) SetPost(p *Post) *MediaCreate {
	return mc.SetPostID(p.ID)
}

// SetReviewID sets the "review" edge to the Review entity by ID.
func (mc *MediaCreate) SetReviewID(id string) *MediaCreate {
	mc.mutation.SetReviewID(id)
	return mc
}

// SetNillableReviewID sets the "review" edge to the Review entity by ID if the given value is not nil.
func (mc *MediaCreate) SetNillableReviewID(id *string) *MediaCreate {
	if id != nil {
		mc = mc.SetReviewID(*id)
	}
	return mc
}

// SetReview sets the "review" edge to the Review entity.
func (mc *MediaCreate) SetReview(r *Review) *MediaCreate {
	return mc.SetReviewID(r.ID)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (mc *MediaCreate) AddCategoryIDs(ids ...string) *MediaCreate {
	mc.mutation.AddCategoryIDs(ids...)
	return mc
}

// AddCategories adds the "categories" edges to the Category entity.
func (mc *MediaCreate) AddCategories(c ...*Category) *MediaCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mc.AddCategoryIDs(ids...)
}

// AddPlaceIDs adds the "place" edge to the Place entity by IDs.
func (mc *MediaCreate) AddPlaceIDs(ids ...string) *MediaCreate {
	mc.mutation.AddPlaceIDs(ids...)
	return mc
}

// AddPlace adds the "place" edges to the Place entity.
func (mc *MediaCreate) AddPlace(p ...*Place) *MediaCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mc.AddPlaceIDs(ids...)
}

// AddPlaceInventoryIDs adds the "place_inventory" edge to the PlaceInventory entity by IDs.
func (mc *MediaCreate) AddPlaceInventoryIDs(ids ...string) *MediaCreate {
	mc.mutation.AddPlaceInventoryIDs(ids...)
	return mc
}

// AddPlaceInventory adds the "place_inventory" edges to the PlaceInventory entity.
func (mc *MediaCreate) AddPlaceInventory(p ...*PlaceInventory) *MediaCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mc.AddPlaceInventoryIDs(ids...)
}

// AddMenuIDs adds the "menu" edge to the Menu entity by IDs.
func (mc *MediaCreate) AddMenuIDs(ids ...string) *MediaCreate {
	mc.mutation.AddMenuIDs(ids...)
	return mc
}

// AddMenu adds the "menu" edges to the Menu entity.
func (mc *MediaCreate) AddMenu(m ...*Menu) *MediaCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddMenuIDs(ids...)
}

// AddRoomCategoryIDs adds the "room_category" edge to the RoomCategory entity by IDs.
func (mc *MediaCreate) AddRoomCategoryIDs(ids ...string) *MediaCreate {
	mc.mutation.AddRoomCategoryIDs(ids...)
	return mc
}

// AddRoomCategory adds the "room_category" edges to the RoomCategory entity.
func (mc *MediaCreate) AddRoomCategory(r ...*RoomCategory) *MediaCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mc.AddRoomCategoryIDs(ids...)
}

// AddRoomIDs adds the "room" edge to the Room entity by IDs.
func (mc *MediaCreate) AddRoomIDs(ids ...string) *MediaCreate {
	mc.mutation.AddRoomIDs(ids...)
	return mc
}

// AddRoom adds the "room" edges to the Room entity.
func (mc *MediaCreate) AddRoom(r ...*Room) *MediaCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mc.AddRoomIDs(ids...)
}

// SetPlanID sets the "plan" edge to the Plan entity by ID.
func (mc *MediaCreate) SetPlanID(id string) *MediaCreate {
	mc.mutation.SetPlanID(id)
	return mc
}

// SetNillablePlanID sets the "plan" edge to the Plan entity by ID if the given value is not nil.
func (mc *MediaCreate) SetNillablePlanID(id *string) *MediaCreate {
	if id != nil {
		mc = mc.SetPlanID(*id)
	}
	return mc
}

// SetPlan sets the "plan" edge to the Plan entity.
func (mc *MediaCreate) SetPlan(p *Plan) *MediaCreate {
	return mc.SetPlanID(p.ID)
}

// Mutation returns the MediaMutation object of the builder.
func (mc *MediaCreate) Mutation() *MediaMutation {
	return mc.mutation
}

// Save creates the Media in the database.
func (mc *MediaCreate) Save(ctx context.Context) (*Media, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MediaCreate) SaveX(ctx context.Context) *Media {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MediaCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MediaCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MediaCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := media.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := media.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.LikeCount(); !ok {
		v := media.DefaultLikeCount
		mc.mutation.SetLikeCount(v)
	}
	if _, ok := mc.mutation.DislikeCount(); !ok {
		v := media.DefaultDislikeCount
		mc.mutation.SetDislikeCount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MediaCreate) check() error {
	if _, ok := mc.mutation.URL(); !ok {
		return &ValidationError{Name: "URL", err: errors.New(`ent: missing required field "Media.URL"`)}
	}
	if _, ok := mc.mutation.MediaType(); !ok {
		return &ValidationError{Name: "MediaType", err: errors.New(`ent: missing required field "Media.MediaType"`)}
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "Media.CreatedAt"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`ent: missing required field "Media.UpdatedAt"`)}
	}
	if _, ok := mc.mutation.LikeCount(); !ok {
		return &ValidationError{Name: "likeCount", err: errors.New(`ent: missing required field "Media.likeCount"`)}
	}
	if _, ok := mc.mutation.DislikeCount(); !ok {
		return &ValidationError{Name: "dislikeCount", err: errors.New(`ent: missing required field "Media.dislikeCount"`)}
	}
	if v, ok := mc.mutation.ID(); ok {
		if err := media.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Media.id": %w`, err)}
		}
	}
	return nil
}

func (mc *MediaCreate) sqlSave(ctx context.Context) (*Media, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Media.ID type: %T", _spec.ID.Value)
		}
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MediaCreate) createSpec() (*Media, *sqlgraph.CreateSpec) {
	var (
		_node = &Media{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(media.Table, sqlgraph.NewFieldSpec(media.FieldID, field.TypeString))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.URL(); ok {
		_spec.SetField(media.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := mc.mutation.MediaType(); ok {
		_spec.SetField(media.FieldMediaType, field.TypeString, value)
		_node.MediaType = value
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(media.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(media.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.LikeCount(); ok {
		_spec.SetField(media.FieldLikeCount, field.TypeInt, value)
		_node.LikeCount = value
	}
	if value, ok := mc.mutation.DislikeCount(); ok {
		_spec.SetField(media.FieldDislikeCount, field.TypeInt, value)
		_node.DislikeCount = value
	}
	if nodes := mc.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   media.PostTable,
			Columns: []string{media.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.post_medias = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ReviewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   media.ReviewTable,
			Columns: []string{media.ReviewColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(review.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.review_medias = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   media.CategoriesTable,
			Columns: media.CategoriesPrimaryKey,
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
	if nodes := mc.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   media.PlaceTable,
			Columns: media.PlacePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.PlaceInventoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   media.PlaceInventoryTable,
			Columns: media.PlaceInventoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placeinventory.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MenuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   media.MenuTable,
			Columns: media.MenuPrimaryKey,
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
	if nodes := mc.mutation.RoomCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   media.RoomCategoryTable,
			Columns: media.RoomCategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(roomcategory.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   media.RoomTable,
			Columns: media.RoomPrimaryKey,
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
	if nodes := mc.mutation.PlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   media.PlanTable,
			Columns: []string{media.PlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(plan.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.plan_media = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MediaCreateBulk is the builder for creating many Media entities in bulk.
type MediaCreateBulk struct {
	config
	err      error
	builders []*MediaCreate
}

// Save creates the Media entities in the database.
func (mcb *MediaCreateBulk) Save(ctx context.Context) ([]*Media, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Media, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MediaMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MediaCreateBulk) SaveX(ctx context.Context) []*Media {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MediaCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MediaCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
