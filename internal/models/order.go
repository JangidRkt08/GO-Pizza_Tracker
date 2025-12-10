package models

import (
	"time"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

var (
	OrderStatuses = []string{
		"Order Placed",
		"Preparing",
		"Baking",
		"Quality Check",
		"Ready"}
	PizzaTypes = []string{
		"Margherita",
		"Pepperoni",
		"Vegetarian",
		"FourCheese",
		"Hawaiian",
		"Truffle Mushroom",
		"Chef's Special"}
	PizzaSizes = []string{
		"Small",
		"Medium",
		"Large",
		"Extra Large"}
)

type OrderModel struct {
	DB *gorm.DB
}

type Order struct { //represent order in DB
	ID           string      `gorm:"primary_key; size:14" json:"id"`
	Status       string      `grom:"not null" json:"status"`
	CustomerName string      `grom:"not null" json:"customerName"`
	Phone        string      `grom:"not null" json:"phone"`
	Address      string      `grom:"not null" json:"address"`
	Items        []OrderItem `grom:"foreignkey:OrderID" json:"items"` //one order contains many items(pizza)
	CreatedAt    time.Time       `json:"created_at"`
}

type OrderItem struct {
	ID           string `gorm:"primary_key; size:14" json:"id"`
	OrderID      string `gorm:"index;" json:"order_id"`
	Size         string `grom:"not null" json:"size"`
	Pizza        string `grom:"not null" json:"pizza"`
	Instructions string `json:"instructions"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.ID == "" {
		o.ID = shortid.MustGenerate()
	}

	return nil
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error {
	if oi.ID == "" {
		oi.ID = shortid.MustGenerate()
	}

	return nil
}

func (o *OrderModel) CreateOrder(order *Order) error {
	return o.DB.Create(order).Error
}

func (o *OrderModel) GetOrder(id string) (*Order, error) {
	var order Order
	err := o.DB.Preload("Items").First(&order, "id = ?", id).Error
	return &order, err
}

func (o *OrderModel) GetAllOrders() ([]Order, error) {
	var orders []Order
	err := o.DB.Preload("Items").Order("created_at desc").Find(&orders).Error
	return orders, err
}

func (o *OrderModel) UpdateOrderStatus(id string, status string) error {
	return o.DB.Model(&Order{}).Where("id = ?", id).Update("status", status).Error 
}

func (o *OrderModel) DeleteOrder(id string) error {
	return o.DB.Select("Items").Delete(&Order{ID: id}).Error
}
