package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type assignTaskRequest struct {
	TaskId  int `json:"task_id"`
	TopicId int `json:"topic_id"`
}

func (h Handler) AssignTaskToTopic(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	var input assignTaskRequest
	if err := c.ShouldBind(&input); err != nil {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	err := h.service.AssignTaskToTopic(ctx, input.TaskId, input.TopicId)
	if err != nil {
		log.Println(err)
		errorText(c.Writer, "Assigning task to topic error", http.StatusBadRequest)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
}

type assignTopicRequest struct {
	StudentId int `json:"student_id"`
	TopicId   int `json:"topic_id"`
}

func (h Handler) AssignTopicToStudent(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	var input assignTopicRequest
	if err := c.ShouldBind(&input); err != nil {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	err := h.service.AssignTopicToStudent(ctx, input.StudentId, input.TopicId)
	if err != nil {
		log.Println(err)
		errorText(c.Writer, "Assigning topic to topic error", http.StatusBadRequest)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
}
