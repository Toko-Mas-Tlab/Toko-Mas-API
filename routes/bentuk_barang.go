package routes

import (
	bentukbarang "toko_mas_api/domain/bentuk_barang"
	"toko_mas_api/handlers"
	"toko_mas_api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func routeBentukBarang(DB *gorm.DB, r *gin.Engine) *gin.RouterGroup {
	bBarangRepo := bentukbarang.NewBentukBarangRepository(DB)
	bBarangService := bentukbarang.NewBentukBarangService(bBarangRepo)
	bBarangHandler := handlers.NewBentukBarangHandler(bBarangService)

	route := r.Group("/bentuk-barang", middleware.AuthMiddleware())
	{
		route.POST("", bBarangHandler.AddNewShape)
		route.GET("", bBarangHandler.ListBentukBarang)
		route.PUT("/:id", bBarangHandler.UpdateBentukBarang)
	}

	return route
}
