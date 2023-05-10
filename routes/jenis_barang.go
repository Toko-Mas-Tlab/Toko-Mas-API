package routes

import (
	jenisbarang "toko_mas_api/domain/jenis_barang"
	"toko_mas_api/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func routeJenisBarang(DB *gorm.DB, r *gin.Engine) *gin.RouterGroup {
	jBarangRepo := jenisbarang.NewJenisBarangRepository(DB)
	jBarangService := jenisbarang.NewJenisBarangService(jBarangRepo)
	jBarangHandler := handlers.NewJenisBarangHandler(jBarangService)

	route := r.Group("/jenis-barang")
	{
		route.POST("", jBarangHandler.AddNewType)
		// route.GET("", jBarangHandler.ListJenisBarang)
		// route.PUT("/:id_jenis_barang", jBarangHandler.UpdateJenisBarang)
	}

	return route
}
