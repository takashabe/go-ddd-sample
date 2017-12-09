package config

import (
	"database/sql"
	"fmt"
	"os"
)

// NewDBConnection returns initialized sql.DB
func NewDBConnection() (*sql.DB, error) {
	user := getEnvWithDefault("DB_USER", "root")
	password := getEnvWithDefault("DB_PASSWORD", "")
	host := getEnvWithDefault("DB_HOST", "localhost")
	port := getEnvWithDefault("DB_PORT", "3306")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/ddd_sample?parseTime=true", user, password, host, port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getEnvWithDefault(name, def string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return def
}
