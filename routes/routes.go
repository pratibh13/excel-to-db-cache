package routes

import (
	"assignment/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/import", controllers.ImportExcel)
	router.GET("/records", controllers.ViewRecords)
	router.PUT("/records/:id", controllers.UpdateRecord)
	return router
}
