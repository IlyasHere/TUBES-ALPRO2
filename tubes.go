package main

import (
	"fmt"
	"tubes/lib"
)

func main() {
	var hitung int
	var opsi int
	var lokasi string
	var fasilitas string
	for {
		fmt.Println("=========================================================")
		fmt.Println("|                   Menu                     			 |")
		fmt.Println("|        1. Tampilkan list data              			 |")
		fmt.Println("|        2. Tambah Data                                 |")
		fmt.Println("|        3. Hapus Data                                  |")
		fmt.Println("|        4. Update Data                                 |")
		fmt.Println("|        5. Cari Data Berdasarkan Lokasi (Binary Search)|")
		fmt.Println("|        6. Cari Data Berdasarkan Fasilitas(Sequential Search)|")
		fmt.Println("|        7. Keluar                                      |")
		fmt.Println("=========================================================")
		fmt.Println("           Pilih menu:                                   ")
		fmt.Scanln(&opsi)

		switch opsi {
		case 1:
			lib.TampilData(lib.Data, hitung)
		case 2:
			lib.TambahData(&lib.Data, &hitung)
		case 3:
			lib.HapusData(&lib.Data, &hitung)
		case 4:
			lib.UpdateData(&lib.Data, hitung)
		case 5:
			fmt.Print("Masukkan lokasi yang ingin dicari: ")
			fmt.Scanln(&lokasi)
			lib.MenampilkanDataSearch(lib.Data, hitung, lokasi)
		case 6:
			fmt.Print("Masukkan nama fasilitas yang dicari: ")
			fmt.Scanln(&fasilitas)
			lib.ListDataFasilits(lib.Data, hitung, fasilitas)
		case 7:
			fmt.Println("Selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
