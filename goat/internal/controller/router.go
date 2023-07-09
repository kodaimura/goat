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
	}

	//response JSON
	api := r.Group("/api")
	{
		uac := NewUserApiController()

		api.POST("/signup", uac.Signup)
		api.POST("/login", uac.Login)
		api.GET("/logout", uac.Logout)


		//response JSON (Authorized request)
		a := api.Group("/", jwt.JwtAuthApiMiddleware())
		{
			a.GET("/profile", uac.GetProfile)
			a.PUT("/username", uac.ChangeUsername)
			a.PUT("/password", uac.ChangePassword)
			a.DELETE("/account", uac.DeleteUser)
		}
	}
}