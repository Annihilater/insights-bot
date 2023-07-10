package welcome

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nekomeowww/fo"
	"go.uber.org/fx"

	"github.com/nekomeowww/insights-bot/internal/models/chathistories"
	"github.com/nekomeowww/insights-bot/internal/models/tgchats"
	"github.com/nekomeowww/insights-bot/pkg/bots/tgbot"
	"github.com/nekomeowww/insights-bot/pkg/logger"
	"github.com/nekomeowww/insights-bot/pkg/types/telegram"
)

func NewModules() fx.Option {
	return fx.Options(
		fx.Provide(NewHandlers()),
	)
}

type NewHandlersParams struct {
	fx.In

	TgChats       *tgchats.Model
	ChatHistories *chathistories.Model
	Logger        *logger.Logger
}

type Handlers struct {
	tgchats       *tgchats.Model
	chatHistories *chathistories.Model
	logger        *logger.Logger
}

func NewHandlers() func(param NewHandlersParams) *Handlers {
	return func(param NewHandlersParams) *Handlers {
		return &Handlers{
			tgchats:       param.TgChats,
			chatHistories: param.ChatHistories,
			logger:        param.Logger,
		}
	}
}

func (h *Handlers) Install(dispatcher *tgbot.Dispatcher) {
	dispatcher.OnMyChatMember(tgbot.NewHandler(h.handleChatMember))
}

func (h *Handlers) handleChatMember(c *tgbot.Context) (tgbot.Response, error) {
	if c.Update.MyChatMember == nil {
		return nil, nil
	}
	if telegram.MemberStatus(c.Update.MyChatMember.NewChatMember.Status) == telegram.MemberStatusLeft {
		h.handleBotLeftChat(c.Update.MyChatMember.Chat.ID)
		return nil, nil
	}
	if telegram.MemberStatus(c.Update.MyChatMember.NewChatMember.Status) == telegram.MemberStatusMember {
		h.handleBotJoinChat(c)
		return nil, nil
	}

	return nil, nil
}

func (h *Handlers) handleBotLeftChat(chatID int64) {
	may := fo.
		NewMay0().
		Use(fo.WithLogFuncHandler(func(a ...any) {
			h.logger.Error(fmt.Sprint(a...))
		}))

	may.Invoke(func() error {
		err := h.tgchats.DeleteAllSubscribersByChatID(chatID)
		if err != nil {
			return fmt.Errorf("failed to delete all subscribers by chat id %d: %w", chatID, err)
		}

		return nil
	}())
	may.Invoke(func() error {
		err := h.tgchats.DeleteOneFeatureFlagByChatID(chatID)
		if err != nil {
			return fmt.Errorf("failed to delete one feature flag by chat id %d: %w", chatID, err)
		}

		return nil
	}())
	may.Invoke(func() error {
		err := h.tgchats.DeleteOneOptionByChatID(chatID)
		if err != nil {
			return fmt.Errorf("failed to delete one option by chat id %d: %w", chatID, err)
		}

		return nil
	}())
	may.Invoke(func() error {
		err := h.chatHistories.DeleteAllChatHistoriesByChatID(chatID)
		if err != nil {
			return fmt.Errorf("failed to delete all chat histories by chat id %d: %w", chatID, err)
		}

		return nil
	}())
}

func (h *Handlers) handleBotJoinChat(c *tgbot.Context) {
	msg := tgbotapi.NewMessage(c.Update.MyChatMember.Chat.ID, fmt.Sprintf(""+
		"欢迎使用 @%s！\n\n"+
		"- 如果要让我帮忙阅读网页文章，请直接使用开箱即用的命令 /smr@%s <code>要阅读"+
		"的链接</code>；\n"+
		"- 如果想要我帮忙总结本群组的聊天记录，请以<b>管理员</b>身份将"+
		"我配置为本群组的管理员（可以关闭所有权限），然后在<b>非匿名和非其他身份的身"+
		"份</b>下（推荐，否则容易出现权限识别错误的情况）发送 /configure_recap@%s "+
		"来开始配置本群组的聊天回顾功能。\n\n"+
		"如果还有疑问的话也可以执行帮助命令 /help@%s 来查看支持的命令，或者前往 Bot "+
		"所在的开源仓库提交 Issue。\n\n"+
		"祝你使用愉快！"+
		"",
		c.Bot.Self.UserName,
		c.Bot.Self.UserName,
		c.Bot.Self.UserName,
		c.Bot.Self.UserName,
	))
	msg.ParseMode = tgbotapi.ModeHTML

	c.Bot.MaySend(msg)
}
