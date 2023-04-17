package jenisbarang

import (
	"testing"
	"time"
	jenisbarang "toko_mas_api/domain/jenis_barang"
)

func TestAdd(t *testing.T) {
	mockRepository := NewMockRepository()
	userService := jenisbarang.NewJenisBarangService(mockRepository)

	inp := jenisbarang.Inputan{Nama: "Perak Logam Mulia"}
	_, err := userService.Add(inp)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	data, err := userService.GetAll("desc")
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if len(data) != 1 {
		t.Errorf("expected data count to be 1, got %d", len(data))
	}
	if data[0].Nama != "Perak Logam Mulia" {
		t.Errorf("expected nama to be Perak Logam Mulia, got %s", data[0].Nama)
	}
}

func TestGetAll(t *testing.T) {
	mockRepository := NewMockRepository()
	mockRepository.data[1] = jenisbarang.JenisBarang{ID: 1, Nama: "Perhiasan Emas", Kode: "PE", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	mockRepository.data[2] = jenisbarang.JenisBarang{ID: 2, Nama: "Perhiasan Perak", Kode: "PP", CreatedAt: time.Now(), UpdatedAt: time.Now()}

	service := jenisbarang.NewJenisBarangService(mockRepository)

	data, err := service.GetAll("desc")
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if len(data) != 2 {
		t.Errorf("expected 2 data, got %d", len(data))
	}
	if data[0].Kode != "PE" {
		t.Errorf("expected first user nama to be John, got %s", data[0].Kode)
	}
	if data[1].Kode != "PP" {
		t.Errorf("expected second user name to be Jane, got %s", data[1].Kode)
	}
}

// func TestUpdate(t *testing.T) {
// 	mockRepository := NewMockRepository()
// 	user := jenisbarang.JenisBarang{
// 		ID: 1,
// 		Nama: "Perhiasan Emas",
// 		Kode: "PE",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// 	mockRepository.data[1] = user

// 	service := jenisbarang.NewJenisBarangService(mockRepository)

// 	// user.Name = "Johnny"
// 	inp := jenisbarang.Inputan{Nama: "Perhiasan Titanium"}
// 	_,err := service.Update(1, inp)
// 	if err != nil {
// 		t.Error("unexpected error:", err)
// 	}

// 	updatedUser, err := service.GetAll("desc")
// 	if err != nil {
// 		t.Error("unexpected error:", err)
// 	}
// 	if updatedUser[0].Nama != "Perhiasan Titanium" {
// 		t.Errorf("expected user name to be Johnny, got %s", updatedUser[0].Nama)
// 	}
// }
