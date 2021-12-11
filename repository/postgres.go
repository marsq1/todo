package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Port string
	Host string
	SSLMode string
	Password string
	DBName string
	Username string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Port, cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
