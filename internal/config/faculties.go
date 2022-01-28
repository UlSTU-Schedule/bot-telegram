package config

import (
	"encoding/json"
	"io/ioutil"
)

const facultiesJSONPath = "configs/faculties.json"

// Faculty represents UlSTU faculty.
type Faculty struct {
	Name   string
	ID     byte
	Groups []string
}

func unmarshalFacultiesTo(cfg *Config) error {
	data, err := ioutil.ReadFile(facultiesJSONPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &cfg.Faculties)
}
