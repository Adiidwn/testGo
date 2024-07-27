package controller

import (
	"net/http"
	models "test/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrderItemController struct {
	DB *gorm.DB
}

func (oic *OrderItemController) CreateOrderItem(c echo.Context) error {
	orderItem := models.OrderItem{}
	if err := c.Bind(&orderItem); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := oic.DB.Create(&orderItem).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, orderItem)
}

func (oic *OrderItemController) GetOrderItems(c echo.Context) error {
	var orderItems []models.OrderItem
	if err := oic.DB.Find(&orderItems).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orderItems)
}
