package repository

import (
	database "be_dbo/database"
	models "be_dbo/models"
)

func GetAllOrders(order *models.Order, pagination *models.Pagination) (*[]models.Order, error) {
	var orders []models.Order
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := database.Instance.Statement.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.Order{}).Where(order).Find(&orders)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &orders, nil
}
