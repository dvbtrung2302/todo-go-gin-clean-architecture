package db

import (
	"log"
	"todo-backend/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	connStr := "postgresql://dvbtrung23:dvbt230220@localhost:5432/todos?sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entity.Todo{})

	return db
}
