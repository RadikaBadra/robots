package robot_routes

import (
	"robots/controllers/robot_controller"

	"github.com/gin-gonic/gin"
)

func RobotsRouter(app *gin.Engine) {
	router := app

	router.GET("/robots", robot_controller.GetAllRobots)
	router.GET("/robots/:id", robot_controller.GetRobot)
	router.POST("/robots", robot_controller.CreateRobot)
	router.PATCH("/robots/:id", robot_controller.UpdateRobot)
	router.DELETE("/robots/:id", robot_controller.DeleteRobot)
}