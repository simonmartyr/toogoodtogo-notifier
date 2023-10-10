package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	EmailConfig EmailConfig            `json:"email_config"`
	Credentials TooGoodToGoCredentials `json:"credentials"`
	Items       []Item                 `json:"items"`
}

type EmailConfig struct {
	To      string `json:"to"`
	Account string `json:"account"`
}

type Item struct {
	Name         string `json:"name"`
	ItemName     string `json:"item_name"`
	ItemId       string `json:"item_id"`
	Notify       bool   `json:"notify"`
	LastNotified string `json:"last_notified"`
}

type TooGoodToGoCredentials struct {
	Email        string `json:"email"`
	UserId       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Cookie       string `json:"cookie"`
}

func (config *Config) WriteConfig() {
	toWrite, writeErr := json.Marshal(config)
	if writeErr != nil {
		log.Fatal(writeErr)
		return
	}
	osWErr := os.WriteFile("tgtgconfig.json", toWrite, 0644)
	if osWErr != nil {
		log.Fatal(osWErr)
		return
	}
}

func LoadConfig() (*Config, error) {
	fileContent, err := os.Open("tgtgconfig.json")
	if err != nil {
		return nil, err
	}
	var config = Config{}
	jsonErr := json.NewDecoder(fileContent).Decode(&config)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &config, nil
}
