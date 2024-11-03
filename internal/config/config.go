package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + configFileName, nil
}

func write(config Config) error {
	jsonData, err := json.Marshal(config)
	if err != nil {
		return err
	}
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error retrieving config's file path: %v", err)
	}
	f, err := os.Open(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("error opening file: %v", err)
	}
	decoder := json.NewDecoder(f)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, fmt.Errorf("error decoding json: %v", err)
	}
	return config, nil
}

func (c Config) SetUser(name string) error {
	c.CurrentUserName = name
	err := write(c)
	if err != nil {
		fmt.Errorf("error writing config to file: %v", err)
	}
	return nil
}
