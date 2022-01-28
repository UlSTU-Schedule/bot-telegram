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
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Second.Groups.Schedule.Command,
				b.commands.Second.Groups.Schedule.Data),
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Second.Groups.Change.Command,
				b.commands.Second.Groups.Change.Data),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Second.Back.Command,
				b.commands.Second.Back.Data),
		),
	)
}

func (b *Bot) teachersKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Second.Teachers.Schedule.Command,
				b.commands.Second.Teachers.Schedule.Data),
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Second.Teachers.Change.Command,
				b.commands.Second.Teachers.Change.Data),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				b.commands.Second.Back.Command,
				b.commands.Second.Back.Data),
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

func (b *Bot) hideInlineKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup()
}

func (b *Bot) hideReplyKeyboard() tgbotapi.ReplyKeyboardRemove {
	return tgbotapi.NewRemoveKeyboard(true)
}
