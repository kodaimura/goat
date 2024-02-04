package controller

import (
	"github.com/gin-gonic/gin"

	"goat-base/config"
	"goat-base/internal/core/jwt"
	"goat-base/internal/core/errs"
	"goat-base/internal/service"
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


//POST /signup
func (uc *UserController) Signup(c *gin.Context) {
	name := c.PostForm("user_name")
	pass := c.PostForm("user_password")

	err := uc.userService.Signup(name, pass)

	if err != nil {
		if _, ok := err.(errs.UniqueConstraintError); ok {
			c.HTML(409, "signup.html", gin.H{
				"user_name": name,
				"user_password": pass,
				"error": "ユーザ名が既に使われています。",
			})
		} else {
			c.HTML(500, "signup.html", gin.H{
				"user_name": name,
				"user_password": pass,
				"error": "登録に失敗しました。",
			})
		}
		c.Abort()
		return
	}

	c.Redirect(303, "/login")
}


//POST /login
func (uc *UserController) Login(c *gin.Context) {
	name := c.PostForm("user_name")
	pass := c.PostForm("user_password")

	user, err := uc.userService.Login(name, pass)

	if err != nil {
		c.HTML(401, "login.html", gin.H{
			"user_name": name,
			"user_password": pass,
			"error": "ユーザ名またはパスワードが異なります。",
		})
		c.Abort()
		return
	}

	jwtStr, err := uc.userService.GenerateJWT(user.Id)

	if err != nil {
		c.HTML(500, "login.html", gin.H{
			"user_name": name,
			"user_password": pass,
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
func (uc *UserController) Logout(c *gin.Context) {
	cf := config.GetConfig()
	c.SetCookie(jwt.COOKIE_KEY_JWT, "", 0, "/", cf.AppHost, false, true)
	c.Redirect(303, "/login")
}


//GET /api/account/profile
func (uc *UserController) GetProfile(c *gin.Context) {
	user, err := uc.userService.GetProfile(jwt.GetUserId(c))

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	c.JSON(200, user)
}


//PUT /api/account/password
func (uc *UserController) UpdatePassword(c *gin.Context) {
	id := jwt.GetUserId(c)
	name := jwt.GetUserName(c)

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
func (uc *UserController) UpdateName(c *gin.Context) {
	id := jwt.GetUserId(c)

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
func (uc *UserController) DeleteAccount(c *gin.Context) {
	id := jwt.GetUserId(c)

	if uc.userService.DeleteUser(id) != nil {
		c.JSON(500, gin.H{"error": "削除に失敗しました。"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}