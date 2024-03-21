package database

import (
	"robots/config/db_config"
	"robots/models/delivery_task_model"
	"robots/models/robot_models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnecDB() {
	database, err := gorm.Open(sqlite.Open(db_config.DB_NAME), &gorm.Config{})
	if err != nil {
		panic("failed connecting to DB")
	}

	err = database.AutoMigrate(&robot_models.Robot{}, &delivery_task_model.Delivery{})
        if err != nil {
                return
        }

	DB = database
}