package config

import "github.com/spf13/viper"

const answersConfigPath = "answers"

// Answers represents the messages that the bot sends to the user.
type Answers struct {
	ToCommand     AnswersToCommands     `mapstructure:"to_command"`
	ToQuery       AnswersToQueries      `mapstructure:"to_query"`
	ToError       AnswersToErrors       `mapstructure:"to_error"`
	ToTextMessage AnswersToTextMessages `mapstructure:"to_text_message"`

	Stickers Stickers `mapstructure:"stickers"`
}

// AnswersToCommands represents the answers that will be sent to the commands.
type AnswersToCommands struct {
	Start   string `mapstructure:"start"`
	About   string `mapstructure:"about"`
	Help    string `mapstructure:"help"`
	Unknown string `mapstructure:"unknown"`
}

// AnswersToQueries represents the answers that will be sent to the queries data.
type AnswersToQueries struct {
	FirstLvlMenu string `mapstructure:"first_lvl_menu"`
}

// AnswersToErrors represents the answers that will be sent to the occurred error.
type AnswersToErrors struct {
	ScheduleIsUnavailable string `mapstructure:"schedule_is_unavailable"`
	ServerError           string `mapstructure:"server_error"`
	IncorrectDateError    string `mapstructure:"incorrect_date_error"`
	UnknownError          string `mapstructure:"unknown_error"`
}

// AnswersToTextMessages represents the answers that will be sent to the text message.
type AnswersToTextMessages struct {
	RedirectToInline string `mapstructure:"redirect_to_inline"`
	ChangeGroup      string `mapstructure:"change_group"`
	Back             string `mapstructure:"back"`
	InfoWithoutGroup string `mapstructure:"info_without_group"`
	InfoWithGroup    string `mapstructure:"info_with_group"`
	IncorrectInput   string `mapstructure:"incorrect_input"`
	GroupNotSelected string `mapstructure:"group_not_selected"`

	ChangesInKEISchedule string `mapstructure:"changes_in_kei_schedule"`
}

// Stickers represents the sticker codes that the bot sends.
type Stickers struct {
	ToExpressGratitude string `mapstructure:"to_express_gratitude"`
	ToSticker          string `mapstructure:"to_sticker"`
	ToVoice            string `mapstructure:"to_voice"`
}

func unmarshalAnswersTo(cfg *Config) error {
	if err := setUpViper(answersConfigPath); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	return viper.UnmarshalKey("answers", &cfg.Answers)
}
