package controllers

type AuthSignInRequestSchema struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthSignInResponseSchema struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type AuthSignUpRequestSchema struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthSignUpResponseSchema struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type AuthMeResponseSchema struct {
	Login     string `json:"login"`
	IsAdmin   bool   `json:"is_admin"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
