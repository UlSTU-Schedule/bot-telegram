package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ulstu-schedule/bot-telegram/internal/config"
	"github.com/ulstu-schedule/bot-telegram/internal/store/postgres"
	"github.com/ulstu-schedule/bot-telegram/internal/telegram"
	"log"
)

// Run runs bot.
func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
	}

	studentDB, err := postgres.NewStudentDB(cfg.StudentDatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	studentRepo := postgres.NewStudentRepository(studentDB)

	api, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	bot := telegram.NewBot(api, cfg.Stickers, cfg.Messages, cfg.Commands, studentRepo, cfg.Faculties)
	bot.Start()
}
