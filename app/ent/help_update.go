



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
			 "placio-app/ent/help"
			 "placio-app/ent/user"

)








// HelpUpdate is the builder for updating Help entities.
type HelpUpdate struct {
	config
	hooks []Hook
	mutation *HelpMutation
}

// Where appends a list predicates to the HelpUpdate builder.
func (hu *HelpUpdate) Where(ps ...predicate.Help) *HelpUpdate {
	hu.mutation.Where(ps...)
	return hu
}


	





	
	
	


	
	
	// SetCategory sets the "category" field.
	func (hu *HelpUpdate) SetCategory(s string) *HelpUpdate {
		hu.mutation.SetCategory(s)
		return hu
	}

	
	
	
	
	
	
		// SetNillableCategory sets the "category" field if the given value is not nil.
		func (hu *HelpUpdate) SetNillableCategory(s *string) *HelpUpdate {
			if s != nil {
				hu.SetCategory(*s)
			}
			return hu
		}
	

	

	

	

	
	
	// SetSubject sets the "subject" field.
	func (hu *HelpUpdate) SetSubject(s string) *HelpUpdate {
		hu.mutation.SetSubject(s)
		return hu
	}

	
	
	
	
	
	
		// SetNillableSubject sets the "subject" field if the given value is not nil.
		func (hu *HelpUpdate) SetNillableSubject(s *string) *HelpUpdate {
			if s != nil {
				hu.SetSubject(*s)
			}
			return hu
		}
	

	

	

	

	
	
	// SetBody sets the "body" field.
	func (hu *HelpUpdate) SetBody(s string) *HelpUpdate {
		hu.mutation.SetBody(s)
		return hu
	}

	
	
	
	
	
	
		// SetNillableBody sets the "body" field if the given value is not nil.
		func (hu *HelpUpdate) SetNillableBody(s *string) *HelpUpdate {
			if s != nil {
				hu.SetBody(*s)
			}
			return hu
		}
	

	

	

	

	
	
	// SetMedia sets the "media" field.
	func (hu *HelpUpdate) SetMedia(s string) *HelpUpdate {
		hu.mutation.SetMedia(s)
		return hu
	}

	
	
	
	
	
	
		// SetNillableMedia sets the "media" field if the given value is not nil.
		func (hu *HelpUpdate) SetNillableMedia(s *string) *HelpUpdate {
			if s != nil {
				hu.SetMedia(*s)
			}
			return hu
		}
	

	

	

	
		
		// ClearMedia clears the value of the "media" field.
		func (hu *HelpUpdate) ClearMedia() *HelpUpdate {
			hu.mutation.ClearMedia()
			return hu
		}
	

	
	
	// SetStatus sets the "status" field.
	func (hu *HelpUpdate) SetStatus(s string) *HelpUpdate {
		hu.mutation.SetStatus(s)
		return hu
	}

	
	
	
	
	
	
		// SetNillableStatus sets the "status" field if the given value is not nil.
		func (hu *HelpUpdate) SetNillableStatus(s *string) *HelpUpdate {
			if s != nil {
				hu.SetStatus(*s)
			}
			return hu
		}
	

	

	

	



	
		
		

// Mutation returns the HelpMutation object of the builder.
func (hu *HelpUpdate) Mutation() *HelpMutation {
	return hu.mutation
}





	





	
		
		



// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HelpUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, hu.sqlSave, hu.mutation, hu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HelpUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HelpUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HelpUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}


	









	// check runs all checks and user-defined validators on the builder.
	func (hu *HelpUpdate) check() error {
				if _, ok := hu.mutation.UserID(); hu.mutation.UserCleared() && !ok {
					return errors.New(`ent: clearing a required unique edge "Help.user"`)
				}
		return nil
	}






	
	





    


func (hu *HelpUpdate) sqlSave(ctx context.Context) (n int, err error) {
		if err := hu.check(); err != nil {
			return n, err
		}
	_spec := sqlgraph.NewUpdateSpec(help.Table, help.Columns,sqlgraph.NewFieldSpec(help.FieldID, field.TypeString))
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
				if value, ok := hu.mutation.Category(); ok {
						_spec.SetField(help.FieldCategory, field.TypeString, value)
				}
				if value, ok := hu.mutation.Subject(); ok {
						_spec.SetField(help.FieldSubject, field.TypeString, value)
				}
				if value, ok := hu.mutation.Body(); ok {
						_spec.SetField(help.FieldBody, field.TypeString, value)
				}
				if value, ok := hu.mutation.Media(); ok {
						_spec.SetField(help.FieldMedia, field.TypeString, value)
				}
				if hu.mutation.MediaCleared() {
					_spec.ClearField(help.FieldMedia, field.TypeString)
				}
				if value, ok := hu.mutation.Status(); ok {
						_spec.SetField(help.FieldStatus, field.TypeString, value)
				}
		if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ help.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hu.mutation.done = true
	return n, nil
}







