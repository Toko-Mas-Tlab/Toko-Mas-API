package main

import (
	"toko_mas_api/config"
	"toko_mas_api/database"

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

	database.AutoMigrate(db)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
