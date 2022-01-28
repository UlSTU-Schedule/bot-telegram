package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case b.commands.WithSlash.About:
		return b.handleAboutCommand(message)
	case b.commands.WithSlash.Help:
		return b.handleHelpCommand(message)
	case b.commands.WithSlash.Start:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	ansText := b.answers.ToCommand.Start
	ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
	ansMsg.ReplyMarkup = b.firstLvlMenu()

	_, err := b.bot.Send(ansMsg)
	return err
}

func (b *Bot) handleAboutCommand(message *tgbotapi.Message) error {
	ansMsg := tgbotapi.NewMessage(message.Chat.ID, b.answers.ToCommand.About)
	ansMsg.DisableWebPagePreview = true

	_, err := b.bot.Send(ansMsg)
	return err
}

func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	ansText := b.answers.ToCommand.Help
	ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
	ansMsg.ReplyMarkup = b.firstLvlMenu()

	_, err := b.bot.Send(ansMsg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	ansMsg := tgbotapi.NewMessage(message.Chat.ID, b.answers.ToCommand.Unknown)

	_, err := b.bot.Send(ansMsg)
	return err
}
