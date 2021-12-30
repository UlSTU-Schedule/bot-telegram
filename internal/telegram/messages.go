package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ulstu-schedule/bot-telegram/internal/schedule"
	"os"
	"regexp"
	"strings"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case b.commands.WithSlash.AboutProject:
		return b.handleAboutProjectCommand(message)
	case b.commands.WithSlash.Help:
		return b.handleHelpCommand(message)
	case b.commands.WithSlash.Start:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	student, err := b.studentStore.Student().GetStudent(int(message.From.ID))
	if err != nil {
		return err
	}

	if student != nil {
		ansText := b.messages.StartWithGroup
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
		ansMsg.ReplyMarkup = getMainMenuKeyboard()

		_, err = b.bot.Send(ansMsg)
	} else {
		ansText := b.messages.StartWithoutGroup
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
		ansMsg.ReplyMarkup = getEmptyKeyboard()

		_, err = b.bot.Send(ansMsg)
	}
	return err
}

func (b *Bot) handleAboutProjectCommand(message *tgbotapi.Message) error {
	ansMsg := tgbotapi.NewMessage(message.Chat.ID, b.messages.AboutProject)
	ansMsg.DisableWebPagePreview = true

	_, err := b.bot.Send(ansMsg)
	return err
}

func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	student, err := b.studentStore.Student().GetStudent(int(message.From.ID))
	if err != nil {
		return err
	}

	if student != nil {
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, b.messages.InfoWithGroup)
		ansMsg.ReplyMarkup = getScheduleMenuKeyboard()

		_, err = b.bot.Send(ansMsg)
	} else {
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, b.messages.ChangeGroup)
		ansMsg.ReplyMarkup = getEmptyKeyboard()

		_, err = b.bot.Send(ansMsg)
	}
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	ansMsg := tgbotapi.NewMessage(message.Chat.ID, b.messages.IncorrectInput)

	_, err := b.bot.Send(ansMsg)
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	userMsg := message.Text

	switch {
	case isUserMsgAmongCommands(userMsg, b.commands.Whole.GetScheduleForDay):
		return b.handleGetScheduleForDayMsg(message)
	case isUserMsgAmongCommands(userMsg, b.commands.Whole.GetScheduleForWeek) ||
		findCommandsAmongUserMsg(userMsg, b.commands.Partial.GetScheduleForWeek):
		return b.handleGetScheduleForWeekMsg(message)
	case isUserMsgAmongCommands(userMsg, b.commands.Whole.ChangeGroup) ||
		findCommandsAmongUserMsg(userMsg, b.commands.Partial.ChangeGroup):
		return b.handleChangeGroupMsg(message)
	case isUserMsgAmongCommands(userMsg, b.commands.Whole.BackToStartMenu) ||
		findCommandsAmongUserMsg(userMsg, b.commands.Partial.BackToStartMenu):
		return b.handleBackToStartMenuMsg(message)
	case isUserMsgAmongCommands(userMsg, b.commands.Whole.GoToScheduleMenu) ||
		findCommandsAmongUserMsg(userMsg, b.commands.Partial.GoToScheduleMenu):
		return b.handleGoToScheduleMenuMsg(message)
	case findCommandsAmongUserMsg(userMsg, b.commands.Partial.ExpressGratitude):
		return b.handleExpressGratitudeMsg(message)
	default:
		return b.handleUnknownMsg(message)
	}
}

func (b *Bot) handleChangeGroupMsg(message *tgbotapi.Message) error {
	ansText := b.messages.ChangeGroup
	ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
	ansMsg.ReplyMarkup = getBackToMenuKeyboard()

	_, err := b.bot.Send(ansMsg)
	return err
}

func (b *Bot) handleGetScheduleForDayMsg(message *tgbotapi.Message) error {
	student, err := b.studentStore.Student().GetStudent(int(message.From.ID))
	if err != nil {
		return err
	}

	if student != nil {
		dailySchedule, err := schedule.GetDailySchedule(student.GroupName, strings.ToLower(message.Text))
		if err != nil {
			return err
		}

		ansText := dailySchedule
		if schedule.IsGroupFromKEI(student.GroupName) {
			ansText += b.messages.ChangesInKEISchedule
		}
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
		ansMsg.ReplyMarkup = getScheduleMenuKeyboard()

		_, err = b.bot.Send(ansMsg)
	} else {
		ansText := b.messages.GroupNotSelected
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
		ansMsg.ReplyMarkup = getEmptyKeyboard()

		_, err = b.bot.Send(ansMsg)
	}
	return err
}

