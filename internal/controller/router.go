package controller

import (
	"github.com/gin-gonic/gin"

	"goat/internal/middleware"
)


func SetRouter(r *gin.Engine) {
	rc := NewRootController()
	uc := NewUserController()

	{
		r.GET("/signup", uc.SignupPage)
		r.GET("/login", uc.LoginPage)
		r.POST("/signup", uc.Signup)
		r.POST("/login", uc.Login)
		r.GET("/logout", uc.Logout)

		auth := r.Group("/", middleware.JwtAuthMiddleware())
		{
			auth.GET("/", rc.IndexPage)
		}
	}

	api := r.Group("/api")
	{
		auth := api.Group("/", middleware.JwtAuthApiMiddleware())
		{
			auth.GET("/account/profile", uc.ApiGetProfile)
			auth.PUT("/account/name", uc.ApiPutName)
			auth.PUT("/account/password", uc.ApiPutPassword)
			auth.DELETE("/account", uc.ApiDeleteAccount)
		}
	}
}