package main

import (
	"MinervaServer/controllers"
	"MinervaServer/docs"
	"MinervaServer/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// gin initialisation
	r := gin.Default()

	// CORS settings
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Authorization"},
		AllowHeaders:     []string{"Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	/***
	 * Swagger settings
	 ***/
	docs.SwaggerInfo.Title = "MINERVA API"
	docs.SwaggerInfo.Description = "This is a REST API for MINERVA"
	docs.SwaggerInfo.BasePath = "/"

	swaggerUrl := ginSwagger.URL("./doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))

	/***
	 * Auth group
	 ***/
	authGroup := r.Group("/auth")

	authGroup.POST("/signIn", controllers.AuthSignInHandler)
	authGroup.POST("/signUp", controllers.AuthSignUpHandler)
	authGroup.GET("/me", controllers.AuthMeHandler)

	/***
	 * Auth manage group
	 ***/

	authManageGroup := authGroup.Group("/manage")
	authManageGroup.Use(middleware.TokenAuthMiddleware)
	authManageGroup.Use(middleware.AdminAuthMiddleware)

	authManageGroup.GET("/users", controllers.AuthManageUsersHandler)
	authManageGroup.POST("/edit", controllers.AuthManageEditHandler)
	authManageGroup.POST("/updatePassword", controllers.AuthManageUpdatePasswordHandler)
	authManageGroup.POST("/add", controllers.AuthManageAddHandler)
	authManageGroup.POST("/delete", controllers.AuthManageDeleteHandler)

	/***
	 * Form builder group
	 ***/

	formBuilderGroup := r.Group("/formBuilder")
	formBuilderGroup.Use(middleware.TokenAuthMiddleware)
	formBuilderGroup.Use(middleware.AdminAuthMiddleware)

	formBuilderGroup.GET("/list", controllers.FormBuilderListHandler)
	formBuilderGroup.POST("/new", controllers.FormBuilderNewHandler)

	/***
	 * Form builder edit group
	 ***/

	formBuilderEditGroup := formBuilderGroup.Group("/edit/:formID")
	formBuilderEditGroup.Use(middleware.FormOwnerOnlyMiddleware)

	formBuilderEditGroup.GET("/list", controllers.FormBuilderEditListHandler)
	formBuilderEditGroup.POST("/add", controllers.FormBuilderEditAddHandler)
	formBuilderEditGroup.POST("/get", controllers.FormBuilderEditGetHandler)
	formBuilderEditGroup.POST("/set", controllers.FormBuilderEditSetHandler)
	formBuilderEditGroup.POST("/moveTo", controllers.FormBuilderEditMoveToHandler)
	formBuilderEditGroup.POST("/setTitle", controllers.FormBuilderEditSetTitleHandler)
	formBuilderEditGroup.POST("/delete", controllers.FormBuilderEditDeleteHandler)

	r.Run(":8080")
}
