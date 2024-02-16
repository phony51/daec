package utils

import (
	"encoding/json"
	"io"
	"os"
)

type address struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}
type user struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type httpConfig struct {
	Address address `json:"address"`
	Timeout string  `json:"timeout"`
}
type database struct {
	Driver             string `json:"driver"`
	Source             string `json:"source"`
	MaxOpenConnections int    `json:"maxOpenConnections"`
}

type rabbitmqConfig struct {
	Address       address `json:"address"`
	User          user    `json:"user"`
	MaxGoroutines int     `json:"maxGoroutines"`
}
type manager struct {
	Databases  []database `json:"databases"`
	HTTPConfig httpConfig `json:"http"`
}
type agent struct {
	RabbitMQConfig rabbitmqConfig `json:"rabbitmq"`
}
type logger struct {
	OutFilename   string `json:"outFilename"`
	FullTimestamp bool   `json:"fullTimestamp"`
	MinLevel      int    `json:"minLevel"`
}
type ProjectConfig struct {
	Logger    logger  `json:"logger"`
	DebugMode bool    `json:"debugMode"`
	Manager   manager `json:"manager"`
	Agent     agent   `json:"agent"`
}

func LoadJSONConfig(cfg *ProjectConfig, fn string) error {
	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	raw, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, cfg)
	return err

}
