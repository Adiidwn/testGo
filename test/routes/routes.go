package routes

import (
	con "test/controller"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	userController := con.UserController{DB: db}
	orderController := con.OrderController{DB: db}
	orderItemController := con.OrderItemController{DB: db}

	api := e.Group("/api/v1")
	{
		api.GET("/users", userController.GetUsers)
		api.GET("/find-high-spending", userController.GetHighSpendingUsers)

		api.GET("/orders", orderController.GetOrders)

		api.GET("/order-items", orderItemController.GetOrderItems)

	}

}
