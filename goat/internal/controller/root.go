package controller

import (
    "github.com/gin-gonic/gin"
    
    "goat/internal/core/jwt"
    "goat/internal/model/repository"
    "goat/internal/constants"
)


func setRootRoute(r *gin.Engine) {
    auth := r.Group("/", jwt.JwtAuthMiddleware())

    rc := newRootController()
    auth.GET("/", rc.indexPage)
}


type rootController struct {
    ur repository.UserRepository
}


func newRootController() *rootController {
    ur := repository.NewUserRepository()
    return &rootController{ur}
}


//GET /
func (ic *rootController) indexPage(c *gin.Context) {
    username, err := jwt.GetUsername(c)

    if err != nil {
        c.Redirect(303, "/login")
        return
    }

    c.HTML(200, "index.html", gin.H{
        "commons": constants.Commons,
        "username": username,
    })
}