// Controllers
package main

import "github.com/jangidRkt08/pizza-tracker-go/internal/models"

type Handler struct {
	orders *models.OrderModel
	users *models.UserModel
	notificationManager *NotificationManager
}

func NewHandler(dbModel *models.DBModel) *Handler{
	return &Handler{
		orders: &dbModel.Order,
		users: &dbModel.User,
		notificationManager: NewNotificationManager(),
	}
}
