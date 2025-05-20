package main

import "fmt"

func main() {

	var hitung int
	var opsi int

	for {
		fmt.Println("==============================================")
		fmt.Println("|                   Menu                     |")
		fmt.Println("|        1. Tampilkan list data              |")
		fmt.Println("|        2. Tambah Data                      |")
		fmt.Println("|        3. Hapus Data                       |")
		fmt.Println("|        4. Update Data                      |")
		fmt.Println("|        5. Keluar                           |")
		fmt.Println("==============================================")
		fmt.Println("           Pilih menu:                          ")
		fmt.Scanln(&opsi)

		switch opsi {
		case 1:
			tampilData(data, hitung)
		case 2:
			tambahData(&data, &hitung)
		case 3:
			hapusData(&data, &hitung)
		case 4:
			updateData(&data, hitung)
		case 5:
			fmt.Println("Selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
