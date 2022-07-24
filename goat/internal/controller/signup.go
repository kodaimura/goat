package controller

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    
    "goat/internal/constants"
    "goat/internal/model/repository"
)


func setSignupRoute(r *gin.Engine) {
    sc := newSignupController()

    r.GET("/signup", sc.signupPage)
    r.POST("/signup", sc.signup)
}


type signupController struct {
    ur repository.UserRepository
}


func newSignupController() *signupController {
    ur := repository.NewUserRepository()
    return &signupController{ur}
}


//GET /signup
func (ctr *signupController) signupPage(c *gin.Context) {
    c.HTML(200, "signup.html", gin.H{
        "commons": constants.Commons,
    })
}


//POST /signup
func (ctr *signupController) signup(c *gin.Context) {
    un := c.PostForm("username")
    pw := c.PostForm("password")

    if _, err := ctr.ur.SelectByUsername(un); err == nil {
        c.HTML(409, "signup.html", gin.H{
            "commons": constants.Commons,
            "error": "Usernameが既に使われています。",
        })
        c.Abort()
        return
    }

    hashed, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)

    if ctr.ur.Signup(un, string(hashed)) != nil {
        c.HTML(500, "signup.html", gin.H{
            "commons": constants.Commons,
            "error": "登録に失敗しました。",
        })
        c.Abort()
        return
    }

    c.Redirect(303, "/login")
}