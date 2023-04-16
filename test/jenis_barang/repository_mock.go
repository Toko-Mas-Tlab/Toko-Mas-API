package jenisbarang

import (
	"errors"
	jenisbarang "toko_mas_api/domain/jenis_barang"
)

type IRepository interface {
	Insert(data jenisbarang.JenisBarang) (jenisbarang.JenisBarang, error)
	ReadAll(sort string) ([]jenisbarang.JenisBarang, error)
	ReadById(id int) (jenisbarang.JenisBarang, error)
	Update(data jenisbarang.JenisBarang) (jenisbarang.JenisBarang, error)
}

type MockRepository struct {
	data map[int]jenisbarang.JenisBarang
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		data: make(map[int]jenisbarang.JenisBarang),
	}
}

func (r *MockRepository) Insert(data jenisbarang.JenisBarang) (jenisbarang.JenisBarang, error) {
	if _, ok := r.data[data.ID]; ok {
		return jenisbarang.JenisBarang{}, errors.New("data already exists")
	}
	r.data[data.ID] = data
	return data, nil
}

func (r *MockRepository) ReadAll(sort string) ([]jenisbarang.JenisBarang, error) {
	jenis := make([]jenisbarang.JenisBarang, 0, len(r.data))
	for _, v := range r.data {
		jenis = append(jenis, v)
	}
	return jenis, nil
}

func (r *MockRepository) ReadById(id int) (jenisbarang.JenisBarang, error) {
	jenis, ok := r.data[id]
	if !ok {
		return jenisbarang.JenisBarang{}, errors.New("jenis not found")
	}
	return jenis, nil
}

func (r *MockRepository) Update(data jenisbarang.JenisBarang) (jenisbarang.JenisBarang, error) {
	if _, ok := r.data[data.ID]; !ok {
		return jenisbarang.JenisBarang{}, errors.New("data not found")
	}
	r.data[data.ID] = data
	return data, nil
}
