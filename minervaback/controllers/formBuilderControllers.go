package controllers

import (
	"MinervaServer/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// @Summary List all forms
// @Description List all forms
// @Tags formBuilder
// @Produce json
// @Success 200 {object} []db.Form
// @Failure 500 {object} DefaultResponse
// @Router /formBuilder/list [get]
func FormBuilderListHandler(c *gin.Context) {
	forms, err := db.LoadAllForms()

	if err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to load forms.", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, forms)
}

// @Summary Create new form
// @Description Create new form
// @Tags formBuilder
// @Accept json
// @Produce json
// @Param input body FormBuilderNewRequest true "New form input"
// @Success 200 {object} FormBuilderNewResponse
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /formBuilder/new [post]
func FormBuilderNewHandler(c *gin.Context) {
	var err error

	form := &db.Form{}
	params := FormBuilderNewRequest{}

	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid input.", Error: err.Error()})
		return
	}

	userIdHex := c.GetString("userId")

	if form.OwnerID, err = primitive.ObjectIDFromHex(userIdHex); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Invalid user ID.", Error: err.Error()})
		return
	}

	form.Title = params.Title

	if err := form.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to save form.", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, FormBuilderNewResponse{ID: form.ID.Hex()})
}
