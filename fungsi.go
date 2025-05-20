package main

import "fmt"

const NMAX = 100

type workSpace struct {
	nama      string
	lokasi    string
	fasilitas [3]string
	jumlahF   int
	hargaSewa float64
}

var data [NMAX]workSpace

func tambahData(A *[NMAX]workSpace, ValueArray *int) {
	var i int
	if *ValueArray >= NMAX {
		fmt.Println("Data sudah penuh.")
		return
	}

	var temp workSpace

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

	A[*ValueArray] = temp
	*ValueArray++

	fmt.Println("Data berhasil ditambahkan.")
}

func tampilData(A [NMAX]workSpace, jumlahData int) {
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

		fmt.Printf(" | Rp%.0f\n", A[i].hargaSewa)
	}
}

func updateData(A *[NMAX]workSpace, jumlahData int) {
	var i int
	if jumlahData == 0 {
		fmt.Println("Tidak ada data untuk diupdate.")
		return
	}

	var index int
	tampilData(*A, jumlahData)
	fmt.Print("Masukkan nomor data yang ingin diupdate: ")
	fmt.Scanln(&index)
	index = index - 1

	if index < 0 || index >= jumlahData {
		fmt.Println("Nomor tidak valid.")
		return
	}

	var temp workSpace

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

	A[index] = temp
	fmt.Println("Data berhasil diupdate.")
}

func hapusData(A *[NMAX]workSpace, jumlahData *int) {
	var i, idx int

	if *jumlahData == 0 {
		fmt.Println("Tidak ada data untuk diupdate.")

	}

	tampilData(*A, *jumlahData)
	fmt.Print("Masukkan nomor data yang ingin dihapus: ")
	fmt.Scanln(&idx)
	idx = idx - 1

	if idx < 0 || idx >= *jumlahData {
		fmt.Println("Nomor tidak valid.")
		return
	}

	for i = idx; i < *jumlahData-1; i++ {
		A[i] = A[i+1]
	}
	*jumlahData = *jumlahData - 1

	fmt.Println("Data berhasil dihapus.")
}
