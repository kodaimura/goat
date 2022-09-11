package controller

import (
    "github.com/gin-gonic/gin"

    "goat/internal/core/jwt"
)


func SetRouter(r *gin.Engine) {
    uc := newUserController()

    //render HTML or redirect
    r.GET("/signup", uc.signupPage)
    r.GET("/login", uc.loginPage)
    r.GET("/logout", uc.logout)
    r.POST("/signup", uc.signup)
    r.POST("/login", uc.login)

    //render HTML or redirect (Authorized request)
    a := r.Group("/", jwt.JwtAuthMiddleware())
    {
        rc := newRootController()
        
        a.GET("/", rc.indexPage)
    }

    //response JSON
    api := r.Group("/api")
    {
        uac := newUserApiController()

        api.POST("/signup", uac.signup)
        api.POST("/login", uac.login)
        api.GET("/logout", uac.logout)


        //response JSON (Authorized request)
        a := api.Group("/", jwt.JwtAuthApiMiddleware())
        {
            a.GET("/profile", uac.getProfile)
            a.PUT("/username", uac.changeUsername)
            a.POST("/username", uac.changeUsername)
            a.PUT("/password", uac.changePassword)
            a.POST("/password", uac.changePassword)
            a.DELETE("/account", uac.deleteUser)
        }
    }
}