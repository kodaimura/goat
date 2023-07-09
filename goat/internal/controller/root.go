package controller

import (
	"github.com/gin-gonic/gin"
	
	"goat/internal/core/jwt"
)


type rootController struct {}


func NewRootController() *rootController {
	return &rootController{}
}


//GET /
func (ctr *rootController) IndexPage(c *gin.Context) {
	username := jwt.GetUserName(c)

	c.HTML(200, "index.html", gin.H{
		"username": username,
	})
}