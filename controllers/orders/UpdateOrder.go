package orders

import (
	database "be_dbo/database"
	"be_dbo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOrder(context *gin.Context) {
	var order models.Order
	if err := database.Instance.Statement.DB.Where("id = ? ", context.Param("id")).First(&order).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := context.ShouldBindJSON(&order); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Instance.Model(&order).Updates(order)
	context.JSON(http.StatusOK, gin.H{"data": order})
}
