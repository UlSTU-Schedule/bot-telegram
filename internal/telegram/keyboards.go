package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) mainMenuKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Расписание групп", "groups"),
			tgbotapi.NewInlineKeyboardButtonData("Расписание учителей", "teachers"),
		),
	)
}

func (b *Bot) teachersKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Моё расписание", "teachers_schedule"),
			tgbotapi.NewInlineKeyboardButtonData("Изменить учителя", "teachers_change"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "first_menu"),
		),
	)
}

func (b *Bot) groupsKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Моё расписание", "groups_schedule"),
			tgbotapi.NewInlineKeyboardButtonData("Изменить группу", "groups_change"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "first_menu"),
		),
	)
}

func (b *Bot) teachersScheduleKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Сегодня", "teachers_schedule_today"),
			tgbotapi.NewInlineKeyboardButtonData("Завтра", "teachers_schedule_tomorrow"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Текущая неделя", "teachers_schedule_curr_week"),
			tgbotapi.NewInlineKeyboardButtonData("Следующая неделя", "teachers_schedule_next_week"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "teachers"),
		),
	)
}

func (b *Bot) groupsScheduleKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Сегодня", "groups_schedule_today"),
			tgbotapi.NewInlineKeyboardButtonData("Завтра", "groups_schedule_tomorrow"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Текущая неделя", "groups_schedule_curr_week"),
			tgbotapi.NewInlineKeyboardButtonData("Следующая неделя", "groups_schedule_next_week"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "groups"),
		),
	)
}

func (b *Bot) emptyKeyboard() tgbotapi.ReplyKeyboardRemove {
	return tgbotapi.NewRemoveKeyboard(true)
}
