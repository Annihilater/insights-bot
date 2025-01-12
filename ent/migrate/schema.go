// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChatHistoriesColumns holds the columns for the "chat_histories" table.
	ChatHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "chat_id", Type: field.TypeInt64, Default: 0},
		{Name: "chat_title", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "chat_type", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "message_id", Type: field.TypeInt64, Default: 0},
		{Name: "user_id", Type: field.TypeInt64, Default: 0},
		{Name: "username", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "full_name", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "text", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "replied_to_message_id", Type: field.TypeInt64, Default: 0},
		{Name: "replied_to_user_id", Type: field.TypeInt64, Default: 0},
		{Name: "replied_to_full_name", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "replied_to_username", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "replied_to_text", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "replied_to_chat_type", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "chatted_at", Type: field.TypeInt64},
		{Name: "embedded", Type: field.TypeBool, Default: false},
		{Name: "from_platform", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeInt64},
	}
	// ChatHistoriesTable holds the schema information for the "chat_histories" table.
	ChatHistoriesTable = &schema.Table{
		Name:       "chat_histories",
		Columns:    ChatHistoriesColumns,
		PrimaryKey: []*schema.Column{ChatHistoriesColumns[0]},
	}
	// FeedbackChatHistoriesRecapsReactionsColumns holds the columns for the "feedback_chat_histories_recaps_reactions" table.
	FeedbackChatHistoriesRecapsReactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "chat_id", Type: field.TypeInt64, Default: 0},
		{Name: "log_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeInt64, Default: 0},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"none", "up_vote", "down_vote", "lmao"}, Default: "none"},
		{Name: "created_at", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeInt64},
	}
	// FeedbackChatHistoriesRecapsReactionsTable holds the schema information for the "feedback_chat_histories_recaps_reactions" table.
	FeedbackChatHistoriesRecapsReactionsTable = &schema.Table{
		Name:       "feedback_chat_histories_recaps_reactions",
		Columns:    FeedbackChatHistoriesRecapsReactionsColumns,
		PrimaryKey: []*schema.Column{FeedbackChatHistoriesRecapsReactionsColumns[0]},
	}
	// FeedbackSummarizationsReactionsColumns holds the columns for the "feedback_summarizations_reactions" table.
	FeedbackSummarizationsReactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "chat_id", Type: field.TypeInt64, Default: 0},
		{Name: "log_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeInt64, Default: 0},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"none", "up_vote", "down_vote", "lmao"}, Default: "none"},
		{Name: "created_at", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeInt64},
	}
	// FeedbackSummarizationsReactionsTable holds the schema information for the "feedback_summarizations_reactions" table.
	FeedbackSummarizationsReactionsTable = &schema.Table{
		Name:       "feedback_summarizations_reactions",
		Columns:    FeedbackSummarizationsReactionsColumns,
		PrimaryKey: []*schema.Column{FeedbackSummarizationsReactionsColumns[0]},
	}
	// LogChatHistoriesRecapsColumns holds the columns for the "log_chat_histories_recaps" table.
	LogChatHistoriesRecapsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "chat_id", Type: field.TypeInt64, Default: 0},
		{Name: "recap_inputs", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "recap_outputs", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "from_platform", Type: field.TypeInt, Default: 0},
		{Name: "prompt_token_usage", Type: field.TypeInt, Default: 0},
		{Name: "completion_token_usage", Type: field.TypeInt, Default: 0},
		{Name: "total_token_usage", Type: field.TypeInt, Default: 0},
		{Name: "recap_type", Type: field.TypeInt, Default: 0},
		{Name: "model_name", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeInt64},
	}
	// LogChatHistoriesRecapsTable holds the schema information for the "log_chat_histories_recaps" table.
	LogChatHistoriesRecapsTable = &schema.Table{
		Name:       "log_chat_histories_recaps",
		Columns:    LogChatHistoriesRecapsColumns,
		PrimaryKey: []*schema.Column{LogChatHistoriesRecapsColumns[0]},
	}
	// LogSummarizationsColumns holds the columns for the "log_summarizations" table.
	LogSummarizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "content_url", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "content_title", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "content_author", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "content_text", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "content_summarized_outputs", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "from_platform", Type: field.TypeInt, Default: 0},
		{Name: "prompt_token_usage", Type: field.TypeInt, Default: 0},
		{Name: "completion_token_usage", Type: field.TypeInt, Default: 0},
		{Name: "total_token_usage", Type: field.TypeInt, Default: 0},
		{Name: "model_name", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeInt64},
	}
	// LogSummarizationsTable holds the schema information for the "log_summarizations" table.
	LogSummarizationsTable = &schema.Table{
		Name:       "log_summarizations",
		Columns:    LogSummarizationsColumns,
		PrimaryKey: []*schema.Column{LogSummarizationsColumns[0]},
	}
	// MetricOpenAiChatCompletionTokenUsagesColumns holds the columns for the "metric_open_ai_chat_completion_token_usages" table.
	MetricOpenAiChatCompletionTokenUsagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "prompt_operation", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "prompt_character_length", Type: field.TypeInt, Default: 0},
		{Name: "prompt_token_usage", Type: field.TypeInt, Default: 0},
		{Name: "completion_character_length", Type: field.TypeInt, Default: 0},
		{Name: "completion_token_usage", Type: field.TypeInt, Default: 0},
		{Name: "total_token_usage", Type: field.TypeInt, Default: 0},
		{Name: "model_name", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeInt64},
	}
	// MetricOpenAiChatCompletionTokenUsagesTable holds the schema information for the "metric_open_ai_chat_completion_token_usages" table.
	MetricOpenAiChatCompletionTokenUsagesTable = &schema.Table{
		Name:       "metric_open_ai_chat_completion_token_usages",
		Columns:    MetricOpenAiChatCompletionTokenUsagesColumns,
		PrimaryKey: []*schema.Column{MetricOpenAiChatCompletionTokenUsagesColumns[0]},
	}
	// SlackOauthCredentialsColumns holds the columns for the "slack_oauth_credentials" table.
	SlackOauthCredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "team_id", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "refresh_token", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "access_token", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeInt64},
	}
	// SlackOauthCredentialsTable holds the schema information for the "slack_oauth_credentials" table.
	SlackOauthCredentialsTable = &schema.Table{
		Name:       "slack_oauth_credentials",
		Columns:    SlackOauthCredentialsColumns,
		PrimaryKey: []*schema.Column{SlackOauthCredentialsColumns[0]},
	}
	// TelegramChatAutoRecapsSubscribersColumns holds the columns for the "telegram_chat_auto_recaps_subscribers" table.
	TelegramChatAutoRecapsSubscribersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "chat_id", Type: field.TypeInt64, Default: 0},
		{Name: "user_id", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeInt64},
	}
	// TelegramChatAutoRecapsSubscribersTable holds the schema information for the "telegram_chat_auto_recaps_subscribers" table.
	TelegramChatAutoRecapsSubscribersTable = &schema.Table{
		Name:       "telegram_chat_auto_recaps_subscribers",
		Columns:    TelegramChatAutoRecapsSubscribersColumns,
		PrimaryKey: []*schema.Column{TelegramChatAutoRecapsSubscribersColumns[0]},
	}
	// TelegramChatFeatureFlagsColumns holds the columns for the "telegram_chat_feature_flags" table.
	TelegramChatFeatureFlagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "chat_id", Type: field.TypeInt64, Unique: true},
		{Name: "chat_type", Type: field.TypeString, Size: 2147483647},
		{Name: "chat_title", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "feature_chat_histories_recap", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeInt64},
	}
	// TelegramChatFeatureFlagsTable holds the schema information for the "telegram_chat_feature_flags" table.
	TelegramChatFeatureFlagsTable = &schema.Table{
		Name:       "telegram_chat_feature_flags",
		Columns:    TelegramChatFeatureFlagsColumns,
		PrimaryKey: []*schema.Column{TelegramChatFeatureFlagsColumns[0]},
	}
	// TelegramChatRecapsOptionsColumns holds the columns for the "telegram_chat_recaps_options" table.
	TelegramChatRecapsOptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "chat_id", Type: field.TypeInt64, Unique: true},
		{Name: "auto_recap_send_mode", Type: field.TypeInt, Default: 0},
		{Name: "manual_recap_rate_per_seconds", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeInt64},
		{Name: "updated_at", Type: field.TypeInt64},
	}
	// TelegramChatRecapsOptionsTable holds the schema information for the "telegram_chat_recaps_options" table.
	TelegramChatRecapsOptionsTable = &schema.Table{
		Name:       "telegram_chat_recaps_options",
		Columns:    TelegramChatRecapsOptionsColumns,
		PrimaryKey: []*schema.Column{TelegramChatRecapsOptionsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChatHistoriesTable,
		FeedbackChatHistoriesRecapsReactionsTable,
		FeedbackSummarizationsReactionsTable,
		LogChatHistoriesRecapsTable,
		LogSummarizationsTable,
		MetricOpenAiChatCompletionTokenUsagesTable,
		SlackOauthCredentialsTable,
		TelegramChatAutoRecapsSubscribersTable,
		TelegramChatFeatureFlagsTable,
		TelegramChatRecapsOptionsTable,
	}
)

func init() {
}
