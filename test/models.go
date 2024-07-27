package main

type User struct {
	ID     int    `gorm:"primaryKey"`
	Name   string `gorm:"size:255"`
	Email  string `gorm:"size:255"`
	Orders []Order
}

type Order struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	Amount    float64
	CreatedAt string `gorm:"type:timestamp"`
	Items     []OrderItem
}

type OrderItem struct {
	ID          int `gorm:"primaryKey"`
	OrderID     int
	ProductName string `gorm:"size:255"`
	Price       float64
	Quantity    int
}
