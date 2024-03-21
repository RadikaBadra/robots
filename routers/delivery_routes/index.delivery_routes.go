package delivery_routes

import (
	"robots/controllers/delivery_task_controller"

	"github.com/gin-gonic/gin"
)

func DeliveriesRouter(app *gin.Engine) {
	router := app

	router.GET("/deliveries", delivery_task_controller.GetAllDeliveries)
	router.GET("/deliveries/:id", delivery_task_controller.GetDelivery)
	router.POST("/deliveries", delivery_task_controller.CreateDelivery)
	router.PATCH("/deliveries/:id", delivery_task_controller.UpdateDelivery)
	router.DELETE("/deliveries/:id", delivery_task_controller.DeleteDelivery)

}