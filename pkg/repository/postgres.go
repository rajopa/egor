package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable   = "users"
	targetsTable = "targets"
)

type Config struct {
	Host     string
	Port     string
	Username string
	DBname   string
	SSLMode  string
	Password string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Open("postgres", dns)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
