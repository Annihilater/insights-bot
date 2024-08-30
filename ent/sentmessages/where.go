// Code generated by ent, DO NOT EDIT.

package sentmessages

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/nekomeowww/insights-bot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLTE(FieldID, id))
}

// ChatID applies equality check predicate on the "chat_id" field. It's identical to ChatIDEQ.
func ChatID(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldChatID, v))
}

// MessageID applies equality check predicate on the "message_id" field. It's identical to MessageIDEQ.
func MessageID(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldMessageID, v))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldText, v))
}

// IsPinned applies equality check predicate on the "is_pinned" field. It's identical to IsPinnedEQ.
func IsPinned(v bool) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldIsPinned, v))
}

// FromPlatform applies equality check predicate on the "from_platform" field. It's identical to FromPlatformEQ.
func FromPlatform(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldFromPlatform, v))
}

// MessageType applies equality check predicate on the "message_type" field. It's identical to MessageTypeEQ.
func MessageType(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldMessageType, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldUpdatedAt, v))
}

// ChatIDEQ applies the EQ predicate on the "chat_id" field.
func ChatIDEQ(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldChatID, v))
}

// ChatIDNEQ applies the NEQ predicate on the "chat_id" field.
func ChatIDNEQ(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNEQ(FieldChatID, v))
}

// ChatIDIn applies the In predicate on the "chat_id" field.
func ChatIDIn(vs ...int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldIn(FieldChatID, vs...))
}

// ChatIDNotIn applies the NotIn predicate on the "chat_id" field.
func ChatIDNotIn(vs ...int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNotIn(FieldChatID, vs...))
}

// ChatIDGT applies the GT predicate on the "chat_id" field.
func ChatIDGT(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGT(FieldChatID, v))
}

// ChatIDGTE applies the GTE predicate on the "chat_id" field.
func ChatIDGTE(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGTE(FieldChatID, v))
}

// ChatIDLT applies the LT predicate on the "chat_id" field.
func ChatIDLT(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLT(FieldChatID, v))
}

// ChatIDLTE applies the LTE predicate on the "chat_id" field.
func ChatIDLTE(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLTE(FieldChatID, v))
}

// MessageIDEQ applies the EQ predicate on the "message_id" field.
func MessageIDEQ(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldMessageID, v))
}

// MessageIDNEQ applies the NEQ predicate on the "message_id" field.
func MessageIDNEQ(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNEQ(FieldMessageID, v))
}

// MessageIDIn applies the In predicate on the "message_id" field.
func MessageIDIn(vs ...int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldIn(FieldMessageID, vs...))
}

// MessageIDNotIn applies the NotIn predicate on the "message_id" field.
func MessageIDNotIn(vs ...int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNotIn(FieldMessageID, vs...))
}

// MessageIDGT applies the GT predicate on the "message_id" field.
func MessageIDGT(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGT(FieldMessageID, v))
}

// MessageIDGTE applies the GTE predicate on the "message_id" field.
func MessageIDGTE(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGTE(FieldMessageID, v))
}

// MessageIDLT applies the LT predicate on the "message_id" field.
func MessageIDLT(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLT(FieldMessageID, v))
}

// MessageIDLTE applies the LTE predicate on the "message_id" field.
func MessageIDLTE(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLTE(FieldMessageID, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldContainsFold(FieldText, v))
}

// IsPinnedEQ applies the EQ predicate on the "is_pinned" field.
func IsPinnedEQ(v bool) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldIsPinned, v))
}

// IsPinnedNEQ applies the NEQ predicate on the "is_pinned" field.
func IsPinnedNEQ(v bool) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNEQ(FieldIsPinned, v))
}

// FromPlatformEQ applies the EQ predicate on the "from_platform" field.
func FromPlatformEQ(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldFromPlatform, v))
}

// FromPlatformNEQ applies the NEQ predicate on the "from_platform" field.
func FromPlatformNEQ(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNEQ(FieldFromPlatform, v))
}

// FromPlatformIn applies the In predicate on the "from_platform" field.
func FromPlatformIn(vs ...int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldIn(FieldFromPlatform, vs...))
}

// FromPlatformNotIn applies the NotIn predicate on the "from_platform" field.
func FromPlatformNotIn(vs ...int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNotIn(FieldFromPlatform, vs...))
}

// FromPlatformGT applies the GT predicate on the "from_platform" field.
func FromPlatformGT(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGT(FieldFromPlatform, v))
}

// FromPlatformGTE applies the GTE predicate on the "from_platform" field.
func FromPlatformGTE(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGTE(FieldFromPlatform, v))
}

// FromPlatformLT applies the LT predicate on the "from_platform" field.
func FromPlatformLT(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLT(FieldFromPlatform, v))
}

// FromPlatformLTE applies the LTE predicate on the "from_platform" field.
func FromPlatformLTE(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLTE(FieldFromPlatform, v))
}

// MessageTypeEQ applies the EQ predicate on the "message_type" field.
func MessageTypeEQ(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldMessageType, v))
}

// MessageTypeNEQ applies the NEQ predicate on the "message_type" field.
func MessageTypeNEQ(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNEQ(FieldMessageType, v))
}

// MessageTypeIn applies the In predicate on the "message_type" field.
func MessageTypeIn(vs ...int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldIn(FieldMessageType, vs...))
}

// MessageTypeNotIn applies the NotIn predicate on the "message_type" field.
func MessageTypeNotIn(vs ...int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNotIn(FieldMessageType, vs...))
}

// MessageTypeGT applies the GT predicate on the "message_type" field.
func MessageTypeGT(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGT(FieldMessageType, v))
}

// MessageTypeGTE applies the GTE predicate on the "message_type" field.
func MessageTypeGTE(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGTE(FieldMessageType, v))
}

// MessageTypeLT applies the LT predicate on the "message_type" field.
func MessageTypeLT(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLT(FieldMessageType, v))
}

// MessageTypeLTE applies the LTE predicate on the "message_type" field.
func MessageTypeLTE(v int) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLTE(FieldMessageType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.SentMessages {
	return predicate.SentMessages(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SentMessages) predicate.SentMessages {
	return predicate.SentMessages(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SentMessages) predicate.SentMessages {
	return predicate.SentMessages(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SentMessages) predicate.SentMessages {
	return predicate.SentMessages(sql.NotPredicates(p))
}
