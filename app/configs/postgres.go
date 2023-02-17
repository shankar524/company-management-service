package configs

import (
	"os"
)

// PostgresConfig object
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// GetPostgresConfig returns PostgresConfig object
func GetPostgresConfig() PostgresConfig {

	return PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
}