func (b *Bot) handleGetScheduleForWeekMsg(message *tgbotapi.Message) error {
	student, err := b.studentStore.Student().GetStudent(int(message.From.ID))
	if err != nil {
		return err
	}

	if student != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Генерирую расписание \U0000231B")
		_, _ = b.bot.Send(msg)

		userMsgLower := strings.ToLower(message.Text)

		caption, weeklySchedulePath, err := schedule.GetWeeklySchedule(student.GroupName, userMsgLower)
		if weeklySchedulePath != "" {
			defer os.Remove(weeklySchedulePath)
		}
		if err != nil {
			return err
		}

		imgMsg := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FilePath(weeklySchedulePath))
		imgMsg.Caption = caption
		if userMsgLower == "5" || userMsgLower == "текущая неделя" {
			if schedule.IsGroupFromKEI(student.GroupName) {
				imgMsg.Caption += b.messages.ChangesInKEISchedule
			}
		}
		imgMsg.ReplyMarkup = getScheduleMenuKeyboard()

		_, err = b.bot.Send(imgMsg)
		return err
	} else {
		ansText := b.messages.GroupNotSelected
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
		ansMsg.ReplyMarkup = getEmptyKeyboard()

		_, err = b.bot.Send(ansMsg)
		return err
	}
}

func (b *Bot) handleBackToStartMenuMsg(message *tgbotapi.Message) error {
	student, err := b.studentStore.Student().GetStudent(int(message.From.ID))
	if err != nil {
		return err
	}

	if student != nil {
		ansText := b.messages.Back
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
		ansMsg.ReplyMarkup = getMainMenuKeyboard()

		_, err = b.bot.Send(ansMsg)
	} else {
		ansText := b.messages.ChangeGroup
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
		ansMsg.ReplyMarkup = getEmptyKeyboard()

		_, err = b.bot.Send(ansMsg)
	}
	return err
}

func (b *Bot) handleGoToScheduleMenuMsg(message *tgbotapi.Message) error {
	student, err := b.studentStore.Student().GetStudent(int(message.From.ID))
	if err != nil {
		return err
	}

	if student != nil {
		ansText := fmt.Sprintf("Твоя группа: %s \U0001F4CC \n\n", student.GroupName)
		ansText += b.messages.InfoWithGroup
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
		ansMsg.ReplyMarkup = getScheduleMenuKeyboard()

		_, err = b.bot.Send(ansMsg)
	} else {
		ansText := b.messages.InfoWithoutGroup

		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
		ansMsg.ReplyMarkup = getEmptyKeyboard()

		_, err = b.bot.Send(ansMsg)
	}
	return err
}

func (b *Bot) handleUnknownMsg(message *tgbotapi.Message) error {
	isUserInputGroup, formattedGroupName, err := schedule.IsGroupExist(message.Text)
	if err != nil {
		return err
	}

	if isUserInputGroup {
		facultyID := b.determineFacultyID(formattedGroupName)

		err = b.studentStore.Student().Information(message.From.FirstName, message.From.LastName, int(message.From.ID), formattedGroupName, facultyID)
		if err != nil {
			return err
		}

		ansText := fmt.Sprintf("Твоя группа обновлена на %s \U00002705\n\n", formattedGroupName)
		ansText += b.messages.InfoWithGroup

		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
		ansMsg.ReplyMarkup = getScheduleMenuKeyboard()

		_, err = b.bot.Send(ansMsg)
	} else {
		ansText := b.messages.IncorrectInput
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)

		_, err = b.bot.Send(ansMsg)
	}
	return err
}

func (b *Bot) handleExpressGratitudeMsg(message *tgbotapi.Message) error {
	msg := tgbotapi.NewSticker(message.Chat.ID, tgbotapi.FileID(b.stickers.ToExpressGratitude))
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleSticker(message *tgbotapi.Message) error {
	msg := tgbotapi.NewSticker(message.Chat.ID, tgbotapi.FileID(b.stickers.ToSticker))
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleVoice(message *tgbotapi.Message) error {
	msg := tgbotapi.NewSticker(message.Chat.ID, tgbotapi.FileID(b.stickers.ToVoice))
	_, err := b.bot.Send(msg)
	return err
}

// determineFacultyID ...
func (b *Bot) determineFacultyID(groupName string) byte {
	for _, faculty := range b.faculties {
		for _, group := range faculty.Groups {
			expr := fmt.Sprintf(`(?i)^%s[\d]+$`, group)
			groupRegexp := regexp.MustCompile(expr)
			if groupRegexp.MatchString(groupName) {
				return faculty.ID
			}
		}
	}
	if schedule.KEIGroupPattern.MatchString(groupName) {
		return 2
	}
	return 12
}

// findCommandsAmongUserMsg ...
func findCommandsAmongUserMsg(userMsg string, commands []string) bool {
	expr := fmt.Sprintf(`(?i)(%s)`, strings.Join(commands, "|"))
	amongRegexp := regexp.MustCompile(expr)
	return amongRegexp.MatchString(userMsg)
}

// isUserMsgAmongCommands ...
func isUserMsgAmongCommands(userMsg string, commands []string) bool {
	for _, message := range commands {
		if strings.EqualFold(userMsg, message) {
			return true
		}
	}
	return false
}
