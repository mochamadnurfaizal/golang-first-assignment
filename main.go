package main

import (
	"fmt"
	"golang-first-assignment/method"
	"golang-first-assignment/models"
	"os"
)

func main() {

	args := os.Args
	// var pesertaList []models.Peserta = function.DataPeserta()

	if len(args) < 2 {
		fmt.Println("Silahkan Inputkan Nama")
		return
	}
	nama := args[1]

	peserta := method.CustomPeserta{
		Peserta: models.Peserta{
			Nama: nama,
		},
	}

	peserta.FindPeserta()
	// method.PrintPeserta(pesertaList)

}
