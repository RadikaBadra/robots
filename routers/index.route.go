package routes

import (
	"net/http"
	"robots/routers/delivery_routes"
	"robots/routers/robot_routes"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	router := app

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Connected to server",
		})
	})

	robot_routes.RobotsRouter(app)
	delivery_routes.DeliveriesRouter(app)
}