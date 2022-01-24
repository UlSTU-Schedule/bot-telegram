package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ulstu-schedule/bot-telegram/internal/config"
	"github.com/ulstu-schedule/bot-telegram/internal/store/postgres"
	"log"
)

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

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
