// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/nekomeowww/insights-bot/ent/feedbacksummarizationsreactions"
)

// FeedbackSummarizationsReactionsCreate is the builder for creating a FeedbackSummarizationsReactions entity.
type FeedbackSummarizationsReactionsCreate struct {
	config
	mutation *FeedbackSummarizationsReactionsMutation
	hooks    []Hook
}

// SetChatID sets the "chat_id" field.
func (fsrc *FeedbackSummarizationsReactionsCreate) SetChatID(i int64) *FeedbackSummarizationsReactionsCreate {
	fsrc.mutation.SetChatID(i)
	return fsrc
}

// SetNillableChatID sets the "chat_id" field if the given value is not nil.
func (fsrc *FeedbackSummarizationsReactionsCreate) SetNillableChatID(i *int64) *FeedbackSummarizationsReactionsCreate {
	if i != nil {
		fsrc.SetChatID(*i)
	}
	return fsrc
}

// SetLogID sets the "log_id" field.
func (fsrc *FeedbackSummarizationsReactionsCreate) SetLogID(u uuid.UUID) *FeedbackSummarizationsReactionsCreate {
	fsrc.mutation.SetLogID(u)
	return fsrc
}

// SetNillableLogID sets the "log_id" field if the given value is not nil.
func (fsrc *FeedbackSummarizationsReactionsCreate) SetNillableLogID(u *uuid.UUID) *FeedbackSummarizationsReactionsCreate {
	if u != nil {
		fsrc.SetLogID(*u)
	}
	return fsrc
}

// SetUserID sets the "user_id" field.
func (fsrc *FeedbackSummarizationsReactionsCreate) SetUserID(i int64) *FeedbackSummarizationsReactionsCreate {
	fsrc.mutation.SetUserID(i)
	return fsrc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (fsrc *FeedbackSummarizationsReactionsCreate) SetNillableUserID(i *int64) *FeedbackSummarizationsReactionsCreate {
	if i != nil {
		fsrc.SetUserID(*i)
	}
	return fsrc
}

// SetType sets the "type" field.
func (fsrc *FeedbackSummarizationsReactionsCreate) SetType(f feedbacksummarizationsreactions.Type) *FeedbackSummarizationsReactionsCreate {
	fsrc.mutation.SetType(f)
	return fsrc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (fsrc *FeedbackSummarizationsReactionsCreate) SetNillableType(f *feedbacksummarizationsreactions.Type) *FeedbackSummarizationsReactionsCreate {
	if f != nil {
		fsrc.SetType(*f)
	}
	return fsrc
}

// SetCreatedAt sets the "created_at" field.
func (fsrc *FeedbackSummarizationsReactionsCreate) SetCreatedAt(i int64) *FeedbackSummarizationsReactionsCreate {
	fsrc.mutation.SetCreatedAt(i)
	return fsrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fsrc *FeedbackSummarizationsReactionsCreate) SetNillableCreatedAt(i *int64) *FeedbackSummarizationsReactionsCreate {
	if i != nil {
		fsrc.SetCreatedAt(*i)
	}
	return fsrc
}

// SetUpdatedAt sets the "updated_at" field.
func (fsrc *FeedbackSummarizationsReactionsCreate) SetUpdatedAt(i int64) *FeedbackSummarizationsReactionsCreate {
	fsrc.mutation.SetUpdatedAt(i)
	return fsrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fsrc *FeedbackSummarizationsReactionsCreate) SetNillableUpdatedAt(i *int64) *FeedbackSummarizationsReactionsCreate {
	if i != nil {
		fsrc.SetUpdatedAt(*i)
	}
	return fsrc
}

// SetID sets the "id" field.
func (fsrc *FeedbackSummarizationsReactionsCreate) SetID(u uuid.UUID) *FeedbackSummarizationsReactionsCreate {
	fsrc.mutation.SetID(u)
	return fsrc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fsrc *FeedbackSummarizationsReactionsCreate) SetNillableID(u *uuid.UUID) *FeedbackSummarizationsReactionsCreate {
	if u != nil {
		fsrc.SetID(*u)
	}
	return fsrc
}

// Mutation returns the FeedbackSummarizationsReactionsMutation object of the builder.
func (fsrc *FeedbackSummarizationsReactionsCreate) Mutation() *FeedbackSummarizationsReactionsMutation {
	return fsrc.mutation
}

// Save creates the FeedbackSummarizationsReactions in the database.
func (fsrc *FeedbackSummarizationsReactionsCreate) Save(ctx context.Context) (*FeedbackSummarizationsReactions, error) {
	fsrc.defaults()
	return withHooks(ctx, fsrc.sqlSave, fsrc.mutation, fsrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fsrc *FeedbackSummarizationsReactionsCreate) SaveX(ctx context.Context) *FeedbackSummarizationsReactions {
	v, err := fsrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fsrc *FeedbackSummarizationsReactionsCreate) Exec(ctx context.Context) error {
	_, err := fsrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsrc *FeedbackSummarizationsReactionsCreate) ExecX(ctx context.Context) {
	if err := fsrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fsrc *FeedbackSummarizationsReactionsCreate) defaults() {
	if _, ok := fsrc.mutation.ChatID(); !ok {
		v := feedbacksummarizationsreactions.DefaultChatID
		fsrc.mutation.SetChatID(v)
	}
	if _, ok := fsrc.mutation.LogID(); !ok {
		v := feedbacksummarizationsreactions.DefaultLogID()
		fsrc.mutation.SetLogID(v)
	}
	if _, ok := fsrc.mutation.UserID(); !ok {
		v := feedbacksummarizationsreactions.DefaultUserID
		fsrc.mutation.SetUserID(v)
	}
	if _, ok := fsrc.mutation.GetType(); !ok {
		v := feedbacksummarizationsreactions.DefaultType
		fsrc.mutation.SetType(v)
	}
	if _, ok := fsrc.mutation.CreatedAt(); !ok {
		v := feedbacksummarizationsreactions.DefaultCreatedAt()
		fsrc.mutation.SetCreatedAt(v)
	}
	if _, ok := fsrc.mutation.UpdatedAt(); !ok {
		v := feedbacksummarizationsreactions.DefaultUpdatedAt()
		fsrc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fsrc.mutation.ID(); !ok {
		v := feedbacksummarizationsreactions.DefaultID()
		fsrc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fsrc *FeedbackSummarizationsReactionsCreate) check() error {
	if _, ok := fsrc.mutation.ChatID(); !ok {
		return &ValidationError{Name: "chat_id", err: errors.New(`ent: missing required field "FeedbackSummarizationsReactions.chat_id"`)}
	}
	if _, ok := fsrc.mutation.LogID(); !ok {
		return &ValidationError{Name: "log_id", err: errors.New(`ent: missing required field "FeedbackSummarizationsReactions.log_id"`)}
	}
	if _, ok := fsrc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "FeedbackSummarizationsReactions.user_id"`)}
	}
	if _, ok := fsrc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "FeedbackSummarizationsReactions.type"`)}
	}
	if v, ok := fsrc.mutation.GetType(); ok {
		if err := feedbacksummarizationsreactions.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "FeedbackSummarizationsReactions.type": %w`, err)}
		}
	}
	if _, ok := fsrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FeedbackSummarizationsReactions.created_at"`)}
	}
	if _, ok := fsrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FeedbackSummarizationsReactions.updated_at"`)}
	}
	return nil
}

