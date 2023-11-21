package method

import (
	"fmt"
	"golang-first-assignment/function"
	"golang-first-assignment/models"
	"strings"
)

func PrintPeserta(listPeserta []models.Peserta) {
	for key, peserta := range listPeserta {
		fmt.Println("ID : ", key+1)
		fmt.Println("Nama : ", peserta.Nama)
		fmt.Println("Alamat : ", peserta.Alamat)
		fmt.Println("Pekerjaan : ", peserta.Pekerjaan)
		fmt.Println("Alasan : ", peserta.Alasan)
	}

}

type CustomPeserta struct {
	models.Peserta
}

func (p CustomPeserta) FindPeserta() {
	var listPeserta = function.DataPeserta()

	for key := range listPeserta {
		if strings.ToLower(listPeserta[key].Nama) == strings.ToLower(p.Nama) {
			fmt.Println("ID : ", key+1)
			fmt.Println("Nama : ", listPeserta[key].Nama)
			fmt.Println("Alamat : ", listPeserta[key].Alamat)
			fmt.Println("Pekerjaan : ", listPeserta[key].Pekerjaan)
			fmt.Println("Alasan : ", listPeserta[key].Alasan)
		}
	}
}
