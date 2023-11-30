package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port      string         `mapstructure:"port"`
	Database  DatabaseConfig `mapstructure:"database"`
	JWTSecret string         `mapstructure:"jwt_secret"`
	SendGrid  SendGridConfig `mapstructure:"sendgrid"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type SendGridConfig struct {
	ApiKey string `mapstructure:"api_key"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to load config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	log.Println("[Config] Loaded configuration:")
	log.Printf("  Port: %s\n", config.Port)
	log.Printf("  Database:\n")
	log.Printf("    Host: %s\n", config.Database.Host)
	log.Printf("    Port: %s\n", config.Database.Port)
	log.Printf("    Name: %s\n", config.Database.Name)
	log.Printf("    User: %s\n", config.Database.User)
	log.Printf("  JWTSecret: %s\n", config.JWTSecret)
	log.Printf("  SendGrid:\n")
	log.Printf("    API Key: %s\n", config.SendGrid.ApiKey)

	return &config, nil
}
