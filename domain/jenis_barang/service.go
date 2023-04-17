package jenisbarang

import (
	"errors"
	"toko_mas_api/helper"
)

type IService interface {
	Add(input Inputan) (JenisBarang, error)
	GetAll(sort string) ([]JenisBarang, error)
	Update(id int, input Inputan) (JenisBarang, error)
}

type service struct {
	repository IRepository
}

func NewJenisBarangService(repository IRepository) *service {
	return &service{repository}
}

func (s *service) Add(input Inputan) (JenisBarang, error) {
	data := JenisBarang{}

	inisial := helper.GenerateTypeCode(input.Nama)

	data.Nama = input.Nama
	data.Kode = inisial

	res, err := s.repository.Insert(data)
	if err != nil {
		return data, err
	}

	return res, nil
}

func (s *service) GetAll(sort string) ([]JenisBarang, error) {
	res, err := s.repository.ReadAll(sort)
	if err != nil {
		return []JenisBarang{}, err
	}

	return res, nil
}

func (s *service) Update(id int, input Inputan) (JenisBarang, error) {
	res, err := s.repository.ReadById(id)
	if err != nil {
		return res, err
	}
	if res.ID == 0 {
		return res, errors.New("ID not found")
	}

	inisial := helper.GenerateTypeCode(input.Nama)

	res.ID = id
	res.Nama = input.Nama
	res.Kode = inisial
	res, err = s.repository.Update(res)
	if err != nil {
		return res, err
	}

	return res, nil
}
