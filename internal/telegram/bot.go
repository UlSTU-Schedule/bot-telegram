package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ulstu-schedule/bot-telegram/internal/config"
	"github.com/ulstu-schedule/bot-telegram/internal/store/postgres"
	"github.com/ulstu-schedule/parser/types"
	"log"
)

// Bot ...
type Bot struct {
	bot           *tgbotapi.BotAPI
	stickers      config.Stickers
	messages      config.Messages
	commands      config.Commands
	studentStore  *postgres.StudentStore
	scheduleStore *postgres.ScheduleStore
	faculties     []config.Faculty
}

func NewBot(api *tgbotapi.BotAPI, stickers config.Stickers, messages config.Messages, commands config.Commands, studentStore *postgres.StudentStore, scheduleStore *postgres.ScheduleStore, faculties []config.Faculty) *Bot {
	return &Bot{bot: api, stickers: stickers, messages: messages, commands: commands, studentStore: studentStore, scheduleStore: scheduleStore, faculties: faculties}
}

func (b *Bot) Start() {
	log.Println("The Telegram bot was launched!")

	updates := b.initUpdatesChannel()
	b.handleUpdates(updates)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // ignores non-messages updates
			continue
		}

		go func(update tgbotapi.Update) {
			err := b.handleMessageUpdate(&update)
			if err != nil {
				b.handleError(&update, err)
			}
		}(update)
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleMessageUpdate(update *tgbotapi.Update) error {
	if update.Message.Sticker != nil {
		return b.handleSticker(update.Message)
	}

	if update.Message.Voice != nil {
		return b.handleVoice(update.Message)
	}

	if update.Message.IsCommand() {
		return b.handleCommand(update.Message)
	}

	return b.handleMsg(update.Message)
}

func (b *Bot) handleError(update *tgbotapi.Update, err error) {
	log.Printf("[TG] @%s: %s", update.Message.From.UserName, update.Message.Text)
	log.Printf("[TG] ERROR: %s", err)

	switch err.(type) {
	case *types.UnavailableScheduleError, *types.IncorrectLinkError:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.ScheduleIsUnavailable)
		_, _ = b.bot.Send(msg)
	case *types.StatusCodeError:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.ServerError)
		_, _ = b.bot.Send(msg)
	case *types.IncorrectDateError:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.IncorrectDateError)
		_, _ = b.bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.UnknownError)
		_, _ = b.bot.Send(msg)
	}
}
