package controllers

import (
	"MinervaServer/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// @Summary List all form blocks
// @Description List all form blocks
// @Tags formBuilder/edit
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Success 200 {object} []db.FormBlock
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /formBuilder/edit/{formId}/list [get]
func FormBuilderEditListHandler(c *gin.Context) {
	formIdHex := c.GetString("formId")
	formId, err := primitive.ObjectIDFromHex(formIdHex)

	if err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid form ID.", Error: err.Error()})
		return
	}

	form := &db.Form{}

	if err := form.LoadById(formId); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to load form.", Error: err.Error()})
		return
	}

	formBlocks, err := db.LoadFormBlocks(formId)

	c.JSON(http.StatusOK, formBlocks)
}

// @Summary Add form block
// @Description Add form block
// @Tags formBuilder/edit
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param input body FormBuilderEditAddRequest true "Add form block input"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /formBuilder/edit/{formId}/add [post]
func FormBuilderEditAddHandler(c *gin.Context) {
	var err error
	var formId primitive.ObjectID

	if formId, err = primitive.ObjectIDFromHex(c.GetString("formId")); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid form ID.", Error: err.Error()})
		return
	}

	params := FormBuilderEditAddRequest{}
	params.Order = -1

	if err = c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid input.", Error: err.Error()})
		return
	}

	formBlock := &db.FormBlock{}
	formBlock.Data = params.Data
	formBlock.FormID = formId

	if err = formBlock.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to save form block.", Error: err.Error()})
		return
	}

	if params.Order >= 0 {
		if err = formBlock.MoveTo(params.Order); err != nil {
			c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to move form block.", Error: err.Error()})
		}
	}

	c.JSON(http.StatusOK, DefaultResponse{Message: "OK"})
}

// @Summary Get form block
// @Description Get form block
// @Tags formBuilder/edit
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param input body FormBuilderEditGetRequest true "Get form block input"
// @Success 200 {object} db.FormBlock
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /formBuilder/edit/{formBlockId}/get [post]
func FormBuilderEditGetHandler(c *gin.Context) {
	var err error
	var formBlockId primitive.ObjectID

	if formBlockId, err = primitive.ObjectIDFromHex(c.GetString("formId")); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid form block ID.", Error: err.Error()})
		return
	}

	params := FormBuilderEditGetRequest{}

	if err = c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid input.", Error: err.Error()})
		return
	}

	formBlock := &db.FormBlock{}
	formBlock.FormID = formBlockId

	if err = formBlock.Load(params.Order); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to load form block.", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, formBlock)
}

// @Summary Set form block
// @Description Set form block
// @Tags formBuilder/edit
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param input body FormBuilderEditSetRequest true "Set form block input"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /formBuilder/edit/{formBlockId}/set [post]
func FormBuilderEditSetHandler(c *gin.Context) {
	var err error
	var formBlockId primitive.ObjectID

	if formBlockId, err = primitive.ObjectIDFromHex(c.GetString("formId")); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid form block ID.", Error: err.Error()})
		return
	}

	params := FormBuilderEditSetRequest{}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid input.", Error: err.Error()})
		return
	}

	formBlock := &db.FormBlock{}
	formBlock.FormID = formBlockId

	if err = formBlock.Load(params.Order); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to load form block.", Error: err.Error()})
		return
	}

	formBlock.Data = params.Data

	if err = formBlock.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to save form block.", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DefaultResponse{Message: "OK"})
}

// @Summary Move form block
// @Description Move form block
// @Tags formBuilder/edit
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param input body FormBuilderEditSetTitleRequest true "Move form block input"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /formBuilder/edit/{formId}/setTitle [post]
func FormBuilderEditSetTitleHandler(c *gin.Context) {
	var err error
	var formBlockId primitive.ObjectID

	if formBlockId, err = primitive.ObjectIDFromHex(c.GetString("formId")); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid form block ID.", Error: err.Error()})
		return
	}

	params := FormBuilderEditSetTitleRequest{}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid input.", Error: err.Error()})
		return
	}

	form := &db.Form{}

	if err = form.LoadById(formBlockId); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to load form.", Error: err.Error()})
		return
	}

	form.Title = params.Title

	if err = form.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to save form block.", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DefaultResponse{Message: "OK"})
}

// @Summary Move form block
// @Description Move form block
// @Tags formBuilder/edit
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Param input body FormBuilderEditMoveToRequest true "Move form block input"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /formBuilder/edit/{formId}/moveTo [post]
func FormBuilderEditMoveToHandler(c *gin.Context) {
	var err error
	var formBlockId primitive.ObjectID

	if formBlockId, err = primitive.ObjectIDFromHex(c.GetString("formId")); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid form block ID.", Error: err.Error()})
		return
	}

	params := FormBuilderEditMoveToRequest{}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid input.", Error: err.Error()})
		return
	}

	formBlock := &db.FormBlock{}
	formBlock.FormID = formBlockId

	if err = formBlock.Load(*params.OldOrder); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to load form block.", Error: err.Error()})
		return
	}

	if err = formBlock.MoveTo(*params.NewOrder); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to move form block.", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DefaultResponse{Message: "OK"})
}

// @Summary Delete form
// @Description Delete form
// @Tags formBuilder/edit
// @Accept json
// @Produce json
// @Param formId path string true "Form ID"
// @Success 200 {object} DefaultResponse
// @Failure 400 {object} DefaultResponse
// @Failure 500 {object} DefaultResponse
// @Router /formBuilder/edit/{formId}/delete [post]
func FormBuilderEditDeleteHandler(c *gin.Context) {
	var err error
	var formId primitive.ObjectID

	if formId, err = primitive.ObjectIDFromHex(c.GetString("formId")); err != nil {
		c.JSON(http.StatusBadRequest, DefaultResponse{Message: "Invalid form ID.", Error: err.Error()})
		return
	}

	form := &db.Form{}

	if err = form.LoadById(formId); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to load form.", Error: err.Error()})
		return
	}

	if err = form.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, DefaultResponse{Message: "Failed to delete form.", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DefaultResponse{Message: "OK"})
}
