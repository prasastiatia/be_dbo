package customers

import (
	database "be_dbo/database"
	"be_dbo/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindCustomerName(context *gin.Context) {
	var customer models.Customer

	name := context.Query("name")
	x := fmt.Sprintf
	query := x(" name LIKE '%%%s%%'", name)
	if err := database.Instance.Statement.DB.Where(query).Find(&customer).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": customer})
}
