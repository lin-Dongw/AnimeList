// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dongwlin/anime-list/internal/ent/anime"
	"github.com/dongwlin/anime-list/internal/ent/predicate"
	"github.com/dongwlin/anime-list/internal/ent/season"
	"github.com/dongwlin/anime-list/internal/ent/theater"
)

// AnimeQuery is the builder for querying Anime entities.
type AnimeQuery struct {
	config
	ctx          *QueryContext
	order        []anime.OrderOption
	inters       []Interceptor
	predicates   []predicate.Anime
	withSeasons  *SeasonQuery
	withTheaters *TheaterQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AnimeQuery builder.
func (aq *AnimeQuery) Where(ps ...predicate.Anime) *AnimeQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AnimeQuery) Limit(limit int) *AnimeQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AnimeQuery) Offset(offset int) *AnimeQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AnimeQuery) Unique(unique bool) *AnimeQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AnimeQuery) Order(o ...anime.OrderOption) *AnimeQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QuerySeasons chains the current query on the "seasons" edge.
func (aq *AnimeQuery) QuerySeasons() *SeasonQuery {
	query := (&SeasonClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(anime.Table, anime.FieldID, selector),
			sqlgraph.To(season.Table, season.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, anime.SeasonsTable, anime.SeasonsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTheaters chains the current query on the "theaters" edge.
func (aq *AnimeQuery) QueryTheaters() *TheaterQuery {
	query := (&TheaterClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(anime.Table, anime.FieldID, selector),
			sqlgraph.To(theater.Table, theater.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, anime.TheatersTable, anime.TheatersColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Anime entity from the query.
// Returns a *NotFoundError when no Anime was found.
func (aq *AnimeQuery) First(ctx context.Context) (*Anime, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{anime.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AnimeQuery) FirstX(ctx context.Context) *Anime {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Anime ID from the query.
// Returns a *NotFoundError when no Anime ID was found.
func (aq *AnimeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{anime.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AnimeQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Anime entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Anime entity is found.
// Returns a *NotFoundError when no Anime entities are found.
func (aq *AnimeQuery) Only(ctx context.Context) (*Anime, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{anime.Label}
	default:
		return nil, &NotSingularError{anime.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AnimeQuery) OnlyX(ctx context.Context) *Anime {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Anime ID in the query.
// Returns a *NotSingularError when more than one Anime ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AnimeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{anime.Label}
	default:
		err = &NotSingularError{anime.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AnimeQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Animes.
func (aq *AnimeQuery) All(ctx context.Context) ([]*Anime, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryAll)
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Anime, *AnimeQuery]()
	return withInterceptors[[]*Anime](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AnimeQuery) AllX(ctx context.Context) []*Anime {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Anime IDs.
func (aq *AnimeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryIDs)
	if err = aq.Select(anime.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AnimeQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AnimeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryCount)
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AnimeQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AnimeQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AnimeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryExist)
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AnimeQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AnimeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AnimeQuery) Clone() *AnimeQuery {
	if aq == nil {
		return nil
	}
	return &AnimeQuery{
		config:       aq.config,
		ctx:          aq.ctx.Clone(),
		order:        append([]anime.OrderOption{}, aq.order...),
		inters:       append([]Interceptor{}, aq.inters...),
		predicates:   append([]predicate.Anime{}, aq.predicates...),
		withSeasons:  aq.withSeasons.Clone(),
		withTheaters: aq.withTheaters.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithSeasons tells the query-builder to eager-load the nodes that are connected to
// the "seasons" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AnimeQuery) WithSeasons(opts ...func(*SeasonQuery)) *AnimeQuery {
	query := (&SeasonClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withSeasons = query
	return aq
}

// WithTheaters tells the query-builder to eager-load the nodes that are connected to
// the "theaters" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AnimeQuery) WithTheaters(opts ...func(*TheaterQuery)) *AnimeQuery {
	query := (&TheaterClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withTheaters = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Anime.Query().
//		GroupBy(anime.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AnimeQuery) GroupBy(field string, fields ...string) *AnimeGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AnimeGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = anime.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Anime.Query().
//		Select(anime.FieldName).
//		Scan(ctx, &v)
func (aq *AnimeQuery) Select(fields ...string) *AnimeSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AnimeSelect{AnimeQuery: aq}
	sbuild.label = anime.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AnimeSelect configured with the given aggregations.
func (aq *AnimeQuery) Aggregate(fns ...AggregateFunc) *AnimeSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AnimeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !anime.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AnimeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Anime, error) {
	var (
		nodes       = []*Anime{}
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withSeasons != nil,
			aq.withTheaters != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Anime).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Anime{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withSeasons; query != nil {
		if err := aq.loadSeasons(ctx, query, nodes,
			func(n *Anime) { n.Edges.Seasons = []*Season{} },
			func(n *Anime, e *Season) { n.Edges.Seasons = append(n.Edges.Seasons, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withTheaters; query != nil {
		if err := aq.loadTheaters(ctx, query, nodes,
			func(n *Anime) { n.Edges.Theaters = []*Theater{} },
			func(n *Anime, e *Theater) { n.Edges.Theaters = append(n.Edges.Theaters, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AnimeQuery) loadSeasons(ctx context.Context, query *SeasonQuery, nodes []*Anime, init func(*Anime), assign func(*Anime, *Season)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Anime)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Season(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(anime.SeasonsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.anime_seasons
		if fk == nil {
			return fmt.Errorf(`foreign-key "anime_seasons" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "anime_seasons" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AnimeQuery) loadTheaters(ctx context.Context, query *TheaterQuery, nodes []*Anime, init func(*Anime), assign func(*Anime, *Theater)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Anime)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Theater(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(anime.TheatersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.anime_theaters
		if fk == nil {
			return fmt.Errorf(`foreign-key "anime_theaters" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "anime_theaters" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aq *AnimeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AnimeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(anime.Table, anime.Columns, sqlgraph.NewFieldSpec(anime.FieldID, field.TypeInt))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, anime.FieldID)
		for i := range fields {
			if fields[i] != anime.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AnimeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(anime.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = anime.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AnimeGroupBy is the group-by builder for Anime entities.
type AnimeGroupBy struct {
	selector
	build *AnimeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AnimeGroupBy) Aggregate(fns ...AggregateFunc) *AnimeGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AnimeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, ent.OpQueryGroupBy)
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AnimeQuery, *AnimeGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AnimeGroupBy) sqlScan(ctx context.Context, root *AnimeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AnimeSelect is the builder for selecting fields of Anime entities.
type AnimeSelect struct {
	*AnimeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AnimeSelect) Aggregate(fns ...AggregateFunc) *AnimeSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AnimeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, ent.OpQuerySelect)
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AnimeQuery, *AnimeSelect](ctx, as.AnimeQuery, as, as.inters, v)
}

func (as *AnimeSelect) sqlScan(ctx context.Context, root *AnimeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
