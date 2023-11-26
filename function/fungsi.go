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

func DataPeserta() []model.Peserta {
	pesertas := []model.Peserta{
		{
			Nama:      "Satrio",
			Alamat:    "London",
			Pekerjaan: "Dev",
			Alasan:    "Mantap",
		},
		{
			Nama:      "Mark",
			Alamat:    "Berlin",
			Pekerjaan: "QA",
			Alasan:    "Rahasia",
		},
	}

	return pesertas
}
