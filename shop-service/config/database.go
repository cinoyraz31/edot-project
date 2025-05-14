package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"shop-service/model"
	"time"
)

func OpenConnection() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database handle!", err)
	}

	// Set connection pool properties
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	err = db.AutoMigrate(
		&model.Shop{},
		&model.ShopProduct{},
	)

	if err != nil {
		panic(err)
	}

	return db
}
