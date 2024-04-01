package controllers

import (
	"MinervaServer/db"
	"MinervaServer/middleware"
	"MinervaServer/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

// @Summary Sign up
// @Description Sign up
// @Tags auth
// @ID auth-sign-up
// @Accept json
// @Produce json
// @Param input body AuthSignInRequestSchema true "Sign up input"
// @Success 200 {object} AuthSignInResponseSchema
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /auth/signIn [post]
func AuthSignInHandler(c *gin.Context) {
	params := AuthSignInRequestSchema{}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid input", Error: err.Error()})
		return
	}

	user := &db.User{}

	err := user.LoadByLogin(params.Login)

	if err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "User not found"})
		return
	}

	if user.Password != params.Password {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid password"})
		return
	}

	token, err := utils.CreateToken(user.ID, time.Now().Add(time.Hour))

	if err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to create token", Error: err.Error()})
		return
	}

	c.SetCookie("token", token, int(time.Now().Add(time.Hour).Unix()), "/", c.Request.Host, false, false)
	c.JSON(http.StatusOK, AuthSignInResponseSchema{Message: "OK", Token: token})
}

// @Summary Sign in
// @Description Sign in
// @Tags auth
// @ID auth-sign-in
// @Accept json
// @Produce json
// @Param input body AuthSignInRequestSchema true "Sign in input"
// @Success 200 {object} AuthSignInResponseSchema
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /auth/signUp [post]
func AuthSignUpHandler(c *gin.Context) {
	params := AuthSignInRequestSchema{}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid input", Error: err.Error()})
		return
	}

	user := &db.User{}

	err := user.LoadByLogin(params.Login)

	if err == nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "User exists"})
		return
	}

	user.Login = params.Login
	user.Password = params.Password

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to create user", Error: err.Error()})
		return
	}

	token, err := utils.CreateToken(user.ID, time.Now().Add(time.Hour))

	if err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to create token", Error: err.Error()})
		return
	}

	c.SetCookie("token", token, int(time.Now().Add(time.Hour).Unix()), "/", c.Request.Host, false, false)
	c.JSON(http.StatusOK, AuthSignInResponseSchema{Message: "OK", Token: token})
}

// @Summary Me
// @Description Me
// @Tags auth
// @ID auth-me
// @Accept json
// @Produce json
// @Success 200 {object} AuthMeResponseSchema
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /auth/me [get]
func AuthMeHandler(c *gin.Context) {
	middleware.TokenAuthMiddleware(c)

	userIdHex := c.GetString("userId")
	userId, err := primitive.ObjectIDFromHex(userIdHex)

	if err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to convert user ID", Error: err.Error()})
		return
	}

	user := &db.User{}

	if err = user.LoadById(userId); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "User not found", Error: err.Error()})
		return
	}

	response := AuthMeResponseSchema{}
	response.Login = user.Login
	response.IsAdmin = user.IsAdmin
	response.FirstName = user.FirstName
	response.LastName = user.LastName
	response.Email = user.Email
	response.Phone = user.Phone

	tokenString, err := utils.CreateToken(user.ID, time.Now().Add(time.Hour))

	if err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to create token", Error: err.Error()})
		return
	}

	c.SetCookie("token", tokenString, int(time.Now().Add(time.Hour).Unix()), "/", c.Request.Host, false, false)
	c.JSON(http.StatusOK, response)
}
