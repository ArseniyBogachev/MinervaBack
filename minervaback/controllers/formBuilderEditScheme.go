package controllers

type FormBuilderEditAddRequest struct {
	Order int         `json:"order" example:"1"`
	Data  interface{} `json:"data" example:"{}"`
}

type FormBuilderEditGetRequest struct {
	Order int `json:"order" example:"1"`
}

type FormBuilderEditSetRequest struct {
	Order int         `json:"order" example:"1"`
	Data  interface{} `json:"data" example:"{}"`
}

type FormBuilderEditSetTitleRequest struct {
	Title string `json:"title" example:"Form title"`
}

type FormBuilderEditMoveToRequest struct {
	OldOrder *int `json:"old_order" example:"1" binding:"required"`
	NewOrder *int `json:"new_order" example:"2" binding:"required"`
}
