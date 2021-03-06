// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/daymenu/gostudy/examples/ent/ent/admin"
	"github.com/daymenu/gostudy/examples/ent/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// AdminUpdate is the builder for updating Admin entities.
type AdminUpdate struct {
	config
	hooks    []Hook
	mutation *AdminMutation
}

// Where adds a new predicate for the builder.
func (au *AdminUpdate) Where(ps ...predicate.Admin) *AdminUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetAge sets the age field.
func (au *AdminUpdate) SetAge(i int) *AdminUpdate {
	au.mutation.ResetAge()
	au.mutation.SetAge(i)
	return au
}

// AddAge adds i to age.
func (au *AdminUpdate) AddAge(i int) *AdminUpdate {
	au.mutation.AddAge(i)
	return au
}

// SetRank sets the rank field.
func (au *AdminUpdate) SetRank(f float64) *AdminUpdate {
	au.mutation.ResetRank()
	au.mutation.SetRank(f)
	return au
}

// SetNillableRank sets the rank field if the given value is not nil.
func (au *AdminUpdate) SetNillableRank(f *float64) *AdminUpdate {
	if f != nil {
		au.SetRank(*f)
	}
	return au
}

// AddRank adds f to rank.
func (au *AdminUpdate) AddRank(f float64) *AdminUpdate {
	au.mutation.AddRank(f)
	return au
}

// ClearRank clears the value of rank.
func (au *AdminUpdate) ClearRank() *AdminUpdate {
	au.mutation.ClearRank()
	return au
}

// SetActive sets the active field.
func (au *AdminUpdate) SetActive(b bool) *AdminUpdate {
	au.mutation.SetActive(b)
	return au
}

// SetNillableActive sets the active field if the given value is not nil.
func (au *AdminUpdate) SetNillableActive(b *bool) *AdminUpdate {
	if b != nil {
		au.SetActive(*b)
	}
	return au
}

// SetName sets the name field.
func (au *AdminUpdate) SetName(s string) *AdminUpdate {
	au.mutation.SetName(s)
	return au
}

// SetCreatedAt sets the created_at field.
func (au *AdminUpdate) SetCreatedAt(t time.Time) *AdminUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (au *AdminUpdate) SetNillableCreatedAt(t *time.Time) *AdminUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetURL sets the url field.
func (au *AdminUpdate) SetURL(u *url.URL) *AdminUpdate {
	au.mutation.SetURL(u)
	return au
}

// ClearURL clears the value of url.
func (au *AdminUpdate) ClearURL() *AdminUpdate {
	au.mutation.ClearURL()
	return au
}

// SetStrings sets the strings field.
func (au *AdminUpdate) SetStrings(s []string) *AdminUpdate {
	au.mutation.SetStrings(s)
	return au
}

// ClearStrings clears the value of strings.
func (au *AdminUpdate) ClearStrings() *AdminUpdate {
	au.mutation.ClearStrings()
	return au
}

// SetState sets the state field.
func (au *AdminUpdate) SetState(a admin.State) *AdminUpdate {
	au.mutation.SetState(a)
	return au
}

// SetNillableState sets the state field if the given value is not nil.
func (au *AdminUpdate) SetNillableState(a *admin.State) *AdminUpdate {
	if a != nil {
		au.SetState(*a)
	}
	return au
}

// ClearState clears the value of state.
func (au *AdminUpdate) ClearState() *AdminUpdate {
	au.mutation.ClearState()
	return au
}

// SetUUID sets the uuid field.
func (au *AdminUpdate) SetUUID(u uuid.UUID) *AdminUpdate {
	au.mutation.SetUUID(u)
	return au
}

// Mutation returns the AdminMutation object of the builder.
func (au *AdminUpdate) Mutation() *AdminMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AdminUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AdminUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AdminUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AdminUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AdminUpdate) check() error {
	if v, ok := au.mutation.Age(); ok {
		if err := admin.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	if v, ok := au.mutation.State(); ok {
		if err := admin.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	return nil
}

func (au *AdminUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   admin.Table,
			Columns: admin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: admin.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: admin.FieldAge,
		})
	}
	if value, ok := au.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: admin.FieldAge,
		})
	}
	if value, ok := au.mutation.Rank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: admin.FieldRank,
		})
	}
	if value, ok := au.mutation.AddedRank(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: admin.FieldRank,
		})
	}
	if au.mutation.RankCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: admin.FieldRank,
		})
	}
	if value, ok := au.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: admin.FieldActive,
		})
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldName,
		})
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admin.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: admin.FieldURL,
		})
	}
	if au.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: admin.FieldURL,
		})
	}
	if value, ok := au.mutation.Strings(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: admin.FieldStrings,
		})
	}
	if au.mutation.StringsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: admin.FieldStrings,
		})
	}
	if value, ok := au.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: admin.FieldState,
		})
	}
	if au.mutation.StateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: admin.FieldState,
		})
	}
	if value, ok := au.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: admin.FieldUUID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AdminUpdateOne is the builder for updating a single Admin entity.
