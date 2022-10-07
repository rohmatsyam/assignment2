package orders

import (
	"assignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) UpdateOrder(c *gin.Context) {
	var (
		order    models.Order
		newOrder models.Order
		result   gin.H
	)
	id := c.Param("orderId")

	err := h.DB.Model(&order).Where("id=?", id).First(&order).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"count":  0,
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	customer_name := c.PostForm("customer_name")
	newOrder.CustomerName = customer_name

	err = h.DB.Model(&order).Updates(&newOrder).Error
	if err != nil {
		result = gin.H{
			"result":  err.Error(),
			"message": "updated failed",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	var items []customItem
	for _, itemOrders := range order.Items {
		detailStruct := customItem{
			ItemCode:    itemOrders.ItemCode,
			Description: itemOrders.Description,
			Quantity:    itemOrders.Quantity,
		}
		items = append(items, detailStruct)
	}
	result = gin.H{
		"customerName": order.CustomerName,
		"orderedAt":    order.OrderedAt,
		"items":        items,
	}
	c.JSON(http.StatusOK, result)
}
