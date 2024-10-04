// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dongwlin/anime-list/internal/ent/anime"
	"github.com/dongwlin/anime-list/internal/ent/predicate"
	"github.com/dongwlin/anime-list/internal/ent/season"
	"github.com/dongwlin/anime-list/internal/ent/theater"
)

// AnimeUpdate is the builder for updating Anime entities.
type AnimeUpdate struct {
	config
	hooks    []Hook
	mutation *AnimeMutation
}

// Where appends a list predicates to the AnimeUpdate builder.
func (au *AnimeUpdate) Where(ps ...predicate.Anime) *AnimeUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AnimeUpdate) SetName(s string) *AnimeUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AnimeUpdate) SetNillableName(s *string) *AnimeUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetDescription sets the "description" field.
func (au *AnimeUpdate) SetDescription(s string) *AnimeUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (au *AnimeUpdate) SetNillableDescription(s *string) *AnimeUpdate {
	if s != nil {
		au.SetDescription(*s)
	}
	return au
}

// SetStatus sets the "status" field.
func (au *AnimeUpdate) SetStatus(i int) *AnimeUpdate {
	au.mutation.ResetStatus()
	au.mutation.SetStatus(i)
	return au
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (au *AnimeUpdate) SetNillableStatus(i *int) *AnimeUpdate {
	if i != nil {
		au.SetStatus(*i)
	}
	return au
}

// AddStatus adds i to the "status" field.
func (au *AnimeUpdate) AddStatus(i int) *AnimeUpdate {
	au.mutation.AddStatus(i)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AnimeUpdate) SetCreatedAt(t time.Time) *AnimeUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AnimeUpdate) SetNillableCreatedAt(t *time.Time) *AnimeUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// AddSeasonIDs adds the "seasons" edge to the Season entity by IDs.
func (au *AnimeUpdate) AddSeasonIDs(ids ...int) *AnimeUpdate {
	au.mutation.AddSeasonIDs(ids...)
	return au
}

// AddSeasons adds the "seasons" edges to the Season entity.
func (au *AnimeUpdate) AddSeasons(s ...*Season) *AnimeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return au.AddSeasonIDs(ids...)
}

// AddTheaterIDs adds the "theaters" edge to the Theater entity by IDs.
func (au *AnimeUpdate) AddTheaterIDs(ids ...int) *AnimeUpdate {
	au.mutation.AddTheaterIDs(ids...)
	return au
}

// AddTheaters adds the "theaters" edges to the Theater entity.
func (au *AnimeUpdate) AddTheaters(t ...*Theater) *AnimeUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTheaterIDs(ids...)
}

// Mutation returns the AnimeMutation object of the builder.
func (au *AnimeUpdate) Mutation() *AnimeMutation {
	return au.mutation
}

// ClearSeasons clears all "seasons" edges to the Season entity.
func (au *AnimeUpdate) ClearSeasons() *AnimeUpdate {
	au.mutation.ClearSeasons()
	return au
}

// RemoveSeasonIDs removes the "seasons" edge to Season entities by IDs.
func (au *AnimeUpdate) RemoveSeasonIDs(ids ...int) *AnimeUpdate {
	au.mutation.RemoveSeasonIDs(ids...)
	return au
}

// RemoveSeasons removes "seasons" edges to Season entities.
func (au *AnimeUpdate) RemoveSeasons(s ...*Season) *AnimeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return au.RemoveSeasonIDs(ids...)
}

// ClearTheaters clears all "theaters" edges to the Theater entity.
func (au *AnimeUpdate) ClearTheaters() *AnimeUpdate {
	au.mutation.ClearTheaters()
	return au
}

// RemoveTheaterIDs removes the "theaters" edge to Theater entities by IDs.
func (au *AnimeUpdate) RemoveTheaterIDs(ids ...int) *AnimeUpdate {
	au.mutation.RemoveTheaterIDs(ids...)
	return au
}

