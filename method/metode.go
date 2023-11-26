package method

import (
	"fmt"
	"projek-pertama/function"
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

func ShowPeserta(pesertaList []model.Peserta) {
	for key, peserta := range pesertaList {
		fmt.Println("ID: ", key+1)
		fmt.Println("Nama: ", peserta.Nama)
		fmt.Println("Alamat: ", peserta.Alamat)
		fmt.Println("Pekerjaan: ", peserta.Pekerjaan)
		fmt.Println("Alasan: ", peserta.Pekerjaan)
	}
}

type CustomPeserta struct {
	model.Peserta
}

func (p CustomPeserta) FindPeserta() {
	var pesertaList = function.DataPeserta()

	for key := range pesertaList {
		if strings.ToLower(pesertaList[key].Nama) == strings.ToLower(p.Nama) {
			fmt.Println("ID : ", key+1)
			fmt.Println("Nama : ", pesertaList[key].Nama)
			fmt.Println("Alamat : ", pesertaList[key].Alamat)
			fmt.Println("Pekerjaan : ", pesertaList[key].Pekerjaan)
			fmt.Println("Alasan : ", pesertaList[key].Alasan)
		}
	}
}
