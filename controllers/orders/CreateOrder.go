package orders

import (
	database "be_dbo/database"
	"be_dbo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(context *gin.Context) {
	var order models.Order

	if err := context.ShouldBindJSON(&order); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Create(&order)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"orderId": order.ID, "customerName": order.CustomerName, "items": order.Items, "quantity": order.Quantity, "total_bayar": order.TotalBayar})
}
