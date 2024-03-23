package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	SERVER_PORT             int
	REQUEST_TIMEOUT         int
	SEND_SLACK_NOTIFICATION bool
	SLACK_API_TOKEN         string
	SLACK_CHANNEL_ID        string
	SLACK_USERS_ID          []string
}

var loadedConfig *Config

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	SERVER_PORT, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("Error parsing SERVER_PORT: %v", err)
	}
	REQUEST_TIMEOUT, err := strconv.Atoi(os.Getenv("REQUEST_TIMEOUT"))
	if err != nil {
		log.Fatalf("Error parsing REQUEST_TIMEOUT: %v", err)
	}

	SEND_SLACK_NOTIFICATION, err := strconv.ParseBool(os.Getenv("SEND_SLACK_NOTIFICATION"))
	if err != nil {
		log.Fatalf("Error parsing SEND_SLACK_NOTIFICATION: %v", err)
	}
	SLACK_API_TOKEN := os.Getenv("SLACK_API_TOKEN")
	SLACK_CHANNEL_ID := os.Getenv("SLACK_CHANNEL_ID")
	SLACK_USERS_ID_ARRAY_STR := os.Getenv("SLACK_USERS_ID")
	SLACK_USERS_ID_VALUES := strings.Split(SLACK_USERS_ID_ARRAY_STR, ",")

	loadedConfig = &Config{
		SERVER_PORT:             SERVER_PORT,
		REQUEST_TIMEOUT:         REQUEST_TIMEOUT,
		SEND_SLACK_NOTIFICATION: SEND_SLACK_NOTIFICATION,
		SLACK_API_TOKEN:         SLACK_API_TOKEN,
		SLACK_CHANNEL_ID:        SLACK_CHANNEL_ID,
		SLACK_USERS_ID:          SLACK_USERS_ID_VALUES,
	}
}

func GetConfig() *Config {
	return loadedConfig
}
