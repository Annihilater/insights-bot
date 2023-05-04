package recap

import (
	"fmt"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nekomeowww/insights-bot/internal/models/tgchats"
	"github.com/nekomeowww/insights-bot/pkg/bots/tgbot"
	"github.com/nekomeowww/insights-bot/pkg/types/telegram"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

func NewModules() fx.Option {
	return fx.Options(
		fx.Provide(NewHandlers()),
		fx.Provide(NewRecapCommandHandler()),
		fx.Provide(NewRecapCallbackQueryHandler()),
		fx.Provide(NewEnableRecapCommandHandler()),
		fx.Provide(NewDisableRecapCommandHandler()),
	)
}

var (
	_ tgbot.CommandHandler = (*RecapCommandHandler)(nil)
)

type NewHandlersParams struct {
	fx.In

	RecapCommand        *RecapCommandHandler
	RecapCallbackQuery  *RecapCallbackQueryHandler
	EnableRecapCommand  *EnableRecapCommandHandler
	DisableRecapCommand *DisableRecapCommandHandler
}

type Handlers struct {
	recapCommand        *RecapCommandHandler
	recapCallbackQuery  *RecapCallbackQueryHandler
	enableRecapCommand  *EnableRecapCommandHandler
	disableRecapCommand *DisableRecapCommandHandler
}

func NewHandlers() func(NewHandlersParams) *Handlers {
	return func(param NewHandlersParams) *Handlers {
		return &Handlers{
			recapCommand:        param.RecapCommand,
			recapCallbackQuery:  param.RecapCallbackQuery,
			enableRecapCommand:  param.EnableRecapCommand,
			disableRecapCommand: param.DisableRecapCommand,
		}
	}
}

func (h *Handlers) Install(dispatcher *tgbot.Dispatcher) {
	dispatcher.OnCommand(h.recapCommand)
	dispatcher.OnCallbackQuery(h.recapCallbackQuery)
	dispatcher.OnCommand(h.enableRecapCommand)
	dispatcher.OnCommand(h.disableRecapCommand)
}

var (
	RecapSelectHourAvailables = []int64{
		1, 2, 4, 6, 12,
	}
	RecapSelectHourAvailableText = lo.SliceToMap(RecapSelectHourAvailables, func(item int64) (int64, string) {
		return item, fmt.Sprintf("%d 小时", item)
	})
	RecapSelectHourAvailableValues = lo.SliceToMap(RecapSelectHourAvailables, func(item int64) (int64, string) {
		return item, fmt.Sprintf("%d", item)
	})
)

type NewRecapCommandHandlerParams struct {
	fx.In

	TgChats *tgchats.Model
}

type RecapCommandHandler struct {
	tgchats *tgchats.Model
}

func NewRecapCommandHandler() func(NewRecapCommandHandlerParams) *RecapCommandHandler {
	return func(param NewRecapCommandHandlerParams) *RecapCommandHandler {
		return &RecapCommandHandler{
			tgchats: param.TgChats,
		}
	}
}

func (h RecapCommandHandler) Command() string {
	return "recap"
}

func (h RecapCommandHandler) CommandHelp() string {
	return "总结过去的聊天记录并生成回顾快报"
}

func (h *RecapCommandHandler) Handle(c *tgbot.Context) (tgbot.Response, error) {
	enabled, err := h.tgchats.HasChatHistoriesRecapEnabled(c.Update.Message.Chat.ID, telegram.ChatType(c.Update.Message.Chat.Type))
	if err != nil {
		return nil, tgbot.NewExceptionError(err).WithMessage("生成失败，请稍后再试。").WithReply(c.Update.Message)
	}
	if !enabled {
		return nil, tgbot.NewMessageError("聊天记录回顾功能在当前群组尚未启用，请使用 /enable_recap 命令启用后再试。").WithReply(c.Update.Message)
	}

	replyMarkupKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			lo.Map(RecapSelectHourAvailables, func(item int64, _ int) tgbotapi.InlineKeyboardButton {
				return tgbotapi.NewInlineKeyboardButtonData(
					RecapSelectHourAvailableText[item],
					tgbot.NewCallbackQueryData("recap", "select_hour", url.Values{"hour": []string{RecapSelectHourAvailableValues[item]}}),
				)
			})...,
		),
	)

	return c.NewMessageReplyTo("要创建过去几个小时内的聊天回顾呢？", c.Update.Message.MessageID).WithReplyMarkup(replyMarkupKeyboard), nil
}