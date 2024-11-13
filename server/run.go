package server

import (
	"DB_course_paper/server/config"
	"DB_course_paper/server/database"
	"DB_course_paper/server/handler"
	"DB_course_paper/server/repository"
	"DB_course_paper/server/service"
	"fmt"
	"log"
	"net/http"
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
	r := repository.NewRepository(db)
	s := service.NewService(r)

	h := handler.NewHandler(s, cfg.Jwt.Secret)
	router := handler.InitRoutes(&h)
	server := newServer(router, cfg.Server)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func newServer(router http.Handler, cfg config.ServerConfig) *http.Server {
	servAddr := fmt.Sprintf("%s:%v", cfg.Host, cfg.Port)
	return &http.Server{
		Addr:    servAddr,
		Handler: router,
	}
}
