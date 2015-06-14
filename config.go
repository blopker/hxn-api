package main

import (
	"encoding/json"
	"os"
)

// Config for your data
type Config struct {
	Token string `json:"token"`
	Port  string `json:"port"`
}

// Load data from config.json
func (c *Config) Load() {
	c.fromFile()
	c.fromEnv()
}

func (c *Config) fromFile() error {
	file, err := os.Open("config.json")

	if err != nil {
		return err
	}

	if err := json.NewDecoder(file).Decode(c); err != nil {
		return err
	}

	return nil
}

func (c *Config) fromEnv() {
	if port := os.Getenv("PORT"); port != "" {
		c.Port = port
	}

	if token := os.Getenv("TOKEN"); token != "" {
		c.Token = token
	}
}
