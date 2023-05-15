package daftarproduk

import (
	"errors"
)

type Service interface {
	Add(input ProdukInput) (DaftarProduk, error)
	GetAll() ([]DaftarProduk, error)
	Update(id int, input ProdukInput) (DaftarProduk, error)
}

type service struct {
	repository Repository
}

func NewDaftarProdukService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Add(input ProdukInput) (DaftarProduk, error) {
	data := DaftarProduk{}
	data.NamaBarang = input.NamaBarang
	data.Jenis = input.Jenis
	data.Bentuk = input.Bentuk

	res, err := s.repository.Insert(data)
	if err != nil {
		return data, err
	}

	return res, nil
}

func (s *service) GetAll() ([]DaftarProduk, error) {
	res, err := s.repository.ReadAll()
	if err != nil {
		return []DaftarProduk{}, err
	}

	return res, nil
}

func (s *service) Update(id int, input ProdukInput) (DaftarProduk, error) {
	res, err := s.repository.ReadById(id)
	if err != nil {
		return res, err
	}
	if res.ID == 0 {
		return res, errors.New("ID not found")
	}

	res.ID = id
	res.NamaBarang = input.NamaBarang
	res.Jenis = input.Jenis
	res.Bentuk = input.Bentuk
	res, err = s.repository.Update(res)
	if err != nil {
		return res, err
	}

	return res, nil
}
