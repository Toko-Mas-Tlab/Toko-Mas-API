package main

import (
	"toko_mas_api/config"
	"toko_mas_api/domain/anggota"
	"toko_mas_api/handlers"
	"toko_mas_api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := config.ConnectionDB()
	if err != nil {
		panic(err)
	}
	DB = db

	// err = db.AutoMigrate(&jenisbarang.JenisBarang{}, &anggota.Anggota{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Migration OK!")
}

func main() {
	config.ConnectionDB()
	anggotaRepo := anggota.NewAnggotaRepository(DB)
	anggotaService := anggota.NewAnggotaService(anggotaRepo)
	authService := middleware.NewService()
	anggotaHandler := handlers.NewAnggotaHandler(anggotaService, authService)

	router := gin.Default()
	api := router.Group("/v1/anggota")
	api.POST("register", anggotaHandler.Register)
	api.POST("login", anggotaHandler.Login)

	// routes.Routes(DB, r)

	router.Run() // listen and serve on 0.0.0.0:8080

}
