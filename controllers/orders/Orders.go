package orders

import (
	model "be_dbo/models"
	"be_dbo/repository"
	util "be_dbo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(context *gin.Context) {
	pagination := util.GeneratePaginationFromRequest(context)
	var order model.Order
	orderLists, err := repository.GetAllOrders(&order, &pagination)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	context.JSON(http.StatusOK, gin.H{
		"data": orderLists,
	})
}
