package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ulstu-schedule/bot-telegram/internal/config"
	"github.com/ulstu-schedule/bot-telegram/internal/store/postgres"
	"log"
)

// Bot ...
type Bot struct {
	bot       *tgbotapi.BotAPI
	stickers  config.Stickers
	messages  config.Messages
	commands  config.Commands
	repos     *postgres.StudentRepository
	faculties []config.Faculty
}

// NewBot ...
func NewBot(api *tgbotapi.BotAPI, stickers config.Stickers, messages config.Messages, commands config.Commands, repos *postgres.StudentRepository, faculties []config.Faculty) *Bot {
	return &Bot{bot: api, stickers: stickers, messages: messages, commands: commands, repos: repos, faculties: faculties}
}

// Start ...
func (b *Bot) Start() {
	log.Println("The Telegram bot was launched!")

	updates := b.initUpdatesChannel()
	b.handleUpdates(updates)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // игнорирует обновления, не связанные с сообщениями
			continue
		}

		if update.Message.IsCommand() {
			go func(update tgbotapi.Update) {
				log.Printf("[TG] @%s: %s", update.Message.From.UserName, update.Message.Text)

				err := b.handleCommand(update.Message)
				if err != nil {
					log.Printf("[TG] ERROR: %s", err)
				}
			}(update)
			continue
		}

		if update.Message.Sticker != nil {
			go func(update tgbotapi.Update) {
				log.Printf("[TG] @%s: %s", update.Message.From.UserName, "sticker")

				err := b.handleSticker(update.Message)
				if err != nil {
					log.Printf("[TG] ERROR: %s", err)
				}
			}(update)
			continue
		}

		if update.Message.Voice != nil {
			go func(update tgbotapi.Update) {
				log.Printf("[TG] @%s: %s", update.Message.From.UserName, "voice")

				err := b.handleVoice(update.Message)
				if err != nil {
					log.Printf("[TG] ERROR: %s", err)
				}
			}(update)
			continue
		}

		go func(update tgbotapi.Update) {
			log.Printf("[TG] @%s: %s", update.Message.From.UserName, update.Message.Text)

			err := b.handleMessage(update.Message)
			if err != nil {
				log.Printf("[TG] ERROR: %s", err)
			}
		}(update)
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
