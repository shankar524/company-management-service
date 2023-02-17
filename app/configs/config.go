package configs

import "os"

// Config object
type Config struct {
	Env       string
	JWTSecret string
	Port      string
	Database  PostgresConfig
}

// GetConfig gets all config for the application
func GetConfig() Config {
	return Config{
		Env:       os.Getenv("ENV"),
		Database:  GetPostgresConfig(),
		JWTSecret: os.Getenv("JWT_SIGN_KEY"),
		Port:      os.Getenv("APP_PORT"),
	}
}
