package controller

import (
    "github.com/gin-gonic/gin"

    "goat/internal/core/jwt"
)


func SetRouter(r *gin.Engine) {
    uc := newUserController()
    rc := newRootController()

    //render HTML or redirect
    r.GET("/signup", uc.signupPage)
    r.GET("/login", uc.loginPage)
    r.GET("/logout", uc.logout)
    r.POST("/signup", uc.signup)
    r.POST("/login", uc.login)

    //render HTML or redirect (Authorized request)
    a := r.Group("/", jwt.JwtAuthMiddleware())
    {
        a.GET("/", rc.indexPage)
    }

    //response JSON
    api := r.Group("/api")
    {

        /* --------------------------------------------*/
        // when you use user.api.go(signup,login,logout)
        /* --------------------------------------------*/
        //api.POST("/signup", uc.signup)
        //api.POST("/login", uc.login)
        //api.GET("/logout", uc.logout)


        //response JSON (Authorized request)
        a := api.Group("/", jwt.JwtAuthApiMiddleware())
        {
            a.GET("/profile", uc.getProfile)
            a.PUT("/username", uc.changeUsername)
            a.POST("/username", uc.changeUsername)
            a.PUT("/password", uc.changePassword)
            a.POST("/password", uc.changePassword)
            a.DELETE("/account", uc.deleteUser)
        }
    }
}