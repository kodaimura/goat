package controller

import (
    "github.com/gin-gonic/gin"
    
    "goat/internal/core/jwt"
    "goat/internal/constant"
)


func setRootRoute(r *gin.Engine) {
    auth := r.Group("/", jwt.JwtAuthMiddleware())

    rc := newRootController()
    auth.GET("/", rc.indexPage)
}


type rootController struct {}


func newRootController() *rootController {
    return &rootController{}
}


//GET /
func (ctr *rootController) indexPage(c *gin.Context) {
    username, err := jwt.GetUsername(c)

    if err != nil {
        c.Redirect(303, "/login")
        return
    }

    c.HTML(200, "index.html", gin.H{
        "commons": constant.Commons,
        "username": username,
    })
}