package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ulstu-schedule/bot-telegram/internal/config"
	"github.com/ulstu-schedule/bot-telegram/internal/store/postgres"
	"log"
)

type Bot struct {
	bot           *tgbotapi.BotAPI
	answers       config.Answers
	commands      config.Commands
	studentStore  *postgres.StudentStore
	scheduleStore *postgres.ScheduleStore
	faculties     []config.Faculty
}

func NewBot(api *tgbotapi.BotAPI, answers config.Answers, commands config.Commands, studentStore *postgres.StudentStore, scheduleStore *postgres.ScheduleStore, faculties []config.Faculty) *Bot {
	return &Bot{bot: api, answers: answers, commands: commands, studentStore: studentStore, scheduleStore: scheduleStore, faculties: faculties}
}

func (b *Bot) Start() {
	log.Println("The Telegram bot was launched!")

	updates := b.initUpdatesChannel()
	b.handleUpdates(updates)
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			go func(message *tgbotapi.Message) {
				err := b.handleMessage(message)
				if err != nil {
					b.handleMessageError(message, err)
				}
			}(update.Message)
		} else if update.CallbackQuery != nil {
			go func(query *tgbotapi.CallbackQuery) {
				err := b.handleCallbackQuery(query)
				if err != nil {
					b.handleCallbackQueryError(query, err)
				}
			}(update.CallbackQuery)
		}
	}
}
