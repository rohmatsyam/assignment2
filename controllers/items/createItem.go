package items

import (
	"assignment2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h handler) CreateItem(c *gin.Context) {
	var (
		item   models.Item
		result gin.H
	)

	item_code := c.PostForm("item_code")
	description := c.PostForm("description")
	quantity := c.PostForm("quantity")
	order_refer := c.PostForm("order_refer")

	quantityInt, err := strconv.Atoi(quantity)
	order_refer_int, err := strconv.Atoi(order_refer)

	item.ItemCode = item_code
	item.Description = description
	item.Quantity = quantityInt
	item.OrderRefer = order_refer_int

	err = h.DB.Create(&item).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	} else {
		result = gin.H{
			"result": item,
		}
	}
	c.JSON(http.StatusOK, result)
}
