// Code generated by ent, DO NOT EDIT.

package placio_api

import (
	"context"
	"errors"
	"fmt"
	"placio_api/featurerelease"
	"placio_api/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FeatureReleaseUpdate is the builder for updating FeatureRelease entities.
type FeatureReleaseUpdate struct {
	config
	hooks    []Hook
	mutation *FeatureReleaseMutation
}

// Where appends a list predicates to the FeatureReleaseUpdate builder.
func (fru *FeatureReleaseUpdate) Where(ps ...predicate.FeatureRelease) *FeatureReleaseUpdate {
	fru.mutation.Where(ps...)
	return fru
}

// SetFeatureName sets the "feature_name" field.
func (fru *FeatureReleaseUpdate) SetFeatureName(s string) *FeatureReleaseUpdate {
	fru.mutation.SetFeatureName(s)
	return fru
}

// SetDescription sets the "description" field.
func (fru *FeatureReleaseUpdate) SetDescription(s string) *FeatureReleaseUpdate {
	fru.mutation.SetDescription(s)
	return fru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fru *FeatureReleaseUpdate) SetNillableDescription(s *string) *FeatureReleaseUpdate {
	if s != nil {
		fru.SetDescription(*s)
	}
	return fru
}

// ClearDescription clears the value of the "description" field.
func (fru *FeatureReleaseUpdate) ClearDescription() *FeatureReleaseUpdate {
	fru.mutation.ClearDescription()
	return fru
}

// SetState sets the "state" field.
func (fru *FeatureReleaseUpdate) SetState(f featurerelease.State) *FeatureReleaseUpdate {
	fru.mutation.SetState(f)
	return fru
}

// SetReleaseDate sets the "release_date" field.
func (fru *FeatureReleaseUpdate) SetReleaseDate(t time.Time) *FeatureReleaseUpdate {
	fru.mutation.SetReleaseDate(t)
	return fru
}

// SetNillableReleaseDate sets the "release_date" field if the given value is not nil.
func (fru *FeatureReleaseUpdate) SetNillableReleaseDate(t *time.Time) *FeatureReleaseUpdate {
	if t != nil {
		fru.SetReleaseDate(*t)
	}
	return fru
}

// SetEligibilityRules sets the "eligibility_rules" field.
func (fru *FeatureReleaseUpdate) SetEligibilityRules(m map[string]interface{}) *FeatureReleaseUpdate {
	fru.mutation.SetEligibilityRules(m)
	return fru
}

// ClearEligibilityRules clears the value of the "eligibility_rules" field.
func (fru *FeatureReleaseUpdate) ClearEligibilityRules() *FeatureReleaseUpdate {
	fru.mutation.ClearEligibilityRules()
	return fru
}

// SetDocumentationLink sets the "documentation_link" field.
func (fru *FeatureReleaseUpdate) SetDocumentationLink(s string) *FeatureReleaseUpdate {
	fru.mutation.SetDocumentationLink(s)
	return fru
}

// SetNillableDocumentationLink sets the "documentation_link" field if the given value is not nil.
func (fru *FeatureReleaseUpdate) SetNillableDocumentationLink(s *string) *FeatureReleaseUpdate {
	if s != nil {
		fru.SetDocumentationLink(*s)
	}
	return fru
}

// ClearDocumentationLink clears the value of the "documentation_link" field.
func (fru *FeatureReleaseUpdate) ClearDocumentationLink() *FeatureReleaseUpdate {
	fru.mutation.ClearDocumentationLink()
	return fru
}

// SetMetadata sets the "metadata" field.
func (fru *FeatureReleaseUpdate) SetMetadata(m map[string]interface{}) *FeatureReleaseUpdate {
	fru.mutation.SetMetadata(m)
	return fru
}

// ClearMetadata clears the value of the "metadata" field.
func (fru *FeatureReleaseUpdate) ClearMetadata() *FeatureReleaseUpdate {
	fru.mutation.ClearMetadata()
	return fru
}

// Mutation returns the FeatureReleaseMutation object of the builder.
func (fru *FeatureReleaseUpdate) Mutation() *FeatureReleaseMutation {
	return fru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fru *FeatureReleaseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fru.sqlSave, fru.mutation, fru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fru *FeatureReleaseUpdate) SaveX(ctx context.Context) int {
	affected, err := fru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fru *FeatureReleaseUpdate) Exec(ctx context.Context) error {
	_, err := fru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fru *FeatureReleaseUpdate) ExecX(ctx context.Context) {
	if err := fru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fru *FeatureReleaseUpdate) check() error {
	if v, ok := fru.mutation.State(); ok {
		if err := featurerelease.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`placio_api: validator failed for field "FeatureRelease.state": %w`, err)}
		}
	}
	return nil
}

