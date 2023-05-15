package daftarproduk

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Insert(produk DaftarProduk) (DaftarProduk, error)
	ReadAll() ([]DaftarProduk, error)
	ReadById(id int) (DaftarProduk, error)
	Update(data DaftarProduk) (DaftarProduk, error)
	ReadAllByDate(Date time.Time) ([]DaftarProduk, error)
}

type repository struct {
	DB *gorm.DB
}

func NewDaftarProdukRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(produk DaftarProduk) (DaftarProduk, error) {
	err := r.DB.Create(&produk).Error
	if err != nil {
		return produk, err
	}
	return produk, nil
}
func (r *repository) ReadAll() ([]DaftarProduk, error) {
	var data []DaftarProduk

	err := r.DB.Order("id DESC").Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) ReadById(id int) (DaftarProduk, error) {
	var data DaftarProduk

	err := r.DB.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) Update(data DaftarProduk) (DaftarProduk, error) {
	err := r.DB.Save(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
func (r *repository) ReadAllByDate(Date time.Time) ([]DaftarProduk, error) {
	var data []DaftarProduk

	err := r.DB.Order("id DESC").Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
