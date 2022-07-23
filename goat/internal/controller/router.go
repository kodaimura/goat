package controller

import (
    "github.com/gin-gonic/gin"
)


func SetRouter(r *gin.Engine) {
    setLoginRoute(r)
    setSignupRoute(r)
    setRootRoute(r)

    setApiAccountRoute(r)
} 