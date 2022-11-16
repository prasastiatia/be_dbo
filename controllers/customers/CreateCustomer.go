package customers

import (
	database "be_dbo/database"
	"be_dbo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(context *gin.Context) {
	var customer models.Customer
	if err := context.ShouldBindJSON(&customer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Create(&customer)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"customerId": customer.ID, "tanggal_lahir": customer.TanggalLahir, "mobile": customer.Mobile})
}
