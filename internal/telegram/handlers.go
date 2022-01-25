package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ulstu-schedule/parser/types"
	"log"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.CallbackQuery != nil {
			go func(query *tgbotapi.CallbackQuery) {
				err := b.handleCallbackQuery(query)
				if err != nil {
					b.handleCallbackQueryError(query, err)
				}
			}(update.CallbackQuery)
		} else if update.Message != nil {
			go func(message *tgbotapi.Message) {
				err := b.handleMessage(message)
				if err != nil {
					b.handleMessageError(message, err)
				}
			}(update.Message)
		}
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	if message.Sticker != nil {
		return b.handleSticker(message)
	}

	if message.Voice != nil {
		return b.handleVoice(message)
	}

	if message.IsCommand() {
		return b.handleCommand(message)
	}

	return b.handleTextMessage(message)
}

func (b *Bot) handleMessageError(message *tgbotapi.Message, err error) {
	log.Printf("[TG] @%s: %s", message.From.UserName, message.Text)
	log.Printf("[TG] MESSAGE ERROR: %s", err.Error())

	switch err.(type) {
	case *types.UnavailableScheduleError, *types.IncorrectLinkError:
		msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.ScheduleIsUnavailable)
		_, _ = b.bot.Send(msg)
	case *types.StatusCodeError:
		msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.ServerError)
		_, _ = b.bot.Send(msg)
	case *types.IncorrectDateError:
		msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.IncorrectDateError)
		_, _ = b.bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownError)
		_, _ = b.bot.Send(msg)
	}
}

func (b *Bot) handleCallbackQuery(query *tgbotapi.CallbackQuery) error {
	switch query.Data {
	case "first_menu":
		ansText := "Главное меню (1)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.firstLvlMenu())
		_, err := b.bot.Request(editedMsg)
		return err
	case b.commands.First.Groups.Data, b.commands.Third.Groups.Back.Data:
		ansText := "Группы (2)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.groupsKeyboard())
		_, err := b.bot.Request(editedMsg)
		return err
	case b.commands.First.Teachers.Data, b.commands.Third.Teachers.Back.Data:
		ansText := "Учителя (2)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.teachersKeyboard())
		_, err := b.bot.Request(editedMsg)
		return err
	case b.commands.Second.Groups.Schedule.Data:
		ansText := "Расписание групп (3)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.thirdLvlMenuGroups())
		_, err := b.bot.Request(editedMsg)
		return err
	case b.commands.Second.Groups.Change.Data:
		ansText := "Изменить группу (2)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.groupsKeyboard())
		_, err := b.bot.Request(editedMsg)
		return err
	case b.commands.Second.Teachers.Schedule.Data:
		ansText := "Расписание учителей (3)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.thirdLvlMenuTeachers())
		_, err := b.bot.Request(editedMsg)
		return err
	case b.commands.Second.Teachers.Change.Data:
		ansText := "Изменить учителя (2)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.teachersKeyboard())
		_, err := b.bot.Request(editedMsg)
		return err
	case b.commands.Third.Groups.Today.Data:
		ansText := "Расписание групп на сегодня (3)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.thirdLvlMenuGroups())
		_, err := b.bot.Request(editedMsg)
		return err
	case b.commands.Third.Groups.Tomorrow.Data:
		ansText := "Расписание групп на завтра (3)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.thirdLvlMenuGroups())
		_, err := b.bot.Request(editedMsg)
		return err
	case b.commands.Third.Groups.CurrWeek.Data:
		ansText := "Расписание групп на текущую неделю (3)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.thirdLvlMenuGroups())
		_, err := b.bot.Request(editedMsg)
		return err
	case b.commands.Third.Groups.NextWeek.Data:
		ansText := "Расписание групп на следующую неделю (3)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.thirdLvlMenuGroups())
		_, err := b.bot.Request(editedMsg)
		return err
	case b.commands.Third.Teachers.Today.Data:
		ansText := "Расписание учителей на сегодня (3)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.thirdLvlMenuTeachers())
		_, err := b.bot.Request(editedMsg)
		return err
	case b.commands.Third.Teachers.Tomorrow.Data:
		ansText := "Расписание учителей на завтра (3)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.thirdLvlMenuTeachers())
		_, err := b.bot.Request(editedMsg)
		return err
	case b.commands.Third.Teachers.CurrWeek.Data:
		ansText := "Расписание учителей на текущую неделю (3)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.thirdLvlMenuTeachers())
		_, err := b.bot.Request(editedMsg)
		return err
	case b.commands.Third.Teachers.NextWeek.Data:
		ansText := "Расписание учителей на следующую неделю (3)"
		editedMsg := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID, ansText, b.thirdLvlMenuTeachers())
		_, err := b.bot.Request(editedMsg)
		return err
	}
	return nil
}

func (b *Bot) handleCallbackQueryError(query *tgbotapi.CallbackQuery, err error) {
	log.Printf("[TG] @%s: %s", query.From.UserName, query.Data)
	log.Printf("[TG] CALLBACK QUERY ERROR: %s", err)
}
