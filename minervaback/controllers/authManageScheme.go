package controllers

type AuthManageDeleteRequestSchema struct {
	Login string `json:"login" binding:"required"`
}

type AuthManageUpdatePasswordRequestSchema struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
