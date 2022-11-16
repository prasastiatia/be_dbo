package repository

import (
	database "be_dbo/database"
	models "be_dbo/models"
)

func GetAllCustomer(customer *models.Customer, pagination *models.Pagination) (*[]models.Customer, error) {
	var customers []models.Customer
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := database.Instance.Statement.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.Customer{}).Where(customer).Find(&customers)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &customers, nil
}
