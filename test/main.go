package main

import (
	"log"
	"net/http"
	"test/resource/db"

	"test/config/"

	"github.com/labstack/echo/v4"
)

func main() {
	var err error

	// Initialize database connection
	db, err := db.NewDBConnection(config.Db)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Migrate the schema
	err = db.AutoMigrate(&User{}, &Order{}, &OrderItem{})
	if err != nil {
		panic("failed to migrate database")
	}

	e := echo.New()

	// Define routes
	e.GET("/users", GetUsers)
	e.POST("/users", CreateUser)
	e.GET("/orders", GetOrders)
	e.POST("/orders", CreateOrder)

	e.Logger.Fatal(e.Start(":8080"))
}

// Handlers

func GetUsers(c echo.Context) error {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, user)
}

func GetOrders(c echo.Context) error {
	var orders []Order
	if err := db.Find(&orders).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orders)
}

func CreateOrder(c echo.Context) error {
	order := Order{}
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := db.Create(&order).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, order)
}
