package main

import (
	"log"
	"test/config"
	models "test/models"
	"test/routes"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	dsn := "user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " host=" + cfg.DBHost + " port=" + cfg.DBPort + " sslmode=" + cfg.DBSSLMode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(models.User{}, models.Order{}, models.OrderItem{})

	e := echo.New()

	routes.SetupRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
