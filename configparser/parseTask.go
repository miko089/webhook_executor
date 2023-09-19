package configparser

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
)

// Task is struct for unmarshalling config.yaml
type Task struct {
	ID      string `yaml:"id"`
	Command string `yaml:"command"`
}

// Config is struct for unmarshalling config.yaml
type Config struct {
	Tasks []*Task `yaml:"tasks"`
}

func parseFile(data []byte) (*Config, error) {

	var config *Config
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// GetCommand gives command from file with some id
func GetCommand(filename, id string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	config, err := parseFile(data)
	if err != nil {
		return "", err
	}
	for _, task := range config.Tasks {
		if task.ID == id {
			return task.Command, nil
		}
	}
	return "", errors.New("id not found")
}
