package controller

import (
	"github.com/gin-gonic/gin"
	
	"goat-base/internal/core/jwt"
)


type RootController struct {}


func NewRootController() *RootController {
	return &RootController{}
}


//GET /
func (rc *RootController) IndexPage(c *gin.Context) {
	name := jwt.GetUserName(c)

	c.HTML(200, "index.html", gin.H{
		"user_name": name,
	})
}