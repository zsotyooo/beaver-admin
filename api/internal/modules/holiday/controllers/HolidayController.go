package controllers

import (
	HolidayRequest "api/internal/modules/holiday/requests"
	HolidayService "api/internal/modules/holiday/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HolidaysController struct {
	holidayService HolidayService.HolidayServiceInterface
}

func New() *HolidaysController {
	return &HolidaysController{
		holidayService: HolidayService.New(),
	}
}

func (controller *HolidaysController) List(c *gin.Context) {
	holidays, err := controller.holidayService.GetHolidays(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, holidays)
}

func (controller *HolidaysController) Create(c *gin.Context) {
	var payload HolidayRequest.HolidayCreatePayload
	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	holiday, err := controller.holidayService.CreateHoliday(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, holiday)
}

func (controller *HolidaysController) Read(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	holiday, err := controller.holidayService.FindHoliday(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, holiday)
}

func (controller *HolidaysController) Update(c *gin.Context) {
	var payload HolidayRequest.HolidayUpdatePayload
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedHoliday, err := controller.holidayService.UpdateHoliday(uint(id), payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, updatedHoliday)
}

func (controller *HolidaysController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	controller.holidayService.DeleteHoliday(uint(id))

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Holiday deleted successfully!"})
}
