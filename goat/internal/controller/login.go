package controller

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    
    "goat/config"
    "goat/internal/core/jwt"
    "goat/internal/model/repository"
    "goat/internal/constants"
    "goat/internal/core/logger"
)


func setLoginRoute(r *gin.Engine) {
    lc := newLoginController()

    r.GET("/login", lc.loginPage)
    r.POST("/login", lc.login)
    r.GET("/logout", lc.logout)
}


type loginController struct {
    ur repository.UserRepository
}


func newLoginController() *loginController {
    ur := repository.NewUserRepository()
    return &loginController{ur}
}


//GET /login
func (ctr *loginController) loginPage(c *gin.Context) {
    c.HTML(200, "login.html", gin.H{
        "commons": constants.Commons,
    })
}


//POST /login
func (ctr *loginController) login(c *gin.Context) {
    un := c.PostForm("username")
    pw := c.PostForm("password")

    user, err := ctr.ur.SelectByUsername(un)

    if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pw)) != nil{
        c.HTML(401, "login.html", gin.H{
            "commons": constants.Commons,
            "error": "UsernameまたはPasswordが異なります。",
        })
        c.Abort()
        return
    }

    var cc jwt.CustomClaims
    cc.UserId = user.UserId
    cc.Username = user.Username
    jwtStr, err := jwt.GenerateJWT(cc)

    if err != nil {
        logger.LogError(err.Error())
        c.HTML(500, "login.html", gin.H{
            "commons": constants.Commons,
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
func (ctr *loginController) logout(c *gin.Context) {
    cf := config.GetConfig()
    c.SetCookie(jwt.COOKIE_KEY_JWT, "", 0, "/", cf.AppHost, false, true)
    c.Redirect(303, "/login")
}