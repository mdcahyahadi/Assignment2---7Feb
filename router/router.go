package router

import (
	"assignment-golang8-7feb/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// user
	router.POST("/order/new-order", controllers.CreateOrder)
	router.GET("/order", controllers.GetAllOrder)
	router.GET("/order/:id", controllers.GetOrderByID)
	router.PUT("/order/:id", controllers.UpdateOrder)
	router.DELETE("/order/:id", controllers.DeleteOrder)

	// item
	router.POST("/item/add-item", controllers.CreateItem)
	router.GET("/order/detail/:id", controllers.GetAllItemByOrderID)
	router.GET("/item/:id", controllers.GetItemByID)
	router.PUT("/item/:id", controllers.UpdateItem)
	router.DELETE("/item/:id", controllers.DeleteItem)

	return router
}
