package server

import (
	"github.com/gin-gonic/gin"

	"goat-base/config"
	"goat-base/internal/core/logger"
	"goat-base/internal/controller"
)

func Run() {
	cf := config.GetConfig()
	logger.SetOutputLogToFile()
	r := router()
	r.Run(":" + cf.AppPort)
}

func router() *gin.Engine {
	r := gin.Default()
	
	//TEMPLATE
	r.LoadHTMLGlob("web/template/*.html")

	//STATIC
	r.Static("/css", "web/static/css")
	r.Static("/js", "web/static/js")
	r.Static("/img", "web/static/img")
	r.StaticFile("/favicon.ico", "web/static/favicon.ico")
	r.StaticFile("/manifest.json", "web/static/manifest.json")

	controller.SetRouter(r)

	return r
}
