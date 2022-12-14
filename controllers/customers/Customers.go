package customers

import (
	model "be_dbo/models"
	"be_dbo/repository"
	util "be_dbo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomers(context *gin.Context) {
	pagination := util.GeneratePaginationFromRequest(context)
	var customer model.Customer
	customerLists, err := repository.GetAllCustomer(&customer, &pagination)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	context.JSON(http.StatusOK, gin.H{
		"data": customerLists,
	})
}