func (fru *FeatureReleaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(featurerelease.Table, featurerelease.Columns, sqlgraph.NewFieldSpec(featurerelease.FieldID, field.TypeString))
	if ps := fru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fru.mutation.FeatureName(); ok {
		_spec.SetField(featurerelease.FieldFeatureName, field.TypeString, value)
	}
	if value, ok := fru.mutation.Description(); ok {
		_spec.SetField(featurerelease.FieldDescription, field.TypeString, value)
	}
	if fru.mutation.DescriptionCleared() {
		_spec.ClearField(featurerelease.FieldDescription, field.TypeString)
	}
	if value, ok := fru.mutation.State(); ok {
		_spec.SetField(featurerelease.FieldState, field.TypeEnum, value)
	}
	if value, ok := fru.mutation.ReleaseDate(); ok {
		_spec.SetField(featurerelease.FieldReleaseDate, field.TypeTime, value)
	}
	if value, ok := fru.mutation.EligibilityRules(); ok {
		_spec.SetField(featurerelease.FieldEligibilityRules, field.TypeJSON, value)
	}
	if fru.mutation.EligibilityRulesCleared() {
		_spec.ClearField(featurerelease.FieldEligibilityRules, field.TypeJSON)
	}
	if value, ok := fru.mutation.DocumentationLink(); ok {
		_spec.SetField(featurerelease.FieldDocumentationLink, field.TypeString, value)
	}
	if fru.mutation.DocumentationLinkCleared() {
		_spec.ClearField(featurerelease.FieldDocumentationLink, field.TypeString)
	}
	if value, ok := fru.mutation.Metadata(); ok {
		_spec.SetField(featurerelease.FieldMetadata, field.TypeJSON, value)
	}
	if fru.mutation.MetadataCleared() {
		_spec.ClearField(featurerelease.FieldMetadata, field.TypeJSON)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{featurerelease.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fru.mutation.done = true
	return n, nil
}

// FeatureReleaseUpdateOne is the builder for updating a single FeatureRelease entity.
type FeatureReleaseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FeatureReleaseMutation
}

// SetFeatureName sets the "feature_name" field.
func (fruo *FeatureReleaseUpdateOne) SetFeatureName(s string) *FeatureReleaseUpdateOne {
	fruo.mutation.SetFeatureName(s)
	return fruo
}

// SetDescription sets the "description" field.
func (fruo *FeatureReleaseUpdateOne) SetDescription(s string) *FeatureReleaseUpdateOne {
	fruo.mutation.SetDescription(s)
	return fruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fruo *FeatureReleaseUpdateOne) SetNillableDescription(s *string) *FeatureReleaseUpdateOne {
	if s != nil {
		fruo.SetDescription(*s)
	}
	return fruo
}

// ClearDescription clears the value of the "description" field.
func (fruo *FeatureReleaseUpdateOne) ClearDescription() *FeatureReleaseUpdateOne {
	fruo.mutation.ClearDescription()
	return fruo
}

// SetState sets the "state" field.
func (fruo *FeatureReleaseUpdateOne) SetState(f featurerelease.State) *FeatureReleaseUpdateOne {
	fruo.mutation.SetState(f)
	return fruo
}

// SetReleaseDate sets the "release_date" field.
func (fruo *FeatureReleaseUpdateOne) SetReleaseDate(t time.Time) *FeatureReleaseUpdateOne {
	fruo.mutation.SetReleaseDate(t)
	return fruo
}

// SetNillableReleaseDate sets the "release_date" field if the given value is not nil.
func (fruo *FeatureReleaseUpdateOne) SetNillableReleaseDate(t *time.Time) *FeatureReleaseUpdateOne {
	if t != nil {
		fruo.SetReleaseDate(*t)
	}
	return fruo
}

// SetEligibilityRules sets the "eligibility_rules" field.
func (fruo *FeatureReleaseUpdateOne) SetEligibilityRules(m map[string]interface{}) *FeatureReleaseUpdateOne {
	fruo.mutation.SetEligibilityRules(m)
	return fruo
}

// ClearEligibilityRules clears the value of the "eligibility_rules" field.
func (fruo *FeatureReleaseUpdateOne) ClearEligibilityRules() *FeatureReleaseUpdateOne {
	fruo.mutation.ClearEligibilityRules()
	return fruo
}

// SetDocumentationLink sets the "documentation_link" field.
func (fruo *FeatureReleaseUpdateOne) SetDocumentationLink(s string) *FeatureReleaseUpdateOne {
	fruo.mutation.SetDocumentationLink(s)
	return fruo
}

// SetNillableDocumentationLink sets the "documentation_link" field if the given value is not nil.
func (fruo *FeatureReleaseUpdateOne) SetNillableDocumentationLink(s *string) *FeatureReleaseUpdateOne {
	if s != nil {
		fruo.SetDocumentationLink(*s)
	}
	return fruo
}

// ClearDocumentationLink clears the value of the "documentation_link" field.
func (fruo *FeatureReleaseUpdateOne) ClearDocumentationLink() *FeatureReleaseUpdateOne {
	fruo.mutation.ClearDocumentationLink()
	return fruo
}

// SetMetadata sets the "metadata" field.
func (fruo *FeatureReleaseUpdateOne) SetMetadata(m map[string]interface{}) *FeatureReleaseUpdateOne {
	fruo.mutation.SetMetadata(m)
	return fruo
}

// ClearMetadata clears the value of the "metadata" field.
func (fruo *FeatureReleaseUpdateOne) ClearMetadata() *FeatureReleaseUpdateOne {
	fruo.mutation.ClearMetadata()
	return fruo
}

// Mutation returns the FeatureReleaseMutation object of the builder.
func (fruo *FeatureReleaseUpdateOne) Mutation() *FeatureReleaseMutation {
	return fruo.mutation
}

// Where appends a list predicates to the FeatureReleaseUpdate builder.
func (fruo *FeatureReleaseUpdateOne) Where(ps ...predicate.FeatureRelease) *FeatureReleaseUpdateOne {
	fruo.mutation.Where(ps...)
	return fruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fruo *FeatureReleaseUpdateOne) Select(field string, fields ...string) *FeatureReleaseUpdateOne {
	fruo.fields = append([]string{field}, fields...)
	return fruo
}

// Save executes the query and returns the updated FeatureRelease entity.
func (fruo *FeatureReleaseUpdateOne) Save(ctx context.Context) (*FeatureRelease, error) {
	return withHooks(ctx, fruo.sqlSave, fruo.mutation, fruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fruo *FeatureReleaseUpdateOne) SaveX(ctx context.Context) *FeatureRelease {
	node, err := fruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fruo *FeatureReleaseUpdateOne) Exec(ctx context.Context) error {
	_, err := fruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fruo *FeatureReleaseUpdateOne) ExecX(ctx context.Context) {
	if err := fruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fruo *FeatureReleaseUpdateOne) check() error {
	if v, ok := fruo.mutation.State(); ok {
		if err := featurerelease.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`placio_api: validator failed for field "FeatureRelease.state": %w`, err)}
		}
	}
	return nil
}

