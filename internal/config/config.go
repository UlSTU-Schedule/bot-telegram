package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Token               string
	StudentDatabaseURL  string
	ScheduleDatabaseURL string

	Answers   Answers
	Commands  Commands
	Faculties []Faculty
}

func New(configPath string) (*Config, error) {
	cfg := &Config{}

	viper.AddConfigPath(configPath)

	if err := unmarshalAnswersTo(cfg); err != nil {
		return nil, err
	}

	if err := unmarshalCommandsTo(cfg); err != nil {
		return nil, err
	}

	if err := unmarshalFacultiesTo(cfg); err != nil {
		return nil, err
	}

	if err := parseEnvVarsTo(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func parseEnvVarsTo(cfg *Config) error {
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

func setUpViper(pathToConfigFile string) error {
	viper.SetConfigName(pathToConfigFile)
	return viper.ReadInConfig()
}
