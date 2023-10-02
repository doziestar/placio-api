// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/notification"
	"placio-app/ent/predicate"
	"placio-app/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotificationUpdate is the builder for updating Notification entities.
type NotificationUpdate struct {
	config
	hooks    []Hook
	mutation *NotificationMutation
}

// Where appends a list predicates to the NotificationUpdate builder.
func (nu *NotificationUpdate) Where(ps ...predicate.Notification) *NotificationUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetTitle sets the "title" field.
func (nu *NotificationUpdate) SetTitle(s string) *NotificationUpdate {
	nu.mutation.SetTitle(s)
	return nu
}

// SetMessage sets the "message" field.
func (nu *NotificationUpdate) SetMessage(s string) *NotificationUpdate {
	nu.mutation.SetMessage(s)
	return nu
}

// SetLink sets the "link" field.
func (nu *NotificationUpdate) SetLink(s string) *NotificationUpdate {
	nu.mutation.SetLink(s)
	return nu
}

// SetIsRead sets the "is_read" field.
func (nu *NotificationUpdate) SetIsRead(b bool) *NotificationUpdate {
	nu.mutation.SetIsRead(b)
	return nu
}

// SetNillableIsRead sets the "is_read" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableIsRead(b *bool) *NotificationUpdate {
	if b != nil {
		nu.SetIsRead(*b)
	}
	return nu
}

