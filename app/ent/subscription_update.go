// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/plan"
	"placio-app/ent/predicate"
	"placio-app/ent/price"
	"placio-app/ent/subscription"
	"placio-app/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubscriptionUpdate is the builder for updating Subscription entities.
type SubscriptionUpdate struct {
	config
	hooks    []Hook
	mutation *SubscriptionMutation
}

// Where appends a list predicates to the SubscriptionUpdate builder.
func (su *SubscriptionUpdate) Where(ps ...predicate.Subscription) *SubscriptionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetStartDate sets the "start_date" field.
func (su *SubscriptionUpdate) SetStartDate(t time.Time) *SubscriptionUpdate {
	su.mutation.SetStartDate(t)
	return su
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableStartDate(t *time.Time) *SubscriptionUpdate {
	if t != nil {
		su.SetStartDate(*t)
	}
	return su
}

// SetEndDate sets the "end_date" field.
func (su *SubscriptionUpdate) SetEndDate(t time.Time) *SubscriptionUpdate {
	su.mutation.SetEndDate(t)
	return su
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableEndDate(t *time.Time) *SubscriptionUpdate {
	if t != nil {
		su.SetEndDate(*t)
	}
	return su
}

// SetFlutterwaveSubscriptionID sets the "flutterwave_subscription_id" field.
func (su *SubscriptionUpdate) SetFlutterwaveSubscriptionID(s string) *SubscriptionUpdate {
	su.mutation.SetFlutterwaveSubscriptionID(s)
	return su
}

// SetNillableFlutterwaveSubscriptionID sets the "flutterwave_subscription_id" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableFlutterwaveSubscriptionID(s *string) *SubscriptionUpdate {
	if s != nil {
		su.SetFlutterwaveSubscriptionID(*s)
	}
	return su
}

// ClearFlutterwaveSubscriptionID clears the value of the "flutterwave_subscription_id" field.
func (su *SubscriptionUpdate) ClearFlutterwaveSubscriptionID() *SubscriptionUpdate {
	su.mutation.ClearFlutterwaveSubscriptionID()
	return su
}

// SetUserID sets the "user" edge to the User entity by ID.
func (su *SubscriptionUpdate) SetUserID(id string) *SubscriptionUpdate {
	su.mutation.SetUserID(id)
	return su
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableUserID(id *string) *SubscriptionUpdate {
	if id != nil {
		su = su.SetUserID(*id)
	}
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *SubscriptionUpdate) SetUser(u *User) *SubscriptionUpdate {
	return su.SetUserID(u.ID)
}

// SetPlanID sets the "plan" edge to the Plan entity by ID.
func (su *SubscriptionUpdate) SetPlanID(id string) *SubscriptionUpdate {
	su.mutation.SetPlanID(id)
	return su
}

// SetNillablePlanID sets the "plan" edge to the Plan entity by ID if the given value is not nil.
func (su *SubscriptionUpdate) SetNillablePlanID(id *string) *SubscriptionUpdate {
	if id != nil {
		su = su.SetPlanID(*id)
	}
	return su
}

// SetPlan sets the "plan" edge to the Plan entity.
func (su *SubscriptionUpdate) SetPlan(p *Plan) *SubscriptionUpdate {
	return su.SetPlanID(p.ID)
}

// SetPriceID sets the "price" edge to the Price entity by ID.
func (su *SubscriptionUpdate) SetPriceID(id string) *SubscriptionUpdate {
	su.mutation.SetPriceID(id)
	return su
}

// SetNillablePriceID sets the "price" edge to the Price entity by ID if the given value is not nil.
func (su *SubscriptionUpdate) SetNillablePriceID(id *string) *SubscriptionUpdate {
	if id != nil {
		su = su.SetPriceID(*id)
	}
	return su
}

// SetPrice sets the "price" edge to the Price entity.
func (su *SubscriptionUpdate) SetPrice(p *Price) *SubscriptionUpdate {
	return su.SetPriceID(p.ID)
}

// Mutation returns the SubscriptionMutation object of the builder.
func (su *SubscriptionUpdate) Mutation() *SubscriptionMutation {
	return su.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (su *SubscriptionUpdate) ClearUser() *SubscriptionUpdate {
	su.mutation.ClearUser()
	return su
}

// ClearPlan clears the "plan" edge to the Plan entity.
func (su *SubscriptionUpdate) ClearPlan() *SubscriptionUpdate {
	su.mutation.ClearPlan()
	return su
}

// ClearPrice clears the "price" edge to the Price entity.
func (su *SubscriptionUpdate) ClearPrice() *SubscriptionUpdate {
	su.mutation.ClearPrice()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubscriptionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubscriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubscriptionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubscriptionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SubscriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(subscription.Table, subscription.Columns, sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.StartDate(); ok {
		_spec.SetField(subscription.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := su.mutation.EndDate(); ok {
		_spec.SetField(subscription.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := su.mutation.FlutterwaveSubscriptionID(); ok {
		_spec.SetField(subscription.FieldFlutterwaveSubscriptionID, field.TypeString, value)
	}
	if su.mutation.FlutterwaveSubscriptionIDCleared() {
		_spec.ClearField(subscription.FieldFlutterwaveSubscriptionID, field.TypeString)
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.UserTable,
			Columns: []string{subscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.UserTable,
			Columns: []string{subscription.UserColumn},
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
	if su.mutation.PlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.PlanTable,
			Columns: []string{subscription.PlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(plan.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.PlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.PlanTable,
			Columns: []string{subscription.PlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(plan.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.PriceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.PriceTable,
			Columns: []string{subscription.PriceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(price.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.PriceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.PriceTable,
			Columns: []string{subscription.PriceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(price.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SubscriptionUpdateOne is the builder for updating a single Subscription entity.
type SubscriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubscriptionMutation
}

// SetStartDate sets the "start_date" field.
func (suo *SubscriptionUpdateOne) SetStartDate(t time.Time) *SubscriptionUpdateOne {
	suo.mutation.SetStartDate(t)
	return suo
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableStartDate(t *time.Time) *SubscriptionUpdateOne {
	if t != nil {
		suo.SetStartDate(*t)
	}
	return suo
}

// SetEndDate sets the "end_date" field.
func (suo *SubscriptionUpdateOne) SetEndDate(t time.Time) *SubscriptionUpdateOne {
	suo.mutation.SetEndDate(t)
	return suo
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableEndDate(t *time.Time) *SubscriptionUpdateOne {
	if t != nil {
		suo.SetEndDate(*t)
	}
	return suo
}

// SetFlutterwaveSubscriptionID sets the "flutterwave_subscription_id" field.
func (suo *SubscriptionUpdateOne) SetFlutterwaveSubscriptionID(s string) *SubscriptionUpdateOne {
	suo.mutation.SetFlutterwaveSubscriptionID(s)
	return suo
}

// SetNillableFlutterwaveSubscriptionID sets the "flutterwave_subscription_id" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableFlutterwaveSubscriptionID(s *string) *SubscriptionUpdateOne {
	if s != nil {
		suo.SetFlutterwaveSubscriptionID(*s)
	}
	return suo
}

// ClearFlutterwaveSubscriptionID clears the value of the "flutterwave_subscription_id" field.
func (suo *SubscriptionUpdateOne) ClearFlutterwaveSubscriptionID() *SubscriptionUpdateOne {
	suo.mutation.ClearFlutterwaveSubscriptionID()
	return suo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (suo *SubscriptionUpdateOne) SetUserID(id string) *SubscriptionUpdateOne {
	suo.mutation.SetUserID(id)
	return suo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableUserID(id *string) *SubscriptionUpdateOne {
	if id != nil {
		suo = suo.SetUserID(*id)
	}
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *SubscriptionUpdateOne) SetUser(u *User) *SubscriptionUpdateOne {
	return suo.SetUserID(u.ID)
}

// SetPlanID sets the "plan" edge to the Plan entity by ID.
func (suo *SubscriptionUpdateOne) SetPlanID(id string) *SubscriptionUpdateOne {
	suo.mutation.SetPlanID(id)
	return suo
}

// SetNillablePlanID sets the "plan" edge to the Plan entity by ID if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillablePlanID(id *string) *SubscriptionUpdateOne {
	if id != nil {
		suo = suo.SetPlanID(*id)
	}
	return suo
}

// SetPlan sets the "plan" edge to the Plan entity.
func (suo *SubscriptionUpdateOne) SetPlan(p *Plan) *SubscriptionUpdateOne {
	return suo.SetPlanID(p.ID)
}

// SetPriceID sets the "price" edge to the Price entity by ID.
func (suo *SubscriptionUpdateOne) SetPriceID(id string) *SubscriptionUpdateOne {
	suo.mutation.SetPriceID(id)
	return suo
}

// SetNillablePriceID sets the "price" edge to the Price entity by ID if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillablePriceID(id *string) *SubscriptionUpdateOne {
	if id != nil {
		suo = suo.SetPriceID(*id)
	}
	return suo
}

// SetPrice sets the "price" edge to the Price entity.
func (suo *SubscriptionUpdateOne) SetPrice(p *Price) *SubscriptionUpdateOne {
	return suo.SetPriceID(p.ID)
}

// Mutation returns the SubscriptionMutation object of the builder.
func (suo *SubscriptionUpdateOne) Mutation() *SubscriptionMutation {
	return suo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (suo *SubscriptionUpdateOne) ClearUser() *SubscriptionUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// ClearPlan clears the "plan" edge to the Plan entity.
func (suo *SubscriptionUpdateOne) ClearPlan() *SubscriptionUpdateOne {
	suo.mutation.ClearPlan()
	return suo
}

// ClearPrice clears the "price" edge to the Price entity.
func (suo *SubscriptionUpdateOne) ClearPrice() *SubscriptionUpdateOne {
	suo.mutation.ClearPrice()
	return suo
}

// Where appends a list predicates to the SubscriptionUpdate builder.
func (suo *SubscriptionUpdateOne) Where(ps ...predicate.Subscription) *SubscriptionUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubscriptionUpdateOne) Select(field string, fields ...string) *SubscriptionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Subscription entity.
func (suo *SubscriptionUpdateOne) Save(ctx context.Context) (*Subscription, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubscriptionUpdateOne) SaveX(ctx context.Context) *Subscription {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SubscriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubscriptionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SubscriptionUpdateOne) sqlSave(ctx context.Context) (_node *Subscription, err error) {
	_spec := sqlgraph.NewUpdateSpec(subscription.Table, subscription.Columns, sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Subscription.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subscription.FieldID)
		for _, f := range fields {
			if !subscription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subscription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.StartDate(); ok {
		_spec.SetField(subscription.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := suo.mutation.EndDate(); ok {
		_spec.SetField(subscription.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := suo.mutation.FlutterwaveSubscriptionID(); ok {
		_spec.SetField(subscription.FieldFlutterwaveSubscriptionID, field.TypeString, value)
	}
	if suo.mutation.FlutterwaveSubscriptionIDCleared() {
		_spec.ClearField(subscription.FieldFlutterwaveSubscriptionID, field.TypeString)
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.UserTable,
			Columns: []string{subscription.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.UserTable,
			Columns: []string{subscription.UserColumn},
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
	if suo.mutation.PlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.PlanTable,
			Columns: []string{subscription.PlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(plan.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.PlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.PlanTable,
			Columns: []string{subscription.PlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(plan.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.PriceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.PriceTable,
			Columns: []string{subscription.PriceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(price.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.PriceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscription.PriceTable,
			Columns: []string{subscription.PriceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(price.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Subscription{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}