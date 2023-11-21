package function

import "golang-first-assignment/models"

func DataPeserta() []models.Peserta {
	pesertaList := []models.Peserta{
		{
			Nama:      "Ijal",
			Alamat:    "Malang",
			Pekerjaan: "Developer",
			Alasan:    "None",
		},
		{
			Nama:      "Tasya",
			Alamat:    "Jakarta",
			Pekerjaan: "Developer",
			Alasan:    "Menarik",
		},
	}

	return pesertaList
}
