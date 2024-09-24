package controller

import (
	"github.com/gin-gonic/gin"

	"goat/internal/core/jwt"
	"goat/internal/core/errs"
	"goat/internal/core/utils"
	"goat/internal/service"
	"goat/internal/dto"
	"goat/internal/request"
	"goat/internal/response"
)

type AccountController struct {
	accountService service.AccountService
}


func NewAccountController() *AccountController {
	return &AccountController{
		accountService: service.NewAccountService(),
	}
}


//GET /signup
func (ctr *AccountController) SignupPage(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}

//GET /login
func (ctr *AccountController) LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}


//GET /logout
func (ctr *AccountController) Logout(c *gin.Context) {
	jwt.RemoveTokenFromCookie(c)
	c.Redirect(303, "/login")
}


//POST /api/signup
func (ctr *AccountController) ApiSignup(c *gin.Context) {
	var req request.Signup
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var input dto.Signup
	utils.MapFields(&input, req)

	accountId, err := ctr.accountService.Signup(input)
	if err != nil {
		if _, ok := err.(errs.UniqueConstraintError); ok {
			c.JSON(409, gin.H{"error": "ユーザ名が既に使われています。"})
		} else {
			c.JSON(500, gin.H{"error": "登録に失敗しました。"})
		}
		c.Abort()
		return
	}

	res := response.Signup{Id: accountId}
	c.JSON(200, res)
}


//POST /api/login
func (ctr *AccountController) ApiLogin(c *gin.Context) {
	var req request.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var input dto.Login
	utils.MapFields(&input, req)

	account, err := ctr.accountService.Login(input)
	if err != nil {
		c.JSON(401, gin.H{"error": "ユーザ名またはパスワードが異なります。"})
		c.Abort()
		return
	}

	pl, err := ctr.accountService.GenerateJwtPayload(account.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": "ログインに失敗しました。"})
		c.Abort()
		return
	}

	jwt.SetTokenToCookie(c, pl)
	c.JSON(200, gin.H{})
}


//GET /api/account
func (ctr *AccountController) ApiGetOne(c *gin.Context) {
	pl := jwt.GetPayload(c)
	account, err := ctr.accountService.GetOne(pl.AccountId)

	if err != nil {
		c.JSON(500, gin.H{})
		c.Abort()
		return
	}

	var res response.GetAccount
	utils.MapFields(&res, account)

	c.JSON(200, res)
}


//PUT /api/account/password
func (ctr *AccountController) ApiPutPassword(c *gin.Context) {
	pl := jwt.GetPayload(c)

	var req request.PutAccountPassword
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	input := dto.Login{Name: pl.AccountName, Password: req.OldPassword}
	_, err := ctr.accountService.Login(input)
	if err != nil {
		c.JSON(400, gin.H{"error": "旧パスワードが異なります。"})
		c.Abort()
		return
	}

	if ctr.accountService.UpdatePassword(pl.AccountId, req.Password) != nil {
		c.JSON(500, gin.H{"error": "変更に失敗しました。"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}


//PUT /api/account/name
func (ctr *AccountController) ApiPutName(c *gin.Context) {
	pl := jwt.GetPayload(c)

	var req request.PutAccountName
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	err := ctr.accountService.UpdateName(pl.AccountId, req.Name)
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
func (ctr *AccountController) ApiDelete(c *gin.Context) {
	pl := jwt.GetPayload(c)

	if ctr.accountService.Delete(pl.AccountId) != nil {
		c.JSON(500, gin.H{"error": "削除に失敗しました。"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{})
}