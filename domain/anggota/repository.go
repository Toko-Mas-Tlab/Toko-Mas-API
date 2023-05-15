package anggota

import "gorm.io/gorm"

type IRepository interface {
	Save(data Anggota) (Anggota, error)
	ReadAll() ([]Anggota, error)
	ReadById(id int) ([]Anggota, error)
	ReadByUsername(username string) (Anggota, error)
	Update(data Anggota) (Anggota, error)
	// Delete(data Anggota) (Anggota, error)
}

type repository struct {
	DB *gorm.DB
}

func NewAnggotaRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(data Anggota) (Anggota, error) {
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

func (r *repository) ReadById(id int) ([]Anggota, error) {
	var data []Anggota

	err := r.DB.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) ReadByUsername(username string) (Anggota, error) {
	var anggota Anggota

	err := r.DB.Where("username = ?", username).Find(&anggota).Error
	if err != nil {
		return Anggota{}, err
	}

	return anggota, nil
}

func (r *repository) Update(data Anggota) (Anggota, error) {
	err := r.DB.Save(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
