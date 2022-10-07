package orders

import (
	"assignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customItem struct {
	ItemCode string
	Description string
	Quantity int
}

func (h handler) GetOrder(c *gin.Context) {
	id := c.Param("orderId")		
	var (
		order models.Order				
		result gin.H
	)
	err := h.DB.Model(&models.Order{}).Preload("Items").Where("id=?",id).First(&order).Error
	if err != nil {
		result = gin.H{
			"result":"data not found",
			"count":0,
		}
		c.JSON(http.StatusNotFound,result)
		return
	}	

	var items []customItem

	for _, itemOrders := range order.Items {
		detailStruct := customItem{
			ItemCode: itemOrders.ItemCode,
			Description: itemOrders.Description,
			Quantity: itemOrders.Quantity,
		}
		items = append(items,detailStruct)
	}

	result = gin.H{
		"customerName":order.CustomerName,
		"orderedAt":order.OrderedAt,
		"items":items,
	}
	
	c.JSON(http.StatusOK,result)
}