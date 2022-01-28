package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	"github.com/ulstu-schedule/bot-telegram/internal/config"
	"github.com/ulstu-schedule/bot-telegram/internal/store/postgres"
	"github.com/ulstu-schedule/bot-telegram/internal/telegram"
	"log"
)

// Run runs the bot.
func Run(configPath string) {
	cfg, err := config.New(configPath)
	if err != nil {
		log.Fatal(err)
	}

	studentDB, err := postgres.NewDB(cfg.StudentDatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(studentDB)
	studentStore := postgres.NewStudentStore(studentDB)

	scheduleDB, err := postgres.NewDB(cfg.ScheduleDatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(scheduleDB)
	scheduleStore := postgres.NewScheduleStore(scheduleDB)

	api, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	bot := telegram.NewBot(api, cfg.Stickers, cfg.Messages, cfg.Commands, studentStore, scheduleStore, cfg.Faculties)
	bot.Start()
}
