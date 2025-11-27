package main

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jangidRkt08/pizza-tracker-go/internal/models"
)

type OrderFormData struct { //When we render OrderForm
	PizzaTypes []string
	PizzaSizes []string
}

type OrderRequest struct {
	Name         string   `form:"name" binding:"required,min=2,max=100"`
	Phone        string   `form:"phone" binding:"required,min = 10,max = 10"`
	Address      string   `form:"address" binding:"required"`
	Sizes        []string `form:"size" binding:"required,min=1,dive,valid_pizza_size"`
	PizzaTypes   []string `form:"pizzaTypes" binding:"required,min=1,dive,valid_pizza_type"`
	Instructions []string `form:"instructions" binding:"max=200"`
}

func (h *Handler) ServeOrderForm(c *gin.Context) {
	c.HTML(http.StatusOK, "order-form.html", OrderFormData{
		PizzaTypes: models.PizzaTypes,
		PizzaSizes: models.PizzaSizes,
	})
}

func (h *Handler) HandleNewOrderPost(c *gin.Context) {
	var form OrderRequest

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	orderItems := make([]models.OrderItem, len(form.Sizes))
	for i := range orderItems {
		orderItems[i] = models.OrderItem{
			Size:         form.Sizes[i],
			Pizza:        form.PizzaTypes[i],
			Instructions: form.Instructions[i],
		}
	}
	order := models.Order{
		CustomerName: form.Name,
		Phone:        form.Phone,
		Address:      form.Address,
		Status:        models.OrderStatuses[0],
		Items:        orderItems,
	}

	if err := h.orders.CreateOrder(&order); err != nil {
		slog.Error("failed to create order","error",err)
		c.String(http.StatusInternalServerError, "Something went wrong")
	}

	slog.Info("order created successfully","orderId",order.ID, "customerName",order.CustomerName)
	c.Redirect(http.StatusSeeOther, "/customer/"+order.ID)
}


// handler to track the order
func (h *Handler) ServeCustomer(c *gin.Context){
	orderID := c.Param("id")
	if orderID == ""{
		c.String(http.StatusBadRequest,"Invalid Order ID")
	}
	order ,err := h.orders.GetOrder(orderID)
	if err != nil{
		// slog.Error("failed to get order","error",err)
		c.String(http.StatusInternalServerError,"Order Not Found")
		return
	}
	c.HTML(http.StatusOK,"customer.tmpl",gin.H{
		"Order":order,
	})
}
