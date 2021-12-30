package config

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
)

const (
	answersConfigPath  = "answers"
	commandsConfigPath = "commands"
	facultiesJSONPath  = "configs/faculties.json"
)

// Config ...
type Config struct {
	// environment variables
	Token               string
	StudentDatabaseURL  string
	ScheduleDatabaseURL string

	Stickers  Stickers
	Messages  Messages
	Commands  Commands
	Faculties []Faculty
}

// Commands ...
type Commands struct {
	WithoutSlash
	WithSlash
}

// WithoutSlash ...
type WithoutSlash struct {
	Whole   // команда должна быть такой же, как и весь текст сообщения (msg='завтра', cmd='завтра')
	Partial // команда может быть среди слов в сообщении (msg='можно моё расписание?', cmd='расписание')
}

// Whole ...
type Whole struct {
	GoToScheduleMenu   []string `mapstructure:"go_to_schedule_menu"`
	ChangeGroup        []string `mapstructure:"change_group"`
	GetScheduleForDay  []string `mapstructure:"get_schedule_for_day"`
	GetScheduleForWeek []string `mapstructure:"get_schedule_for_week"`
	BackToStartMenu    []string `mapstructure:"back_to_start_menu"`
}

// Partial ...
type Partial struct {
	GoToScheduleMenu   []string `mapstructure:"go_to_schedule_menu"`
	GetScheduleForWeek []string `mapstructure:"get_schedule_for_week"`
	ChangeGroup        []string `mapstructure:"change_group"`
	BackToStartMenu    []string `mapstructure:"back_to_start_menu"`
	ExpressGratitude   []string `mapstructure:"express_gratitude"`
}

// WithSlash ...
type WithSlash struct {
	Start        string `mapstructure:"start"`
	Help         string `mapstructure:"help"`
	AboutProject string `mapstructure:"about_project"`
}

// Faculty ...
type Faculty struct {
	Name   string
	ID     byte
	Groups []string
}

// Messages ...
type Messages struct {
	// Answers
	StartWithGroup    string `mapstructure:"start_with_group"`
	StartWithoutGroup string `mapstructure:"start_without_group"`
	ChangeGroup       string `mapstructure:"change_group"`
	Back              string `mapstructure:"back"`
	InfoWithoutGroup  string `mapstructure:"info_without_group"`
	InfoWithGroup     string `mapstructure:"info_with_group"`
	AboutProject      string `mapstructure:"about_project"`
	IncorrectInput    string `mapstructure:"incorrect_input"`
	GroupNotSelected  string `mapstructure:"group_not_selected"`

	// Additions
	ChangesInKEISchedule string `mapstructure:"changes_in_kei_schedule"`

	// Errors
	ScheduleIsUnavailable string `mapstructure:"schedule_is_unavailable"`
	ServerError           string `mapstructure:"server_error"`
	IncorrectDateError    string `mapstructure:"incorrect_date_error"`
	UnknownError          string `mapstructure:"unknown_error"`
}

// Stickers ...
type Stickers struct {
	ToExpressGratitude string `mapstructure:"to_express_gratitude"`
	ToSticker          string `mapstructure:"to_sticker"`
	ToVoice            string `mapstructure:"to_voice"`
}

// New ...
func New(configPath string) (*Config, error) {
	cfg := &Config{}

	viper.AddConfigPath(configPath)

	if err := fromAnswers(cfg); err != nil {
		return nil, err
	}

	if err := fromCommands(cfg); err != nil {
		return nil, err
	}

	if err := fromEnv(cfg); err != nil {
		return nil, err
	}

	if err := fromJson(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func fromAnswers(cfg *Config) error {
	if err := setUpViper(answersConfigPath); err != nil {
		return err
	}

	return unmarshalAnswers(cfg)
}

func unmarshalAnswers(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("answers", &cfg.Messages); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("additions", &cfg.Messages); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("errors", &cfg.Messages); err != nil {
		return err
	}

	return viper.UnmarshalKey("stickers", &cfg.Stickers)
}

func fromCommands(cfg *Config) error {
	if err := setUpViper(commandsConfigPath); err != nil {
		return err
	}

	return unmarshalCommands(cfg)
}

func unmarshalCommands(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("without_slash", &cfg.Commands.WithoutSlash); err != nil {
		return err
	}

	return viper.UnmarshalKey("with_slash", &cfg.Commands.WithSlash)
}

func fromEnv(cfg *Config) error {
	if err := viper.BindEnv("TOKEN"); err != nil {
		return err
	}
	cfg.Token = viper.GetString("TOKEN")

	if err := viper.BindEnv("STUDENT_DATABASE_URL"); err != nil {
		return err
	}
	cfg.StudentDatabaseURL = viper.GetString("STUDENT_DATABASE_URL")

	if err := viper.BindEnv("SCHEDULE_DATABASE_URL"); err != nil {
		return err
	}
	cfg.ScheduleDatabaseURL = viper.GetString("SCHEDULE_DATABASE_URL")

	return nil
}

func fromJson(cfg *Config) error {
	data, err := ioutil.ReadFile(facultiesJSONPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &cfg.Faculties)
}

func setUpViper(pathToConfigFile string) error {
	viper.SetConfigName(pathToConfigFile)
	return viper.ReadInConfig()
}
