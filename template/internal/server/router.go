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
	ic := controller.NewIndexController()
	ac := controller.NewAccountController()

	r.GET("/signup", ac.SignupPage)
	r.GET("/login", ac.LoginPage)
	r.GET("/logout", ac.Logout)

	auth := r.Group("", middleware.JwtAuthMiddleware())
	{
		auth.GET("/", ic.IndexPage)
	}
}


/*
 Routing for "/api"
*/
func SetApiRouter(r *gin.RouterGroup) {
	ac := controller.NewAccountController()

	r.POST("/signup", ac.ApiSignup)
	r.POST("/login", ac.ApiLogin)

	auth := r.Group("", middleware.JwtAuthApiMiddleware())
	{
		auth.GET("/account", ac.ApiGetProfile)
		auth.PUT("/account/name", ac.ApiPutName)
		auth.PUT("/account/password", ac.ApiPutPassword)
		auth.DELETE("/account", ac.ApiDeleteAccount)
	}
}