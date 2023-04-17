package anggota

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	Register(input Inputan) (Anggota, error)
	Login(input InpLogin) (Anggota, error)
	GetAll() ([]Anggota, error)
	Update(id int, input Inputan) (Anggota, error)
}

type service struct {
	repository IRepository
}

func NewAnggotaService(repository IRepository) *service {
	return &service{repository}
}

func (s *service) Register(input Inputan) (Anggota, error) {
	res := Anggota{}

	//enkripsi password
	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if errHash != nil {
		return res, errHash
	}

	res.NamaLengkap = input.NamaLengkap
	res.Username = input.Username
	res.Password = string(passwordHash)
	res.NoHp = input.NoHp
	res.TanggalMasuk = time.Now()
	res.Status = input.Status

	res, err := s.repository.Insert(res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *service) Login(input InpLogin) (Anggota, error) {
	username := input.Username
	password := input.Password

	user, err := s.repository.ReadByUsername(username)
	if err != nil {
		return user, err
	}

	//cek jika user tidak ada
	if user.ID == 0 {
		return user, errors.New("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetAll() ([]Anggota, error) {
	res, err := s.repository.ReadAll()
	if err != nil {
		return []Anggota{}, err
	}

	return res, nil
}

func (s *service) Update(id int, input Inputan) (Anggota, error) {
	res, err := s.repository.ReadById(id)
	if err != nil {
		return res, err
	}
	if res.ID == 0 {
		return res, errors.New("ID not found")
	}

	res.ID = id
	res.NamaLengkap = input.NamaLengkap
	res.Username = input.Username
	res.Password = input.Password
	res.NoHp = input.NoHp
	res.TanggalMasuk = time.Now()
	res.Status = input.Status

	res, err = s.repository.Update(res)
	if err != nil {
		return res, err
	}

	return res, nil
}
