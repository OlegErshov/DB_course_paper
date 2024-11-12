package handler

import (
	"DB_course_paper/server/entity"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	requestExpiredInSeconds     = 5
	authHeader                  = "Authorization"
	refreshTokenCookie          = "refresh_token"
	refreshCookieExpiredInHours = 24 * 30
	accessExpiredInMinutes      = 1
	refreshExpiredInMinutes     = 5
)

func (h Handler) authMiddleware(jwtSecret []byte) gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
		defer cancel()

		accessToken := c.GetHeader(authHeader)
		if accessToken == "" {
			log.Println("auth middleware: access token not found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		claims := &Claims{}

		log.Println("jwt secret bytes:", jwtSecret)

		parsedToken, err := jwt.ParseWithClaims(accessToken, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			if !errors.Is(err, jwt.ErrTokenExpired) {
				log.Println("auth middleware: token  invalid", err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
				return
			}
			log.Println("auth middleware: access token expired")
		}

		id, err := h.GetUserIdFromJwt(accessToken)
		if err != nil {
			log.Println("auth middleware: invalid access token", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		tokens, err := h.service.GetTokens(ctx, id)
		if err != nil {
			log.Println("auth middleware:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		if !parsedToken.Valid {
			log.Println("auth middleware: token invalid", err)
			h.processInvalidAccessToken(c, ctx, tokens, id)
			c.Next()
			return
		}

		if accessToken != tokens.AccessToken {
			log.Println("auth middleware: access token doesnt match", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		}
		c.Next()
		return
	}
}

func (h Handler) processInvalidAccessToken(c *gin.Context, ctx context.Context, tokens entity.Token, userId int) {
	refresh, err := c.Cookie(refreshTokenCookie)
	if err != nil {
		log.Println("auth middleware: refresh token cookie invalid", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	if refresh != tokens.RefreshToken {
		log.Println("auth middleware: refresh token doesnt match", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	tokens, err = h.createTokens(userId)
	if err != nil {
		log.Println("auth middleware: cannot create token", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	_, err = h.service.UpdateTokens(ctx, tokens)
	if err != nil {
		log.Println("auth middleware: cannot update token in postgres", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	c.SetCookie(refreshTokenCookie, tokens.RefreshToken, int(time.Now().Add(refreshCookieExpiredInHours*time.Hour).Unix()), "/", "/", false, true)
	c.Header(authHeader, tokens.AccessToken)
}

type Claims struct {
	Id  string `json:"id"`
	Exp int64  `json:"exp"`
	jwt.RegisteredClaims
}

func (h Handler) GetUserIdFromJwt(token string) (int, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return h.jwtSecret, nil
	})
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(claims.Id)
}

func (h Handler) createTokens(id int) (entity.Token, error) {
	accessToken, err := h.createToken(id, accessExpiredInMinutes)
	if err != nil {
		return entity.Token{}, err
	}

	refreshToken, err := h.createToken(id, refreshExpiredInMinutes)
	if err != nil {
		return entity.Token{}, err
	}

	return entity.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (h Handler) createToken(id, expired int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  strconv.Itoa(id),
			"exp": jwt.NewNumericDate(time.Now().Add(time.Duration(expired) * time.Minute)),
		})

	log.Println("jwt secret bytes", h.jwtSecret)
	tokenString, err := token.SignedString(h.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
