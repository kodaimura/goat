package controller

import (
	"github.com/gin-gonic/gin"

	"goat/config"
	"goat/internal/core/jwt"
	"goat/internal/service"
	"goat/internal/model/entity"
)


type UserService interface {
	Signup(username, password string) error
	Login(username, password string) (entity.User, error)
	GenerateJWT(id int) (string, error)
	GetProfile(id int) (entity.User, error)
	ChangeUsername(id int, username string) error
	ChangePassword(id int, password string) error
	DeleteUser(id int) error
}

type userController struct {
	uServ UserService
}


func NewUserController() *userController {
	uServ := service.NewUserService()
	return &userController{uServ}
}


//GET /signup
func (ctr *userController) SignupPage(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}

//GET /login
func (ctr *userController) LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}


//POST /signup
func (ctr *userController) Signup(c *gin.Context) {
	name := c.PostForm("username")
	pass := c.PostForm("password")

	err := ctr.uServ.Signup(name, pass)

	if err != nil {
		if _, ok := err.(*service.SignupConflictError); ok {
			c.HTML(409, "signup.html", gin.H{
				"error": "ユーザ名が既に使われています。",
			})
		} else {
			c.HTML(500, "signup.html", gin.H{
				"error": "登録に失敗しました。",
			})
		}
		c.Abort()
		return
	}

	c.Redirect(303, "/login")
}


//POST /login
func (ctr *userController) Login(c *gin.Context) {
	name := c.PostForm("username")
	pass := c.PostForm("password")

	user, err := ctr.uServ.Login(name, pass)

	if err != nil {
		c.HTML(401, "login.html", gin.H{
			"error": "ユーザ名またはパスワードが異なります。",
		})
		c.Abort()
		return
	}

	jwtStr, err := ctr.uServ.GenerateJWT(user.UserId)

	if err != nil {
		c.HTML(500, "login.html", gin.H{
			"error": "ログインに失敗しました。",
		})
		c.Abort()
		return
	}

	cf := config.GetConfig()
	c.SetCookie(jwt.COOKIE_KEY_JWT, jwtStr, int(jwt.JWT_EXPIRES), "/", cf.AppHost, false, true)
	c.Redirect(303, "/")
}


//GET /logout
func (ctr *userController) Logout(c *gin.Context) {
	cf := config.GetConfig()
	c.SetCookie(jwt.COOKIE_KEY_JWT, "", 0, "/", cf.AppHost, false, true)
	c.Redirect(303, "/login")
}


//GET /api/account/profile
func (ctr *userController) GetAccountProfile(c *gin.Context) {
	user, err := ctr.uServ.GetProfile(jwt.GetUserId(c))

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, user)
}


//PUT /api/account/password
func (ctr *userController) ChangePassword(c *gin.Context) {
	id := jwt.GetUserId(c)

	m := map[string]string{}
	c.BindJSON(&m)
	pass := m["password"]

	if ctr.uServ.ChangePassword(id, pass) != nil {
		c.JSON(500, gin.H{"error": "登録に失敗しました。"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}


//PUT /api/account/username
func (ctr *userController) ChangeUsername(c *gin.Context) {
	id := jwt.GetUserId(c)

	m := map[string]string{}
	c.BindJSON(&m)
	name := m["username"]

	if ctr.uServ.ChangeUsername(id, name) != nil {
		c.JSON(500, gin.H{"error": "登録に失敗しました。"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}


//DELETE /api/account
func (ctr *userController) DeleteAccount(c *gin.Context) {
	id := jwt.GetUserId(c)

	if ctr.uServ.DeleteUser(id) != nil {
		c.JSON(500, gin.H{"error": "削除に失敗しました。"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}