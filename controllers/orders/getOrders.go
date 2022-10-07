package orders

import (
	"assignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func(h handler) GetOrders(c *gin.Context) {
	var 	result gin.H
	var orders []models.Order	
	err := h.DB.Model(&models.Order{}).Preload("Items").Find(&orders).Error
	if err != nil {
		result = gin.H{
			"result":"data not found",
			"count":0,
		}
		c.JSON(http.StatusNotFound,result)
		return
	}else{
		result = gin.H{
			"result":orders,
			"count":len(orders),
		}
	}
	c.JSON(http.StatusOK,result)
}