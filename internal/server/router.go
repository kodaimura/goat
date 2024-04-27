package server

import (
	"github.com/gin-gonic/gin"

	"goat/internal/middleware"
	"goat/internal/controller"
)

/*
 Routing for "/" 
*/
func SetWebRouter(r *gin.RouterGroup) {
	rc := controller.NewRootController()
	uc := controller.NewUserController()

	r.GET("/signup", uc.SignupPage)
	r.GET("/login", uc.LoginPage)
	r.POST("/signup", uc.Signup)
	r.POST("/login", uc.Login)
	r.GET("/logout", uc.Logout)

	auth := r.Group("", middleware.JwtAuthMiddleware())
	{
		auth.GET("/", rc.IndexPage)
	}
}

/*
 Routing for "/api"
*/
func SetApiRouter(r *gin.RouterGroup) {
	uc := controller.NewUserController()

	auth := r.Group("", middleware.JwtAuthApiMiddleware())
	{
		auth.GET("/account/profile", uc.ApiGetProfile)
		auth.PUT("/account/name", uc.ApiPutName)
		auth.PUT("/account/password", uc.ApiPutPassword)
		auth.DELETE("/account", uc.ApiDeleteAccount)
	}
}