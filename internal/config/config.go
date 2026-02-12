// Package config - Config
package config

import (
	"log"

	"github.com/versenilvis/sendio/internal/utils"
)

type Config struct {
	Port      string         `json:"port"`
	Env       string         `json:"env"`
	ServerURL string         `json:"server_url"`
	WebURL    string         `json:"web_url"`
	Database  DatabaseConfig `json:"database"`
	Redis     RedisConfig    `json:"redis"`
	SMTP      SMTPConfig     `json:"smtp"`
	Security  SecurityConfig `json:"security"`
	OAuth     OAuthConfig    `json:"oauth"`
}

type DatabaseConfig struct {
	URL      string `json:"-"` // sensitive information please use "-"
	MaxConns int    `json:"max_conns"`
	MinConns int    `json:"min_conns"`
}

type RedisConfig struct {
	URL string `json:"-"`
}

type SMTPConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"-"`
}

type SecurityConfig struct {
	APIKeySalt    string `json:"-"`
	EncryptionKey string `json:"-"`
	SessionSecret string `json:"-"`
}

type OAuthConfig struct {
	GoogleClientID     string `json:"google_client_id"`
	GoogleClientSecret string `json:"-"`
	GitHubClientID     string `json:"github_client_id"`
	GitHubClientSecret string `json:"-"`
}

func Load() *Config {
	cfg := &Config{
		Port:      utils.GetEnv("PORT", "8080"),
		Env:       utils.GetEnv("APP_ENV", "development"),
		ServerURL: utils.GetEnv("SERVER_URL", "http://localhost:8080"),
		WebURL:    utils.GetEnv("WEB_URL", "http://localhost:5173"),

		Database: DatabaseConfig{
			URL:      utils.GetEnv("DATABASE_URL", ""),
			MaxConns: utils.GetEnvAsInt("DB_MAX_CONNS", 25),
			MinConns: utils.GetEnvAsInt("DB_MIN_CONNS", 5),
		},

		Redis: RedisConfig{
			URL: utils.GetEnv("REDIS_URL", "redis://localhost:6379"),
		},

		SMTP: SMTPConfig{
			Host:     utils.GetEnv("SMTP_HOST", "localhost"),
			Port:     utils.GetEnv("SMTP_PORT", "1025"),
			User:     utils.GetEnv("SMTP_USER", ""),
			Password: utils.GetEnv("SMTP_PASS", ""),
		},

		Security: SecurityConfig{
			APIKeySalt:    utils.GetEnv("API_KEY_SALT", ""),
			EncryptionKey: utils.GetEnv("ENCRYPTION_KEY", ""),
			SessionSecret: utils.GetEnv("SESSION_SECRET", ""),
		},

		OAuth: OAuthConfig{
			GoogleClientID:     utils.GetEnv("GOOGLE_CLIENT_ID", ""),
			GoogleClientSecret: utils.GetEnv("GOOGLE_CLIENT_SECRET", ""),
			GitHubClientID:     utils.GetEnv("GITHUB_CLIENT_ID", ""),
			GitHubClientSecret: utils.GetEnv("GITHUB_CLIENT_SECRET", ""),
		},
	}

	// if cfg.Env == "production" {
	// 	cfg.Validate()
	// }
	cfg.Validate()

	return cfg
}

func (c *Config) Validate() {
	if c.Database.URL == "" {
		log.Fatal("CRITICAL: DATABASE_URL is not set")
	}

	if c.Security.APIKeySalt == "" {
		log.Fatal("CRITICAL: API_KEY_SALT is not set. It is required for secure authentication")
	}

	if c.Security.SessionSecret == "" {
		log.Fatal("CRITICAL: SESSION_SECRET is not set. It is required for secure sessions")
	}

	if c.Security.EncryptionKey == "" {
		log.Fatal("CRITICAL: ENCRYPTION_KEY is not set. It is required for encrypting sensitive data")
	}

	if len(c.Security.EncryptionKey) != 32 {
		log.Fatalf("CRITICAL: ENCRYPTION_KEY must be exactly 32 characters long for AES-256 (current length: %d)", len(c.Security.EncryptionKey))
	}

	// OPTIONAL: only if you expect these to be always present
	if c.OAuth.GoogleClientID == "" || c.OAuth.GoogleClientSecret == "" {
		log.Println("WARNING: Google OAuth credentials are missing. Google Auth will not work")
	}

	if c.OAuth.GitHubClientID == "" || c.OAuth.GitHubClientSecret == "" {
		log.Println("WARNING: GitHub OAuth credentials are missing. GitHub Auth will not work")
	}
}
