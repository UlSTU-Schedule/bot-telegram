package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ulstu-schedule/bot-telegram/internal/schedule"
	"github.com/ulstu-schedule/parser/types"
	"log"
	"os"
	"regexp"
	"strings"
)

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
	log.Printf("@%s: %s", message.From.UserName, message.Text)
	log.Printf("MESSAGE ERROR: %s", err.Error())

	switch err.(type) {
	case *types.UnavailableScheduleError, *types.IncorrectLinkError:
		msg := tgbotapi.NewMessage(message.Chat.ID, b.answers.ToError.ScheduleIsUnavailable)
		_, _ = b.bot.Send(msg)
	case *types.StatusCodeError:
		msg := tgbotapi.NewMessage(message.Chat.ID, b.answers.ToError.ServerError)
		_, _ = b.bot.Send(msg)
	case *types.IncorrectDateError:
		msg := tgbotapi.NewMessage(message.Chat.ID, b.answers.ToError.IncorrectDateError)
		_, _ = b.bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, b.answers.ToError.UnknownError)
		_, _ = b.bot.Send(msg)
	}
}

func (b *Bot) handleTextMessage(message *tgbotapi.Message) error {
	userMsgLowered := strings.ToLower(message.Text)

	switch {
	case containsPartial(b.commands.ExpressGratitude, userMsgLowered):
		return b.handleExpressGratitudeMsg(message)
	case contains(b.commands.OldWhole, userMsgLowered) ||
		containsPartial(b.commands.OldPartial, userMsgLowered):
		return b.handleOldTextCommandMsg(message)
	default:
		return b.handleUnknownMsg(message)
	}
}

func (b *Bot) handleOldTextCommandMsg(message *tgbotapi.Message) error {
	ansText := b.answers.ToTextMessage.RedirectToInline
	ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
	ansMsg.ReplyMarkup = b.hideReplyKeyboard()

	_, err := b.bot.Send(ansMsg)
	return err
}

func (b *Bot) handleChangeGroupMsg(message *tgbotapi.Message) error {
	ansText := b.answers.ToTextMessage.ChangeGroup
	ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
	ansMsg.ReplyMarkup = b.hideReplyKeyboard()

	_, err := b.bot.Send(ansMsg)
	return err
}

func (b *Bot) handleGetScheduleForDayMsg(message *tgbotapi.Message) error {
	student, err := b.studentStore.Student().GetStudent(int(message.From.ID))
	if err != nil {
		return err
	}

	if student != nil {
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, "Генерирую расписание \U0000231B")
		_, _ = b.bot.Send(ansMsg)

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
			daySchedule += b.answers.ToTextMessage.ChangesInKEISchedule
		}

		ansMsg = tgbotapi.NewMessage(message.Chat.ID, daySchedule)
		ansMsg.ReplyMarkup = b.hideReplyKeyboard()

		_, err = b.bot.Send(ansMsg)
	} else {
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, b.answers.ToTextMessage.GroupNotSelected)
		ansMsg.ReplyMarkup = b.hideReplyKeyboard()

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
			defer func(path string) {
				err = os.Remove(path)
				if err != nil {
					log.Printf("error occured while removing week schedule image: %s", err.Error())
				}
			}(weekSchedulePath)
		}

		imgMsg := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FilePath(weekSchedulePath))
		imgMsg.Caption = caption
		if (loweredUserMsg == "5" || loweredUserMsg == "текущая неделя") && schedule.IsKEIGroup(student.GroupName) {
			imgMsg.Caption += b.answers.ToTextMessage.ChangesInKEISchedule
		}
		imgMsg.ReplyMarkup = b.hideReplyKeyboard()

		_, err = b.bot.Send(imgMsg)
		return err
	} else {
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, b.answers.ToTextMessage.GroupNotSelected)
		ansMsg.ReplyMarkup = b.hideReplyKeyboard()

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
		ansText := b.answers.ToTextMessage.Back
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
		ansMsg.ReplyMarkup = b.hideReplyKeyboard()

		_, err = b.bot.Send(ansMsg)
	} else {
		ansText := b.answers.ToTextMessage.ChangeGroup
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
		ansMsg.ReplyMarkup = b.hideReplyKeyboard()

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
		ansText += b.answers.ToTextMessage.InfoWithGroup
		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
		ansMsg.ReplyMarkup = b.hideReplyKeyboard()

		_, err = b.bot.Send(ansMsg)
	} else {
		ansText := b.answers.ToTextMessage.InfoWithoutGroup

		ansMsg := tgbotapi.NewMessage(message.Chat.ID, ansText)
		ansMsg.ReplyMarkup = b.hideReplyKeyboard()

		_, err = b.bot.Send(ansMsg)
	}
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
			ansMsg := tgbotapi.NewMessage(message.Chat.ID, b.answers.ToTextMessage.IncorrectInput)

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
	ansText += b.answers.ToTextMessage.InfoWithGroup

	ansMsg := tgbotapi.NewMessage(chatID, ansText)
	ansMsg.ReplyMarkup = b.hideReplyKeyboard()

	_, err = b.bot.Send(ansMsg)
	return err
}

func (b *Bot) handleExpressGratitudeMsg(message *tgbotapi.Message) error {
	msg := tgbotapi.NewSticker(message.Chat.ID, tgbotapi.FileID(b.answers.Stickers.ToExpressGratitude))
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleSticker(message *tgbotapi.Message) error {
	msg := tgbotapi.NewSticker(message.Chat.ID, tgbotapi.FileID(b.answers.Stickers.ToSticker))
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleVoice(message *tgbotapi.Message) error {
	msg := tgbotapi.NewSticker(message.Chat.ID, tgbotapi.FileID(b.answers.Stickers.ToVoice))
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
