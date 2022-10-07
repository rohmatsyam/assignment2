package orders

import (
	"assignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteOrder(c *gin.Context) {
	id := c.Param("orderId")
	var (
		order  models.Order
		result gin.H
	)

	err := h.DB.First(&order, "id=?", id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	err = h.DB.Delete(&order).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	} else {
		result = gin.H{
			"result": "sucessfully deleted data",
		}
	}
	c.JSON(http.StatusOK, result)
}
