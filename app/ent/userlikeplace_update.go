



// Code generated by ent, DO NOT EDIT.



package ent



	
import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
		"placio-app/ent/predicate"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
			"database/sql/driver"
			"entgo.io/ent/dialect/sql"
			"entgo.io/ent/dialect/sql/sqlgraph"
			"entgo.io/ent/dialect/sql/sqljson"
			"entgo.io/ent/schema/field"
			 "placio-app/ent/userlikeplace"
			 "placio-app/ent/user"
			 "placio-app/ent/place"

)








// UserLikePlaceUpdate is the builder for updating UserLikePlace entities.
type UserLikePlaceUpdate struct {
	config
	hooks []Hook
	mutation *UserLikePlaceMutation
}

// Where appends a list predicates to the UserLikePlaceUpdate builder.
func (ulpu *UserLikePlaceUpdate) Where(ps ...predicate.UserLikePlace) *UserLikePlaceUpdate {
	ulpu.mutation.Where(ps...)
	return ulpu
}


	





	
	
	


	
	
	// SetCreatedAt sets the "CreatedAt" field.
	func (ulpu *UserLikePlaceUpdate) SetCreatedAt(t time.Time) *UserLikePlaceUpdate {
		ulpu.mutation.SetCreatedAt(t)
		return ulpu
	}

	
	
	
	
	
	
		// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
		func (ulpu *UserLikePlaceUpdate) SetNillableCreatedAt(t *time.Time) *UserLikePlaceUpdate {
			if t != nil {
				ulpu.SetCreatedAt(*t)
			}
			return ulpu
		}
	

	

	

	

	
	
	// SetUpdatedAt sets the "UpdatedAt" field.
	func (ulpu *UserLikePlaceUpdate) SetUpdatedAt(t time.Time) *UserLikePlaceUpdate {
		ulpu.mutation.SetUpdatedAt(t)
		return ulpu
	}

	
	
	
	
	
	

	

	

	



	
	
	
	
	
		// SetUserID sets the "user" edge to the User entity by ID.
		func (ulpu *UserLikePlaceUpdate) SetUserID(id string) *UserLikePlaceUpdate {
			ulpu.mutation.SetUserID(id)
			return ulpu
		}
	
	
		
		// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
		func (ulpu *UserLikePlaceUpdate) SetNillableUserID(id *string) *UserLikePlaceUpdate {
			if id != nil {
				ulpu = ulpu.SetUserID(*id)
			}
			return ulpu
		}
	
	
	
	
	// SetUser sets the "user" edge to the User entity.
	func (ulpu *UserLikePlaceUpdate) SetUser(u *User) *UserLikePlaceUpdate {
		return ulpu.SetUserID(u.ID)
	}

	
	
	
	
	
		// SetPlaceID sets the "place" edge to the Place entity by ID.
		func (ulpu *UserLikePlaceUpdate) SetPlaceID(id string) *UserLikePlaceUpdate {
			ulpu.mutation.SetPlaceID(id)
			return ulpu
		}
	
	
		
		// SetNillablePlaceID sets the "place" edge to the Place entity by ID if the given value is not nil.
		func (ulpu *UserLikePlaceUpdate) SetNillablePlaceID(id *string) *UserLikePlaceUpdate {
			if id != nil {
				ulpu = ulpu.SetPlaceID(*id)
			}
			return ulpu
		}
	
	
	
	
	// SetPlace sets the "place" edge to the Place entity.
	func (ulpu *UserLikePlaceUpdate) SetPlace(p *Place) *UserLikePlaceUpdate {
		return ulpu.SetPlaceID(p.ID)
	}


// Mutation returns the UserLikePlaceMutation object of the builder.
func (ulpu *UserLikePlaceUpdate) Mutation() *UserLikePlaceMutation {
	return ulpu.mutation
}





	





	
	
	// ClearUser clears the "user" edge to the User entity.
	func (ulpu *UserLikePlaceUpdate) ClearUser() *UserLikePlaceUpdate {
		ulpu.mutation.ClearUser()
		return ulpu
	}
	

	
	
	// ClearPlace clears the "place" edge to the Place entity.
	func (ulpu *UserLikePlaceUpdate) ClearPlace() *UserLikePlaceUpdate {
		ulpu.mutation.ClearPlace()
		return ulpu
	}
	




