package items

import (
	"assignment2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h handler) UpdateItem(c *gin.Context) {
	var (
		item    models.Item
		newItem models.Item
		result  gin.H
	)
	id := c.Param("itemId")

	err := h.DB.Model(&item).Where("id=?", id).First(&item).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"count":  0,
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	item_code := c.PostForm("item_code")
	description := c.PostForm("description")
	quantity := c.PostForm("quantity")
	order_refer := c.PostForm("order_refer")

	quantityInt, err := strconv.Atoi(quantity)
	order_refer_int, err := strconv.Atoi(order_refer)

	newItem.ItemCode = item_code
	newItem.Description = description
	newItem.Quantity = quantityInt
	newItem.OrderRefer = order_refer_int

	err = h.DB.Model(&item).Updates(&newItem).Error
	if err != nil {
		result = gin.H{
			"result":  err.Error(),
			"message": "updated failed",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	result = gin.H{
		"result": "sucessfully updated data",
	}
	c.JSON(http.StatusOK, result)
}
