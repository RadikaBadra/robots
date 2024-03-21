package bootstrap

import (
	"robots/config/app_config"
	"robots/database"
	routes "robots/routers"

	"github.com/gin-gonic/gin"
)

func BootStrap() {
	database.ConnecDB()
	app := gin.Default()
	app.LoadHTMLGlob("views/**/*.html")
	app.Static("/static", "./static/")
	routes.InitRoute(app)
	app.Run(app_config.PORT)
}