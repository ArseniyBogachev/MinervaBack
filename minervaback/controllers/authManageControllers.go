package controllers

import (
	"MinervaServer/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Manage users
// @Description Manage users
// @Tags auth/manage
// @ID auth-manage-users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []db.User
// @Failure 500 {object} DefaultResponse
// @Router /auth/manage/users [get]
func AuthManageUsersHandler(c *gin.Context) {
	users, err := db.LoadUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Edit user
// @Description Edit user
// @Tags auth/manage
// @ID auth-manage-edit
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body db.User true "User input"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /auth/manage/edit [post]
func AuthManageEditHandler(c *gin.Context) {
	user := db.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	fmt.Printf("User: %v\n", user)

	oldUser := &db.User{}

	if err := oldUser.LoadById(user.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User not found", "error": err.Error()})
		return
	}

	user.Password = oldUser.Password

	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// @Summary Update password
// @Description Update password
// @Tags auth/manage
// @ID auth-manage-update-password
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body AuthManageUpdatePasswordRequestSchema true "Update password input"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /auth/manage/updatePassword [post]
func AuthManageUpdatePasswordHandler(c *gin.Context) {
	params := AuthManageUpdatePasswordRequestSchema{}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid input", Error: err.Error()})
		return
	}

	user := &db.User{}

	if err := user.LoadByLogin(params.Login); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "User not found"})
		return
	}

	user.Password = params.Password

	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to update password", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DefaultResponse{Message: "OK"})
}

// @Summary Add user
// @Description Add user
// @Tags auth/manage
// @ID auth-manage-add
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body db.User true "User input"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /auth/manage/add [post]
func AuthManageAddHandler(c *gin.Context) {
	user := db.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// @Summary Delete user
// @Description Delete user
// @Tags auth/manage
// @ID auth-manage-delete
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body AuthManageDeleteRequestSchema true "Delete user input"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /auth/manage/delete [post]
func AuthManageDeleteHandler(c *gin.Context) {
	params := AuthManageDeleteRequestSchema{}

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

	if err = user.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to delete user", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DefaultResponse{Message: "OK"})
}