// Save executes the query and returns the number of nodes affected by the update operation.
func (ulpu *UserLikePlaceUpdate) Save(ctx context.Context) (int, error) {
			ulpu.defaults()
	return withHooks(ctx, ulpu.sqlSave, ulpu.mutation, ulpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ulpu *UserLikePlaceUpdate) SaveX(ctx context.Context) int {
	affected, err := ulpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ulpu *UserLikePlaceUpdate) Exec(ctx context.Context) error {
	_, err := ulpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulpu *UserLikePlaceUpdate) ExecX(ctx context.Context) {
	if err := ulpu.Exec(ctx); err != nil {
		panic(err)
	}
}


	







	// defaults sets the default values of the builder before save.
	func (ulpu *UserLikePlaceUpdate) defaults() {
				if _, ok := ulpu.mutation.UpdatedAt(); !ok  {
					v := userlikeplace.UpdateDefaultUpdatedAt()
					ulpu.mutation.SetUpdatedAt(v)
				}
	}








	
	





    


func (ulpu *UserLikePlaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userlikeplace.Table, userlikeplace.Columns,sqlgraph.NewFieldSpec(userlikeplace.FieldID, field.TypeString))
	if ps := ulpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
				if value, ok := ulpu.mutation.CreatedAt(); ok {
						_spec.SetField(userlikeplace.FieldCreatedAt, field.TypeTime, value)
				}
				if value, ok := ulpu.mutation.UpdatedAt(); ok {
						_spec.SetField(userlikeplace.FieldUpdatedAt, field.TypeTime, value)
				}
		if ulpu.mutation.UserCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: userlikeplace.UserTable,
		Columns: []string{ userlikeplace.UserColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
		if nodes := ulpu.mutation.UserIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: userlikeplace.UserTable,
		Columns: []string{ userlikeplace.UserColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges.Add = append(_spec.Edges.Add, edge)
		}
		if ulpu.mutation.PlaceCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: false,
		Table: userlikeplace.PlaceTable,
		Columns: []string{ userlikeplace.PlaceColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
		if nodes := ulpu.mutation.PlaceIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: false,
		Table: userlikeplace.PlaceTable,
		Columns: []string{ userlikeplace.PlaceColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges.Add = append(_spec.Edges.Add, edge)
		}
		if n, err = sqlgraph.UpdateNodes(ctx, ulpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ userlikeplace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ulpu.mutation.done = true
	return n, nil
}







// UserLikePlaceUpdateOne is the builder for updating a single UserLikePlace entity.
type UserLikePlaceUpdateOne struct {
	config
	fields []string
	hooks []Hook
	mutation *UserLikePlaceMutation

}


	





	
	
	


	
	
	// SetCreatedAt sets the "CreatedAt" field.
	func (ulpuo *UserLikePlaceUpdateOne) SetCreatedAt(t time.Time) *UserLikePlaceUpdateOne {
		ulpuo.mutation.SetCreatedAt(t)
		return ulpuo
	}

	
	
	
	
	
	
		// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
		func (ulpuo *UserLikePlaceUpdateOne) SetNillableCreatedAt(t *time.Time) *UserLikePlaceUpdateOne {
			if t != nil {
				ulpuo.SetCreatedAt(*t)
			}
			return ulpuo
		}
	

	

	

	

	
	
	// SetUpdatedAt sets the "UpdatedAt" field.
	func (ulpuo *UserLikePlaceUpdateOne) SetUpdatedAt(t time.Time) *UserLikePlaceUpdateOne {
		ulpuo.mutation.SetUpdatedAt(t)
		return ulpuo
	}

	
	
	
	
	
	

	

	

	



	
	
	
	
	
		// SetUserID sets the "user" edge to the User entity by ID.
		func (ulpuo *UserLikePlaceUpdateOne) SetUserID(id string) *UserLikePlaceUpdateOne {
			ulpuo.mutation.SetUserID(id)
			return ulpuo
		}
	
	
		
		// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
		func (ulpuo *UserLikePlaceUpdateOne) SetNillableUserID(id *string) *UserLikePlaceUpdateOne {
			if id != nil {
				ulpuo = ulpuo.SetUserID(*id)
			}
			return ulpuo
		}
	
	
	
	
	// SetUser sets the "user" edge to the User entity.
	func (ulpuo *UserLikePlaceUpdateOne) SetUser(u *User) *UserLikePlaceUpdateOne {
		return ulpuo.SetUserID(u.ID)
	}

	
	
	
	
	
		// SetPlaceID sets the "place" edge to the Place entity by ID.
		func (ulpuo *UserLikePlaceUpdateOne) SetPlaceID(id string) *UserLikePlaceUpdateOne {
			ulpuo.mutation.SetPlaceID(id)
			return ulpuo
		}
	
	
		
		// SetNillablePlaceID sets the "place" edge to the Place entity by ID if the given value is not nil.
		func (ulpuo *UserLikePlaceUpdateOne) SetNillablePlaceID(id *string) *UserLikePlaceUpdateOne {
			if id != nil {
				ulpuo = ulpuo.SetPlaceID(*id)
			}
			return ulpuo
		}
	
	
	
	
	// SetPlace sets the "place" edge to the Place entity.
	func (ulpuo *UserLikePlaceUpdateOne) SetPlace(p *Place) *UserLikePlaceUpdateOne {
		return ulpuo.SetPlaceID(p.ID)
	}


// Mutation returns the UserLikePlaceMutation object of the builder.
func (ulpuo *UserLikePlaceUpdateOne) Mutation() *UserLikePlaceMutation {
	return ulpuo.mutation
}






	





	
	
	// ClearUser clears the "user" edge to the User entity.
	func (ulpuo *UserLikePlaceUpdateOne) ClearUser() *UserLikePlaceUpdateOne {
		ulpuo.mutation.ClearUser()
		return ulpuo
	}
	

	
	
	// ClearPlace clears the "place" edge to the Place entity.
	func (ulpuo *UserLikePlaceUpdateOne) ClearPlace() *UserLikePlaceUpdateOne {
		ulpuo.mutation.ClearPlace()
		return ulpuo
	}
	




// Where appends a list predicates to the UserLikePlaceUpdate builder.
func (ulpuo *UserLikePlaceUpdateOne) Where(ps ...predicate.UserLikePlace) *UserLikePlaceUpdateOne {
	ulpuo.mutation.Where(ps...)
	return ulpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ulpuo *UserLikePlaceUpdateOne) Select(field string, fields ...string) *UserLikePlaceUpdateOne {
	ulpuo.fields = append([]string{field}, fields...)
	return ulpuo
}

// Save executes the query and returns the updated UserLikePlace entity.
func (ulpuo *UserLikePlaceUpdateOne ) Save(ctx context.Context) (*UserLikePlace, error) {
			ulpuo.defaults()
	return withHooks(ctx, ulpuo.sqlSave, ulpuo.mutation, ulpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ulpuo *UserLikePlaceUpdateOne) SaveX(ctx context.Context) *UserLikePlace {
	node, err := ulpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ulpuo *UserLikePlaceUpdateOne) Exec(ctx context.Context) error {
	_, err := ulpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulpuo *UserLikePlaceUpdateOne) ExecX(ctx context.Context) {
	if err := ulpuo.Exec(ctx); err != nil {
		panic(err)
	}
}


	







	// defaults sets the default values of the builder before save.
	func (ulpuo *UserLikePlaceUpdateOne) defaults() {
				if _, ok := ulpuo.mutation.UpdatedAt(); !ok  {
					v := userlikeplace.UpdateDefaultUpdatedAt()
					ulpuo.mutation.SetUpdatedAt(v)
				}
	}








	
	





    


func (ulpuo *UserLikePlaceUpdateOne) sqlSave(ctx context.Context) (_node *UserLikePlace, err error) {
	_spec := sqlgraph.NewUpdateSpec(userlikeplace.Table, userlikeplace.Columns,sqlgraph.NewFieldSpec(userlikeplace.FieldID, field.TypeString))
			id, ok := ulpuo.mutation.ID()
			if !ok {
				return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserLikePlace.id" for update`)}
			}
			_spec.Node.ID.Value = id
			if fields := ulpuo.fields; len(fields) > 0 {
				_spec.Node.Columns = make([]string, 0, len(fields))
				_spec.Node.Columns = append(_spec.Node.Columns, userlikeplace.FieldID)
				for _, f := range fields {
					if !userlikeplace.ValidColumn(f) {
						return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
					}
					if f != userlikeplace.FieldID {
						_spec.Node.Columns = append(_spec.Node.Columns, f)
					}
				}
			}
	if ps := ulpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
				if value, ok := ulpuo.mutation.CreatedAt(); ok {
						_spec.SetField(userlikeplace.FieldCreatedAt, field.TypeTime, value)
				}
				if value, ok := ulpuo.mutation.UpdatedAt(); ok {
						_spec.SetField(userlikeplace.FieldUpdatedAt, field.TypeTime, value)
				}
		if ulpuo.mutation.UserCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: userlikeplace.UserTable,
		Columns: []string{ userlikeplace.UserColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
		if nodes := ulpuo.mutation.UserIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: userlikeplace.UserTable,
		Columns: []string{ userlikeplace.UserColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges.Add = append(_spec.Edges.Add, edge)
		}
		if ulpuo.mutation.PlaceCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: false,
		Table: userlikeplace.PlaceTable,
		Columns: []string{ userlikeplace.PlaceColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
		if nodes := ulpuo.mutation.PlaceIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: false,
		Table: userlikeplace.PlaceTable,
		Columns: []string{ userlikeplace.PlaceColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges.Add = append(_spec.Edges.Add, edge)
		}
		_node = &UserLikePlace{config: ulpuo.config}
		_spec.Assign = _node.assignValues
		_spec.ScanValues = _node.scanValues
		if err = sqlgraph.UpdateNode(ctx, ulpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ userlikeplace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ulpuo.mutation.done = true
	return _node, nil
}