// RemoveTheaters removes "theaters" edges to Theater entities.
func (au *AnimeUpdate) RemoveTheaters(t ...*Theater) *AnimeUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTheaterIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AnimeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AnimeUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AnimeUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AnimeUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AnimeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(anime.Table, anime.Columns, sqlgraph.NewFieldSpec(anime.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(anime.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.SetField(anime.FieldDescription, field.TypeString, value)
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.SetField(anime.FieldStatus, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedStatus(); ok {
		_spec.AddField(anime.FieldStatus, field.TypeInt, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(anime.FieldCreatedAt, field.TypeTime, value)
	}
	if au.mutation.SeasonsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   anime.SeasonsTable,
			Columns: []string{anime.SeasonsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedSeasonsIDs(); len(nodes) > 0 && !au.mutation.SeasonsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   anime.SeasonsTable,
			Columns: []string{anime.SeasonsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.SeasonsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   anime.SeasonsTable,
			Columns: []string{anime.SeasonsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.TheatersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   anime.TheatersTable,
			Columns: []string{anime.TheatersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(theater.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedTheatersIDs(); len(nodes) > 0 && !au.mutation.TheatersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   anime.TheatersTable,
			Columns: []string{anime.TheatersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(theater.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TheatersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   anime.TheatersTable,
			Columns: []string{anime.TheatersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(theater.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{anime.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AnimeUpdateOne is the builder for updating a single Anime entity.
type AnimeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AnimeMutation
}

// SetName sets the "name" field.
func (auo *AnimeUpdateOne) SetName(s string) *AnimeUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AnimeUpdateOne) SetNillableName(s *string) *AnimeUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetDescription sets the "description" field.
func (auo *AnimeUpdateOne) SetDescription(s string) *AnimeUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (auo *AnimeUpdateOne) SetNillableDescription(s *string) *AnimeUpdateOne {
	if s != nil {
		auo.SetDescription(*s)
	}
	return auo
}

// SetStatus sets the "status" field.
func (auo *AnimeUpdateOne) SetStatus(i int) *AnimeUpdateOne {
	auo.mutation.ResetStatus()
	auo.mutation.SetStatus(i)
	return auo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auo *AnimeUpdateOne) SetNillableStatus(i *int) *AnimeUpdateOne {
	if i != nil {
		auo.SetStatus(*i)
	}
	return auo
}

// AddStatus adds i to the "status" field.
func (auo *AnimeUpdateOne) AddStatus(i int) *AnimeUpdateOne {
	auo.mutation.AddStatus(i)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AnimeUpdateOne) SetCreatedAt(t time.Time) *AnimeUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AnimeUpdateOne) SetNillableCreatedAt(t *time.Time) *AnimeUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// AddSeasonIDs adds the "seasons" edge to the Season entity by IDs.
func (auo *AnimeUpdateOne) AddSeasonIDs(ids ...int) *AnimeUpdateOne {
	auo.mutation.AddSeasonIDs(ids...)
	return auo
}

// AddSeasons adds the "seasons" edges to the Season entity.
func (auo *AnimeUpdateOne) AddSeasons(s ...*Season) *AnimeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return auo.AddSeasonIDs(ids...)
}

// AddTheaterIDs adds the "theaters" edge to the Theater entity by IDs.
func (auo *AnimeUpdateOne) AddTheaterIDs(ids ...int) *AnimeUpdateOne {
	auo.mutation.AddTheaterIDs(ids...)
	return auo
}

// AddTheaters adds the "theaters" edges to the Theater entity.
func (auo *AnimeUpdateOne) AddTheaters(t ...*Theater) *AnimeUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTheaterIDs(ids...)
}

// Mutation returns the AnimeMutation object of the builder.
func (auo *AnimeUpdateOne) Mutation() *AnimeMutation {
	return auo.mutation
}

// ClearSeasons clears all "seasons" edges to the Season entity.
func (auo *AnimeUpdateOne) ClearSeasons() *AnimeUpdateOne {
	auo.mutation.ClearSeasons()
	return auo
}

// RemoveSeasonIDs removes the "seasons" edge to Season entities by IDs.
func (auo *AnimeUpdateOne) RemoveSeasonIDs(ids ...int) *AnimeUpdateOne {
	auo.mutation.RemoveSeasonIDs(ids...)
	return auo
}

// RemoveSeasons removes "seasons" edges to Season entities.
func (auo *AnimeUpdateOne) RemoveSeasons(s ...*Season) *AnimeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return auo.RemoveSeasonIDs(ids...)
}

// ClearTheaters clears all "theaters" edges to the Theater entity.
func (auo *AnimeUpdateOne) ClearTheaters() *AnimeUpdateOne {
	auo.mutation.ClearTheaters()
	return auo
}

// RemoveTheaterIDs removes the "theaters" edge to Theater entities by IDs.
func (auo *AnimeUpdateOne) RemoveTheaterIDs(ids ...int) *AnimeUpdateOne {
	auo.mutation.RemoveTheaterIDs(ids...)
	return auo
}

// RemoveTheaters removes "theaters" edges to Theater entities.
func (auo *AnimeUpdateOne) RemoveTheaters(t ...*Theater) *AnimeUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTheaterIDs(ids...)
}

// Where appends a list predicates to the AnimeUpdate builder.
func (auo *AnimeUpdateOne) Where(ps ...predicate.Anime) *AnimeUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AnimeUpdateOne) Select(field string, fields ...string) *AnimeUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Anime entity.
func (auo *AnimeUpdateOne) Save(ctx context.Context) (*Anime, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AnimeUpdateOne) SaveX(ctx context.Context) *Anime {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AnimeUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AnimeUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AnimeUpdateOne) sqlSave(ctx context.Context) (_node *Anime, err error) {
	_spec := sqlgraph.NewUpdateSpec(anime.Table, anime.Columns, sqlgraph.NewFieldSpec(anime.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Anime.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, anime.FieldID)
		for _, f := range fields {
			if !anime.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != anime.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(anime.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.SetField(anime.FieldDescription, field.TypeString, value)
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.SetField(anime.FieldStatus, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedStatus(); ok {
		_spec.AddField(anime.FieldStatus, field.TypeInt, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(anime.FieldCreatedAt, field.TypeTime, value)
	}
	if auo.mutation.SeasonsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   anime.SeasonsTable,
			Columns: []string{anime.SeasonsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedSeasonsIDs(); len(nodes) > 0 && !auo.mutation.SeasonsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   anime.SeasonsTable,
			Columns: []string{anime.SeasonsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.SeasonsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   anime.SeasonsTable,
			Columns: []string{anime.SeasonsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.TheatersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   anime.TheatersTable,
			Columns: []string{anime.TheatersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(theater.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedTheatersIDs(); len(nodes) > 0 && !auo.mutation.TheatersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   anime.TheatersTable,
			Columns: []string{anime.TheatersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(theater.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TheatersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   anime.TheatersTable,
			Columns: []string{anime.TheatersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(theater.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Anime{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{anime.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
