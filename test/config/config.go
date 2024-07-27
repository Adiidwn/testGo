package config

import (
	"fmt"
	"os"
	"test/helpers"
	"test/resource/db"

	"github.com/pkg/errors"
)

type Configuration struct {
	App AppConfig
	Db  db.DBConfiguration
}

type AppConfig struct {
	Environment string
	Debug       bool
	Timezone    string
	Port        string
}

func ReadConfiguration() (Configuration, error) {
	if err := helpers.LoadEnv(".env"); err != nil {
		causer := errors.Cause(err)
		if os.IsNotExist(causer) {
			fmt.Println("using default env config")
		} else {
			panic(causer)
		}
	}

	cfg := Configuration{
		App: AppConfig{
			Environment: helpers.EnvString("ENVIRONMENT", ""),
			Debug:       helpers.EnvBool("DEBUG", true),
			Timezone:    helpers.EnvString("TIMEZONE", "Asia/Jakarta"),
			Port:        helpers.EnvString("PORT", "8080"),
		},
		Db: db.DBConfiguration{
			Host:           helpers.EnvString("DB_HOST", "127.0.0.1"),
			DBName:         helpers.EnvString("DB_NAME", "postgres"),
			Username:       helpers.EnvString("DB_USERNAME", "postgres"),
			Password:       helpers.EnvString("DB_PASSWORD", "postgres"),
			Port:           helpers.EnvString("DB_PORT", "5432"),
			Logging:        helpers.EnvBool("DB_LOGGING", false),
			Schema:         helpers.EnvString("DB_SCHEMA", ""),
			ConnectTimeout: helpers.EnvInt("DB_CONNECT_TIMEOUT", 30),
			MaxOpenConn:    helpers.EnvInt("DB_MAX_OPEN_CONN", 50),
			MaxIdleConn:    helpers.EnvInt("DB_MAX_IDDLE_CONN", 10),
		},
	}
	return cfg, nil
}
