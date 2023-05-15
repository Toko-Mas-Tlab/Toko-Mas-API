package daftarproduk

import "time"

type DaftarProduk struct {
	ID         int `json:"id"`
	Tanggal    time.Time
	NamaBarang string `json:"nama_barang"`
	Jenis      string `json:"jenis"`
	Bentuk     string `json:"bentuk"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ProdukInput struct {
	NamaBarang string `json:"nama_barang"`
	Jenis      string `json:"jenis"`
	Bentuk     string `json:"bentuk"`
}
