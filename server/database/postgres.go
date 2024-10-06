package database

import (
	"DB_course_paper/server/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(cfg config.DBConfig) (*sqlx.DB, error) {
	conStr := fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Connect("postgres", conStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
