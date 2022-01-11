package config

import "os"

type Config struct {
	DBUser     string
	DBPass     string
	DBBase     string
	SigningKey string
}

func LoadConfig() Config {
	config := Config{
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPass:     os.Getenv("POSTGRES_PASS"),
		DBBase:     os.Getenv("POSTGRES_BASE"),
		SigningKey: os.Getenv("SIGN_KEY"),
	}
	return config
}
