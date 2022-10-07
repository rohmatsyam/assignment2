package items

import (
	"assignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetItem(c *gin.Context)  {
	var (
		item models.Item
		result gin.H
	)
	id := c.Param("itemId")

	err := h.DB.Model(&item).Where("id=?",id).First(&item).Error
	if err != nil {
		result = gin.H{
			"result":err.Error(),
			"count":0,
		}
		c.JSON(http.StatusNotFound,result)
		return
	}else{
		result = gin.H{
			"result":item,
			"count":1,
		}
	}
	c.JSON(http.StatusNotFound,result)
}