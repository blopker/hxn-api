package main

import (
	"encoding/json"
	"os"
)

// Config for your data
type Config struct {
	Token string `json:"token"`
}

// Load data from config.json
func (c *Config) Load() error {
	file, err := os.Open("config.json")

	if err != nil {
		return err
	}

	if err := json.NewDecoder(file).Decode(c); err != nil {
		return err
	}

	return nil
}
