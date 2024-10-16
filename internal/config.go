package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

// Declare public and private variables for the configuration
var (
	// Public variables
	Token     string // Token to authenticate the bot
	BotPrefix string // Prefix to recognize bot commands
	// Private variables
	config *configStruct // Pointer to configStruct to hold configuration
)

// configStruct holds the structure of the configuration file
type configStruct struct {
	Token     string `json:"Token"`     // Token field in JSON
	BotPrefix string `json:"BotPrefix"` // BotPrefix field in JSON
}

// ReadConfig reads the configuration from the config file
func ReadConfig() error {
	fmt.Println("Reading config file...")

	// Read the config file
	file, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	// Unmarshal the JSON content into the configStruct
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Assign the read values to the public variables
	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
