package controller

import (
	"github.com/gin-gonic/gin"

	"goat-base/internal/core/jwt"
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
	a := r.Group("/", jwt.JwtAuthMiddleware())
	{
		a.GET("/", rc.IndexPage)
	}

	//response JSON (Authorized request)
	api := r.Group("/api", jwt.JwtAuthApiMiddleware())
	{
		api.GET("/account/profile", uc.GetProfile)
		api.PUT("/account/name", uc.UpdateName)
		api.PUT("/account/password", uc.UpdatePassword)
		api.DELETE("/account", uc.DeleteAccount)
	}
}