package main

import (
	"fmt"
	"tubes/lib"
)

func main() {
	var hitung int
	var opsi int
	var lokasiAtauNama string
	var fasilitas string
	var tinggiRendah string
	var rating string
	var i int
	var lokasi string
	lib.InitDummyData(&lib.Data, &hitung)
	for {
		fmt.Println("===============================================================================")
		fmt.Println("|         Aplikasi Manajemen dan Review Co-Working Space                     |")
		fmt.Println("|                                MENU                                        |")
		fmt.Println("|----------------------------------------------------------------------------|")
		fmt.Println("|  1. Tampilkan List Data                                                    |")
		fmt.Println("|  2. Tambah Data                                                            |")
		fmt.Println("|  3. Hapus Data                                                             |")
		fmt.Println("|  4. Update Data                                                            |")
		fmt.Println("|  5. Cari Data Berdasarkan Nama (Sequential Search)                         |")
		fmt.Println("|  6. Cari Data Berdasarkan Lokasi (Binary Search)                           |")
		fmt.Println("|  7. Cari Data Berdasarkan Fasilitas (Sequential Search)                    |")
		fmt.Println("|  8. Urutkan Data Berdasarkan Harga Sewa (Selection Sort)                   |")
		fmt.Println("|  9. Urutkan Data Berdasarkan Rating (Insertion Sort)                       |")
		fmt.Println("| 10. Keluar                                                                 |")
		fmt.Println("===============================================================================")
		fmt.Print("Pilih menu: ")
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
			fmt.Print("Masukan Nama Co Working yang ingin anda cari : ")
			fmt.Scanln(&lokasiAtauNama)
			lib.MenampilkanDataSearchByNama(lib.Data, hitung, lokasiAtauNama)
		case 6:
			fmt.Print("Masukkan lokasi Co working yang ingin dicari: ")
			fmt.Scanln(&lokasi)
			lib.SelectionSortByLokasiAsc(&lib.Data, hitung)
			fmt.Println("\n=== DATA DITEMUKAN ===")
			for i = 0; i < hitung; i++ {
				if lib.Data[i].Lokasi == lokasi {
					lib.TampilDataDarilokasi(lib.Data, i)
				}
			}
		case 7:
			fmt.Print("Masukkan nama fasilitas yang dicari: ")
			fmt.Scanln(&fasilitas)
			lib.ListDataFasilits(lib.Data, hitung, fasilitas)
		case 8:
			fmt.Print("Pilih urutan harga Sewa ( tertinggi atau terendah ): ")
			fmt.Scan(&tinggiRendah)
			if tinggiRendah == "tertinggi" {
				lib.SelectionSortHargaDes(&lib.Data, hitung)
				lib.TampilData(lib.Data, hitung)
			} else if tinggiRendah == "terendah" {
				lib.SelectionSortHargaAsc(&lib.Data, hitung)
				lib.TampilData(lib.Data, hitung)
			}
		case 9:
			fmt.Print("Pilih urutan rating ( tertinggi atau terendah ): ")
			fmt.Scan(&rating)
			if rating == "tertinggi" {
				lib.InsertionSortRatingDesc(&lib.Data, hitung)
				lib.TampilData(lib.Data, hitung)
			} else if rating == "terendah" {
				lib.InsertionSortRatingAsc(&lib.Data, hitung)
				lib.TampilData(lib.Data, hitung)
			}
		case 10:
			fmt.Println("Selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}

}
