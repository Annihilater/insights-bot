// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nekomeowww/insights-bot/ent/feedbackchathistoriesrecapsreactions"
	"github.com/nekomeowww/insights-bot/ent/internal"
	"github.com/nekomeowww/insights-bot/ent/predicate"
)

// FeedbackChatHistoriesRecapsReactionsDelete is the builder for deleting a FeedbackChatHistoriesRecapsReactions entity.
type FeedbackChatHistoriesRecapsReactionsDelete struct {
	config
	hooks    []Hook
	mutation *FeedbackChatHistoriesRecapsReactionsMutation
}

// Where appends a list predicates to the FeedbackChatHistoriesRecapsReactionsDelete builder.
func (fchrrd *FeedbackChatHistoriesRecapsReactionsDelete) Where(ps ...predicate.FeedbackChatHistoriesRecapsReactions) *FeedbackChatHistoriesRecapsReactionsDelete {
	fchrrd.mutation.Where(ps...)
	return fchrrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fchrrd *FeedbackChatHistoriesRecapsReactionsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fchrrd.sqlExec, fchrrd.mutation, fchrrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fchrrd *FeedbackChatHistoriesRecapsReactionsDelete) ExecX(ctx context.Context) int {
	n, err := fchrrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fchrrd *FeedbackChatHistoriesRecapsReactionsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(feedbackchathistoriesrecapsreactions.Table, sqlgraph.NewFieldSpec(feedbackchathistoriesrecapsreactions.FieldID, field.TypeUUID))
	_spec.Node.Schema = fchrrd.schemaConfig.FeedbackChatHistoriesRecapsReactions
	ctx = internal.NewSchemaConfigContext(ctx, fchrrd.schemaConfig)
	if ps := fchrrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fchrrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fchrrd.mutation.done = true
	return affected, err
}

// FeedbackChatHistoriesRecapsReactionsDeleteOne is the builder for deleting a single FeedbackChatHistoriesRecapsReactions entity.
type FeedbackChatHistoriesRecapsReactionsDeleteOne struct {
	fchrrd *FeedbackChatHistoriesRecapsReactionsDelete
}

// Where appends a list predicates to the FeedbackChatHistoriesRecapsReactionsDelete builder.
func (fchrrdo *FeedbackChatHistoriesRecapsReactionsDeleteOne) Where(ps ...predicate.FeedbackChatHistoriesRecapsReactions) *FeedbackChatHistoriesRecapsReactionsDeleteOne {
	fchrrdo.fchrrd.mutation.Where(ps...)
	return fchrrdo
}

// Exec executes the deletion query.
func (fchrrdo *FeedbackChatHistoriesRecapsReactionsDeleteOne) Exec(ctx context.Context) error {
	n, err := fchrrdo.fchrrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{feedbackchathistoriesrecapsreactions.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fchrrdo *FeedbackChatHistoriesRecapsReactionsDeleteOne) ExecX(ctx context.Context) {
	if err := fchrrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
