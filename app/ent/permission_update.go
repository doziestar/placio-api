// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/permission"
	"placio-app/ent/predicate"
	"placio-app/ent/staff"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks    []Hook
	mutation *PermissionMutation
}

// Where appends a list predicates to the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PermissionUpdate) SetName(s string) *PermissionUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetDescription sets the "description" field.
func (pu *PermissionUpdate) SetDescription(s string) *PermissionUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableDescription(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *PermissionUpdate) ClearDescription() *PermissionUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// AddStaffIDs adds the "staffs" edge to the Staff entity by IDs.
func (pu *PermissionUpdate) AddStaffIDs(ids ...string) *PermissionUpdate {
	pu.mutation.AddStaffIDs(ids...)
	return pu
}

// AddStaffs adds the "staffs" edges to the Staff entity.
func (pu *PermissionUpdate) AddStaffs(s ...*Staff) *PermissionUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddStaffIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// ClearStaffs clears all "staffs" edges to the Staff entity.
func (pu *PermissionUpdate) ClearStaffs() *PermissionUpdate {
	pu.mutation.ClearStaffs()
	return pu
}

// RemoveStaffIDs removes the "staffs" edge to Staff entities by IDs.
func (pu *PermissionUpdate) RemoveStaffIDs(ids ...string) *PermissionUpdate {
	pu.mutation.RemoveStaffIDs(ids...)
	return pu
}

// RemoveStaffs removes "staffs" edges to Staff entities.
func (pu *PermissionUpdate) RemoveStaffs(s ...*Staff) *PermissionUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveStaffIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeString))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(permission.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(permission.FieldDescription, field.TypeString)
	}
	if pu.mutation.StaffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.StaffsTable,
			Columns: permission.StaffsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedStaffsIDs(); len(nodes) > 0 && !pu.mutation.StaffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.StaffsTable,
			Columns: permission.StaffsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.StaffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.StaffsTable,
			Columns: permission.StaffsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PermissionMutation
}

// SetName sets the "name" field.
func (puo *PermissionUpdateOne) SetName(s string) *PermissionUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetDescription sets the "description" field.
func (puo *PermissionUpdateOne) SetDescription(s string) *PermissionUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableDescription(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *PermissionUpdateOne) ClearDescription() *PermissionUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// AddStaffIDs adds the "staffs" edge to the Staff entity by IDs.
func (puo *PermissionUpdateOne) AddStaffIDs(ids ...string) *PermissionUpdateOne {
	puo.mutation.AddStaffIDs(ids...)
	return puo
}

// AddStaffs adds the "staffs" edges to the Staff entity.
func (puo *PermissionUpdateOne) AddStaffs(s ...*Staff) *PermissionUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddStaffIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// ClearStaffs clears all "staffs" edges to the Staff entity.
func (puo *PermissionUpdateOne) ClearStaffs() *PermissionUpdateOne {
	puo.mutation.ClearStaffs()
	return puo
}

// RemoveStaffIDs removes the "staffs" edge to Staff entities by IDs.
func (puo *PermissionUpdateOne) RemoveStaffIDs(ids ...string) *PermissionUpdateOne {
	puo.mutation.RemoveStaffIDs(ids...)
	return puo
}

// RemoveStaffs removes "staffs" edges to Staff entities.
func (puo *PermissionUpdateOne) RemoveStaffs(s ...*Staff) *PermissionUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveStaffIDs(ids...)
}

// Where appends a list predicates to the PermissionUpdate builder.
func (puo *PermissionUpdateOne) Where(ps ...predicate.Permission) *PermissionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeString))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Permission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(permission.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(permission.FieldDescription, field.TypeString)
	}
	if puo.mutation.StaffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.StaffsTable,
			Columns: permission.StaffsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedStaffsIDs(); len(nodes) > 0 && !puo.mutation.StaffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.StaffsTable,
			Columns: permission.StaffsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.StaffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.StaffsTable,
			Columns: permission.StaffsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
