package handler

import (
	"DB_course_paper/server/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service   service.Servicer
	jwtSecret []byte
}

func NewHandler(service service.Servicer, jwtSecret []byte) Handler {
	return Handler{service: service, jwtSecret: jwtSecret}
}

func InitRoutes(h *Handler) *gin.Engine {
	router := gin.New()
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
			auth.POST("/log-out", h.authMiddleware(h.jwtSecret), h.logOut)
		}

	}
	return router
}
