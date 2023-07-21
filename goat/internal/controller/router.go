package controller

import (
	"github.com/gin-gonic/gin"

	"goat/internal/core/jwt"
)


func SetRouter(r *gin.Engine) {
	uc := NewUserController()

	//render HTML or redirect
	r.GET("/signup", uc.SignupPage)
	r.GET("/login", uc.LoginPage)
	r.GET("/logout", uc.Logout)
	r.POST("/signup", uc.Signup)
	r.POST("/login", uc.Login)

	//render HTML or redirect (Authorized request)
	a := r.Group("/", jwt.JwtAuthMiddleware())
	{
		rc := NewRootController()
		
		a.GET("/", rc.IndexPage)

		a.GET("/api/account/profile", uc.GetAccountProfile)
		a.PUT("/api/account/username", uc.ChangeUsername)
		a.PUT("/api/account/password", uc.ChangePassword)
		a.DELETE("/api/account", uc.DeleteAccount)
	}
}