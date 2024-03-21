package robot_models

import "github.com/google/uuid"

type Robot struct {
	Id       uuid.UUID `gorm:"primaryKey"`
	Name     string
	VendorId string
	Battery  float32
	Status   int
}