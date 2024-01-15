package routes

import (
	HolidaysController "api/internal/modules/holiday/controllers"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	holidaysController := HolidaysController.New()
	router.GET("/holidays", holidaysController.List)
	router.GET("/holidays/:id", holidaysController.Read)
	router.POST("/holidays", holidaysController.Create)
	router.PATCH("/holidays/:id", holidaysController.Update)
	router.DELETE("/holidays/:id", holidaysController.Delete)
}
