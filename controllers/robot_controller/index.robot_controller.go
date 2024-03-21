package robot_controller

import (
	"net/http"
	"robots/database"
	"robots/models/robot_models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllRobots(ctx *gin.Context) {
	var robots []robot_models.Robot

	database.DB.Find(&robots)
	ctx.HTML(http.StatusOK, "index.robots.html", gin.H{
		"data": robots,
	})
}

func GetRobot(ctx *gin.Context) {
	var robots robot_models.Robot

	if err := database.DB.Where("id = ?", ctx.Param("id")).First(&robots).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	ctx.HTML(http.StatusOK, "detail.robots.html", gin.H{
		"id":        robots.Id,
		"name":      robots.Name,
		"vendor_id": robots.VendorId,
		"battery":   robots.Battery,
		"status":    robots.Status,
	})
}

type CreateRobotInput struct {
	Id       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	VendorId string  `json:"vendor_id" binding:"required"`
	Battery  float32 `json:"battery" binding:"required"`
	Status   int     `json:"status" binding:"required"`
}

func CreateRobot(ctx *gin.Context) {
	var input CreateRobotInput
	robot_id := uuid.New()

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	robot := robot_models.Robot{Id: robot_id, Name: input.Name, VendorId: input.VendorId, Battery: input.Battery, Status: input.Status}

	database.DB.Create(&robot)

	ctx.JSON(http.StatusOK, gin.H{"data": robot})
}

type UpdateRobotInput struct {
	Name     string  `json:"name"`
	VendorId string  `json:"vendor_id"`
	Battery  float32 `json:"battery"`
	Status   int     `json:"status"`
}

func UpdateRobot(ctx *gin.Context) {
	var input UpdateRobotInput
	var robot robot_models.Robot

	if err := database.DB.Where("id = ?", ctx.Param("id")).First(&robot).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&robot).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": robot})
}

func DeleteRobot(ctx *gin.Context) {
	var robot robot_models.Robot

	if err := database.DB.Where("id = ?", ctx.Param("id")).First(&robot).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	database.DB.Delete(&robot)
	ctx.JSON(http.StatusOK, gin.H{"data": true})

}
