package main

import (
	"fmt"
	"os"
	"projek-pertama/method"
	"projek-pertama/model"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Inputkan Nama")
		return
	}
	nama := args[1]

	peserta := method.CustomPeserta{
		Peserta: model.Peserta{
			Nama: nama,
		},
	}

	peserta.FindPeserta()
}
