package base

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Conn *Connection
}

func (c *Config) Load() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	c.Conn = &Connection{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
}
