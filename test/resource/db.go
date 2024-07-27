package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfiguration struct {
	Host           string
	Port           string
	Schema         string
	DBName         string
	Username       string
	Password       string
	Logging        bool
	ConnectTimeout int
	MaxOpenConn    int
	MaxIdleConn    int
	SessionName    string
}

func NewDBConnection(config DBConfiguration) (*gorm.DB, error) {
	dbCfg := DBConfiguration{
		Host:           config.Host,
		Port:           config.Port,
		Schema:         config.Schema,
		DBName:         config.DBName,
		Username:       config.Username,
		Password:       config.Password,
		Logging:        config.Logging,
		ConnectTimeout: config.ConnectTimeout,
		MaxOpenConn:    config.MaxOpenConn,
		MaxIdleConn:    config.MaxIdleConn,
		SessionName:    config.SessionName,
	}

	sqlConn, err := sql.Open("postgres", makePostgresString(dbCfg))
	if err != nil {
		return nil, errors.Wrap(err, "can't establish db connection")
	}
	sqlConn.SetMaxIdleConns(dbCfg.MaxIdleConn)
	sqlConn.SetMaxOpenConns(dbCfg.MaxOpenConn)
	sqlConn.SetConnMaxLifetime(time.Hour)

	DB, err := gorm.Open(postgres.New(
		postgres.Config{
			Conn: sqlConn}))
	if err != nil {
		return nil, errors.Wrap(err, "can't open db connection")
	}

	return DB, nil
}

func makePostgresString(p DBConfiguration) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s connect_timeout=%d application_name=%s",
		p.Host, p.Port, p.Username, p.DBName, p.Password, p.ConnectTimeout, p.SessionName)
}
