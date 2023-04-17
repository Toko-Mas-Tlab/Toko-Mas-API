package anggota

import (
	"time"
)

type Anggota struct {
	ID           int       `gorm:"primary_key;auto_increment;not_null"`
	NamaLengkap  string    `gorm:"type:varchar(255);not null"`
	Username     string    `gorm:"type:varchar(50);not null;unique"`
	Password     string    `gorm:"type:varchar(255)"`
	NoHp         string    `gorm:"type:varchar(20)"`
	TanggalMasuk time.Time `gorm:"type:timestamp"`
	Status       string    `sql:"type:ENUM('Aktif', 'No-Aktif')" gorm:"column:status_type"`
	CreatedAt    time.Time `gorm:"type:timestamp"`
	CreatedBy    string    `gorm:"type:varchar(50)"`
	UpdatedAt    time.Time `gorm:"type:timestamp"`
	UpdateBy     string    `gorm:"type:varchar(50)"`
}

type Inputan struct {
	NamaLengkap  string    `json:"nama_lengkap" binding:"required"`
	Username     string    `json:"username" binding:"required"`
	Password     string    `json:"password" binding:"required"`
	NoHp         string    `json:"no_hp" binding:"required"`
	TanggalMasuk time.Time `json:"tanggal_masuk" binding:"required"`
	Status       string    `json:"status" binding:"required"`
}
