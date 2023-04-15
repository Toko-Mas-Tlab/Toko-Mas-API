package helper

import "strings"

func GenerateTypeCode(nama string) string {
	// Split kalimat menjadi array kata
	kata := strings.Split(nama, " ")

	// Looping setiap kata dan ambil inisial huruf depannya
	var inisial string
	for _, k := range kata {
		inisial += string(k[0])
	}

	// Ubah inisial menjadi huruf besar semua
	inisial = strings.ToUpper(inisial)

	return inisial
}
