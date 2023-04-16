// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"socialite/ent/follow"
	"socialite/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FollowCreate is the builder for creating a Follow entity.
type FollowCreate struct {
	config
	mutation *FollowMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (fc *FollowCreate) SetCreatedAt(t time.Time) *FollowCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FollowCreate) SetNillableCreatedAt(t *time.Time) *FollowCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FollowCreate) SetUpdatedAt(t time.Time) *FollowCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FollowCreate) SetNillableUpdatedAt(t *time.Time) *FollowCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FollowCreate) SetID(u uuid.UUID) *FollowCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FollowCreate) SetNillableID(u *uuid.UUID) *FollowCreate {
	if u != nil {
		fc.SetID(*u)
	}
	return fc
}

// SetFollowerID sets the "follower" edge to the User entity by ID.
func (fc *FollowCreate) SetFollowerID(id uuid.UUID) *FollowCreate {
	fc.mutation.SetFollowerID(id)
	return fc
}

// SetFollower sets the "follower" edge to the User entity.
func (fc *FollowCreate) SetFollower(u *User) *FollowCreate {
	return fc.SetFollowerID(u.ID)
}

// SetFolloweeID sets the "followee" edge to the User entity by ID.
func (fc *FollowCreate) SetFolloweeID(id uuid.UUID) *FollowCreate {
	fc.mutation.SetFolloweeID(id)
	return fc
}

// SetFollowee sets the "followee" edge to the User entity.
func (fc *FollowCreate) SetFollowee(u *User) *FollowCreate {
	return fc.SetFolloweeID(u.ID)
}

// Mutation returns the FollowMutation object of the builder.
func (fc *FollowCreate) Mutation() *FollowMutation {
	return fc.mutation
}

// Save creates the Follow in the database.
func (fc *FollowCreate) Save(ctx context.Context) (*Follow, error) {
	fc.defaults()
	return withHooks[*Follow, FollowMutation](ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FollowCreate) SaveX(ctx context.Context) *Follow {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FollowCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FollowCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FollowCreate) defaults() {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := follow.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		v := follow.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		v := follow.DefaultID()
		fc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FollowCreate) check() error {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Follow.created_at"`)}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Follow.updated_at"`)}
	}
	if _, ok := fc.mutation.FollowerID(); !ok {
		return &ValidationError{Name: "follower", err: errors.New(`ent: missing required edge "Follow.follower"`)}
	}
	if _, ok := fc.mutation.FolloweeID(); !ok {
		return &ValidationError{Name: "followee", err: errors.New(`ent: missing required edge "Follow.followee"`)}
	}
	return nil
}

func (fc *FollowCreate) sqlSave(ctx context.Context) (*Follow, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FollowCreate) createSpec() (*Follow, *sqlgraph.CreateSpec) {
	var (
		_node = &Follow{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(follow.Table, sqlgraph.NewFieldSpec(follow.FieldID, field.TypeUUID))
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(follow.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.SetField(follow.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := fc.mutation.FollowerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.FollowerTable,
			Columns: []string{follow.FollowerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.follow_follower = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FolloweeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.FolloweeTable,
			Columns: []string{follow.FolloweeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.follow_followee = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FollowCreateBulk is the builder for creating many Follow entities in bulk.
type FollowCreateBulk struct {
	config
	builders []*FollowCreate
}

// Save creates the Follow entities in the database.
func (fcb *FollowCreateBulk) Save(ctx context.Context) ([]*Follow, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Follow, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FollowMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FollowCreateBulk) SaveX(ctx context.Context) []*Follow {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FollowCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FollowCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
