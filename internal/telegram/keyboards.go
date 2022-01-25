package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) firstLvlMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Inline.First.Groups.Command,
				b.commands.Inline.First.Groups.Data,
			),
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Inline.First.Teachers.Command,
				b.commands.Inline.First.Teachers.Data,
			),
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

func (b *Bot) thirdLvlMenuGroups() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Third.Groups.Today.Command,
				b.commands.Third.Groups.Today.Data,
			),
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Third.Groups.Tomorrow.Command,
				b.commands.Third.Groups.Tomorrow.Data,
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Third.Groups.CurrWeek.Command,
				b.commands.Third.Groups.CurrWeek.Data,
			),
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Third.Groups.NextWeek.Command,
				b.commands.Third.Groups.NextWeek.Data,
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Third.Groups.Back.Command,
				b.commands.Third.Groups.Back.Data,
			),
		),
	)
}

func (b *Bot) thirdLvlMenuTeachers() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Third.Teachers.Today.Command,
				b.commands.Third.Teachers.Today.Data,
			),
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Third.Teachers.Tomorrow.Command,
				b.commands.Third.Teachers.Tomorrow.Data,
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Third.Teachers.CurrWeek.Command,
				b.commands.Third.Teachers.CurrWeek.Data,
			),
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Third.Teachers.NextWeek.Command,
				b.commands.Third.Teachers.NextWeek.Data,
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Third.Teachers.Back.Command,
				b.commands.Third.Teachers.Back.Data,
			),
		),
	)
}

func (b *Bot) hideReplyKeyboard() tgbotapi.ReplyKeyboardRemove {
	return tgbotapi.NewRemoveKeyboard(true)
}
