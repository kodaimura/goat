package controller

import (
	"github.com/gin-gonic/gin"
	
	"goat/internal/core/jwt"
)


type RootController struct {}


func NewRootController() *RootController {
	return &RootController{}
}


//GET /
func (rc *RootController) IndexPage(c *gin.Context) {
	username := jwt.GetUsername(c)

	c.HTML(200, "index.html", gin.H{
		"username": username,
	})
}