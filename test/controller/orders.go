package controller

import (
	"net/http"
	models "test/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrderController struct {
	DB *gorm.DB
}

func (oc *OrderController) CreateOrder(c echo.Context) error {
	order := models.Order{}
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := oc.DB.Create(&order).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, order)
}

func (oc *OrderController) GetOrders(c echo.Context) error {
	var orders []models.Order
	if err := oc.DB.Find(&orders).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orders)
}
