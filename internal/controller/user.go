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
func (uc *UserController) SignupPage(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}

//GET /login
func (uc *UserController) LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}


//GET /logout
func (uc *UserController) Logout(c *gin.Context) {
	jwt.RemoveTokenFromCookie(c)
	c.Redirect(303, "/login")
}


//POST /api/signup
func (uc *UserController) ApiSignup(c *gin.Context) {
	m := map[string]string{}
	c.BindJSON(&m)
	name := m["user_name"]
	pass := m["user_password"]

	err := uc.userService.Signup(name, pass)
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
func (uc *UserController) ApiLogin(c *gin.Context) {
	m := map[string]string{}
	c.BindJSON(&m)
	name := m["user_name"]
	pass := m["user_password"]

	user, err := uc.userService.Login(name, pass)
	if err != nil {
		c.JSON(401, gin.H{"error": "ユーザ名またはパスワードが異なります。"})
		c.Abort()
		return
	}

	pl, err := uc.userService.GenerateJwtPayload(user.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": "ログインに失敗しました。"})
		c.Abort()
		return
	}

	jwt.SetTokenToCookie(c, pl)
	c.JSON(200, gin.H{})
}


//GET /api/account/profile
func (uc *UserController) ApiGetProfile(c *gin.Context) {
	pl := jwt.GetPayload(c)
	user, err := uc.userService.GetProfile(pl.UserId)

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, user)
}


//PUT /api/account/password
func (uc *UserController) ApiPutPassword(c *gin.Context) {
	pl := jwt.GetPayload(c)
	id := pl.UserId
	name := pl.UserName

	m := map[string]string{}
	c.BindJSON(&m)
	oldPass := m["old_user_password"]
	pass := m["user_password"]

	_, err := uc.userService.Login(name, oldPass)
	if err != nil {
		c.JSON(400, gin.H{"error": "旧パスワードが異なります。"})
		c.Abort()
		return
	}

	if uc.userService.UpdatePassword(id, pass) != nil {
		c.JSON(500, gin.H{"error": "変更に失敗しました。"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}


//PUT /api/account/name
func (uc *UserController) ApiPutName(c *gin.Context) {
	pl := jwt.GetPayload(c)
	id := pl.UserId

	m := map[string]string{}
	c.BindJSON(&m)
	name := m["user_name"]

	err := uc.userService.UpdateName(id, name)
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
func (uc *UserController) ApiDeleteAccount(c *gin.Context) {
	pl := jwt.GetPayload(c)
	id := pl.UserId

	if uc.userService.DeleteUser(id) != nil {
		c.JSON(500, gin.H{"error": "削除に失敗しました。"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}