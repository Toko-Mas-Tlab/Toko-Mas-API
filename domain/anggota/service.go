package anggota

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	Register(input Inputan) (Anggota, error)
	Login(input InputLogin) (Anggota, error)
	GetAll() ([]Anggota, error)
	// Update(id int, input Inputan) (Anggota, error)
}

type service struct {
	repository IRepository
}

func NewAnggotaService(repository IRepository) *service {
	return &service{repository}
}

func (s *service) Register(input Inputan) (Anggota, error) {
	user := Anggota{}

	user.NamaLengkap = input.NamaLengkap
	user.Username = input.Username

	//enkripsi password
	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if errHash != nil {
		return user, errHash
	}
	user.Password = string(passwordHash)
	user.NoHp = input.NoHp
	user.Status = input.Status

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input InputLogin) (Anggota, error) {
	username := input.Username
	password := input.Password

	user, err := s.repository.ReadByUsername(username)
	if err != nil {
		return user, err
	}

	//cek jika user tidak ada
	if user.ID == 0 {
		return user, errors.New("No user found")
	}

	//cek password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetAll() ([]Anggota, error) {
	user, err := s.repository.ReadAll()
	if err != nil {
		return []Anggota{}, err
	}

	return user, nil
}

// func (s *service) Update(id int, input Inputan) (Anggota, error) {
// 	user, err := s.repository.ReadById(id)
// 	if err != nil {
// 		return user, err
// 	}
// 	if user.ID == 0 {
// 		return user, errors.New("ID not found")
// 	}

// 	user.ID = id
// 	user.NamaLengkap = input.NamaLengkap
// 	user.Username = input.Username
// 	user.Password = input.Password
// 	user.NoHp = input.NoHp
// 	user.TanggalMasuk = time.Now()
// 	user.Status = input.Status

// 	user, err = s.repository.Update(user)
// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }
