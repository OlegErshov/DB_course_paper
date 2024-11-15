package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (h Handler) GetAllTeachersTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	accessToken := c.GetHeader(authHeader)

	teacherId, err := h.GetUserIdFromJwt(accessToken)
	if err != nil {
		errorText(c.Writer, "Something went wrong", http.StatusUnauthorized)
		return
	}

	tasks, err := h.service.GetTeacherTopics(ctx, teacherId)

	j, err := json.Marshal(tasks)
	if err != nil {
		log.Println("SignUp handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write(j)
	if err != nil {
		log.Println("CreateCart handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
