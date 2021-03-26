package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	DB DB `json:"db"`
}

type DB struct {
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DBName string `json:"dbName"`
}

var Conf Config

func GetConfig(inputConfPath string) error {
	configPath := "./data.json"
	if inputConfPath != "" {
		configPath = inputConfPath
	}
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	var data Config
	json.Unmarshal(file, &data)
	Conf = data
	return nil
}
