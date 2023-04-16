package jenisbarang

import "gorm.io/gorm"

type IRepository interface {
	Insert(data JenisBarang) (JenisBarang, error)
	ReadAll() ([]JenisBarang, error)
	ReadById(id int) (JenisBarang, error)
	Update(data JenisBarang) (JenisBarang, error)
	// Delete(data JenisBarang) (JenisBarang, error)
}

type repository struct {
	DB *gorm.DB
}

func NewJenisBarangRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(data JenisBarang) (JenisBarang, error) {
	err := r.DB.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) ReadAll() ([]JenisBarang, error) {
	var data []JenisBarang

	err := r.DB.Order("id DESC").Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) ReadById(id int) (JenisBarang, error) {
	var data JenisBarang

	err := r.DB.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) Update(data JenisBarang) (JenisBarang, error) {
	err := r.DB.Save(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
