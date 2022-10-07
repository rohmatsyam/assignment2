package items

import (
	"assignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteItem(c *gin.Context) {
	id := c.Param("itemId")
	var (
		item   models.Item
		result gin.H
	)

	err := h.DB.First(&item, "id=?", id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	err = h.DB.Delete(&item).Error
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
