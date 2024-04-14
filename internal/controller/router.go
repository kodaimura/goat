package controller

import (
	"github.com/gin-gonic/gin"

	"goat/internal/middleware"
)


func SetRouter(r *gin.Engine) {
	rc := NewRootController()
	uc := NewUserController()

	//render HTML or redirect
	r.GET("/signup", uc.SignupPage)
	r.POST("/signup", uc.Signup)
	r.GET("/login", uc.LoginPage)
	r.POST("/login", uc.Login)
	r.GET("/logout", uc.Logout)

	//render HTML or redirect (Authorized request)
	a := r.Group("/", middleware.JwtAuthMiddleware())
	{
		a.GET("/", rc.IndexPage)
	}

	//response JSON (Authorized request)
	api := r.Group("/api", middleware.JwtAuthApiMiddleware())
	{
		api.GET("/account/profile", uc.ApiGetProfile)
		api.PUT("/account/name", uc.ApiPutName)
		api.PUT("/account/password", uc.ApiPutPassword)
		api.DELETE("/account", uc.ApiDeleteAccount)
	}
}