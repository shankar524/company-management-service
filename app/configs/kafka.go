package configs

import "os"

// Config object
type KafkaConfig struct {
	Host string
	Port string
}

// GetConfig gets all config for the application
func GetKafkaConfig() KafkaConfig {
	return KafkaConfig{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),
	}
}
