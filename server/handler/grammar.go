package handler

import (
	"DB_course_paper/server/entity"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type createGrammarTask struct {
	Sentence    string `json:"sentence"`
	RightAnswer string `json:"right_answer"`
	Hint        string `json:"hint"`
	Explanation string `json:"explanation"`
}

func (h Handler) CreateGrammarTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	var input createGrammarTask
	if err := c.ShouldBind(&input); err != nil {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	taskId, err := h.service.CreateGrammarTask(ctx, entity.GrammarTask{Sentence: input.Sentence, RightAnswer: input.RightAnswer, Hint: input.Hint, Explanation: input.Explanation})
	if err != nil {
		log.Println(err)
		errorText(c.Writer, "error of creating grammar task", http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(taskId)
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
