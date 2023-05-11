// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/nekomeowww/insights-bot/ent/chathistories"
)

// ChatHistories is the model entity for the ChatHistories schema.
type ChatHistories struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ChatID holds the value of the "chat_id" field.
	ChatID int64 `json:"chat_id,omitempty"`
	// ChatTitle holds the value of the "chat_title" field.
	ChatTitle string `json:"chat_title,omitempty"`
	// MessageID holds the value of the "message_id" field.
	MessageID int64 `json:"message_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// FullName holds the value of the "full_name" field.
	FullName string `json:"full_name,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// RepliedToMessageID holds the value of the "replied_to_message_id" field.
	RepliedToMessageID int64 `json:"replied_to_message_id,omitempty"`
	// RepliedToUserID holds the value of the "replied_to_user_id" field.
	RepliedToUserID int64 `json:"replied_to_user_id,omitempty"`
	// RepliedToFullName holds the value of the "replied_to_full_name" field.
	RepliedToFullName string `json:"replied_to_full_name,omitempty"`
	// RepliedToUsername holds the value of the "replied_to_username" field.
	RepliedToUsername string `json:"replied_to_username,omitempty"`
	// RepliedToText holds the value of the "replied_to_text" field.
	RepliedToText string `json:"replied_to_text,omitempty"`
	// ChattedAt holds the value of the "chatted_at" field.
	ChattedAt int64 `json:"chatted_at,omitempty"`
	// Embedded holds the value of the "embedded" field.
	Embedded bool `json:"embedded,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    int64 `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ChatHistories) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case chathistories.FieldEmbedded:
			values[i] = new(sql.NullBool)
		case chathistories.FieldChatID, chathistories.FieldMessageID, chathistories.FieldUserID, chathistories.FieldRepliedToMessageID, chathistories.FieldRepliedToUserID, chathistories.FieldChattedAt, chathistories.FieldCreatedAt, chathistories.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case chathistories.FieldChatTitle, chathistories.FieldUsername, chathistories.FieldFullName, chathistories.FieldText, chathistories.FieldRepliedToFullName, chathistories.FieldRepliedToUsername, chathistories.FieldRepliedToText:
			values[i] = new(sql.NullString)
		case chathistories.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChatHistories fields.
func (ch *ChatHistories) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chathistories.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ch.ID = *value
			}
		case chathistories.FieldChatID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field chat_id", values[i])
			} else if value.Valid {
				ch.ChatID = value.Int64
			}
		case chathistories.FieldChatTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chat_title", values[i])
			} else if value.Valid {
				ch.ChatTitle = value.String
			}
		case chathistories.FieldMessageID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field message_id", values[i])
			} else if value.Valid {
				ch.MessageID = value.Int64
			}
		case chathistories.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ch.UserID = value.Int64
			}
		case chathistories.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				ch.Username = value.String
			}
		case chathistories.FieldFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_name", values[i])
			} else if value.Valid {
				ch.FullName = value.String
			}
		case chathistories.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				ch.Text = value.String
			}
		case chathistories.FieldRepliedToMessageID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field replied_to_message_id", values[i])
			} else if value.Valid {
				ch.RepliedToMessageID = value.Int64
			}
		case chathistories.FieldRepliedToUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field replied_to_user_id", values[i])
			} else if value.Valid {
				ch.RepliedToUserID = value.Int64
			}
		case chathistories.FieldRepliedToFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field replied_to_full_name", values[i])
			} else if value.Valid {
				ch.RepliedToFullName = value.String
			}
		case chathistories.FieldRepliedToUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field replied_to_username", values[i])
			} else if value.Valid {
				ch.RepliedToUsername = value.String
			}
		case chathistories.FieldRepliedToText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field replied_to_text", values[i])
			} else if value.Valid {
				ch.RepliedToText = value.String
			}
		case chathistories.FieldChattedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field chatted_at", values[i])
			} else if value.Valid {
				ch.ChattedAt = value.Int64
			}
		case chathistories.FieldEmbedded:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field embedded", values[i])
			} else if value.Valid {
				ch.Embedded = value.Bool
			}
		case chathistories.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ch.CreatedAt = value.Int64
			}
		case chathistories.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ch.UpdatedAt = value.Int64
			}
		default:
			ch.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ChatHistories.
// This includes values selected through modifiers, order, etc.
func (ch *ChatHistories) Value(name string) (ent.Value, error) {
	return ch.selectValues.Get(name)
}

// Update returns a builder for updating this ChatHistories.
// Note that you need to call ChatHistories.Unwrap() before calling this method if this ChatHistories
// was returned from a transaction, and the transaction was committed or rolled back.
func (ch *ChatHistories) Update() *ChatHistoriesUpdateOne {
	return NewChatHistoriesClient(ch.config).UpdateOne(ch)
}

// Unwrap unwraps the ChatHistories entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ch *ChatHistories) Unwrap() *ChatHistories {
	_tx, ok := ch.config.driver.(*txDriver)
	if !ok {
		panic("ent: ChatHistories is not a transactional entity")
	}
	ch.config.driver = _tx.drv
	return ch
}

// String implements the fmt.Stringer.
func (ch *ChatHistories) String() string {
	var builder strings.Builder
	builder.WriteString("ChatHistories(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ch.ID))
	builder.WriteString("chat_id=")
	builder.WriteString(fmt.Sprintf("%v", ch.ChatID))
	builder.WriteString(", ")
	builder.WriteString("chat_title=")
	builder.WriteString(ch.ChatTitle)
	builder.WriteString(", ")
	builder.WriteString("message_id=")
	builder.WriteString(fmt.Sprintf("%v", ch.MessageID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ch.UserID))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(ch.Username)
	builder.WriteString(", ")
	builder.WriteString("full_name=")
	builder.WriteString(ch.FullName)
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(ch.Text)
	builder.WriteString(", ")
	builder.WriteString("replied_to_message_id=")
	builder.WriteString(fmt.Sprintf("%v", ch.RepliedToMessageID))
	builder.WriteString(", ")
	builder.WriteString("replied_to_user_id=")
	builder.WriteString(fmt.Sprintf("%v", ch.RepliedToUserID))
	builder.WriteString(", ")
	builder.WriteString("replied_to_full_name=")
	builder.WriteString(ch.RepliedToFullName)
	builder.WriteString(", ")
	builder.WriteString("replied_to_username=")
	builder.WriteString(ch.RepliedToUsername)
	builder.WriteString(", ")
	builder.WriteString("replied_to_text=")
	builder.WriteString(ch.RepliedToText)
	builder.WriteString(", ")
	builder.WriteString("chatted_at=")
	builder.WriteString(fmt.Sprintf("%v", ch.ChattedAt))
	builder.WriteString(", ")
	builder.WriteString("embedded=")
	builder.WriteString(fmt.Sprintf("%v", ch.Embedded))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ch.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ch.UpdatedAt))
	builder.WriteByte(')')
	return builder.String()
}

// ChatHistoriesSlice is a parsable slice of ChatHistories.
type ChatHistoriesSlice []*ChatHistories