// SetType sets the "type" field.
func (nu *NotificationUpdate) SetType(i int) *NotificationUpdate {
	nu.mutation.ResetType()
	nu.mutation.SetType(i)
	return nu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableType(i *int) *NotificationUpdate {
	if i != nil {
		nu.SetType(*i)
	}
	return nu
}

// AddType adds i to the "type" field.
func (nu *NotificationUpdate) AddType(i int) *NotificationUpdate {
	nu.mutation.AddType(i)
	return nu
}

// SetCreatedAt sets the "created_at" field.
func (nu *NotificationUpdate) SetCreatedAt(t time.Time) *NotificationUpdate {
	nu.mutation.SetCreatedAt(t)
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NotificationUpdate) SetUpdatedAt(t time.Time) *NotificationUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// SetNotifiableType sets the "notifiable_type" field.
func (nu *NotificationUpdate) SetNotifiableType(s string) *NotificationUpdate {
	nu.mutation.SetNotifiableType(s)
	return nu
}

// SetNotifiableID sets the "notifiable_id" field.
func (nu *NotificationUpdate) SetNotifiableID(s string) *NotificationUpdate {
	nu.mutation.SetNotifiableID(s)
	return nu
}

// SetTriggeredBy sets the "triggered_by" field.
func (nu *NotificationUpdate) SetTriggeredBy(s string) *NotificationUpdate {
	nu.mutation.SetTriggeredBy(s)
	return nu
}

// SetTriggeredTo sets the "triggered_to" field.
func (nu *NotificationUpdate) SetTriggeredTo(s string) *NotificationUpdate {
	nu.mutation.SetTriggeredTo(s)
	return nu
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (nu *NotificationUpdate) AddUserIDs(ids ...string) *NotificationUpdate {
	nu.mutation.AddUserIDs(ids...)
	return nu
}

// AddUser adds the "user" edges to the User entity.
func (nu *NotificationUpdate) AddUser(u ...*User) *NotificationUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nu.AddUserIDs(ids...)
}

// AddBusinessAccountIDs adds the "business_account" edge to the Business entity by IDs.
func (nu *NotificationUpdate) AddBusinessAccountIDs(ids ...string) *NotificationUpdate {
	nu.mutation.AddBusinessAccountIDs(ids...)
	return nu
}

// AddBusinessAccount adds the "business_account" edges to the Business entity.
func (nu *NotificationUpdate) AddBusinessAccount(b ...*Business) *NotificationUpdate {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return nu.AddBusinessAccountIDs(ids...)
}

// Mutation returns the NotificationMutation object of the builder.
func (nu *NotificationUpdate) Mutation() *NotificationMutation {
	return nu.mutation
}

// ClearUser clears all "user" edges to the User entity.
func (nu *NotificationUpdate) ClearUser() *NotificationUpdate {
	nu.mutation.ClearUser()
	return nu
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (nu *NotificationUpdate) RemoveUserIDs(ids ...string) *NotificationUpdate {
	nu.mutation.RemoveUserIDs(ids...)
	return nu
}

// RemoveUser removes "user" edges to User entities.
func (nu *NotificationUpdate) RemoveUser(u ...*User) *NotificationUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nu.RemoveUserIDs(ids...)
}

// ClearBusinessAccount clears all "business_account" edges to the Business entity.
func (nu *NotificationUpdate) ClearBusinessAccount() *NotificationUpdate {
	nu.mutation.ClearBusinessAccount()
	return nu
}

// RemoveBusinessAccountIDs removes the "business_account" edge to Business entities by IDs.
func (nu *NotificationUpdate) RemoveBusinessAccountIDs(ids ...string) *NotificationUpdate {
	nu.mutation.RemoveBusinessAccountIDs(ids...)
	return nu
}

// RemoveBusinessAccount removes "business_account" edges to Business entities.
func (nu *NotificationUpdate) RemoveBusinessAccount(b ...*Business) *NotificationUpdate {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return nu.RemoveBusinessAccountIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NotificationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NotificationUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NotificationUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NotificationUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nu *NotificationUpdate) check() error {
	if v, ok := nu.mutation.Title(); ok {
		if err := notification.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Notification.title": %w`, err)}
		}
	}
	if v, ok := nu.mutation.Message(); ok {
		if err := notification.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "Notification.message": %w`, err)}
		}
	}
	if v, ok := nu.mutation.Link(); ok {
		if err := notification.LinkValidator(v); err != nil {
			return &ValidationError{Name: "link", err: fmt.Errorf(`ent: validator failed for field "Notification.link": %w`, err)}
		}
	}
	if v, ok := nu.mutation.NotifiableType(); ok {
		if err := notification.NotifiableTypeValidator(v); err != nil {
			return &ValidationError{Name: "notifiable_type", err: fmt.Errorf(`ent: validator failed for field "Notification.notifiable_type": %w`, err)}
		}
	}
	if v, ok := nu.mutation.NotifiableID(); ok {
		if err := notification.NotifiableIDValidator(v); err != nil {
			return &ValidationError{Name: "notifiable_id", err: fmt.Errorf(`ent: validator failed for field "Notification.notifiable_id": %w`, err)}
		}
	}
	if v, ok := nu.mutation.TriggeredBy(); ok {
		if err := notification.TriggeredByValidator(v); err != nil {
			return &ValidationError{Name: "triggered_by", err: fmt.Errorf(`ent: validator failed for field "Notification.triggered_by": %w`, err)}
		}
	}
	if v, ok := nu.mutation.TriggeredTo(); ok {
		if err := notification.TriggeredToValidator(v); err != nil {
			return &ValidationError{Name: "triggered_to", err: fmt.Errorf(`ent: validator failed for field "Notification.triggered_to": %w`, err)}
		}
	}
	return nil
}

func (nu *NotificationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeString))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Title(); ok {
		_spec.SetField(notification.FieldTitle, field.TypeString, value)
	}
	if value, ok := nu.mutation.Message(); ok {
		_spec.SetField(notification.FieldMessage, field.TypeString, value)
	}
	if value, ok := nu.mutation.Link(); ok {
		_spec.SetField(notification.FieldLink, field.TypeString, value)
	}
	if value, ok := nu.mutation.IsRead(); ok {
		_spec.SetField(notification.FieldIsRead, field.TypeBool, value)
	}
	if value, ok := nu.mutation.GetType(); ok {
		_spec.SetField(notification.FieldType, field.TypeInt, value)
	}
	if value, ok := nu.mutation.AddedType(); ok {
		_spec.AddField(notification.FieldType, field.TypeInt, value)
	}
	if value, ok := nu.mutation.CreatedAt(); ok {
		_spec.SetField(notification.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.SetField(notification.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nu.mutation.NotifiableType(); ok {
		_spec.SetField(notification.FieldNotifiableType, field.TypeString, value)
	}
	if value, ok := nu.mutation.NotifiableID(); ok {
		_spec.SetField(notification.FieldNotifiableID, field.TypeString, value)
	}
	if value, ok := nu.mutation.TriggeredBy(); ok {
		_spec.SetField(notification.FieldTriggeredBy, field.TypeString, value)
	}
	if value, ok := nu.mutation.TriggeredTo(); ok {
		_spec.SetField(notification.FieldTriggeredTo, field.TypeString, value)
	}
	if nu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notification.UserTable,
			Columns: notification.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedUserIDs(); len(nodes) > 0 && !nu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notification.UserTable,
			Columns: notification.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notification.UserTable,
			Columns: notification.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.BusinessAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notification.BusinessAccountTable,
			Columns: notification.BusinessAccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedBusinessAccountIDs(); len(nodes) > 0 && !nu.mutation.BusinessAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notification.BusinessAccountTable,
			Columns: notification.BusinessAccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.BusinessAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notification.BusinessAccountTable,
			Columns: notification.BusinessAccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NotificationUpdateOne is the builder for updating a single Notification entity.
type NotificationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotificationMutation
}

// SetTitle sets the "title" field.
func (nuo *NotificationUpdateOne) SetTitle(s string) *NotificationUpdateOne {
	nuo.mutation.SetTitle(s)
	return nuo
}

// SetMessage sets the "message" field.
func (nuo *NotificationUpdateOne) SetMessage(s string) *NotificationUpdateOne {
	nuo.mutation.SetMessage(s)
	return nuo
}

// SetLink sets the "link" field.
func (nuo *NotificationUpdateOne) SetLink(s string) *NotificationUpdateOne {
	nuo.mutation.SetLink(s)
	return nuo
}

// SetIsRead sets the "is_read" field.
func (nuo *NotificationUpdateOne) SetIsRead(b bool) *NotificationUpdateOne {
	nuo.mutation.SetIsRead(b)
	return nuo
}

// SetNillableIsRead sets the "is_read" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableIsRead(b *bool) *NotificationUpdateOne {
	if b != nil {
		nuo.SetIsRead(*b)
	}
	return nuo
}

// SetType sets the "type" field.
func (nuo *NotificationUpdateOne) SetType(i int) *NotificationUpdateOne {
	nuo.mutation.ResetType()
	nuo.mutation.SetType(i)
	return nuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableType(i *int) *NotificationUpdateOne {
	if i != nil {
		nuo.SetType(*i)
	}
	return nuo
}

// AddType adds i to the "type" field.
func (nuo *NotificationUpdateOne) AddType(i int) *NotificationUpdateOne {
	nuo.mutation.AddType(i)
	return nuo
}

// SetCreatedAt sets the "created_at" field.
func (nuo *NotificationUpdateOne) SetCreatedAt(t time.Time) *NotificationUpdateOne {
	nuo.mutation.SetCreatedAt(t)
	return nuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NotificationUpdateOne) SetUpdatedAt(t time.Time) *NotificationUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// SetNotifiableType sets the "notifiable_type" field.
func (nuo *NotificationUpdateOne) SetNotifiableType(s string) *NotificationUpdateOne {
	nuo.mutation.SetNotifiableType(s)
	return nuo
}

// SetNotifiableID sets the "notifiable_id" field.
func (nuo *NotificationUpdateOne) SetNotifiableID(s string) *NotificationUpdateOne {
	nuo.mutation.SetNotifiableID(s)
	return nuo
}

// SetTriggeredBy sets the "triggered_by" field.
func (nuo *NotificationUpdateOne) SetTriggeredBy(s string) *NotificationUpdateOne {
	nuo.mutation.SetTriggeredBy(s)
	return nuo
}

// SetTriggeredTo sets the "triggered_to" field.
func (nuo *NotificationUpdateOne) SetTriggeredTo(s string) *NotificationUpdateOne {
	nuo.mutation.SetTriggeredTo(s)
	return nuo
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (nuo *NotificationUpdateOne) AddUserIDs(ids ...string) *NotificationUpdateOne {
	nuo.mutation.AddUserIDs(ids...)
	return nuo
}

// AddUser adds the "user" edges to the User entity.
func (nuo *NotificationUpdateOne) AddUser(u ...*User) *NotificationUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nuo.AddUserIDs(ids...)
}

// AddBusinessAccountIDs adds the "business_account" edge to the Business entity by IDs.
func (nuo *NotificationUpdateOne) AddBusinessAccountIDs(ids ...string) *NotificationUpdateOne {
	nuo.mutation.AddBusinessAccountIDs(ids...)
	return nuo
}

// AddBusinessAccount adds the "business_account" edges to the Business entity.
func (nuo *NotificationUpdateOne) AddBusinessAccount(b ...*Business) *NotificationUpdateOne {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return nuo.AddBusinessAccountIDs(ids...)
}

// Mutation returns the NotificationMutation object of the builder.
func (nuo *NotificationUpdateOne) Mutation() *NotificationMutation {
	return nuo.mutation
}

// ClearUser clears all "user" edges to the User entity.
func (nuo *NotificationUpdateOne) ClearUser() *NotificationUpdateOne {
	nuo.mutation.ClearUser()
	return nuo
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (nuo *NotificationUpdateOne) RemoveUserIDs(ids ...string) *NotificationUpdateOne {
	nuo.mutation.RemoveUserIDs(ids...)
	return nuo
}

// RemoveUser removes "user" edges to User entities.
func (nuo *NotificationUpdateOne) RemoveUser(u ...*User) *NotificationUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return nuo.RemoveUserIDs(ids...)
}

// ClearBusinessAccount clears all "business_account" edges to the Business entity.
func (nuo *NotificationUpdateOne) ClearBusinessAccount() *NotificationUpdateOne {
	nuo.mutation.ClearBusinessAccount()
	return nuo
}

// RemoveBusinessAccountIDs removes the "business_account" edge to Business entities by IDs.
func (nuo *NotificationUpdateOne) RemoveBusinessAccountIDs(ids ...string) *NotificationUpdateOne {
	nuo.mutation.RemoveBusinessAccountIDs(ids...)
	return nuo
}

// RemoveBusinessAccount removes "business_account" edges to Business entities.
func (nuo *NotificationUpdateOne) RemoveBusinessAccount(b ...*Business) *NotificationUpdateOne {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return nuo.RemoveBusinessAccountIDs(ids...)
}

// Where appends a list predicates to the NotificationUpdate builder.
func (nuo *NotificationUpdateOne) Where(ps ...predicate.Notification) *NotificationUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NotificationUpdateOne) Select(field string, fields ...string) *NotificationUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Notification entity.
func (nuo *NotificationUpdateOne) Save(ctx context.Context) (*Notification, error) {
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NotificationUpdateOne) SaveX(ctx context.Context) *Notification {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NotificationUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NotificationUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nuo *NotificationUpdateOne) check() error {
	if v, ok := nuo.mutation.Title(); ok {
		if err := notification.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Notification.title": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.Message(); ok {
		if err := notification.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "Notification.message": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.Link(); ok {
		if err := notification.LinkValidator(v); err != nil {
			return &ValidationError{Name: "link", err: fmt.Errorf(`ent: validator failed for field "Notification.link": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.NotifiableType(); ok {
		if err := notification.NotifiableTypeValidator(v); err != nil {
			return &ValidationError{Name: "notifiable_type", err: fmt.Errorf(`ent: validator failed for field "Notification.notifiable_type": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.NotifiableID(); ok {
		if err := notification.NotifiableIDValidator(v); err != nil {
			return &ValidationError{Name: "notifiable_id", err: fmt.Errorf(`ent: validator failed for field "Notification.notifiable_id": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.TriggeredBy(); ok {
		if err := notification.TriggeredByValidator(v); err != nil {
			return &ValidationError{Name: "triggered_by", err: fmt.Errorf(`ent: validator failed for field "Notification.triggered_by": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.TriggeredTo(); ok {
		if err := notification.TriggeredToValidator(v); err != nil {
			return &ValidationError{Name: "triggered_to", err: fmt.Errorf(`ent: validator failed for field "Notification.triggered_to": %w`, err)}
		}
	}
	return nil
}

func (nuo *NotificationUpdateOne) sqlSave(ctx context.Context) (_node *Notification, err error) {
	if err := nuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeString))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Notification.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notification.FieldID)
		for _, f := range fields {
			if !notification.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.Title(); ok {
		_spec.SetField(notification.FieldTitle, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Message(); ok {
		_spec.SetField(notification.FieldMessage, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Link(); ok {
		_spec.SetField(notification.FieldLink, field.TypeString, value)
	}
	if value, ok := nuo.mutation.IsRead(); ok {
		_spec.SetField(notification.FieldIsRead, field.TypeBool, value)
	}
	if value, ok := nuo.mutation.GetType(); ok {
		_spec.SetField(notification.FieldType, field.TypeInt, value)
	}
	if value, ok := nuo.mutation.AddedType(); ok {
		_spec.AddField(notification.FieldType, field.TypeInt, value)
	}
	if value, ok := nuo.mutation.CreatedAt(); ok {
		_spec.SetField(notification.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.SetField(notification.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.NotifiableType(); ok {
		_spec.SetField(notification.FieldNotifiableType, field.TypeString, value)
	}
	if value, ok := nuo.mutation.NotifiableID(); ok {
		_spec.SetField(notification.FieldNotifiableID, field.TypeString, value)
	}
	if value, ok := nuo.mutation.TriggeredBy(); ok {
		_spec.SetField(notification.FieldTriggeredBy, field.TypeString, value)
	}
	if value, ok := nuo.mutation.TriggeredTo(); ok {
		_spec.SetField(notification.FieldTriggeredTo, field.TypeString, value)
	}
	if nuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notification.UserTable,
			Columns: notification.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedUserIDs(); len(nodes) > 0 && !nuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notification.UserTable,
			Columns: notification.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notification.UserTable,
			Columns: notification.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.BusinessAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notification.BusinessAccountTable,
			Columns: notification.BusinessAccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedBusinessAccountIDs(); len(nodes) > 0 && !nuo.mutation.BusinessAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notification.BusinessAccountTable,
			Columns: notification.BusinessAccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.BusinessAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notification.BusinessAccountTable,
			Columns: notification.BusinessAccountPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Notification{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}