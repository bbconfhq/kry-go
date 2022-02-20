package config

import "os"

type DbConfig struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

func LoadDbConfig() DbConfig {
	return DbConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
}
