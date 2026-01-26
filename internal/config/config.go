// Package config - Config
package config

import "github.com/versenilvis/sendio/internal/utils"

type Config struct {
	Port string
	Env  string

	DatabaseURL string

	RedisURL string

	SMTPHost string
	SMTPPort string
	SMTPUser string
	SMTPPass string

	APIKeySalt    string
	EncryptionKey string
}

func Load() *Config {
	return &Config{
		Port:          utils.GetEnv("PORT", "3000"),
		Env:           utils.GetEnv("ENV", "development"),
		DatabaseURL:   utils.GetEnv("DATABASE_URL", "postgres://sendio:sendio@localhost:5432/sendio?sslmode=disable"),
		RedisURL:      utils.GetEnv("REDIS_URL", "redis://localhost:6379"),
		SMTPHost:      utils.GetEnv("SMTP_HOST", "localhost"),
		SMTPPort:      utils.GetEnv("SMTP_PORT", "1025"),
		SMTPUser:      utils.GetEnv("SMTP_USER", ""),
		SMTPPass:      utils.GetEnv("SMTP_PASS", ""),
		APIKeySalt:    utils.GetEnv("API_KEY_SALT", ""),   // IMPORTANT: change this in production
		EncryptionKey: utils.GetEnv("ENCRYPTION_KEY", ""), // 32 byte key for aes encryption!
	}
}
