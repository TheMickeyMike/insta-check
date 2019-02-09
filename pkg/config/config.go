package config

import (
	"encoding/json"
	"log"
	"os"
)

const configEnv = "INSTA_CHECK"
const defaultConfig = `
{
    "instagram": {
        "url": "https://www.instagram.com/accounts/web_create_ajax/attempt/"
    }
}`

type InstagramConfig struct {
	URL string `yaml:"url"`
}

type AppConfig struct {
	Instagram *InstagramConfig `yaml:"instagram"`
}

func LoadConfig() *AppConfig {
	var appConfig AppConfig
	log.Println("Loading application config...")
	rawConf := getEnvOrDefault(configEnv, defaultConfig)
	err := json.Unmarshal([]byte(rawConf), &appConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return &appConfig
}

func getEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
