package server

import (
	"DB_course_paper/server/config"
	"DB_course_paper/server/database"
	"log"
)

func Run() {
	log.Println("starting server")
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDB(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
}
