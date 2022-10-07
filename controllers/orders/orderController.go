package orders

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

	router := r.Group("/orders")
	router.POST("/", h.CreateOrder)
	router.GET("/:orderId", h.GetOrder)
	router.GET("/", h.GetOrders)
	router.PUT("/:orderId", h.UpdateOrder)
	router.DELETE("/:orderId", h.DeleteOrder)
}
