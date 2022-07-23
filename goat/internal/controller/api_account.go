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


func newAccountController() accountController {
	ur := repository.NewUserRepository()
	return accountController{ur}
}


//PUT[POST] /api/account/password
func (ac accountController) changePassword(c *gin.Context) {
	uid, _ := jwt.GetUId(c)

	var body map[string]interface{}
	c.BindJSON(&body)
	newPw := body["password"].(string)
	hashed, err := bcrypt.GenerateFromPassword([]byte(newPw), bcrypt.DefaultCost)

	if err != nil || ac.ur.UpdatePasswordByUId(uid, string(hashed)) != nil {
		logger.LogError(err.Error())
		c.JSON(500, gin.H{"error": http.StatusText(500)})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}


//PUT[POST] /api/account/username
func (ac accountController) changeUsername(c *gin.Context) {
	uid, _ := jwt.GetUId(c)
	var body map[string]interface{}
	c.BindJSON(&body)
	newUn := body["username"].(string)

	if err := ac.ur.UpdateUsernameByUId(uid, newUn); err != nil {
		logger.LogError(err.Error())
		c.JSON(500, gin.H{"error": http.StatusText(500)})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}


//GET /api/account/profile
func (ac accountController) getProfile(c *gin.Context) {
	uid, _ := jwt.GetUId(c)
	user, err := ac.ur.SelectByUId(uid)

	if err != nil {
		logger.LogError(err.Error())
		c.JSON(500, gin.H{"error": http.StatusText(500)})
		c.Abort()
		return
	}

	c.JSON(200, user)
}


//DELETE /api/account
func (ac accountController) delete(c *gin.Context) {
	uid, _ := jwt.GetUId(c)

	if err := ac.ur.DeleteByUId(uid); err != nil {
		logger.LogError(err.Error())
		c.JSON(500, gin.H{"error": http.StatusText(500)})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}