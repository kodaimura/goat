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

	auth := r.Group("", middleware.JwtAuth())
	{
		auth.GET("/", ic.IndexPage)
	}
}


/*
 Routing for "/api"
*/
func SetApiRouter(r *gin.RouterGroup) {
	r.Use(middleware.ApiResponse())

	ac := controller.NewAccountController()

	r.POST("/signup", ac.ApiSignup)
	r.POST("/login", ac.ApiLogin)

	auth := r.Group("", middleware.ApiJwtAuth())
	{
		auth.GET("/accounts/me", ac.ApiGetOne)
		auth.PUT("/accounts/me/name", ac.ApiPutName)
		auth.PUT("/accounts/me/password", ac.ApiPutPassword)
		auth.DELETE("/accounts/me", ac.ApiDelete)
	}
}