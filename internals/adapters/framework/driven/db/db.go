package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"shortener/internals/config"
	"shortener/internals/core/domains/url/models"
	"shortener/internals/ports"
)

var DB *gorm.DB

type dbAdapter struct {
}

func (d dbAdapter) Migration() {
	err := DB.AutoMigrate(&models.Url{})
	if err != nil {
		return
	}
	fmt.Println("👍 Migration complete")
}

func (d dbAdapter) Connect() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to get Config")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.AppConfig.DbHost, config.AppConfig.DbUser, config.AppConfig.DbPassword, config.AppConfig.DbName, config.AppConfig.DbPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")
}

func DbAdapter() ports.DB {
	return dbAdapter{}
}
