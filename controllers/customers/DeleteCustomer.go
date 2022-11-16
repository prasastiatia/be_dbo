package customers

import (
	database "be_dbo/database"
	"be_dbo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCustomer(context *gin.Context) {
	var customer models.Customer

	if err := database.Instance.Statement.DB.Where("id = ? ", context.Param("id")).First(&customer).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.Instance.Statement.DB.Delete(&customer)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
