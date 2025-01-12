package tgchats

import (
	"context"
	"time"

	"github.com/nekomeowww/insights-bot/ent"
	"github.com/nekomeowww/insights-bot/ent/telegramchatrecapsoptions"
	"github.com/nekomeowww/insights-bot/pkg/types/tgchat"
	"go.uber.org/zap"
)

func (m *Model) FindOneRecapsOption(chatID int64) (*ent.TelegramChatRecapsOptions, error) {
	option, err := m.ent.TelegramChatRecapsOptions.
		Query().
		Where(telegramchatrecapsoptions.ChatID(chatID)).
		First(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}

		return nil, err
	}

	return option, nil
}

func (m *Model) SetRecapsRecapMode(chatID int64, recapMode tgchat.AutoRecapSendMode) error {
	option, err := m.FindOneRecapsOption(chatID)
	if err != nil {
		return err
	}
	if option == nil {
		err := m.ent.TelegramChatRecapsOptions.
			Create().
			SetChatID(chatID).
			SetAutoRecapSendMode(int(recapMode)).
			Exec(context.Background())
		if err != nil {
			return err
		}

		return nil
	}
	if option.AutoRecapSendMode == int(recapMode) {
		return nil
	}

	return m.ent.TelegramChatRecapsOptions.
		UpdateOne(option).
		SetAutoRecapSendMode(int(recapMode)).
		Exec(context.Background())
}

func (m *Model) ManualRecapRatePerSeconds(option *ent.TelegramChatRecapsOptions) time.Duration {
	seconds := m.config.HardLimit.ManualRecapRatePerSeconds
	if option != nil && option.ManualRecapRatePerSeconds > seconds { // if chat has a seconds for one rate config that is greater than default, use it, otherwise use default
		seconds = option.ManualRecapRatePerSeconds
	}

	return time.Duration(seconds) * time.Second
}

func (m *Model) DeleteOneOptionByChatID(chatID int64) error {
	_, err := m.ent.TelegramChatRecapsOptions.
		Delete().
		Where(telegramchatrecapsoptions.ChatID(chatID)).
		Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) MigrateOptionOfChatFromChatIDToChatID(fromChatID int64, toChatID int64) error {
	affectedRows, err := m.ent.TelegramChatRecapsOptions.
		Update().
		Where(
			telegramchatrecapsoptions.ChatIDEQ(fromChatID),
		).
		SetChatID(toChatID).
		Save(context.Background())

	if err != nil {
		return err
	}

	m.logger.Info("successfully migrated options of chat",
		zap.Int64("from_chat_id", fromChatID),
		zap.Int64("to_chat_id", toChatID),
		zap.Int("affected_rows", affectedRows),
	)

	return nil
}
