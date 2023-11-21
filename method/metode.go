package method

import (
	"projek-pertama/model"
	"strings"
)

func GetPeserta(data []model.Peserta, nama string) (index int, biodata interface{}) {
	for i, v := range data {
		if strings.ToLower(nama) == strings.ToLower(v.Nama) {
			return i, v
		}
	}

	return 0, nil
}
