package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "goat/config"
    "goat/internal/core/jwt"
    "goat/internal/constant"
    "goat/internal/service"
    "goat/internal/model/queryservice"
)


func setUserRoute(r *gin.Engine) {
    uc := newUserController()

    r.GET("/signup", uc.signupPage)
    r.GET("/login", uc.loginPage)
    r.GET("/logout", uc.logout)

    r.POST("/signup", uc.signup)
    r.POST("/login", uc.login)

    api := r.Group("/api", jwt.JwtAuthMiddleware())
    {
        api.GET("/profile", uc.getProfile)
    }
}


type userController struct {
    uServ service.UserService
    uQue queryservice.UserQueryService
}


func newUserController() *userController {
    uServ := service.NewUserService()
    uQue := queryservice.NewUserQueryService()
    return &userController{uServ, uQue}
}


//GET /signup
func (ctr *userController) signupPage(c *gin.Context) {
    c.HTML(200, "signup.html", gin.H{
        "commons": constant.Commons,
    })
}

//GET /login
func (ctr *userController) loginPage(c *gin.Context) {
    c.HTML(200, "login.html", gin.H{
        "commons": constant.Commons,
    })
}


//POST /signup
func (ctr *userController) signup(c *gin.Context) {
    name := c.PostForm("user_name")
    pass := c.PostForm("password")

    result := ctr.uServ.Signup(name, pass)

    if result == service.SIGNUP_SUCCESS_INT {
        c.Redirect(303, "/login")

    } else if result == service.SIGNUP_CONFLICT_INT {
        c.HTML(409, "signup.html", gin.H{
            "commons": constant.Commons,
            "error": "Usernameが既に使われています。",
        })

    } else {
        c.HTML(500, "signup.html", gin.H{
            "commons": constant.Commons,
            "error": "登録に失敗しました。",
        })
    }
}


//POST /login
func (ctr *userController) login(c *gin.Context) {
    name := c.PostForm("user_name")
    pass := c.PostForm("password")

    userId := ctr.uServ.Login(name, pass)

    if userId == service.LOGIN_FAILURE_INT {
        c.HTML(401, "login.html", gin.H{
            "commons": constant.Commons,
            "error": "UserNameまたはPasswordが異なります。",
        })
        c.Abort()
        return
    }

    jwtStr := ctr.uServ.GenerateJWT(userId)

    if jwtStr == service.GENERATE_JWT_FAILURE_STR {
        c.HTML(500, "login.html", gin.H{
            "commons": constant.Commons,
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
func (ctr *userController) logout(c *gin.Context) {
    cf := config.GetConfig()
    c.SetCookie(jwt.COOKIE_KEY_JWT, "", 0, "/", cf.AppHost, false, true)
    c.Redirect(303, "/login")
}


//GET /api/profile
func (ctr *userController) getProfile(c *gin.Context) {
    userId, _ := jwt.GetUserId(c)
    user, err := ctr.uQue.QueryUser(userId)

    if err != nil {
        c.JSON(500, gin.H{"error": http.StatusText(500)})
        c.Abort()
        return
    }

    c.JSON(200, user)
}