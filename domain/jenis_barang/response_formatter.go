package jenisbarang

import (
	"toko_mas_api/helper"
)

type jenisbarangFormatter struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Kode      string `json:"kode"`
	CreatedAt string `json:"tgl_dibuat"`
}

func ResponseFormatter(data []JenisBarang) []jenisbarangFormatter {
	var format jenisbarangFormatter
	var result []jenisbarangFormatter

	for _, v := range data {
		tglBuat := helper.DateParseToBahasa(v.CreatedAt)

		format.ID = v.ID
		format.Nama = v.Nama
		format.Kode = v.Kode
		format.CreatedAt = tglBuat

		result = append(result, format)
	}

	return result
}
