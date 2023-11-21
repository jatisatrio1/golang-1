package function

import (
	"projek-pertama/model"
	"strings"
)

type PopulatedData struct {
	Datas []model.Peserta
}

func (p PopulatedData) GenerateData(nama string) (idx int, biodata interface{}) {
	for i, v := range p.Datas {
		if strings.ToLower(nama) == strings.ToLower(v.Nama) {
			return i, v
		}
	}

	return 0, nil
}
