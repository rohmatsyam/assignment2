package items

import (
	"assignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetItems(c *gin.Context) {
	var (
		item []models.Item
		result gin.H
	)	

	err := h.DB.Model(&item).Find(&item).Error
	if err != nil {
		result = gin.H{
			"result":"data not found",
			"count":0,
		}
		c.JSON(http.StatusNotFound,result)
		return
	}else{
		result = gin.H{
			"result":item,
			"count":len(item),
		}
	}
	c.JSON(http.StatusOK,result)
}