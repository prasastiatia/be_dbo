package Routes

import (
	controllers "be_dbo/controllers"
	customers "be_dbo/controllers/customers"
	"be_dbo/middlewares"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)

		}
		customer := api.Group("/customers").Use(middlewares.Auth())
		{
			customer.POST("/create", customers.CreateCustomer)
			customer.GET("/customers", customers.GetCustomers)
			customer.GET("/customer/:id", customers.FindCustomerId)
			customer.PATCH("/customer/:id", customers.UpdateCustomer)
			customer.DELETE("/customer/:id", customers.DeleteCustomer)

		}
	}
	return router
}
