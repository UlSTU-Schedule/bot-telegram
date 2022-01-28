package config

import "github.com/spf13/viper"

const commandsConfigPath = "commands"

// Commands represents the bot commands.
type Commands struct {
	WithoutSlash `mapstructure:"without_slash"`
	WithSlash    `mapstructure:"with_slash"`
}

// WithoutSlash represents commands that do not start with a slash character.
type WithoutSlash struct {
	Inline `mapstructure:"inline"`

	// OldWhole represents commands, which should be a whole message (е.g. msg='завтра', cmd='завтра').
	// Using to redirect users on inline keyboard messages
	OldWhole []string `mapstructure:"old_whole"`

	// OldPartial represents commands that can be inside a message (msg='можно моё расписание?', cmd='расписание').
	// Using to redirect users on inline keyboard messages
	OldPartial []string `mapstructure:"old_partial"`

	ExpressGratitude []string `mapstructure:"express_gratitude"`
}

// Inline represents the inline keyboard commands and data.
type Inline struct {
	First  FirstLvlKeyboard  `mapstructure:"first_lvl"`
	Second SecondLvlKeyboard `mapstructure:"second_lvl"`
	Third  ThirdLvlKeyboard  `mapstructure:"third_lvl"`
}

// InlineButtonInfo represents the information that will be contained in the inline keyboard button.
type InlineButtonInfo struct {
	Command string `mapstructure:"command"`
	Data    string `mapstructure:"data"`
}

// FirstLvlKeyboard represents inline keyboard commands and data that are on the first level of the keyboard.
type FirstLvlKeyboard struct {
	Groups   InlineButtonInfo `mapstructure:"groups"`
	Teachers InlineButtonInfo `mapstructure:"teachers"`
}

// SecondLvlKeyboard represents inline keyboard commands and data that are on the second level of the keyboard.
type SecondLvlKeyboard struct {
	Groups   SecondLvlKeyboardSection `mapstructure:"groups"`
	Teachers SecondLvlKeyboardSection `mapstructure:"teachers"`
	Back     InlineButtonInfo         `mapstructure:"back"`
}

// SecondLvlKeyboardSection represents schedule and change the inline keyboard section of teachers or groups.
type SecondLvlKeyboardSection struct {
	Schedule InlineButtonInfo `mapstructure:"schedule"`
	Change   InlineButtonInfo `mapstructure:"change"`
}

// ThirdLvlKeyboard represents inline keyboard commands and data that are on the third level of the keyboard.
type ThirdLvlKeyboard struct {
	Groups   ThirdLvlKeyboardSection `mapstructure:"groups"`
	Teachers ThirdLvlKeyboardSection `mapstructure:"teachers"`
}

// ThirdLvlKeyboardSection represents teachers or groups section of the inline keyboard.
type ThirdLvlKeyboardSection struct {
	Today    InlineButtonInfo `mapstructure:"today"`
	Tomorrow InlineButtonInfo `mapstructure:"tomorrow"`
	CurrWeek InlineButtonInfo `mapstructure:"curr_week"`
	NextWeek InlineButtonInfo `mapstructure:"next_week"`
	Back     InlineButtonInfo `mapstructure:"back"`
}

// WithSlash represents commands that start with a slash character.
type WithSlash struct {
	Start string `mapstructure:"start"`
	Help  string `mapstructure:"help"`
	About string `mapstructure:"about"`
}

func unmarshalCommandsTo(cfg *Config) error {
	if err := setUpViper(commandsConfigPath); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("without_slash", &cfg.Commands.WithoutSlash); err != nil {
		return err
	}

	return viper.UnmarshalKey("with_slash", &cfg.Commands.WithSlash)
}
