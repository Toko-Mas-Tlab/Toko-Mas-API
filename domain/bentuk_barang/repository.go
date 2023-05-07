package bentukbarang

import "gorm.io/gorm"

type IRepository interface {
	Insert(data BentukBarang) (BentukBarang, error)
	ReadAll(sort string) ([]BentukBarang, error)
	ReadById(id int) (BentukBarang, error)
	Update(data BentukBarang) (BentukBarang, error)
	// Delete(data JenisBarang) (JenisBarang, error)
}

type repository struct {
	DB *gorm.DB
}

func NewBentukBarangRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(data BentukBarang) (BentukBarang, error) {
	err := r.DB.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) ReadAll(sort string) ([]BentukBarang, error) {
	var data []BentukBarang

	// err := r.DB.Order("id DESC").Find(&data).Error
	sql := "SELECT * FROM bentuk_barangs ORDER BY created_at " + sort
	err := r.DB.Raw(sql).Scan(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) ReadById(id int) (BentukBarang, error) {
	var data BentukBarang

	err := r.DB.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) Update(data BentukBarang) (BentukBarang, error) {
	err := r.DB.Save(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
