package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ridakaddir/go-fiber-api/employee/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	DB *gorm.DB
}

var Database DBInstance

func ConnectDB() {

	DBHost:=os.Getenv("DB_HOST")
	if DBHost == "" {
		DBHost = "localhost"
	}
	DBPort:=os.Getenv("DB_PORT")
	if DBPort == "" {
		DBPort = "5432"
	}
	DBUser:=os.Getenv("DB_USER")
	if DBUser == "" {
		DBUser = "api_db_user"
	}
	DBPass:=os.Getenv("DB_PASS")
	if DBPass == "" {
		DBPass = "P@ssword"
	}
	
	

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=api_db port=%s sslmode=disable TimeZone=Africa/Casablanca", DBHost, DBUser, DBPass, DBPort)
	

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database \n", err.Error())
		
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	// TODO: Add migrations
	db.AutoMigrate(&models.Employee{})

	Database.DB = db
}