package config

import "os"

type Config struct {
	DBUser     string
	DBPass     string
	DBBase     string
	DBHost     string
	DBPort     string
	DBSSLMode  string
	SigningKey string
}

func LoadConfig() Config {
	config := Config{
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPass:     os.Getenv("POSTGRES_PASS"),
		DBBase:     os.Getenv("POSTGRES_BASE"),
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
		DBSSLMode:  os.Getenv("POSTGRES_SSLMODE"),
		SigningKey: os.Getenv("SIGN_KEY"),
	}
	return config
}