func (fruo *FeatureReleaseUpdateOne) sqlSave(ctx context.Context) (_node *FeatureRelease, err error) {
	if err := fruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(featurerelease.Table, featurerelease.Columns, sqlgraph.NewFieldSpec(featurerelease.FieldID, field.TypeString))
	id, ok := fruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`placio_api: missing "FeatureRelease.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, featurerelease.FieldID)
		for _, f := range fields {
			if !featurerelease.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("placio_api: invalid field %q for query", f)}
			}
			if f != featurerelease.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fruo.mutation.FeatureName(); ok {
		_spec.SetField(featurerelease.FieldFeatureName, field.TypeString, value)
	}
	if value, ok := fruo.mutation.Description(); ok {
		_spec.SetField(featurerelease.FieldDescription, field.TypeString, value)
	}
	if fruo.mutation.DescriptionCleared() {
		_spec.ClearField(featurerelease.FieldDescription, field.TypeString)
	}
	if value, ok := fruo.mutation.State(); ok {
		_spec.SetField(featurerelease.FieldState, field.TypeEnum, value)
	}
	if value, ok := fruo.mutation.ReleaseDate(); ok {
		_spec.SetField(featurerelease.FieldReleaseDate, field.TypeTime, value)
	}
	if value, ok := fruo.mutation.EligibilityRules(); ok {
		_spec.SetField(featurerelease.FieldEligibilityRules, field.TypeJSON, value)
	}
	if fruo.mutation.EligibilityRulesCleared() {
		_spec.ClearField(featurerelease.FieldEligibilityRules, field.TypeJSON)
	}
	if value, ok := fruo.mutation.DocumentationLink(); ok {
		_spec.SetField(featurerelease.FieldDocumentationLink, field.TypeString, value)
	}
	if fruo.mutation.DocumentationLinkCleared() {
		_spec.ClearField(featurerelease.FieldDocumentationLink, field.TypeString)
	}
	if value, ok := fruo.mutation.Metadata(); ok {
		_spec.SetField(featurerelease.FieldMetadata, field.TypeJSON, value)
	}
	if fruo.mutation.MetadataCleared() {
		_spec.ClearField(featurerelease.FieldMetadata, field.TypeJSON)
	}
	_node = &FeatureRelease{config: fruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{featurerelease.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fruo.mutation.done = true
	return _node, nil
}
