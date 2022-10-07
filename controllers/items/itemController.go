package items

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	router := r.Group("/items")
	router.POST("/", h.CreateItem)
	router.GET("/:itemId", h.GetItem)
	router.GET("/", h.GetItems)
	router.PUT("/:itemId", h.UpdateItem)
	router.DELETE("/:itemId", h.DeleteItem)
}
