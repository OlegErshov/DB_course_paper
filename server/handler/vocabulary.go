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

type OptionsInput struct {
	Sentence      string `json:"sentence"`
	AnswerOptions string `json:"answer_options"`
	RightAnswer   string `json:"right_answer"`
	Explanation   string `json:"explanation"`
}

func (h Handler) CreateVocabularyOptionsTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	var input OptionsInput
	if err := c.ShouldBind(&input); err != nil {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	taskId, err := h.service.CreateOptionsTask(ctx, entity.VocabularyOptionsTask{Sentence: input.Sentence, AnswerOptions: input.AnswerOptions, RightAnswer: input.RightAnswer, Explanation: input.Explanation})
	if err != nil {
		log.Println(err)
		errorText(c.Writer, "Error in Creating Options Task", http.StatusBadRequest)
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

type SentenceInput struct {
	FirstPart   string `json:"first_part"`
	SecondPart  string `json:"second_part"`
	Explanation string `json:"explanation"`
}

func (h Handler) CreateVocabularySentenceTask(c *gin.Context) {
	_, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()
}

type WordInput struct {
	Sentence    string `json:"sentence"`
	Answer      string `json:"answer"`
	Explanation string `json:"explanation"`
}

func (h Handler) CreateVocabularyWordTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	var input WordInput
	if err := c.ShouldBind(&input); err != nil {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	taskId, err := h.service.CreateWordTask(ctx, entity.VocabularyWordTask{Sentence: input.Sentence, Answer: input.Answer, Explanation: input.Explanation})
	if err != nil {
		log.Println(err)
		errorText(c.Writer, "Error in Creating Options Task", http.StatusBadRequest)
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
