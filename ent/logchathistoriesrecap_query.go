// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/nekomeowww/insights-bot/ent/internal"
	"github.com/nekomeowww/insights-bot/ent/logchathistoriesrecap"
	"github.com/nekomeowww/insights-bot/ent/predicate"
)

// LogChatHistoriesRecapQuery is the builder for querying LogChatHistoriesRecap entities.
type LogChatHistoriesRecapQuery struct {
	config
	ctx        *QueryContext
	order      []logchathistoriesrecap.OrderOption
	inters     []Interceptor
	predicates []predicate.LogChatHistoriesRecap
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LogChatHistoriesRecapQuery builder.
func (lchrq *LogChatHistoriesRecapQuery) Where(ps ...predicate.LogChatHistoriesRecap) *LogChatHistoriesRecapQuery {
	lchrq.predicates = append(lchrq.predicates, ps...)
	return lchrq
}

// Limit the number of records to be returned by this query.
func (lchrq *LogChatHistoriesRecapQuery) Limit(limit int) *LogChatHistoriesRecapQuery {
	lchrq.ctx.Limit = &limit
	return lchrq
}

// Offset to start from.
func (lchrq *LogChatHistoriesRecapQuery) Offset(offset int) *LogChatHistoriesRecapQuery {
	lchrq.ctx.Offset = &offset
	return lchrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lchrq *LogChatHistoriesRecapQuery) Unique(unique bool) *LogChatHistoriesRecapQuery {
	lchrq.ctx.Unique = &unique
	return lchrq
}

// Order specifies how the records should be ordered.
func (lchrq *LogChatHistoriesRecapQuery) Order(o ...logchathistoriesrecap.OrderOption) *LogChatHistoriesRecapQuery {
	lchrq.order = append(lchrq.order, o...)
	return lchrq
}

// First returns the first LogChatHistoriesRecap entity from the query.
// Returns a *NotFoundError when no LogChatHistoriesRecap was found.
func (lchrq *LogChatHistoriesRecapQuery) First(ctx context.Context) (*LogChatHistoriesRecap, error) {
	nodes, err := lchrq.Limit(1).All(setContextOp(ctx, lchrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{logchathistoriesrecap.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lchrq *LogChatHistoriesRecapQuery) FirstX(ctx context.Context) *LogChatHistoriesRecap {
	node, err := lchrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LogChatHistoriesRecap ID from the query.
// Returns a *NotFoundError when no LogChatHistoriesRecap ID was found.
func (lchrq *LogChatHistoriesRecapQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lchrq.Limit(1).IDs(setContextOp(ctx, lchrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{logchathistoriesrecap.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lchrq *LogChatHistoriesRecapQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := lchrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LogChatHistoriesRecap entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LogChatHistoriesRecap entity is found.
// Returns a *NotFoundError when no LogChatHistoriesRecap entities are found.
func (lchrq *LogChatHistoriesRecapQuery) Only(ctx context.Context) (*LogChatHistoriesRecap, error) {
	nodes, err := lchrq.Limit(2).All(setContextOp(ctx, lchrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{logchathistoriesrecap.Label}
	default:
		return nil, &NotSingularError{logchathistoriesrecap.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lchrq *LogChatHistoriesRecapQuery) OnlyX(ctx context.Context) *LogChatHistoriesRecap {
	node, err := lchrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LogChatHistoriesRecap ID in the query.
// Returns a *NotSingularError when more than one LogChatHistoriesRecap ID is found.
// Returns a *NotFoundError when no entities are found.
func (lchrq *LogChatHistoriesRecapQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lchrq.Limit(2).IDs(setContextOp(ctx, lchrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{logchathistoriesrecap.Label}
	default:
		err = &NotSingularError{logchathistoriesrecap.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lchrq *LogChatHistoriesRecapQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := lchrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LogChatHistoriesRecaps.
func (lchrq *LogChatHistoriesRecapQuery) All(ctx context.Context) ([]*LogChatHistoriesRecap, error) {
	ctx = setContextOp(ctx, lchrq.ctx, "All")
	if err := lchrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LogChatHistoriesRecap, *LogChatHistoriesRecapQuery]()
	return withInterceptors[[]*LogChatHistoriesRecap](ctx, lchrq, qr, lchrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lchrq *LogChatHistoriesRecapQuery) AllX(ctx context.Context) []*LogChatHistoriesRecap {
	nodes, err := lchrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LogChatHistoriesRecap IDs.
func (lchrq *LogChatHistoriesRecapQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if lchrq.ctx.Unique == nil && lchrq.path != nil {
		lchrq.Unique(true)
	}
	ctx = setContextOp(ctx, lchrq.ctx, "IDs")
	if err = lchrq.Select(logchathistoriesrecap.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lchrq *LogChatHistoriesRecapQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := lchrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lchrq *LogChatHistoriesRecapQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lchrq.ctx, "Count")
	if err := lchrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lchrq, querierCount[*LogChatHistoriesRecapQuery](), lchrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lchrq *LogChatHistoriesRecapQuery) CountX(ctx context.Context) int {
	count, err := lchrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lchrq *LogChatHistoriesRecapQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lchrq.ctx, "Exist")
	switch _, err := lchrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lchrq *LogChatHistoriesRecapQuery) ExistX(ctx context.Context) bool {
	exist, err := lchrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LogChatHistoriesRecapQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lchrq *LogChatHistoriesRecapQuery) Clone() *LogChatHistoriesRecapQuery {
	if lchrq == nil {
		return nil
	}
	return &LogChatHistoriesRecapQuery{
		config:     lchrq.config,
		ctx:        lchrq.ctx.Clone(),
		order:      append([]logchathistoriesrecap.OrderOption{}, lchrq.order...),
		inters:     append([]Interceptor{}, lchrq.inters...),
		predicates: append([]predicate.LogChatHistoriesRecap{}, lchrq.predicates...),
		// clone intermediate query.
		sql:  lchrq.sql.Clone(),
		path: lchrq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ChatID int64 `json:"chat_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LogChatHistoriesRecap.Query().
//		GroupBy(logchathistoriesrecap.FieldChatID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lchrq *LogChatHistoriesRecapQuery) GroupBy(field string, fields ...string) *LogChatHistoriesRecapGroupBy {
	lchrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LogChatHistoriesRecapGroupBy{build: lchrq}
	grbuild.flds = &lchrq.ctx.Fields
	grbuild.label = logchathistoriesrecap.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ChatID int64 `json:"chat_id,omitempty"`
//	}
//
//	client.LogChatHistoriesRecap.Query().
//		Select(logchathistoriesrecap.FieldChatID).
//		Scan(ctx, &v)
func (lchrq *LogChatHistoriesRecapQuery) Select(fields ...string) *LogChatHistoriesRecapSelect {
	lchrq.ctx.Fields = append(lchrq.ctx.Fields, fields...)
	sbuild := &LogChatHistoriesRecapSelect{LogChatHistoriesRecapQuery: lchrq}
	sbuild.label = logchathistoriesrecap.Label
	sbuild.flds, sbuild.scan = &lchrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LogChatHistoriesRecapSelect configured with the given aggregations.
func (lchrq *LogChatHistoriesRecapQuery) Aggregate(fns ...AggregateFunc) *LogChatHistoriesRecapSelect {
	return lchrq.Select().Aggregate(fns...)
}

func (lchrq *LogChatHistoriesRecapQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lchrq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lchrq); err != nil {
				return err
			}
		}
	}
	for _, f := range lchrq.ctx.Fields {
		if !logchathistoriesrecap.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lchrq.path != nil {
		prev, err := lchrq.path(ctx)
		if err != nil {
			return err
		}
		lchrq.sql = prev
	}
	return nil
}

func (lchrq *LogChatHistoriesRecapQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LogChatHistoriesRecap, error) {
	var (
		nodes = []*LogChatHistoriesRecap{}
		_spec = lchrq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LogChatHistoriesRecap).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LogChatHistoriesRecap{config: lchrq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = lchrq.schemaConfig.LogChatHistoriesRecap
	ctx = internal.NewSchemaConfigContext(ctx, lchrq.schemaConfig)
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lchrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lchrq *LogChatHistoriesRecapQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lchrq.querySpec()
	_spec.Node.Schema = lchrq.schemaConfig.LogChatHistoriesRecap
	ctx = internal.NewSchemaConfigContext(ctx, lchrq.schemaConfig)
	_spec.Node.Columns = lchrq.ctx.Fields
	if len(lchrq.ctx.Fields) > 0 {
		_spec.Unique = lchrq.ctx.Unique != nil && *lchrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lchrq.driver, _spec)
}

func (lchrq *LogChatHistoriesRecapQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(logchathistoriesrecap.Table, logchathistoriesrecap.Columns, sqlgraph.NewFieldSpec(logchathistoriesrecap.FieldID, field.TypeUUID))
	_spec.From = lchrq.sql
	if unique := lchrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lchrq.path != nil {
		_spec.Unique = true
	}
	if fields := lchrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logchathistoriesrecap.FieldID)
		for i := range fields {
			if fields[i] != logchathistoriesrecap.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lchrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lchrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lchrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lchrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lchrq *LogChatHistoriesRecapQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lchrq.driver.Dialect())
	t1 := builder.Table(logchathistoriesrecap.Table)
	columns := lchrq.ctx.Fields
	if len(columns) == 0 {
		columns = logchathistoriesrecap.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lchrq.sql != nil {
		selector = lchrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lchrq.ctx.Unique != nil && *lchrq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(lchrq.schemaConfig.LogChatHistoriesRecap)
	ctx = internal.NewSchemaConfigContext(ctx, lchrq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range lchrq.predicates {
		p(selector)
	}
	for _, p := range lchrq.order {
		p(selector)
	}
	if offset := lchrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lchrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LogChatHistoriesRecapGroupBy is the group-by builder for LogChatHistoriesRecap entities.
type LogChatHistoriesRecapGroupBy struct {
	selector
	build *LogChatHistoriesRecapQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lchrgb *LogChatHistoriesRecapGroupBy) Aggregate(fns ...AggregateFunc) *LogChatHistoriesRecapGroupBy {
	lchrgb.fns = append(lchrgb.fns, fns...)
	return lchrgb
}

// Scan applies the selector query and scans the result into the given value.
func (lchrgb *LogChatHistoriesRecapGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lchrgb.build.ctx, "GroupBy")
	if err := lchrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LogChatHistoriesRecapQuery, *LogChatHistoriesRecapGroupBy](ctx, lchrgb.build, lchrgb, lchrgb.build.inters, v)
}

func (lchrgb *LogChatHistoriesRecapGroupBy) sqlScan(ctx context.Context, root *LogChatHistoriesRecapQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lchrgb.fns))
	for _, fn := range lchrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lchrgb.flds)+len(lchrgb.fns))
		for _, f := range *lchrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lchrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lchrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LogChatHistoriesRecapSelect is the builder for selecting fields of LogChatHistoriesRecap entities.
type LogChatHistoriesRecapSelect struct {
	*LogChatHistoriesRecapQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lchrs *LogChatHistoriesRecapSelect) Aggregate(fns ...AggregateFunc) *LogChatHistoriesRecapSelect {
	lchrs.fns = append(lchrs.fns, fns...)
	return lchrs
}

// Scan applies the selector query and scans the result into the given value.
func (lchrs *LogChatHistoriesRecapSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lchrs.ctx, "Select")
	if err := lchrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LogChatHistoriesRecapQuery, *LogChatHistoriesRecapSelect](ctx, lchrs.LogChatHistoriesRecapQuery, lchrs, lchrs.inters, v)
}

func (lchrs *LogChatHistoriesRecapSelect) sqlScan(ctx context.Context, root *LogChatHistoriesRecapQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lchrs.fns))
	for _, fn := range lchrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lchrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lchrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
