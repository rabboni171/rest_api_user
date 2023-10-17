package repository

import (
	"fmt"
	"log"
	"restApi/configs"
	"restApi/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB(cfg configs.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Dbname, cfg.Password, cfg.SSLMode)))
	if err != nil {
		return nil, err
	}
	db.Logger.LogMode(logger.Info)
	db.Logger = db.Logger.LogMode(logger.Info)

	fmt.Println("Connected TO DB!!!")

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to automigrate !!!")
	}

	return db, nil
}
