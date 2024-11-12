package handler

import (
	"DB_course_paper/server/entity"
	"DB_course_paper/server/errorsx"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type signUpRequest struct {
	Role     string
	Name     string
	Email    string
	Password string
	Phone    string
}

func (h Handler) signUp(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	var user signUpRequest
	if err := c.ShouldBind(&user); err != nil {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	var (
		userId int
		err    error
	)

	if user.Role == "student" {
		userId, err = h.service.CreateStudent(ctx, entity.Student{Name: user.Name, Email: user.Email, Password: user.Password, Phone: user.Phone})
	} else {
		userId, err = h.service.CreateTeacher(ctx, entity.Teacher{Name: user.Name, Email: user.Email, Password: user.Password, Phone: user.Phone})
	}

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			errorText(c.Writer, "time limit exceeded", http.StatusInternalServerError)
			return
		}
		log.Println("SignUp handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(userId)
	if err != nil {
		log.Println("SignUp handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusCreated)
	_, err = c.Writer.Write(j)
	if err != nil {
		log.Println("CreateCart handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

type userSignIn struct {
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (h Handler) signIn(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	var input userSignIn
	if err := c.ShouldBind(&input); err != nil {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	var (
		userId int
		err    error
	)

	if input.Role == "student" {
		userId, err = h.service.GetStudentByCreds(ctx, input.Phone, input.Password)
	} else {
		userId, err = h.service.GetTeacherByCreds(ctx, input.Phone, input.Password)
	}

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			errorText(c.Writer, "time limit exceeded", http.StatusInternalServerError)
			return
		}

		switch err {
		case errorsx.UserDoesNotExistError, errorsx.PasswordMismatchError:
			errorText(c.Writer, "User not found", http.StatusInternalServerError)
			return
		default:
			log.Println("SignIn handler error:", err)
			errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}

	tokens, err := h.createTokens(userId)
	if err != nil {
		log.Println("SignIn handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	_, err = h.service.CreateTokens(ctx, tokens)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			errorText(c.Writer, "time limit exceeded", http.StatusInternalServerError)
			return
		}
		log.Println("SignIn handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(tokens)
	if err != nil {
		log.Println("SignUp handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")

	c.Writer.Header().Set(authHeader, tokens.AccessToken)
	c.SetCookie(refreshTokenCookie, tokens.RefreshToken, int(time.Now().Add(refreshCookieExpiredInHours*time.Hour).Unix()), "/", "/", false, true)

	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write(j)
	if err != nil {
		log.Println("CreateCart handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func (h Handler) logOut(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()
	accessToken := c.GetHeader(authHeader)
	if accessToken == "" {
		log.Println("LogOut handler error: Authentication header not found")
		errorText(c.Writer, "Authentication error", http.StatusUnauthorized)
		return
	}

	userId, err := h.GetUserIdFromJwt(accessToken)
	if err != nil {
		log.Println("LogOut handler error:", err)
		errorText(c.Writer, "Authentication error", http.StatusUnauthorized)
		return
	}

	err = h.service.LogOutUser(ctx, userId)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			errorText(c.Writer, "time limit exceeded", http.StatusInternalServerError)
			return
		}
		log.Println("LogOut handler error:", err)
		errorText(c.Writer, "LogOut handler error", http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
}
