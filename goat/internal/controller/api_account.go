package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"goat/internal/model/repository"
	"goat/internal/core/jwt"
	"goat/internal/core/logger"

)


func setApiAccountRoute(r *gin.Engine) {
    ac := newAccountController()

    r.PUT("/account/password", ac.changePassword)
    r.POST("/account/password", ac.changePassword)
    r.PUT("/account/username", ac.changeUsername)
    r.POST("/account/username", ac.changeUsername)
    r.GET("/account/profile", ac.getProfile)
    r.DELETE("/account", ac.delete)
}


type accountController struct {
	ur repository.UserRepository
}


func newAccountController() *accountController {
	ur := repository.NewUserRepository()
	return &accountController{ur}
}


//PUT[POST] /api/account/password
func (ctr *accountController) changePassword(c *gin.Context) {
	userId, _ := jwt.GetUserId(c)

	var body map[string]interface{}
	c.BindJSON(&body)
	newPw := body["password"].(string)
	hashed, err := bcrypt.GenerateFromPassword([]byte(newPw), bcrypt.DefaultCost)

	if err != nil || ctr.ur.UpdatePasswordByUserId(userId, string(hashed)) != nil {
		logger.LogError(err.Error())
		c.JSON(500, gin.H{"error": http.StatusText(500)})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}


//PUT[POST] /api/account/username
func (ctr *accountController) changeUsername(c *gin.Context) {
	userId, _ := jwt.GetUserId(c)
	var body map[string]interface{}
	c.BindJSON(&body)
	newUn := body["username"].(string)

	if err := ctr.ur.UpdateUsernameByUserId(userId, newUn); err != nil {
		logger.LogError(err.Error())
		c.JSON(500, gin.H{"error": http.StatusText(500)})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}


//GET /api/account/profile
func (ctr *accountController) getProfile(c *gin.Context) {
	userId, _ := jwt.GetUserId(c)
	user, err := ctr.ur.SelectByUserId(userId)

	if err != nil {
		logger.LogError(err.Error())
		c.JSON(500, gin.H{"error": http.StatusText(500)})
		c.Abort()
		return
	}

	c.JSON(200, user)
}


//DELETE /api/account
func (ctr *accountController) delete(c *gin.Context) {
	userId, _ := jwt.GetUserId(c)

	if err := ctr.ur.DeleteByUserId(userId); err != nil {
		logger.LogError(err.Error())
		c.JSON(500, gin.H{"error": http.StatusText(500)})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}