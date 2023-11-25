




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
			 "placio-app/ent/template"
			 "placio-app/ent/website"

)







// TemplateCreate is the builder for creating a Template entity.
type TemplateCreate struct {
	config
	mutation *TemplateMutation
	hooks []Hook
}


	







	
	
	// SetName sets the "name" field.
	func (tc *TemplateCreate) SetName(s string) *TemplateCreate {
		tc.mutation.SetName(s)
		return tc
	}

	
	
	
	
	
	

	

	

	

	
	
	// SetDefaultHTML sets the "defaultHTML" field.
	func (tc *TemplateCreate) SetDefaultHTML(s string) *TemplateCreate {
		tc.mutation.SetDefaultHTML(s)
		return tc
	}

	
	
	
	
	
	

	

	

	

	
	
	// SetDefaultCSS sets the "defaultCSS" field.
	func (tc *TemplateCreate) SetDefaultCSS(s string) *TemplateCreate {
		tc.mutation.SetDefaultCSS(s)
		return tc
	}

	
	
	
	
	
	

	

	

	



	
	
	
	
	
		// AddWebsiteIDs adds the "websites" edge to the Website entity by IDs.
		func (tc *TemplateCreate) AddWebsiteIDs(ids ... string) *TemplateCreate {
			tc.mutation.AddWebsiteIDs(ids ...)
			return tc
		}
	
	
	
	
	
	// AddWebsites adds the "websites" edges to the Website entity.
	func (tc *TemplateCreate) AddWebsites(w ...*Website) *TemplateCreate {
		ids := make([]string, len(w))
			for i := range w {
				ids[i] = w[i].ID
			}
			return tc.AddWebsiteIDs(ids...)
	}


// Mutation returns the TemplateMutation object of the builder.
func (tc *TemplateCreate) Mutation() *TemplateMutation {
	return tc.mutation
}




// Save creates the Template in the database.
func (tc *TemplateCreate) Save(ctx context.Context) (*Template, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TemplateCreate) SaveX(ctx context.Context) *Template {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TemplateCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TemplateCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}


// check runs all checks and user-defined validators on the builder.
func (tc *TemplateCreate) check() error {
					if _, ok := tc.mutation.Name(); !ok {
						return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Template.name"`)}
					}
					if _, ok := tc.mutation.DefaultHTML(); !ok {
						return &ValidationError{Name: "defaultHTML", err: errors.New(`ent: missing required field "Template.defaultHTML"`)}
					}
					if _, ok := tc.mutation.DefaultCSS(); !ok {
						return &ValidationError{Name: "defaultCSS", err: errors.New(`ent: missing required field "Template.defaultCSS"`)}
					}
	return nil
}


	
	




func (tc *TemplateCreate) sqlSave(ctx context.Context) (*Template, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec  := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
		if _spec.ID.Value != nil {
				if id, ok := _spec.ID.Value.(string); ok {
					_node.ID = id
				} else {
					return nil, fmt.Errorf("unexpected Template.ID type: %T", _spec.ID.Value)
				}
		}
		tc.mutation.id = &_node.ID
		tc.mutation.done = true
	return _node, nil
}

func (tc *TemplateCreate) createSpec() (*Template, *sqlgraph.CreateSpec) {
	var (
		_node = &Template{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(template.Table, sqlgraph.NewFieldSpec(template.FieldID, field.TypeString))
	)
		if value, ok := tc.mutation.Name(); ok {
				_spec.SetField(template.FieldName, field.TypeString, value)
			_node.Name = value
		}
		if value, ok := tc.mutation.DefaultHTML(); ok {
				_spec.SetField(template.FieldDefaultHTML, field.TypeString, value)
			_node.DefaultHTML = value
		}
		if value, ok := tc.mutation.DefaultCSS(); ok {
				_spec.SetField(template.FieldDefaultCSS, field.TypeString, value)
			_node.DefaultCSS = value
		}
		if nodes := tc.mutation.WebsitesIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.O2M,
		Inverse: false,
		Table: template.WebsitesTable,
		Columns: []string{ template.WebsitesColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(website.FieldID, field.TypeString),
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges = append(_spec.Edges, edge)
		}
	return _node, _spec
}
	








// TemplateCreateBulk is the builder for creating many Template entities in bulk.
type TemplateCreateBulk struct {
	config
	err error
	builders []*TemplateCreate
}




	
		



// Save creates the Template entities in the database.
func (tcb *TemplateCreateBulk) Save(ctx context.Context) ([]*Template, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Template, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TemplateMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TemplateCreateBulk) SaveX(ctx context.Context) []*Template {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TemplateCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
	


	

