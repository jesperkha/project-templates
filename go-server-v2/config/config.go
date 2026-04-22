package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Environment string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:        toGoPort(getOrDefault("PORT", ":8080")),
		Environment: getOrDefault("ENVIRONMENT", "development"),
	}
}

func getOrDefault(key string, defaltV string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaltV
	}
	return v
}

func toGoPort(port string) string {
	if port == "" {
		return ":8080"
	}
	if port[0] != ':' {
		return ":" + port
	}
	return port
}