type AdminUpdateOne struct {
	config
	hooks    []Hook
	mutation *AdminMutation
}

// SetAge sets the age field.
func (auo *AdminUpdateOne) SetAge(i int) *AdminUpdateOne {
	auo.mutation.ResetAge()
	auo.mutation.SetAge(i)
	return auo
}

// AddAge adds i to age.
func (auo *AdminUpdateOne) AddAge(i int) *AdminUpdateOne {
	auo.mutation.AddAge(i)
	return auo
}

// SetRank sets the rank field.
func (auo *AdminUpdateOne) SetRank(f float64) *AdminUpdateOne {
	auo.mutation.ResetRank()
	auo.mutation.SetRank(f)
	return auo
}

// SetNillableRank sets the rank field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableRank(f *float64) *AdminUpdateOne {
	if f != nil {
		auo.SetRank(*f)
	}
	return auo
}

// AddRank adds f to rank.
func (auo *AdminUpdateOne) AddRank(f float64) *AdminUpdateOne {
	auo.mutation.AddRank(f)
	return auo
}

// ClearRank clears the value of rank.
func (auo *AdminUpdateOne) ClearRank() *AdminUpdateOne {
	auo.mutation.ClearRank()
	return auo
}

// SetActive sets the active field.
func (auo *AdminUpdateOne) SetActive(b bool) *AdminUpdateOne {
	auo.mutation.SetActive(b)
	return auo
}

// SetNillableActive sets the active field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableActive(b *bool) *AdminUpdateOne {
	if b != nil {
		auo.SetActive(*b)
	}
	return auo
}

// SetName sets the name field.
func (auo *AdminUpdateOne) SetName(s string) *AdminUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetCreatedAt sets the created_at field.
func (auo *AdminUpdateOne) SetCreatedAt(t time.Time) *AdminUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableCreatedAt(t *time.Time) *AdminUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetURL sets the url field.
func (auo *AdminUpdateOne) SetURL(u *url.URL) *AdminUpdateOne {
	auo.mutation.SetURL(u)
	return auo
}

// ClearURL clears the value of url.
func (auo *AdminUpdateOne) ClearURL() *AdminUpdateOne {
	auo.mutation.ClearURL()
	return auo
}

// SetStrings sets the strings field.
func (auo *AdminUpdateOne) SetStrings(s []string) *AdminUpdateOne {
	auo.mutation.SetStrings(s)
	return auo
}

// ClearStrings clears the value of strings.
func (auo *AdminUpdateOne) ClearStrings() *AdminUpdateOne {
	auo.mutation.ClearStrings()
	return auo
}

// SetState sets the state field.
func (auo *AdminUpdateOne) SetState(a admin.State) *AdminUpdateOne {
	auo.mutation.SetState(a)
	return auo
}

// SetNillableState sets the state field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableState(a *admin.State) *AdminUpdateOne {
	if a != nil {
		auo.SetState(*a)
	}
	return auo
}

// ClearState clears the value of state.
func (auo *AdminUpdateOne) ClearState() *AdminUpdateOne {
	auo.mutation.ClearState()
	return auo
}

// SetUUID sets the uuid field.
func (auo *AdminUpdateOne) SetUUID(u uuid.UUID) *AdminUpdateOne {
	auo.mutation.SetUUID(u)
	return auo
}

// Mutation returns the AdminMutation object of the builder.
func (auo *AdminUpdateOne) Mutation() *AdminMutation {
	return auo.mutation
}

// Save executes the query and returns the updated entity.
func (auo *AdminUpdateOne) Save(ctx context.Context) (*Admin, error) {
	var (
		err  error
		node *Admin
	)
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AdminUpdateOne) SaveX(ctx context.Context) *Admin {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AdminUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AdminUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AdminUpdateOne) check() error {
	if v, ok := auo.mutation.Age(); ok {
		if err := admin.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	if v, ok := auo.mutation.State(); ok {
		if err := admin.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	return nil
}

func (auo *AdminUpdateOne) sqlSave(ctx context.Context) (_node *Admin, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   admin.Table,
			Columns: admin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: admin.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Admin.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: admin.FieldAge,
		})
	}
	if value, ok := auo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: admin.FieldAge,
		})
	}
	if value, ok := auo.mutation.Rank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: admin.FieldRank,
		})
	}
	if value, ok := auo.mutation.AddedRank(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: admin.FieldRank,
		})
	}
	if auo.mutation.RankCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: admin.FieldRank,
		})
	}
	if value, ok := auo.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: admin.FieldActive,
		})
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldName,
		})
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admin.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: admin.FieldURL,
		})
	}
	if auo.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: admin.FieldURL,
		})
	}
	if value, ok := auo.mutation.Strings(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: admin.FieldStrings,
		})
	}
	if auo.mutation.StringsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: admin.FieldStrings,
		})
	}
	if value, ok := auo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: admin.FieldState,
		})
	}
	if auo.mutation.StateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: admin.FieldState,
		})
	}
	if value, ok := auo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: admin.FieldUUID,
		})
	}
	_node = &Admin{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
