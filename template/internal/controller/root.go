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
	pl := jwt.GetPayload(c)
	name := pl.UserName

	c.HTML(200, "index.html", gin.H{
		"user_name": name,
	})
}