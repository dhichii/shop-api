package env

import (
	"os"
)

type sqlEnv struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func GetSQLEnv() sqlEnv {
	return sqlEnv{
		Host:     os.Getenv("SQL_HOST"),
		Port:     os.Getenv("SQL_PORT"),
		User:     os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		DBName:   os.Getenv("SQL_DB_NAME"),
	}
}
