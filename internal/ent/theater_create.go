// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dongwlin/anime-list/internal/ent/anime"
	"github.com/dongwlin/anime-list/internal/ent/theater"
)

// TheaterCreate is the builder for creating a Theater entity.
type TheaterCreate struct {
	config
	mutation *TheaterMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TheaterCreate) SetName(s string) *TheaterCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetCover sets the "cover" field.
func (tc *TheaterCreate) SetCover(s string) *TheaterCreate {
	tc.mutation.SetCover(s)
	return tc
}

// SetReleasedAt sets the "released_at" field.
func (tc *TheaterCreate) SetReleasedAt(t time.Time) *TheaterCreate {
	tc.mutation.SetReleasedAt(t)
	return tc
}

// SetNillableReleasedAt sets the "released_at" field if the given value is not nil.
func (tc *TheaterCreate) SetNillableReleasedAt(t *time.Time) *TheaterCreate {
	if t != nil {
		tc.SetReleasedAt(*t)
	}
	return tc
}

// SetDescription sets the "description" field.
func (tc *TheaterCreate) SetDescription(s string) *TheaterCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TheaterCreate) SetNillableDescription(s *string) *TheaterCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetStatus sets the "status" field.
func (tc *TheaterCreate) SetStatus(i int) *TheaterCreate {
	tc.mutation.SetStatus(i)
	return tc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tc *TheaterCreate) SetNillableStatus(i *int) *TheaterCreate {
	if i != nil {
		tc.SetStatus(*i)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TheaterCreate) SetCreatedAt(t time.Time) *TheaterCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TheaterCreate) SetNillableCreatedAt(t *time.Time) *TheaterCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetAnimeID sets the "anime" edge to the Anime entity by ID.
func (tc *TheaterCreate) SetAnimeID(id int) *TheaterCreate {
	tc.mutation.SetAnimeID(id)
	return tc
}

// SetNillableAnimeID sets the "anime" edge to the Anime entity by ID if the given value is not nil.
func (tc *TheaterCreate) SetNillableAnimeID(id *int) *TheaterCreate {
	if id != nil {
		tc = tc.SetAnimeID(*id)
	}
	return tc
}

// SetAnime sets the "anime" edge to the Anime entity.
func (tc *TheaterCreate) SetAnime(a *Anime) *TheaterCreate {
	return tc.SetAnimeID(a.ID)
}

// Mutation returns the TheaterMutation object of the builder.
func (tc *TheaterCreate) Mutation() *TheaterMutation {
	return tc.mutation
}

// Save creates the Theater in the database.
func (tc *TheaterCreate) Save(ctx context.Context) (*Theater, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TheaterCreate) SaveX(ctx context.Context) *Theater {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TheaterCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TheaterCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TheaterCreate) defaults() {
	if _, ok := tc.mutation.ReleasedAt(); !ok {
		v := theater.DefaultReleasedAt
		tc.mutation.SetReleasedAt(v)
	}
	if _, ok := tc.mutation.Description(); !ok {
		v := theater.DefaultDescription
		tc.mutation.SetDescription(v)
	}
	if _, ok := tc.mutation.Status(); !ok {
		v := theater.DefaultStatus
		tc.mutation.SetStatus(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := theater.DefaultCreatedAt
		tc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TheaterCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Theater.name"`)}
	}
	if _, ok := tc.mutation.Cover(); !ok {
		return &ValidationError{Name: "cover", err: errors.New(`ent: missing required field "Theater.cover"`)}
	}
	if _, ok := tc.mutation.ReleasedAt(); !ok {
		return &ValidationError{Name: "released_at", err: errors.New(`ent: missing required field "Theater.released_at"`)}
	}
	if _, ok := tc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Theater.description"`)}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Theater.status"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Theater.created_at"`)}
	}
	return nil
}

func (tc *TheaterCreate) sqlSave(ctx context.Context) (*Theater, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TheaterCreate) createSpec() (*Theater, *sqlgraph.CreateSpec) {
	var (
		_node = &Theater{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(theater.Table, sqlgraph.NewFieldSpec(theater.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(theater.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Cover(); ok {
		_spec.SetField(theater.FieldCover, field.TypeString, value)
		_node.Cover = value
	}
	if value, ok := tc.mutation.ReleasedAt(); ok {
		_spec.SetField(theater.FieldReleasedAt, field.TypeTime, value)
		_node.ReleasedAt = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(theater.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(theater.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(theater.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := tc.mutation.AnimeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   theater.AnimeTable,
			Columns: []string{theater.AnimeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(anime.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.anime_theaters = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TheaterCreateBulk is the builder for creating many Theater entities in bulk.
type TheaterCreateBulk struct {
	config
	err      error
	builders []*TheaterCreate
}

// Save creates the Theater entities in the database.
func (tcb *TheaterCreateBulk) Save(ctx context.Context) ([]*Theater, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Theater, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TheaterMutation)
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
func (tcb *TheaterCreateBulk) SaveX(ctx context.Context) []*Theater {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TheaterCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TheaterCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
