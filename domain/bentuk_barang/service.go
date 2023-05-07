package bentukbarang

import (
	"errors"
	"toko_mas_api/helper"
)

type IService interface {
	Add(input Inputan, userID int) (BentukBarang, error)
	GetAll(sort string) ([]BentukBarang, error)
	Update(id int, input Inputan) (BentukBarang, error)
}

type service struct {
	repository IRepository
}

func NewBentukBarangService(repository IRepository) *service {
	return &service{repository}
}

func (s *service) Add(input Inputan, userID int) (BentukBarang, error) {
	data := BentukBarang{}

	inisial := helper.GenerateShapeCode(input.Nama)

	data.Nama = input.Nama
	data.Kode = inisial
	data.CreateBy = uint(userID)
	data.EditBy = data.CreateBy

	res, err := s.repository.Insert(data)
	if err != nil {
		return data, err
	}

	return res, nil
}

func (s *service) GetAll(sort string) ([]BentukBarang, error) {
	res, err := s.repository.ReadAll(sort)
	if err != nil {
		return []BentukBarang{}, err
	}

	return res, nil
}

func (s *service) Update(id int, input Inputan) (BentukBarang, error) {
	res, err := s.repository.ReadById(id)
	if err != nil {
		return res, err
	}
	if res.ID == 0 {
		return res, errors.New("ID not found")
	}

	inisial := helper.GenerateShapeCode(input.Nama)

	res.ID = id
	res.Nama = input.Nama
	res.Kode = inisial
	res, err = s.repository.Update(res)
	if err != nil {
		return res, err
	}

	return res, nil
}
