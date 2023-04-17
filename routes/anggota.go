package routes

import (
	"toko_mas_api/domain/anggota"
	"toko_mas_api/handlers"
	"toko_mas_api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func routeAnggota(DB *gorm.DB, r *gin.Engine) *gin.RouterGroup {
	jwtService := middleware.NewService()

	anggotaRepo := anggota.NewAnggotaRepository(DB)
	anggotaService := anggota.NewAnggotaService(anggotaRepo)
	anggotaHandler := handlers.NewAnggotaHandler(anggotaService, jwtService)

	r.POST("/login", anggotaHandler.Login)
	route := r.Group("/anggota")
	{
		route.POST("", anggotaHandler.Register)
		route.POST("/login", anggotaHandler.Login)
		// route.GET("", jBarangHandler.ListAnggota)
		// route.PUT("/:id_jenis_barang", jBarangHandler.UpdateAnggota)
	}

	return route
}
