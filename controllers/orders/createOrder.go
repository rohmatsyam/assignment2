package orders

import (
	"assignment2/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// POST
func (h handler) CreateOrder(c *gin.Context)  {	
	var (
		order models.Order
		result gin.H
	)

	customer_name := c.PostForm("customer_name")

	order.CustomerName = customer_name
	order.OrderedAt = time.Now()

	err := h.DB.Create(&order).Error
	if err != nil {
		result = gin.H{
			"result":err.Error(),
		}
		c.JSON(http.StatusInternalServerError,result)
		return
	}else{
		result = gin.H{
			"result":order,
		}
	}
	c.JSON(http.StatusOK,result)
}