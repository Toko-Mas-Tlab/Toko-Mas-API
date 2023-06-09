package routes

import (
	"toko_mas_api/domain/anggota"
	"toko_mas_api/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func routeAnggota(DB *gorm.DB, r *gin.Engine) *gin.RouterGroup {
	anggotaRepo := anggota.NewAnggotaRepository(DB)
	anggotaService := anggota.NewAnggotaService(anggotaRepo)
	anggotaHandler := handlers.NewAnggotaHandler(anggotaService)

	route := r.Group("/anggota")
	{
		route.POST("", anggotaHandler.Register)
		// route.GET("", jBarangHandler.ListAnggota)
		// route.PUT("/:id_jenis_barang", jBarangHandler.UpdateAnggota)
	}

	return route
}
