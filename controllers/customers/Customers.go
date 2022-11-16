package customers

import (
	util "be_dbo/Utils"
	model "be_dbo/models"
	"be_dbo/repository"
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
