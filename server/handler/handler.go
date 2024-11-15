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
		students := api.Group("/students")
		{
			students.GET("/tasks", h.GetAllTasks)
			students.POST("/tasks/{id}", h.CheckTask)
			students.GET("/tasks/{id}", h.GetTask)
		}
		teachers := api.Group("/teachers")
		{
			topic := teachers.Group("/topic")
			{
				topic.POST("/", h.CreateTopic)
			}
			assign := teachers.Group("/assign")
			{
				assign.POST("/task", h.AssignTaskToTopic)
				assign.POST("/topic", h.AssignTopicToStudent)
			}
			tasks := teachers.Group("/tasks")
			{
				tasks.GET("/", h.GetAllTeachersTask)
				tasks.POST("/create_grammar", h.CreateGrammarTask)
				vocabluary := tasks.Group("/vocabulary")
				{
					vocabluary.POST("/options", h.CreateVocabularyOptionsTask)
					vocabluary.POST("/sentence", h.CreateVocabularySentenceTask)
					vocabluary.POST("/word", h.CreateVocabularyWordTask)
				}
			}

		}
		journal := api.Group("/journal")
		{
			journal.GET("/marks", h.GetStudentMarks)
		}
	}
	return router
}
