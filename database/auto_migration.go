package database

import (
	"log"
	jenisbarang "toko_mas_api/domain/jenis_barang"

	"gorm.io/gorm"
)

type model struct {
	Model interface{}
}

func AutoMigrate(db *gorm.DB) {
	//DB.AutoMigrate(&user.User{}, &item.Item{})
	for _, model := range registerModel() {
		errModel := db.Debug().AutoMigrate(model)

		if errModel != nil {
			log.Fatal(errModel)
			panic(errModel)
		}
	}
}

func registerModel() []model {
	return []model{
		{Model: jenisbarang.JenisBarang{}},
	}
}
