package config

import (
	"fmt"
	"log"
	"os"
	"task_management_service/database"
	"task_management_service/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase() {
	var err error

	PSQL_HOST := Config("PSQL_HOST")
	PSQL_DB_NAME := Config("PSQL_DB_NAME")
	PSQL_USER := Config("PSQL_USER")
	PSQL_PASS := Config("PSQL_PASS")
	PSQL_PORT := Config("PSQL_PORT")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s",
		PSQL_HOST, PSQL_USER, PSQL_DB_NAME, PSQL_PORT, PSQL_PASS)

	log.Print("Connecting to PostgreSQL DB....")
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("connected")
	log.Println("running migrations")

	// กำหนดค่า connection pool
	sqlDB, err := database.DBConn.DB()
	if err != nil {
		log.Fatal("Failed to get database connection. \n", err)
		os.Exit(2)
	}

	// ตั้งค่าการเชื่อมต่อ connection pool
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Println("running migrations")

	defer database.DBConn.AutoMigrate(
		&models.Task{},
		&models.User{},
	)

}
