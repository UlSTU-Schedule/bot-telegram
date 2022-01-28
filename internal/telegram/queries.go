package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b *Bot) handleCallbackQuery(query *tgbotapi.CallbackQuery) error {
	switch query.Data {
	case "first_menu":
		return b.handleFirstLvlMenuQuery(query.Message)
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
	log.Printf("@%s: %s", query.From.UserName, query.Data)
	log.Printf("CALLBACK QUERY ERROR: %s", err)
}

func (b *Bot) handleFirstLvlMenuQuery(message *tgbotapi.Message) error {
	editMsgText := b.answers.ToQuery.FirstLvlMenu
	editMsgCfg := tgbotapi.NewEditMessageTextAndMarkup(
		message.Chat.ID,
		message.MessageID,
		editMsgText,
		b.firstLvlMenu(),
	)

	_, err := b.bot.Request(editMsgCfg)
	return err
}
