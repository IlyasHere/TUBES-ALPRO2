package lib

import "fmt"

const NMAX = 100

type WorkSpace struct {
	nama      string
	lokasi    string
	fasilitas [3]string
	jumlahF   int
	hargaSewa float64
	ulasan    string
	rating    string
}

var Data [NMAX]WorkSpace

func TambahData(A *[NMAX]WorkSpace, ValueArray *int) {
	var i int
	if *ValueArray >= NMAX {
		fmt.Println("Data sudah penuh.")
		return
	}

	var temp WorkSpace

	fmt.Print("Nama: ")
	fmt.Scanln(&temp.nama)
	fmt.Print("Lokasi: ")
	fmt.Scanln(&temp.lokasi)
	fmt.Print("Jumlah fasilitas(max 3): ")
	fmt.Scanln(&temp.jumlahF)

	if temp.jumlahF > 3 {
		temp.jumlahF = 3
	}

	for i = 0; i < temp.jumlahF; i++ {
		fmt.Printf("Fasilitas ke-%d: ", i+1)
		fmt.Scanln(&temp.fasilitas[i])
	}

	fmt.Print("Harga sewa: ")
	fmt.Scanln(&temp.hargaSewa)

	fmt.Print("Rating pengguna")
	fmt.Scanln(&temp.rating)
	fmt.Print("Ulasan singkat: ")
	fmt.Scanf("%s\n", &temp.ulasan)

	A[*ValueArray] = temp
	*ValueArray++

	fmt.Println("Data berhasil ditambahkan.")
}

func TampilData(A [NMAX]WorkSpace, jumlahData int) {
	var i int
	var indexFasilitas int

	if jumlahData == 0 {
		fmt.Println("Data kosong.")
		return
	}

	fmt.Println("\n=== DAFTAR DATA ===")
	for i = 0; i < jumlahData; i++ {
		fmt.Printf("%d. %s | %s | ", i+1, A[i].nama, A[i].lokasi)

		for indexFasilitas = 0; indexFasilitas < A[i].jumlahF; indexFasilitas++ {
			fmt.Print(A[i].fasilitas[indexFasilitas])
		}

		fmt.Printf(" | Rp%.0f %s %s\n", A[i].hargaSewa, A[i].rating, A[i].ulasan)
	}
}

func UpdateData(A *[NMAX]WorkSpace, jumlahData int) {
	var i int
	if jumlahData == 0 {
		fmt.Println("Tidak ada data untuk diupdate.")
		return
	}

	var namaCari string
	fmt.Print("Masukkan nama data yang ingin diupdate: ")
	fmt.Scanln(&namaCari)

	var index = -1
	for i = 0; i < jumlahData; i++ {
		if A[i].nama == namaCari {
			index = i
		}
	}

	if index == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	var temp WorkSpace
	fmt.Print("Nama baru: ")
	fmt.Scanln(&temp.nama)
	fmt.Print("Lokasi baru: ")
	fmt.Scanln(&temp.lokasi)
	fmt.Print("Jumlah fasilitas (max 3): ")
	fmt.Scanln(&temp.jumlahF)

	if temp.jumlahF > 3 {
		temp.jumlahF = 3
	}
	for i = 0; i < temp.jumlahF; i++ {
		fmt.Printf("Fasilitas ke-%d: ", i+1)
		fmt.Scanln(&temp.fasilitas[i])
	}

	fmt.Print("Harga sewa baru: ")
	fmt.Scanln(&temp.hargaSewa)
	fmt.Print("Rating pengguna (0.0 - 5.0): ")
	fmt.Scanln(&temp.rating)
	fmt.Print("Ulasan singkat: ")
	fmt.Scanln(&temp.ulasan)

	A[index] = temp
	fmt.Println("Data berhasil diupdate.")
}

func HapusData(A *[NMAX]WorkSpace, jumlahData *int) {
	var i int
	var index = -1
	if *jumlahData == 0 {
		fmt.Println("Tidak ada data untuk dihapus.")
		return
	}

	var namaCari string
	fmt.Print("Masukkan nama data yang ingin dihapus: ")
	fmt.Scanln(&namaCari)

	for i = 0; i < *jumlahData; i++ {
		if A[i].nama == namaCari {
			index = i
		}
	}

	if index == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	for i = index; i < *jumlahData-1; i++ {
		A[i] = A[i+1]
	}
	*jumlahData--

	fmt.Println("Data berhasil dihapus.")
}

func BinarySearchLokasi(A [NMAX]WorkSpace, jumlah int, target string) int {
	var kiri, kanan, tengah int
	kiri = 0
	kanan = jumlah - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if A[tengah].lokasi == target {
			return tengah
		} else if A[tengah].lokasi < target {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func MenampilkanDataSearch(A [NMAX]WorkSpace, jumlahData int, lokasiYangDicari string) {
	var Ditemukan bool = false
	var i int
	var indexFasilitas int

	fmt.Printf("\n=== Data dengan lokasi \"%s\" ===\n", lokasiYangDicari)

	for i = 0; i < jumlahData; i++ {
		if A[i].lokasi == lokasiYangDicari {
			Ditemukan = true
			fmt.Printf("%d. %s | ", i+1, A[i].nama)

			for indexFasilitas = 0; indexFasilitas < A[i].jumlahF; indexFasilitas++ {
				fmt.Print(A[i].fasilitas[indexFasilitas])
				if indexFasilitas < A[i].jumlahF-1 {
					fmt.Print(", ")
				}
			}

			fmt.Printf(" | Rp%.0f | Rating: %s | Ulasan: %s\n", A[i].hargaSewa, A[i].rating, A[i].ulasan)
		}
	}

	if !Ditemukan {
		fmt.Println(" Tidak ada data dengan lokasi tersebut.")
	}
}

func ListDataFasilits(A [NMAX]WorkSpace, jumlahData int, fasilitasCari string) {
	var dataDitemukan bool = false
	var i, indeksFasilitas int
	var cocok bool

	fmt.Printf("\n=== Data dengan fasilitas \"%s\" ===\n", fasilitasCari)

	for i = 0; i < jumlahData; i++ {
		cocok = false
		for indeksFasilitas = 0; indeksFasilitas < A[i].jumlahF; indeksFasilitas++ {
			if A[i].fasilitas[indeksFasilitas] == fasilitasCari {
				cocok = true
			}
		}

		if cocok {
			dataDitemukan = true
			fmt.Printf("%d. %s | %s | ", i+1, A[i].nama, A[i].lokasi)

			for indeksFasilitas = 0; indeksFasilitas < A[i].jumlahF; indeksFasilitas++ {
				fmt.Print(A[i].fasilitas[indeksFasilitas])
				if indeksFasilitas < A[i].jumlahF-1 {
					fmt.Print(", ")
				}
			}
			fmt.Printf(" | Rp%.0f | Rating: %s | Ulasan: %s\n", A[i].hargaSewa, A[i].rating, A[i].ulasan)
		}
	}

	if !dataDitemukan {
		fmt.Println("Tidak ada data dengan fasilitas tersebut.")
	}
}
