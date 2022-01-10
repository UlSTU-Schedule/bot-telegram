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

func (b *Bot) handleMsg(message *tgbotapi.Message) error {
	userMsgLowered := strings.ToLower(message.Text)

	switch {
	case contains(b.commands.Whole.GetScheduleForDay, userMsgLowered):
		return b.handleGetScheduleForDayMsg(message)
	case contains(b.commands.Whole.GetScheduleForWeek, userMsgLowered) ||
		containsPartial(b.commands.Partial.GetScheduleForWeek, userMsgLowered):
		return b.handleGetScheduleForWeekMsg(message)
	case contains(b.commands.Whole.ChangeGroup, userMsgLowered) ||
		containsPartial(b.commands.Partial.ChangeGroup, userMsgLowered):
		return b.handleChangeGroupMsg(message)
	case contains(b.commands.Whole.BackToStartMenu, userMsgLowered) ||
		containsPartial(b.commands.Partial.BackToStartMenu, userMsgLowered):
		return b.handleBackToStartMenuMsg(message)
	case contains(b.commands.Whole.GoToScheduleMenu, userMsgLowered) ||
		containsPartial(b.commands.Partial.GoToScheduleMenu, userMsgLowered):
		return b.handleGoToScheduleMenuMsg(message)
	case containsPartial(b.commands.Partial.ExpressGratitude, userMsgLowered):
		return b.handleExpressGratitudeMsg(message)
	case contains(b.commands.Whole.Session, userMsgLowered) ||
		containsPartial(b.commands.Partial.Session, userMsgLowered):
		return b.handleSessionMsg(message)
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
		loweredUserMsg := strings.ToLower(message.Text)

		daySchedule, err := schedule.GetDayGroupSchedule(student.GroupName, loweredUserMsg)
		if err != nil {
			groupScheduleJSON, err := b.scheduleStore.GroupSchedule().GetSchedule(student.GroupName)
			if err != nil {
				return err
			}

			updateTimeFmt := groupScheduleJSON.UpdateTime.Format("15:04 02.01.2006")

			daySchedule, err = schedule.ParseDayGroupSchedule(groupScheduleJSON.Info, updateTimeFmt, student.GroupName, loweredUserMsg)
			if err != nil {
				return err
			}
		}

		if schedule.IsKEIGroup(student.GroupName) {
			daySchedule += b.messages.ChangesInKEISchedule
		}

		ansMsg := tgbotapi.NewMessage(message.Chat.ID, daySchedule)
		ansMsg.ReplyMarkup = getScheduleMenuKeyboard()

		_, err = b.bot.Send(ansMsg)
	} else {
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, b.messages.GroupNotSelected)
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
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, "Генерирую расписание \U0000231B")
		_, _ = b.bot.Send(ansMsg)

		loweredUserMsg := strings.ToLower(message.Text)

		caption, weekSchedulePath, err := schedule.GetWeekGroupSchedule(student.GroupName, loweredUserMsg)
		if err != nil {
			groupScheduleJSON, err := b.scheduleStore.GroupSchedule().GetSchedule(student.GroupName)
			if err != nil {
				return err
			}

			updateTimeFmt := groupScheduleJSON.UpdateTime.Format("15:04 02.01.2006")

			caption, weekSchedulePath, err = schedule.ParseWeekGroupSchedule(groupScheduleJSON.Info, updateTimeFmt, student.GroupName, loweredUserMsg)
			if err != nil {
				return err
			}
		}
		if weekSchedulePath != "" {
			defer os.Remove(weekSchedulePath)
		}

		imgMsg := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FilePath(weekSchedulePath))
		imgMsg.Caption = caption
		if (loweredUserMsg == "5" || loweredUserMsg == "текущая неделя") && schedule.IsKEIGroup(student.GroupName) {
			imgMsg.Caption += b.messages.ChangesInKEISchedule
		}
		imgMsg.ReplyMarkup = getScheduleMenuKeyboard()

		_, err = b.bot.Send(imgMsg)
		return err
	} else {
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, b.messages.GroupNotSelected)
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
		ansText := fmt.Sprintf("Твоя группа: %s \U0001F4CC\n\n", student.GroupName)
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

func (b *Bot) handleSessionMsg(message *tgbotapi.Message) error {
	ansText := b.messages.Session
	ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)

	_, err := b.bot.Send(ansMsg)
	return err
}

func (b *Bot) handleUnknownMsg(message *tgbotapi.Message) error {
	if isGroup, groupName := schedule.IsGroupParser(message.Text); isGroup {
		return b.updateGroup(message.From.FirstName, message.From.LastName, message.From.ID, message.Chat.ID, groupName)
	} else {
		groups, err := b.scheduleStore.GroupSchedule().GetGroups()
		if err != nil {
			return err
		}

		if isGroup, groupName = schedule.IsGroupReserver(groups, message.Text); isGroup {
			return b.updateGroup(message.From.FirstName, message.From.LastName, message.From.ID, message.Chat.ID, groupName)
		} else {
			ansMsg := tgbotapi.NewMessage(message.Chat.ID, b.messages.IncorrectInput)

			_, err = b.bot.Send(ansMsg)
			return err
		}
	}
}

func (b *Bot) updateGroup(firstName, lastName string, userID, chatID int64, groupName string) error {
	facultyID := b.determineFacultyID(groupName)

	err := b.studentStore.Student().Information(firstName, lastName, int(userID), groupName, facultyID)
	if err != nil {
		return err
	}

	ansText := fmt.Sprintf("Твоя группа обновлена на %s \U00002705\n\n", groupName)
	ansText += b.messages.InfoWithGroup

	ansMsg := tgbotapi.NewMessage(chatID, ansText)
	ansMsg.ReplyMarkup = getScheduleMenuKeyboard()

	_, err = b.bot.Send(ansMsg)
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

func containsPartial(s []string, e string) bool {
	amongRegexp := regexp.MustCompile(strings.Join(s, "|"))
	return amongRegexp.MatchString(e)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if e == a {
			return true
		}
	}
	return false
}
