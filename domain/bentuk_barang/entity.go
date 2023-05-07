package bentukbarang

import (
	"time"
	"toko_mas_api/domain/anggota"
)

type BentukBarang struct {
	ID        int             `gorm:"primary_key;auto_increment;not null"`
	Nama      string          `gorm:"type:varchar(100);not null"`
	Kode      string          `gorm:"type:varchar(10);not null"`
	CreateBy  uint            `gorm:"not null"`
	Creator   anggota.Anggota `gorm:"foreignKey:CreateBy"`
	EditBy    uint            `gorm:"not null"`
	Editor    anggota.Anggota `gorm:"foreignKey:EditBy"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Inputan struct {
	Nama string `json:"nama" binding:"required"`
}
