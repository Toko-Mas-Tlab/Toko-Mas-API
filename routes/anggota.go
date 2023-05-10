package routes

import (
	"toko_mas_api/domain/anggota"
	"toko_mas_api/handlers"
	"toko_mas_api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func routeAnggota(DB *gorm.DB, r *gin.Engine) *gin.RouterGroup {
	// jwtService := middleware.NewService()

	anggotaRepo := anggota.NewAnggotaRepository(DB)
	anggotaService := anggota.NewAnggotaService(anggotaRepo)
	authService := middleware.NewService()
	anggotaHandler := handlers.NewAnggotaHandler(anggotaService, authService)

	router := gin.Default()
	// r.POST("/login", anggotaHandler.Login)
	route := router.Group("/anggota")
	{
		route.POST("", anggotaHandler.Register)
		// route.GET("", jBarangHandler.ListAnggota)
		// route.PUT("/:id_jenis_barang", jBarangHandler.UpdateAnggota)
	}

	return route
}
