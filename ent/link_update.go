// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kindlyfire/longa/ent/link"
	"github.com/kindlyfire/longa/ent/predicate"
)

// LinkUpdate is the builder for updating Link entities.
type LinkUpdate struct {
	config
	hooks    []Hook
	mutation *LinkMutation
}

// Where appends a list predicates to the LinkUpdate builder.
func (lu *LinkUpdate) Where(ps ...predicate.Link) *LinkUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetLink sets the "link" field.
func (lu *LinkUpdate) SetLink(s string) *LinkUpdate {
	lu.mutation.SetLink(s)
	return lu
}

// SetTarget sets the "target" field.
func (lu *LinkUpdate) SetTarget(s string) *LinkUpdate {
	lu.mutation.SetTarget(s)
	return lu
}

// SetCreatedAt sets the "createdAt" field.
func (lu *LinkUpdate) SetCreatedAt(t time.Time) *LinkUpdate {
	lu.mutation.SetCreatedAt(t)
	return lu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (lu *LinkUpdate) SetNillableCreatedAt(t *time.Time) *LinkUpdate {
	if t != nil {
		lu.SetCreatedAt(*t)
	}
	return lu
}

// Mutation returns the LinkMutation object of the builder.
func (lu *LinkUpdate) Mutation() *LinkMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LinkUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LinkUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LinkUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LinkUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LinkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   link.Table,
			Columns: link.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: link.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Link(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: link.FieldLink,
		})
	}
	if value, ok := lu.mutation.Target(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: link.FieldTarget,
		})
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: link.FieldCreatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// LinkUpdateOne is the builder for updating a single Link entity.
type LinkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LinkMutation
}

// SetLink sets the "link" field.
func (luo *LinkUpdateOne) SetLink(s string) *LinkUpdateOne {
	luo.mutation.SetLink(s)
	return luo
}

// SetTarget sets the "target" field.
func (luo *LinkUpdateOne) SetTarget(s string) *LinkUpdateOne {
	luo.mutation.SetTarget(s)
	return luo
}

// SetCreatedAt sets the "createdAt" field.
func (luo *LinkUpdateOne) SetCreatedAt(t time.Time) *LinkUpdateOne {
	luo.mutation.SetCreatedAt(t)
	return luo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (luo *LinkUpdateOne) SetNillableCreatedAt(t *time.Time) *LinkUpdateOne {
	if t != nil {
		luo.SetCreatedAt(*t)
	}
	return luo
}

// Mutation returns the LinkMutation object of the builder.
func (luo *LinkUpdateOne) Mutation() *LinkMutation {
	return luo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LinkUpdateOne) Select(field string, fields ...string) *LinkUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Link entity.
func (luo *LinkUpdateOne) Save(ctx context.Context) (*Link, error) {
	var (
		err  error
		node *Link
	)
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LinkUpdateOne) SaveX(ctx context.Context) *Link {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LinkUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LinkUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LinkUpdateOne) sqlSave(ctx context.Context) (_node *Link, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   link.Table,
			Columns: link.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: link.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Link.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, link.FieldID)
		for _, f := range fields {
			if !link.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != link.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Link(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: link.FieldLink,
		})
	}
	if value, ok := luo.mutation.Target(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: link.FieldTarget,
		})
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: link.FieldCreatedAt,
		})
	}
	_node = &Link{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
