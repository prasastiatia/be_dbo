package customers

import (
	database "be_dbo/database"
	"be_dbo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCustomer(context *gin.Context) {
	var customer models.Customer
	if err := database.Instance.Statement.DB.Where("id = ? ", context.Param("id")).First(&customer).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := context.ShouldBindJSON(&customer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Instance.Model(&customer).Updates(customer)
	context.JSON(http.StatusOK, gin.H{"data": customer})
}