func (fsrc *FeedbackSummarizationsReactionsCreate) sqlSave(ctx context.Context) (*FeedbackSummarizationsReactions, error) {
	if err := fsrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fsrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fsrc.driver, _spec); err != nil {
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
	fsrc.mutation.id = &_node.ID
	fsrc.mutation.done = true
	return _node, nil
}

func (fsrc *FeedbackSummarizationsReactionsCreate) createSpec() (*FeedbackSummarizationsReactions, *sqlgraph.CreateSpec) {
	var (
		_node = &FeedbackSummarizationsReactions{config: fsrc.config}
		_spec = sqlgraph.NewCreateSpec(feedbacksummarizationsreactions.Table, sqlgraph.NewFieldSpec(feedbacksummarizationsreactions.FieldID, field.TypeUUID))
	)
	_spec.Schema = fsrc.schemaConfig.FeedbackSummarizationsReactions
	if id, ok := fsrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fsrc.mutation.ChatID(); ok {
		_spec.SetField(feedbacksummarizationsreactions.FieldChatID, field.TypeInt64, value)
		_node.ChatID = value
	}
	if value, ok := fsrc.mutation.LogID(); ok {
		_spec.SetField(feedbacksummarizationsreactions.FieldLogID, field.TypeUUID, value)
		_node.LogID = value
	}
	if value, ok := fsrc.mutation.UserID(); ok {
		_spec.SetField(feedbacksummarizationsreactions.FieldUserID, field.TypeInt64, value)
		_node.UserID = value
	}
	if value, ok := fsrc.mutation.GetType(); ok {
		_spec.SetField(feedbacksummarizationsreactions.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := fsrc.mutation.CreatedAt(); ok {
		_spec.SetField(feedbacksummarizationsreactions.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := fsrc.mutation.UpdatedAt(); ok {
		_spec.SetField(feedbacksummarizationsreactions.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// FeedbackSummarizationsReactionsCreateBulk is the builder for creating many FeedbackSummarizationsReactions entities in bulk.
type FeedbackSummarizationsReactionsCreateBulk struct {
	config
	builders []*FeedbackSummarizationsReactionsCreate
}

// Save creates the FeedbackSummarizationsReactions entities in the database.
func (fsrcb *FeedbackSummarizationsReactionsCreateBulk) Save(ctx context.Context) ([]*FeedbackSummarizationsReactions, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fsrcb.builders))
	nodes := make([]*FeedbackSummarizationsReactions, len(fsrcb.builders))
	mutators := make([]Mutator, len(fsrcb.builders))
	for i := range fsrcb.builders {
		func(i int, root context.Context) {
			builder := fsrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeedbackSummarizationsReactionsMutation)
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
					_, err = mutators[i+1].Mutate(root, fsrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fsrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fsrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fsrcb *FeedbackSummarizationsReactionsCreateBulk) SaveX(ctx context.Context) []*FeedbackSummarizationsReactions {
	v, err := fsrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fsrcb *FeedbackSummarizationsReactionsCreateBulk) Exec(ctx context.Context) error {
	_, err := fsrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsrcb *FeedbackSummarizationsReactionsCreateBulk) ExecX(ctx context.Context) {
	if err := fsrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
