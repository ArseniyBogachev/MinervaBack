package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	UserId   primitive.ObjectID `json:"user_id" example:"user"`
	IssuedAt int64              `json:"issued_at" example:"1708540714"`

	jwt.RegisteredClaims
}

func ExtractToken(c *gin.Context) string {
	cookie, err := c.Cookie("token")

	if err == nil {
		return cookie
	}

	token := c.Query("token")

	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func CreateToken(userId primitive.ObjectID, expiresAt time.Time) (string, error) {
	claims := &Claims{
		UserId:   userId,
		IssuedAt: time.Now().Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func VerifyToken(t string) (*jwt.Token, error) {
	jwtToken, err := jwt.ParseWithClaims(t, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, errors.New("invalid token")
	}

	return jwtToken, nil
}

func ExtractTokenClaims(jwt *jwt.Token) (*Claims, error) {
	claims, ok := jwt.Claims.(*Claims)

	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
