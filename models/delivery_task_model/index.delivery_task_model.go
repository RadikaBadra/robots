package delivery_task_model

import "github.com/google/uuid"

type Delivery struct {
	Id               uuid.UUID `gorm:"primaryKey"`
	RobotId          string
	VendorId         string
	TaskId           string
	TaskType         string
	DestinationPoint string
}