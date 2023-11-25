



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
			 "placio-app/ent/categoryassignment"
			 "placio-app/ent/user"
			 "placio-app/ent/business"
			 "placio-app/ent/place"
			 "placio-app/ent/category"

)








// CategoryAssignmentUpdate is the builder for updating CategoryAssignment entities.
type CategoryAssignmentUpdate struct {
	config
	hooks []Hook
	mutation *CategoryAssignmentMutation
}

// Where appends a list predicates to the CategoryAssignmentUpdate builder.
func (cau *CategoryAssignmentUpdate) Where(ps ...predicate.CategoryAssignment) *CategoryAssignmentUpdate {
	cau.mutation.Where(ps...)
	return cau
}


	





	
	
	


	
	
	// SetEntityID sets the "entity_id" field.
	func (cau *CategoryAssignmentUpdate) SetEntityID(s string) *CategoryAssignmentUpdate {
		cau.mutation.SetEntityID(s)
		return cau
	}

	
	
	
	
	
	
		// SetNillableEntityID sets the "entity_id" field if the given value is not nil.
		func (cau *CategoryAssignmentUpdate) SetNillableEntityID(s *string) *CategoryAssignmentUpdate {
			if s != nil {
				cau.SetEntityID(*s)
			}
			return cau
		}
	

	

	

	
		
		// ClearEntityID clears the value of the "entity_id" field.
		func (cau *CategoryAssignmentUpdate) ClearEntityID() *CategoryAssignmentUpdate {
			cau.mutation.ClearEntityID()
			return cau
		}
	

	
	
	// SetEntityType sets the "entity_type" field.
	func (cau *CategoryAssignmentUpdate) SetEntityType(s string) *CategoryAssignmentUpdate {
		cau.mutation.SetEntityType(s)
		return cau
	}

	
	
	
	
	
	
		// SetNillableEntityType sets the "entity_type" field if the given value is not nil.
		func (cau *CategoryAssignmentUpdate) SetNillableEntityType(s *string) *CategoryAssignmentUpdate {
			if s != nil {
				cau.SetEntityType(*s)
			}
			return cau
		}
	

	

	

	
		
		// ClearEntityType clears the value of the "entity_type" field.
		func (cau *CategoryAssignmentUpdate) ClearEntityType() *CategoryAssignmentUpdate {
			cau.mutation.ClearEntityType()
			return cau
		}
	

	
	
	// SetCategoryID sets the "category_id" field.
	func (cau *CategoryAssignmentUpdate) SetCategoryID(s string) *CategoryAssignmentUpdate {
		cau.mutation.SetCategoryID(s)
		return cau
	}

	
	
	
	
	
	
		// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
		func (cau *CategoryAssignmentUpdate) SetNillableCategoryID(s *string) *CategoryAssignmentUpdate {
			if s != nil {
				cau.SetCategoryID(*s)
			}
			return cau
		}
	

	

	

	
		
		// ClearCategoryID clears the value of the "category_id" field.
		func (cau *CategoryAssignmentUpdate) ClearCategoryID() *CategoryAssignmentUpdate {
			cau.mutation.ClearCategoryID()
			return cau
		}
	



	
	
	
	
	
		// SetUserID sets the "user" edge to the User entity by ID.
		func (cau *CategoryAssignmentUpdate) SetUserID(id string) *CategoryAssignmentUpdate {
			cau.mutation.SetUserID(id)
			return cau
		}
	
	
		
		// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
		func (cau *CategoryAssignmentUpdate) SetNillableUserID(id *string) *CategoryAssignmentUpdate {
			if id != nil {
				cau = cau.SetUserID(*id)
			}
			return cau
		}
	
	
	
	
	// SetUser sets the "user" edge to the User entity.
	func (cau *CategoryAssignmentUpdate) SetUser(u *User) *CategoryAssignmentUpdate {
		return cau.SetUserID(u.ID)
	}

	
	
	
	
	
		// SetBusinessID sets the "business" edge to the Business entity by ID.
		func (cau *CategoryAssignmentUpdate) SetBusinessID(id string) *CategoryAssignmentUpdate {
			cau.mutation.SetBusinessID(id)
			return cau
		}
	
	
		
		// SetNillableBusinessID sets the "business" edge to the Business entity by ID if the given value is not nil.
		func (cau *CategoryAssignmentUpdate) SetNillableBusinessID(id *string) *CategoryAssignmentUpdate {
			if id != nil {
				cau = cau.SetBusinessID(*id)
			}
			return cau
		}
	
	
	
	
	// SetBusiness sets the "business" edge to the Business entity.
	func (cau *CategoryAssignmentUpdate) SetBusiness(b *Business) *CategoryAssignmentUpdate {
		return cau.SetBusinessID(b.ID)
	}

	
	
	
	
	
		// SetPlaceID sets the "place" edge to the Place entity by ID.
		func (cau *CategoryAssignmentUpdate) SetPlaceID(id string) *CategoryAssignmentUpdate {
			cau.mutation.SetPlaceID(id)
			return cau
		}
	
	
		
		// SetNillablePlaceID sets the "place" edge to the Place entity by ID if the given value is not nil.
		func (cau *CategoryAssignmentUpdate) SetNillablePlaceID(id *string) *CategoryAssignmentUpdate {
			if id != nil {
				cau = cau.SetPlaceID(*id)
			}
			return cau
		}
	
	
	
	
	// SetPlace sets the "place" edge to the Place entity.
	func (cau *CategoryAssignmentUpdate) SetPlace(p *Place) *CategoryAssignmentUpdate {
		return cau.SetPlaceID(p.ID)
	}

	
	
	
	
	
	
	
	
	
	// SetCategory sets the "category" edge to the Category entity.
	func (cau *CategoryAssignmentUpdate) SetCategory(c *Category) *CategoryAssignmentUpdate {
		return cau.SetCategoryID(c.ID)
	}


