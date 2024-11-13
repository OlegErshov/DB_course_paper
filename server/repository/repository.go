package repository

import (
	"DB_course_paper/server/service"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) service.Repositorier {
	return repository{
		db: db,
	}
}
