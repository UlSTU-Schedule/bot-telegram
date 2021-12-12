package schedule

import (
	"fmt"
	"github.com/ulstu-schedule/parser/schedule"
	"log"
	"regexp"
	"strings"
)

var (
	keiGroupPattern   = regexp.MustCompile(`^[А-Я]+[до]+[-0-9]+$`)
	groupCharsPattern = regexp.MustCompile(`[А-Яа-я0-9-]+`)
)

func GetDailySchedule(userGroup string, userMsg string) (string, error) {
	if userMsg == "3" || userMsg == "сегодня" {
		dailySchedule, err := schedule.GetTextDailyGroupSchedule(userGroup, 0)
		if err != nil {
			log.Printf("Ошибка при получении расписания на сегодня: %v, %T", err, err)
			return "", err
		}

		return dailySchedule, nil
	} else {
		dailySchedule, err := schedule.GetTextDailyGroupSchedule(userGroup, 1)
		if err != nil {
			log.Printf("Ошибка при получении расписания на завтра: %v, %T", err, err)
			return "", err
		}

		return dailySchedule, nil
	}
}

func GetWeeklySchedule(groupName, userMsg string) (string, string, error) {
	if userMsg == "5" || userMsg == "текущая неделя" {
		weeklySchedule, err := schedule.GetCurrWeekGroupScheduleImg(groupName)
		if err != nil {
			return "", "", err
		}

		caption := fmt.Sprintf("Расписание %s на текущую неделю\U0001F446\n\n", groupName)

		return caption, weeklySchedule, nil
	} else {
		weeklySchedule, err := schedule.GetNextWeekGroupScheduleImg(groupName)
		if err != nil {
			return "", "", err
		}

		caption := fmt.Sprintf("Расписание %s на следующую неделю\U0001F446", groupName)

		return caption, weeklySchedule, nil
	}
}

func IsGroupExist(input string) (bool, string, error) {
	lowerInput := strings.ToLower(input)
	convertedInput := convertToGroupName(lowerInput)

	groups := schedule.GetGroups()
	for _, group := range groups {
		if strings.Contains(group, ", ") {
			splitGroups := strings.Split(group, ", ")
			for _, splitGroup := range splitGroups {
				if convertedInput == strings.ToLower(splitGroup) {
					return true, splitGroup, nil
				}
			}
		} else {
			if convertedInput == strings.ToLower(group) {
				return true, group, nil
			}
		}
	}

	return false, "", nil
}

func convertToGroupName(input string) string {
	inputWithoutExcessSymbols := deleteExcessSymbols(input)
	corrGroupNameInRunes := make([]rune, 0, len(inputWithoutExcessSymbols)+2)

	var afterNum, quantityNum int
	for _, character := range inputWithoutExcessSymbols {
		switch {
		case character == '-':
			afterNum, quantityNum = 0, 0
		case 48 <= character && character <= 57 && afterNum == 1 && quantityNum != 2:
			afterNum = 0
			quantityNum++
			corrGroupNameInRunes = append(corrGroupNameInRunes, '-')
		case 48 <= character && character <= 57 && afterNum == 0 && quantityNum != 2:
			quantityNum++
		case quantityNum == 2:
			quantityNum = 0
			corrGroupNameInRunes = append(corrGroupNameInRunes, '-')
		default:
			afterNum = 1
		}
		corrGroupNameInRunes = append(corrGroupNameInRunes, character)
	}

	return string(corrGroupNameInRunes)
}

func deleteExcessSymbols(s string) string {
	res := groupCharsPattern.FindAllString(s, -1)
	return strings.Join(res, "")
}

func IsGroupFromKEI(groupName string) bool {
	return keiGroupPattern.MatchString(groupName) && !strings.Contains(groupName, "РОНд")
}