// Mutation returns the CategoryAssignmentMutation object of the builder.
func (cau *CategoryAssignmentUpdate) Mutation() *CategoryAssignmentMutation {
	return cau.mutation
}





	





	
	
	// ClearUser clears the "user" edge to the User entity.
	func (cau *CategoryAssignmentUpdate) ClearUser() *CategoryAssignmentUpdate {
		cau.mutation.ClearUser()
		return cau
	}
	

	
	
	// ClearBusiness clears the "business" edge to the Business entity.
	func (cau *CategoryAssignmentUpdate) ClearBusiness() *CategoryAssignmentUpdate {
		cau.mutation.ClearBusiness()
		return cau
	}
	

	
	
	// ClearPlace clears the "place" edge to the Place entity.
	func (cau *CategoryAssignmentUpdate) ClearPlace() *CategoryAssignmentUpdate {
		cau.mutation.ClearPlace()
		return cau
	}
	

	
	
	// ClearCategory clears the "category" edge to the Category entity.
	func (cau *CategoryAssignmentUpdate) ClearCategory() *CategoryAssignmentUpdate {
		cau.mutation.ClearCategory()
		return cau
	}
	




// Save executes the query and returns the number of nodes affected by the update operation.
func (cau *CategoryAssignmentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cau.sqlSave, cau.mutation, cau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cau *CategoryAssignmentUpdate) SaveX(ctx context.Context) int {
	affected, err := cau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cau *CategoryAssignmentUpdate) Exec(ctx context.Context) error {
	_, err := cau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cau *CategoryAssignmentUpdate) ExecX(ctx context.Context) {
	if err := cau.Exec(ctx); err != nil {
		panic(err)
	}
}


	














	
	





    


