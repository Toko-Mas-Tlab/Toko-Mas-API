package anggota

import (
	"time"
)

type Anggota struct {
	ID           int       `gorm:"primary_key;auto_increment;not null"`
	NamaLengkap  string    `gorm:"type:varchar(255);not null"`
	Username     string    `gorm:"type:varchar(50);not null;unique"`
	Password     string    `gorm:"type:varchar(255)"`
	NoHp         string    `gorm:"type:varchar(20)"`
	TanggalMasuk time.Time `gorm:"type:timestamp;not null"`
	Status       string    `sql:"type:ENUM('Aktif', 'No-Aktif')" gorm:"column:status_type;not null"`
	CreatedAt    time.Time
	CreatedBy    string `gorm:"type:varchar(50)"`
	UpdatedAt    time.Time
	UpdateBy     string `gorm:"type:varchar(50)"`
}

type Inputan struct {
	NamaLengkap string `json:"nama_lengkap" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	NoHp        string `json:"no_hp" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type InputLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
