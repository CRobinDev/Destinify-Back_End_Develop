package postgres

import (
	"github.com/CRobin69/Destinify-Back_End_Develop/pkg/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDB() *gorm.DB {

	db, err := gorm.Open(postgres.Open(config.LoadDatabaseConfig()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	return db
}
