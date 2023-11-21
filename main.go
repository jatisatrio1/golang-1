package main

import (
	"bufio"
	"fmt"
	"os"
	"projek-pertama/function"
	"projek-pertama/method"
	"projek-pertama/model"
)

func main() {
	// Console inout
	fmt.Print("Search by Name: [ Enter Name ]\n")
	var name string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		name = scanner.Text()
	}

	// Geerate Data
	data := function.GenerateData

	// Search Data and Print
	idx, searchedData := method.GetPeserta(data(), name)

	if searchedData == nil {
		fmt.Println("Data Not Found")
		return
	}

	resultData := searchedData.(model.Peserta)

	fmt.Println("Index: ", idx)
	fmt.Println("Nama: ", resultData.Nama)
	fmt.Println("Alamat: ", resultData.Alamat)
	fmt.Println("Pekerjaan: ", resultData.Pekerjaan)
	fmt.Println("Alasan: ", resultData.Alasan)
}
