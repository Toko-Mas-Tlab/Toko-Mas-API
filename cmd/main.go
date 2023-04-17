package main

import (
	"fmt"
	"toko_mas_api/config"
	"toko_mas_api/domain/anggota"
	jenisbarang "toko_mas_api/domain/jenis_barang"
	"toko_mas_api/routes"

	// jenisbarang "toko_mas_api/domain/jenis_barang"
	// "toko_mas_api/routes"

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

	err = db.AutoMigrate(&jenisbarang.JenisBarang{}, &anggota.Anggota{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Migration OK!")
}

func main() {
	r := gin.Default()

	routes.Routes(DB, r)

	r.Run() // listen and serve on 0.0.0.0:8080

	// tokenString, err := anggota.Login("admin", "password123")
	// if err != nil {
	// 	fmt.Println("Login gagal:", err)
	// 	return
	// }

	// Cetak token JWT
	// fmt.Println("Token JWT:", tokenString)

}
