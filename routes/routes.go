package Routes

import (
	controllers "be_dbo/controllers"
	customers "be_dbo/controllers/customers"
	orders "be_dbo/controllers/orders"
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
			customer.GET("/customer/", customers.FindCustomerName)

		}
		order := api.Group("/orders").Use(middlewares.Auth())
		{
			order.POST("/create", orders.CreateOrder)
			order.GET("/orders", orders.GetOrders)
			order.GET("/order/:id", orders.FindOrderId)
			order.PATCH("/order/:id", orders.UpdateOrder)
			order.DELETE("/order/:id", orders.DeleteOrder)
		}
	}
	return router
}
