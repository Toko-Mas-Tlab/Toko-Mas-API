package anggota

import "gorm.io/gorm"

type IRepository interface {
	Insert(data Anggota) (Anggota, error)
	ReadAll() ([]Anggota, error)
	ReadById(id int) (Anggota, error)
	Update(data Anggota) (Anggota, error)
	// Delete(data Anggota) (Anggota, error)
}

type repository struct {
	DB *gorm.DB
}

func NewAnggotaRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(data Anggota) (Anggota, error) {
	err := r.DB.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) ReadAll() ([]Anggota, error) {
	var data []Anggota

	err := r.DB.Order("id DESC").Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) ReadById(id int) (Anggota, error) {
	var data Anggota

	err := r.DB.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) Update(data Anggota) (Anggota, error) {
	err := r.DB.Save(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
