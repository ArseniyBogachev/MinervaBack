package controllers

type FormBuilderNewRequest struct {
	Title string `json:"title" binding:"required"`
}

type FormBuilderNewResponse struct {
	ID string `json:"id"`
}
