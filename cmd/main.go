package main

import (
	"toko_mas_api/config"
	"toko_mas_api/domain/anggota"
	daftarproduk "toko_mas_api/domain/daftar_produk"
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

	// err = db.AutoMigrate(&daftarproduk.DaftarProduk{})
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

	//produk
	produkRepo := daftarproduk.NewDaftarProdukRepository(DB)
	produkService := daftarproduk.NewDaftarProdukService(produkRepo)
	produkHandler := handlers.NewDaftarProdukHandler(produkService)

	router := gin.Default()
	//router anggota
	api := router.Group("/v1/anggota")
	api.POST("register", anggotaHandler.Register)
	api.POST("login", anggotaHandler.Login)
	api.GET("list", anggotaHandler.ListAnggota)

	//router produk
	api = router.Group("/v1/produk")
	api.GET("list", produkHandler.ListProduk)
	api.POST("add", produkHandler.AddProduk)

	// routes.Routes(DB, router)

	router.Run() // listen and serve on 0.0.0.0:8080

}
