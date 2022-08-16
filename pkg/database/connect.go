package database

import (
	"fmt"
	"os"
)

type Config struct {
	DB       string
	Host     string
	User     string
	Password string
}

func BuildDSN(connType string, conf Config) string {
	user := conf.User
	if user == "" {
		user = os.Getenv("DB_USER")
	}
	host := conf.Host
	if host == "" {
		host = os.Getenv("DB_HOST")
	}
	dbName := conf.DB
	if dbName == "" {
		dbName = os.Getenv("DB_NAME")
	}
	return fmt.Sprintf("%s://%s:%s@%s:5432/%s?sslmode=disable", connType, user, conf.Password, host, dbName)
}
