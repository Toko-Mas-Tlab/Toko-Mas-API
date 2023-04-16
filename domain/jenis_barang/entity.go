package jenisbarang

import "time"

type JenisBarang struct {
	ID        int    `gorm:"primary_key;auto_increment;not_null"`
	Nama      string `gorm:"type:varchar(100);not null"`
	Kode      string `gorm:"type:varchar(10);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Inputan struct {
	Nama string `json:"nama" binding:"required"`
}