func (cau *CategoryAssignmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(categoryassignment.Table, categoryassignment.Columns,sqlgraph.NewFieldSpec(categoryassignment.FieldID, field.TypeString))
	if ps := cau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
				if value, ok := cau.mutation.EntityType(); ok {
						_spec.SetField(categoryassignment.FieldEntityType, field.TypeString, value)
				}
				if cau.mutation.EntityTypeCleared() {
					_spec.ClearField(categoryassignment.FieldEntityType, field.TypeString)
				}
		if cau.mutation.UserCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.UserTable,
		Columns: []string{ categoryassignment.UserColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
		if nodes := cau.mutation.UserIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.UserTable,
		Columns: []string{ categoryassignment.UserColumn },
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
		if cau.mutation.BusinessCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.BusinessTable,
		Columns: []string{ categoryassignment.BusinessColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
		if nodes := cau.mutation.BusinessIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.BusinessTable,
		Columns: []string{ categoryassignment.BusinessColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges.Add = append(_spec.Edges.Add, edge)
		}
		if cau.mutation.PlaceCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.PlaceTable,
		Columns: []string{ categoryassignment.PlaceColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
		if nodes := cau.mutation.PlaceIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.PlaceTable,
		Columns: []string{ categoryassignment.PlaceColumn },
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
		if cau.mutation.CategoryCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.CategoryTable,
		Columns: []string{ categoryassignment.CategoryColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
		if nodes := cau.mutation.CategoryIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.CategoryTable,
		Columns: []string{ categoryassignment.CategoryColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges.Add = append(_spec.Edges.Add, edge)
		}
		if n, err = sqlgraph.UpdateNodes(ctx, cau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ categoryassignment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cau.mutation.done = true
	return n, nil
}







// CategoryAssignmentUpdateOne is the builder for updating a single CategoryAssignment entity.
type CategoryAssignmentUpdateOne struct {
	config
	fields []string
	hooks []Hook
	mutation *CategoryAssignmentMutation

}


	





	
	
	


	
	
	// SetEntityID sets the "entity_id" field.
	func (cauo *CategoryAssignmentUpdateOne) SetEntityID(s string) *CategoryAssignmentUpdateOne {
		cauo.mutation.SetEntityID(s)
		return cauo
	}

	
	
	
	
	
	
		// SetNillableEntityID sets the "entity_id" field if the given value is not nil.
		func (cauo *CategoryAssignmentUpdateOne) SetNillableEntityID(s *string) *CategoryAssignmentUpdateOne {
			if s != nil {
				cauo.SetEntityID(*s)
			}
			return cauo
		}
	

	

	

	
		
		// ClearEntityID clears the value of the "entity_id" field.
		func (cauo *CategoryAssignmentUpdateOne) ClearEntityID() *CategoryAssignmentUpdateOne {
			cauo.mutation.ClearEntityID()
			return cauo
		}
	

	
	
	// SetEntityType sets the "entity_type" field.
	func (cauo *CategoryAssignmentUpdateOne) SetEntityType(s string) *CategoryAssignmentUpdateOne {
		cauo.mutation.SetEntityType(s)
		return cauo
	}

	
	
	
	
	
	
		// SetNillableEntityType sets the "entity_type" field if the given value is not nil.
		func (cauo *CategoryAssignmentUpdateOne) SetNillableEntityType(s *string) *CategoryAssignmentUpdateOne {
			if s != nil {
				cauo.SetEntityType(*s)
			}
			return cauo
		}
	

	

	

	
		
		// ClearEntityType clears the value of the "entity_type" field.
		func (cauo *CategoryAssignmentUpdateOne) ClearEntityType() *CategoryAssignmentUpdateOne {
			cauo.mutation.ClearEntityType()
			return cauo
		}
	

	
	
	// SetCategoryID sets the "category_id" field.
	func (cauo *CategoryAssignmentUpdateOne) SetCategoryID(s string) *CategoryAssignmentUpdateOne {
		cauo.mutation.SetCategoryID(s)
		return cauo
	}

	
	
	
	
	
	
		// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
		func (cauo *CategoryAssignmentUpdateOne) SetNillableCategoryID(s *string) *CategoryAssignmentUpdateOne {
			if s != nil {
				cauo.SetCategoryID(*s)
			}
			return cauo
		}
	

	

	

	
		
		// ClearCategoryID clears the value of the "category_id" field.
		func (cauo *CategoryAssignmentUpdateOne) ClearCategoryID() *CategoryAssignmentUpdateOne {
			cauo.mutation.ClearCategoryID()
			return cauo
		}
	



	
	
	
	
	
		// SetUserID sets the "user" edge to the User entity by ID.
		func (cauo *CategoryAssignmentUpdateOne) SetUserID(id string) *CategoryAssignmentUpdateOne {
			cauo.mutation.SetUserID(id)
			return cauo
		}
	
	
		
		// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
		func (cauo *CategoryAssignmentUpdateOne) SetNillableUserID(id *string) *CategoryAssignmentUpdateOne {
			if id != nil {
				cauo = cauo.SetUserID(*id)
			}
			return cauo
		}
	
	
	
	
	// SetUser sets the "user" edge to the User entity.
	func (cauo *CategoryAssignmentUpdateOne) SetUser(u *User) *CategoryAssignmentUpdateOne {
		return cauo.SetUserID(u.ID)
	}

	
	
	
	
	
		// SetBusinessID sets the "business" edge to the Business entity by ID.
		func (cauo *CategoryAssignmentUpdateOne) SetBusinessID(id string) *CategoryAssignmentUpdateOne {
			cauo.mutation.SetBusinessID(id)
			return cauo
		}
	
	
		
		// SetNillableBusinessID sets the "business" edge to the Business entity by ID if the given value is not nil.
		func (cauo *CategoryAssignmentUpdateOne) SetNillableBusinessID(id *string) *CategoryAssignmentUpdateOne {
			if id != nil {
				cauo = cauo.SetBusinessID(*id)
			}
			return cauo
		}
	
	
	
	
	// SetBusiness sets the "business" edge to the Business entity.
	func (cauo *CategoryAssignmentUpdateOne) SetBusiness(b *Business) *CategoryAssignmentUpdateOne {
		return cauo.SetBusinessID(b.ID)
	}

	
	
	
	
	
		// SetPlaceID sets the "place" edge to the Place entity by ID.
		func (cauo *CategoryAssignmentUpdateOne) SetPlaceID(id string) *CategoryAssignmentUpdateOne {
			cauo.mutation.SetPlaceID(id)
			return cauo
		}
	
	
		
		// SetNillablePlaceID sets the "place" edge to the Place entity by ID if the given value is not nil.
		func (cauo *CategoryAssignmentUpdateOne) SetNillablePlaceID(id *string) *CategoryAssignmentUpdateOne {
			if id != nil {
				cauo = cauo.SetPlaceID(*id)
			}
			return cauo
		}
	
	
	
	
	// SetPlace sets the "place" edge to the Place entity.
	func (cauo *CategoryAssignmentUpdateOne) SetPlace(p *Place) *CategoryAssignmentUpdateOne {
		return cauo.SetPlaceID(p.ID)
	}

	
	
	
	
	
	
	
	
	
	// SetCategory sets the "category" edge to the Category entity.
	func (cauo *CategoryAssignmentUpdateOne) SetCategory(c *Category) *CategoryAssignmentUpdateOne {
		return cauo.SetCategoryID(c.ID)
	}


// Mutation returns the CategoryAssignmentMutation object of the builder.
func (cauo *CategoryAssignmentUpdateOne) Mutation() *CategoryAssignmentMutation {
	return cauo.mutation
}






	





	
	
	// ClearUser clears the "user" edge to the User entity.
	func (cauo *CategoryAssignmentUpdateOne) ClearUser() *CategoryAssignmentUpdateOne {
		cauo.mutation.ClearUser()
		return cauo
	}
	

	
	
	// ClearBusiness clears the "business" edge to the Business entity.
	func (cauo *CategoryAssignmentUpdateOne) ClearBusiness() *CategoryAssignmentUpdateOne {
		cauo.mutation.ClearBusiness()
		return cauo
	}
	

	
	
	// ClearPlace clears the "place" edge to the Place entity.
	func (cauo *CategoryAssignmentUpdateOne) ClearPlace() *CategoryAssignmentUpdateOne {
		cauo.mutation.ClearPlace()
		return cauo
	}
	

	
	
	// ClearCategory clears the "category" edge to the Category entity.
	func (cauo *CategoryAssignmentUpdateOne) ClearCategory() *CategoryAssignmentUpdateOne {
		cauo.mutation.ClearCategory()
		return cauo
	}
	




// Where appends a list predicates to the CategoryAssignmentUpdate builder.
func (cauo *CategoryAssignmentUpdateOne) Where(ps ...predicate.CategoryAssignment) *CategoryAssignmentUpdateOne {
	cauo.mutation.Where(ps...)
	return cauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cauo *CategoryAssignmentUpdateOne) Select(field string, fields ...string) *CategoryAssignmentUpdateOne {
	cauo.fields = append([]string{field}, fields...)
	return cauo
}

// Save executes the query and returns the updated CategoryAssignment entity.
func (cauo *CategoryAssignmentUpdateOne ) Save(ctx context.Context) (*CategoryAssignment, error) {
	return withHooks(ctx, cauo.sqlSave, cauo.mutation, cauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cauo *CategoryAssignmentUpdateOne) SaveX(ctx context.Context) *CategoryAssignment {
	node, err := cauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cauo *CategoryAssignmentUpdateOne) Exec(ctx context.Context) error {
	_, err := cauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cauo *CategoryAssignmentUpdateOne) ExecX(ctx context.Context) {
	if err := cauo.Exec(ctx); err != nil {
		panic(err)
	}
}


	














	
	





    


func (cauo *CategoryAssignmentUpdateOne) sqlSave(ctx context.Context) (_node *CategoryAssignment, err error) {
	_spec := sqlgraph.NewUpdateSpec(categoryassignment.Table, categoryassignment.Columns,sqlgraph.NewFieldSpec(categoryassignment.FieldID, field.TypeString))
			id, ok := cauo.mutation.ID()
			if !ok {
				return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CategoryAssignment.id" for update`)}
			}
			_spec.Node.ID.Value = id
			if fields := cauo.fields; len(fields) > 0 {
				_spec.Node.Columns = make([]string, 0, len(fields))
				_spec.Node.Columns = append(_spec.Node.Columns, categoryassignment.FieldID)
				for _, f := range fields {
					if !categoryassignment.ValidColumn(f) {
						return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
					}
					if f != categoryassignment.FieldID {
						_spec.Node.Columns = append(_spec.Node.Columns, f)
					}
				}
			}
	if ps := cauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
				if value, ok := cauo.mutation.EntityType(); ok {
						_spec.SetField(categoryassignment.FieldEntityType, field.TypeString, value)
				}
				if cauo.mutation.EntityTypeCleared() {
					_spec.ClearField(categoryassignment.FieldEntityType, field.TypeString)
				}
		if cauo.mutation.UserCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.UserTable,
		Columns: []string{ categoryassignment.UserColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
		if nodes := cauo.mutation.UserIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.UserTable,
		Columns: []string{ categoryassignment.UserColumn },
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
		if cauo.mutation.BusinessCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.BusinessTable,
		Columns: []string{ categoryassignment.BusinessColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
		if nodes := cauo.mutation.BusinessIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.BusinessTable,
		Columns: []string{ categoryassignment.BusinessColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges.Add = append(_spec.Edges.Add, edge)
		}
		if cauo.mutation.PlaceCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.PlaceTable,
		Columns: []string{ categoryassignment.PlaceColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
		if nodes := cauo.mutation.PlaceIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.PlaceTable,
		Columns: []string{ categoryassignment.PlaceColumn },
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
		if cauo.mutation.CategoryCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.CategoryTable,
		Columns: []string{ categoryassignment.CategoryColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
		if nodes := cauo.mutation.CategoryIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: categoryassignment.CategoryTable,
		Columns: []string{ categoryassignment.CategoryColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges.Add = append(_spec.Edges.Add, edge)
		}
		_node = &CategoryAssignment{config: cauo.config}
		_spec.Assign = _node.assignValues
		_spec.ScanValues = _node.scanValues
		if err = sqlgraph.UpdateNode(ctx, cauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ categoryassignment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cauo.mutation.done = true
	return _node, nil
}



