package delivery_task_controller

import (
	"net/http"
	"robots/database"
	"robots/models/delivery_task_model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllDeliveries(ctx *gin.Context) {
	var deliveries []delivery_task_model.Delivery

	database.DB.Find(&deliveries)
	ctx.HTML(http.StatusOK, "index.delivery.html", gin.H {
		"data" : deliveries,
	})
}

func GetDelivery(ctx *gin.Context) {
	var delivery delivery_task_model.Delivery

  if err := database.DB.Where("id = ?", ctx.Param("id")).First(&delivery).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	ctx.HTML(http.StatusOK, "detail.delivery.html", gin.H{
		"id" : delivery.Id,
		"robot_id" : delivery.RobotId,
		"vendor_id" : delivery.VendorId,
		"task_id" : delivery.TaskId,
		"task_type"  : delivery.TaskType,
		"destination_point" : delivery.DestinationPoint,
	})
}

type CreateDeliveryInput struct {
	Id               int	`json:"id"` 
	RobotId          string	`json:"robot_id" binding:"required"`
	VendorId         string `json:"vendor_id" binding:"required"`
	TaskId           string `json:"task_id" binding:"required"`
	TaskType         string `json:"task_type" binding:"required"`
	DestinationPoint string `json:"destination_point" binding:"required"`
}

func CreateDelivery(ctx *gin.Context) {
	var input CreateDeliveryInput
	delivery_id := uuid.New()

	if err := ctx.ShouldBindJSON(&input); err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

	delivery := delivery_task_model.Delivery{Id: delivery_id, RobotId: input.RobotId, VendorId: input.VendorId, TaskId: input.TaskId, TaskType: input.TaskType, DestinationPoint: input.DestinationPoint}

	database.DB.Create(&delivery)

	ctx.JSON(http.StatusOK, gin.H{"data": delivery})
}

type UpdateRobotInput struct {
	RobotId          string	`json:"robot_id"`
	VendorId         string `json:"vendor_id"`
	TaskId           string `json:"task_id"`
	TaskType         string `json:"task_type"`
	DestinationPoint string `json:"destination_point"`
}

func UpdateDelivery(ctx *gin.Context) {
	var input UpdateRobotInput
	var delivery delivery_task_model.Delivery
	
	if err := database.DB.Where("id = ?", ctx.Param("id")).First(&delivery).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

	database.DB.Model(&delivery).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": delivery})
}

func DeleteDelivery(ctx *gin.Context) {
	var delivery delivery_task_model.Delivery

	if err := database.DB.Where("id = ?", ctx.Param("id")).First(&delivery).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	database.DB.Delete(&delivery)
	ctx.JSON(http.StatusOK, gin.H{"data": true})

}
