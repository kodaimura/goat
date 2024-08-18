package controller

import (
	"github.com/gin-gonic/gin"

	"goat/internal/core/jwt"
	"goat/internal/core/errs"
	"goat/internal/service"
)

type UserController struct {
	userService service.UserService
}


func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}


//GET /signup
func (ctr *UserController) SignupPage(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}

//GET /login
func (ctr *UserController) LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}


//GET /logout
func (ctr *UserController) Logout(c *gin.Context) {
	jwt.RemoveTokenFromCookie(c)
	c.Redirect(303, "/login")
}


//POST /api/signup
func (ctr *UserController) ApiSignup(c *gin.Context) {
	m := map[string]string{}
	c.BindJSON(&m)
	name := m["user_name"]
	pass := m["user_password"]

	err := ctr.userService.Signup(name, pass)
	if err != nil {
		if _, ok := err.(errs.UniqueConstraintError); ok {
			c.JSON(409, gin.H{"error": "ユーザ名が既に使われています。"})
		} else {
			c.JSON(500, gin.H{"error": "登録に失敗しました。"})
		}
		c.Abort()
		return
	}
	c.JSON(200, gin.H{})
}


//POST /api/login
func (ctr *UserController) ApiLogin(c *gin.Context) {
	m := map[string]string{}
	c.BindJSON(&m)
	name := m["user_name"]
	pass := m["user_password"]

	user, err := ctr.userService.Login(name, pass)
	if err != nil {
		c.JSON(401, gin.H{"error": "ユーザ名またはパスワードが異なります。"})
		c.Abort()
		return
	}

	pl, err := ctr.userService.GenerateJwtPayload(user.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": "ログインに失敗しました。"})
		c.Abort()
		return
	}

	jwt.SetTokenToCookie(c, pl)
	c.JSON(200, gin.H{})
}


//GET /api/account/profile
func (ctr *UserController) ApiGetProfile(c *gin.Context) {
	pl := jwt.GetPayload(c)
	user, err := ctr.userService.GetProfile(pl.UserId)

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, user)
}


//PUT /api/account/password
func (ctr *UserController) ApiPutPassword(c *gin.Context) {
	pl := jwt.GetPayload(c)
	id := pl.UserId
	name := pl.UserName

	m := map[string]string{}
	c.BindJSON(&m)
	oldPass := m["old_user_password"]
	pass := m["user_password"]

	_, err := ctr.userService.Login(name, oldPass)
	if err != nil {
		c.JSON(400, gin.H{"error": "旧パスワードが異なります。"})
		c.Abort()
		return
	}

	if ctr.userService.UpdatePassword(id, pass) != nil {
		c.JSON(500, gin.H{"error": "変更に失敗しました。"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}


//PUT /api/account/name
func (ctr *UserController) ApiPutName(c *gin.Context) {
	pl := jwt.GetPayload(c)
	id := pl.UserId

	m := map[string]string{}
	c.BindJSON(&m)
	name := m["user_name"]

	err := ctr.userService.UpdateName(id, name)
	if err != nil {
		if _, ok := err.(errs.UniqueConstraintError); ok {
			c.JSON(409, gin.H{"error": "ユーザ名が既に使われています。"})
		} else {
			c.JSON(500, gin.H{"error": "変更に失敗しました。"})
		}
		c.Abort()
		return
	}
	c.JSON(200, gin.H{})
}


//DELETE /api/account
func (ctr *UserController) ApiDeleteAccount(c *gin.Context) {
	pl := jwt.GetPayload(c)
	id := pl.UserId

	if ctr.userService.DeleteUser(id) != nil {
		c.JSON(500, gin.H{"error": "削除に失敗しました。"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}