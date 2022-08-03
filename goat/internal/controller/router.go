package controller

import (
    "github.com/gin-gonic/gin"

    "goat/internal/core/jwt"
)


func SetRouter(r *gin.Engine) {
    uc := newUserController()
    rc := newRootController()

    r.GET("/signup", uc.signupPage)
    r.GET("/login", uc.loginPage)
    r.GET("/logout", uc.logout)

    r.POST("/signup", uc.signup)
    r.POST("/login", uc.login)

    a := r.Group("/", jwt.JwtAuthMiddleware())
    {
        a.GET("/", rc.indexPage)
    }

    api := r.Group("/api", jwt.JwtAuthMiddleware())
    {
        api.GET("/profile", uc.getProfile)
    }
} 