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

// GET /signup
func (ctr *AccountController) SignupPage(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}

// GET /login
func (ctr *AccountController) LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

// GET /logout
func (ctr *AccountController) Logout(c *gin.Context) {
	jwt.RemoveTokenFromCookie(c)
	c.Redirect(303, "/login")
}

// POST /api/signup
func (ctr *AccountController) ApiSignup(c *gin.Context) {
	var req request.Signup
	if err := c.ShouldBindJSON(&req); err != nil {
		JsonError(c, 400, "不正なリクエストです。")
		return
	}

	var input dto.Signup
	utils.MapFields(&input, req)

	pk, err := ctr.accountService.Signup(input)
	if err != nil {
		if _, ok := err.(errs.UniqueConstraintError); ok {
			JsonError(c, 409, "ユーザ名が既に使われています。")
		} else {
			JsonError(c, 500, "登録に失敗しました。")
		}
		return
	}

	var res response.AccountPK
	utils.MapFields(&res, pk)
	c.JSON(200, res)
}

// POST /api/login
func (ctr *AccountController) ApiLogin(c *gin.Context) {
	var req request.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		JsonError(c, 400, "不正なリクエストです。")
		return
	}

	var input dto.Login
	utils.MapFields(&input, req)

	account, err := ctr.accountService.Login(input)
	if err != nil {
		JsonError(c, 401, "ユーザ名またはパスワードが異なります。")
		return
	}

	pl, err := ctr.accountService.GenerateJwtPayload(dto.AccountPK{Id: account.Id})
	if err != nil {
		JsonError(c, 500, "ログインに失敗しました。")
		return
	}

	jwt.SetTokenToCookie(c, pl)
	c.JSON(200, gin.H{})
}

// GET /api/account
func (ctr *AccountController) ApiGetOne(c *gin.Context) {
	pl := jwt.GetPayload(c)
	account, err := ctr.accountService.GetOne(dto.AccountPK{Id: pl.AccountId})

	if err != nil {
		JsonError(c, 500, "アカウント情報の取得に失敗しました。")
		return
	}

	var res response.GetAccount
	utils.MapFields(&res, account)

	c.JSON(200, res)
}

// PUT /api/account/password
func (ctr *AccountController) ApiPutPassword(c *gin.Context) {
	pl := jwt.GetPayload(c)

	var req request.PutAccountPassword
	if err := c.ShouldBindJSON(&req); err != nil {
		JsonError(c, 400, "不正なリクエストです。")
		return
	}

	_, err := ctr.accountService.Login(dto.Login{Name: pl.AccountName, Password: req.OldPassword})
	if err != nil {
		JsonError(c, 400, "旧パスワードが異なります。")
		return
	}

	var input dto.UpdateAccountPassword
	utils.MapFields(&input, req)
	input.Id = pl.AccountId

	if err := ctr.accountService.UpdatePassword(input); err != nil {
		JsonError(c, 500, "変更に失敗しました。")
		return
	}

	c.JSON(200, gin.H{})
}

// PUT /api/account/name
func (ctr *AccountController) ApiPutName(c *gin.Context) {
	pl := jwt.GetPayload(c)

	var req request.PutAccountName
	if err := c.ShouldBindJSON(&req); err != nil {
		JsonError(c, 400, "不正なリクエストです。")
		return
	}

	var input dto.UpdateAccountName
	utils.MapFields(&input, req)
	input.Id = pl.AccountId

	err := ctr.accountService.UpdateName(input)
	if err != nil {
		if _, ok := err.(errs.UniqueConstraintError); ok {
			JsonError(c, 409, "ユーザ名が既に使われています。")
		} else {
			JsonError(c, 500, "変更に失敗しました。")
		}
		return
	}
	c.JSON(200, gin.H{})
}

// DELETE /api/account
func (ctr *AccountController) ApiDelete(c *gin.Context) {
	pl := jwt.GetPayload(c)

	if err := ctr.accountService.Delete(dto.AccountPK{Id: pl.AccountId}); err != nil {
		JsonError(c, 500, "削除に失敗しました。")
		return
	}

	c.JSON(200, gin.H{})
}
