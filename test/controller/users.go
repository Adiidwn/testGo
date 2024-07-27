package controller

import (
	"fmt"
	"net/http"
	models "test/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func (uc *UserController) CreateUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := uc.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, user)
}

func (uc *UserController) GetUsers(c echo.Context) error {
	var users []models.User
	if err := uc.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println("users", users)
	return c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetHighSpendingUsers(c echo.Context) error {
	var result []struct {
		Username   string  `json:"username"`
		TotalSpent float64 `json:"total_spent"`
	}

	query := `
	SELECT u.name AS username, SUM(o.amount) AS total_spent
	FROM users u
	JOIN orders o ON u.id = o.user_id
	WHERE o.created_at >= '2022-01-01'
	GROUP BY u.name
	HAVING SUM(o.amount) >= 1000;
	`

	// Log the query
	fmt.Println("Executing SQL Query:", query)

	// Execute the query and scan the results into the result slice
	if err := uc.DB.Raw(query).Scan(&result).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Return the results as a JSON response
	return c.JSON(http.StatusOK, result)
}