// HelpUpdateOne is the builder for updating a single Help entity.
type HelpUpdateOne struct {
	config
	fields []string
	hooks []Hook
	mutation *HelpMutation

}


	





	
	
	


	
	
	// SetCategory sets the "category" field.
	func (huo *HelpUpdateOne) SetCategory(s string) *HelpUpdateOne {
		huo.mutation.SetCategory(s)
		return huo
	}

	
	
	
	
	
	
		// SetNillableCategory sets the "category" field if the given value is not nil.
		func (huo *HelpUpdateOne) SetNillableCategory(s *string) *HelpUpdateOne {
			if s != nil {
				huo.SetCategory(*s)
			}
			return huo
		}
	

	

	

	

	
	
	// SetSubject sets the "subject" field.
	func (huo *HelpUpdateOne) SetSubject(s string) *HelpUpdateOne {
		huo.mutation.SetSubject(s)
		return huo
	}

	
	
	
	
	
	
		// SetNillableSubject sets the "subject" field if the given value is not nil.
		func (huo *HelpUpdateOne) SetNillableSubject(s *string) *HelpUpdateOne {
			if s != nil {
				huo.SetSubject(*s)
			}
			return huo
		}
	

	

	

	

	
	
	// SetBody sets the "body" field.
	func (huo *HelpUpdateOne) SetBody(s string) *HelpUpdateOne {
		huo.mutation.SetBody(s)
		return huo
	}

	
	
	
	
	
	
		// SetNillableBody sets the "body" field if the given value is not nil.
		func (huo *HelpUpdateOne) SetNillableBody(s *string) *HelpUpdateOne {
			if s != nil {
				huo.SetBody(*s)
			}
			return huo
		}
	

	

	

	

	
	
	// SetMedia sets the "media" field.
	func (huo *HelpUpdateOne) SetMedia(s string) *HelpUpdateOne {
		huo.mutation.SetMedia(s)
		return huo
	}

	
	
	
	
	
	
		// SetNillableMedia sets the "media" field if the given value is not nil.
		func (huo *HelpUpdateOne) SetNillableMedia(s *string) *HelpUpdateOne {
			if s != nil {
				huo.SetMedia(*s)
			}
			return huo
		}
	

	

	

	
		
		// ClearMedia clears the value of the "media" field.
		func (huo *HelpUpdateOne) ClearMedia() *HelpUpdateOne {
			huo.mutation.ClearMedia()
			return huo
		}
	

	
	
	// SetStatus sets the "status" field.
	func (huo *HelpUpdateOne) SetStatus(s string) *HelpUpdateOne {
		huo.mutation.SetStatus(s)
		return huo
	}

	
	
	
	
	
	
		// SetNillableStatus sets the "status" field if the given value is not nil.
		func (huo *HelpUpdateOne) SetNillableStatus(s *string) *HelpUpdateOne {
			if s != nil {
				huo.SetStatus(*s)
			}
			return huo
		}
	

	

	

	



	
		
		

// Mutation returns the HelpMutation object of the builder.
func (huo *HelpUpdateOne) Mutation() *HelpMutation {
	return huo.mutation
}






	





	
		
		



// Where appends a list predicates to the HelpUpdate builder.
func (huo *HelpUpdateOne) Where(ps ...predicate.Help) *HelpUpdateOne {
	huo.mutation.Where(ps...)
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HelpUpdateOne) Select(field string, fields ...string) *HelpUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated Help entity.
func (huo *HelpUpdateOne ) Save(ctx context.Context) (*Help, error) {
	return withHooks(ctx, huo.sqlSave, huo.mutation, huo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HelpUpdateOne) SaveX(ctx context.Context) *Help {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HelpUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HelpUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}


	









	// check runs all checks and user-defined validators on the builder.
	func (huo *HelpUpdateOne) check() error {
				if _, ok := huo.mutation.UserID(); huo.mutation.UserCleared() && !ok {
					return errors.New(`ent: clearing a required unique edge "Help.user"`)
				}
		return nil
	}






	
	





    


func (huo *HelpUpdateOne) sqlSave(ctx context.Context) (_node *Help, err error) {
		if err := huo.check(); err != nil {
			return _node, err
		}
	_spec := sqlgraph.NewUpdateSpec(help.Table, help.Columns,sqlgraph.NewFieldSpec(help.FieldID, field.TypeString))
			id, ok := huo.mutation.ID()
			if !ok {
				return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Help.id" for update`)}
			}
			_spec.Node.ID.Value = id
			if fields := huo.fields; len(fields) > 0 {
				_spec.Node.Columns = make([]string, 0, len(fields))
				_spec.Node.Columns = append(_spec.Node.Columns, help.FieldID)
				for _, f := range fields {
					if !help.ValidColumn(f) {
						return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
					}
					if f != help.FieldID {
						_spec.Node.Columns = append(_spec.Node.Columns, f)
					}
				}
			}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
				if value, ok := huo.mutation.Category(); ok {
						_spec.SetField(help.FieldCategory, field.TypeString, value)
				}
				if value, ok := huo.mutation.Subject(); ok {
						_spec.SetField(help.FieldSubject, field.TypeString, value)
				}
				if value, ok := huo.mutation.Body(); ok {
						_spec.SetField(help.FieldBody, field.TypeString, value)
				}
				if value, ok := huo.mutation.Media(); ok {
						_spec.SetField(help.FieldMedia, field.TypeString, value)
				}
				if huo.mutation.MediaCleared() {
					_spec.ClearField(help.FieldMedia, field.TypeString)
				}
				if value, ok := huo.mutation.Status(); ok {
						_spec.SetField(help.FieldStatus, field.TypeString, value)
				}
		_node = &Help{config: huo.config}
		_spec.Assign = _node.assignValues
		_spec.ScanValues = _node.scanValues
		if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ help.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	huo.mutation.done = true
	return _node, nil
}



