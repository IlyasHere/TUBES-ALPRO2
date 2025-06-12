package lib

import "fmt"

const NMAX = 100

type WorkSpace struct {
	nama      string
	Lokasi    string
	fasilitas [3]string
	jumlahF   int
	hargaSewa float64
	ulasan    string
	rating    string
}

var Data [NMAX]WorkSpace

func TambahData(A *[NMAX]WorkSpace, n *int) {
	var i int
	var temp WorkSpace

	if *n >= NMAX {
		fmt.Println("Data sudah penuh.")
	} else {
		fmt.Print("Nama: ")
		fmt.Scanln(&temp.nama)
		fmt.Print("Lokasi: ")
		fmt.Scanln(&temp.Lokasi)
		fmt.Print("Jumlah fasilitas (max 3): ")
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

		fmt.Print("Rating pengguna: ")
		fmt.Scanln(&temp.rating)

		fmt.Print("Ulasan singkat: ")
		fmt.Scanln(&temp.ulasan)

		A[*n] = temp
		*n++

		fmt.Println("Data berhasil ditambahkan.")
	}
}

func TampilData(A [NMAX]WorkSpace, n int) {
	var i, idxFasilitas int
	var fasilitas string

	fmt.Println("\n=== DAFTAR DATA ===")
	fmt.Printf("%-3s %-15s %-12s %-30s %-10s %-8s %-s\n", "No", "Nama", "Lokasi", "Fasilitas", "Harga", "Rating", "Ulasan")

	for i = 0; i < n; i++ {
		fasilitas = ""
		for idxFasilitas = 0; idxFasilitas < A[i].jumlahF; idxFasilitas++ {
			fasilitas += A[i].fasilitas[idxFasilitas]
			if idxFasilitas < A[i].jumlahF-1 {
				fasilitas += ", "
			}
		}

		fmt.Printf("%-3d %-15s %-12s %-30s Rp%-9.0f %-8s %-s\n",
			i+1, A[i].nama, A[i].Lokasi, fasilitas, A[i].hargaSewa, A[i].rating, A[i].ulasan)
	}
}

func UpdateData(A *[NMAX]WorkSpace, n int) {
	var i int
	var temp WorkSpace
	var namaCari string
	var idx int

	idx = -1
	fmt.Print("Masukkan nama data yang ingin diupdate: ")
	fmt.Scanln(&namaCari)

	for i = 0; i < n; i++ {
		if A[i].nama == namaCari {
			idx = i
		}
	}

	if n == 0 {
		fmt.Println("Tidak ada data untuk diupdate.")
	} else if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
	} else {
		fmt.Print("Nama baru: ")
		fmt.Scanln(&temp.nama)
		fmt.Print("Lokasi baru: ")
		fmt.Scanln(&temp.Lokasi)
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

		A[idx] = temp
		fmt.Println("Data berhasil diupdate.")
	}
}

func HapusData(A *[NMAX]WorkSpace, jumlahData *int) {
	var namaCari string
	var ketemu, k int

	fmt.Print("Masukkan nama data yang ingin dihapus: ")
	fmt.Scanln(&namaCari)

	fmt.Println()
	fmt.Println("=== HAPUS DATA ===")

	ketemu = -1
	k = 0

	if *jumlahData == 0 {
		fmt.Println("Tidak ada data untuk dihapus.")
	} else {
		for ketemu == -1 && k < *jumlahData {
			if A[k].nama == namaCari {
				ketemu = k
			}
			k++
		}

		if ketemu == -1 {
			fmt.Println("Data tidak ditemukan.")
		} else {
			for i := ketemu; i < *jumlahData-1; i++ {
				A[i] = A[i+1]
			}
			*jumlahData--
			fmt.Println("Data berhasil dihapus.")
		}
	}
}

func BinarySearchLokasi(A [NMAX]WorkSpace, n int, target string) int {
	var left, right, mid int
	var found int

	left = 0
	right = n - 1
	found = -1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if target < A[mid].Lokasi {
			right = mid - 1
		} else if target > A[mid].Lokasi {
			left = mid + 1
		} else {
			found = mid
		}
	}

	return found
}

func SelectionSortByLokasiAsc(A *[NMAX]WorkSpace, jumlah int) {
	var pass, idx, i int
	var temp WorkSpace

	pass = 1
	for pass <= jumlah-1 {
		idx = pass - 1
		i = pass
		for i < jumlah {
			if A[idx].Lokasi > A[i].Lokasi {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func TampilDataDarilokasi(A [NMAX]WorkSpace, i int) {
	var fasilitas string
	for j := 0; j < A[i].jumlahF; j++ {
		fasilitas += A[i].fasilitas[j]
		if j < A[i].jumlahF-1 {
			fasilitas += ", "
		}
	}

	fmt.Printf("%-3d %-15s %-12s %-30s Rp%-10.0f %-8s %-s\n",
		i+1, A[i].nama, A[i].Lokasi, fasilitas, A[i].hargaSewa, A[i].rating, A[i].ulasan)
}

func MenampilkanDataSearchByNama(A [NMAX]WorkSpace, jumlahData int, namaCari string) {
	var ketemu, k int
	var fasilitas string
	ketemu = -1
	k = 0

	for ketemu == -1 && k < jumlahData {
		if A[k].nama == namaCari {
			ketemu = k
		}
		k++
	}
	if ketemu != -1 {

		for j := 0; j < A[ketemu].jumlahF; j++ {
			fasilitas += A[ketemu].fasilitas[j]
			if j < A[ketemu].jumlahF-1 {
				fasilitas += ", "
			}
		}

		fmt.Printf("\n=== DATA DITEMUKAN ===\n")
		fmt.Printf("%-3s %-15s %-12s %-30s %-10s %-8s %-s\n", "No", "Nama", "Lokasi", "Fasilitas", "Harga", "Rating", "Ulasan")
		fmt.Printf("%-3d %-15s %-12s %-30s Rp%-10.0f %-8s %-s\n",
			ketemu+1, A[ketemu].nama, A[ketemu].Lokasi, fasilitas, A[ketemu].hargaSewa, A[ketemu].rating, A[ketemu].ulasan)
	} else {
		fmt.Println("Tidak ada data dengan nama tersebut.")
	}
}

func ListDataFasilits(A [NMAX]WorkSpace, jumlahData int, fasilitasCari string) {
	var ketemu int
	var i int
	var indeksF int
	var ditemukan bool

	fmt.Printf("\n=== Data dengan fasilitas \"%s\" ===\n", fasilitasCari)
	ketemu = -1
	i = 0
	for i < jumlahData {
		ketemu = -1
		indeksF = 0
		for ketemu == -1 && indeksF < A[i].jumlahF {
			if A[i].fasilitas[indeksF] == fasilitasCari {
				ketemu = indeksF
			}
			indeksF++
		}

		if ketemu != -1 {
			ditemukan = true
			fmt.Printf("%d. %s | %s | ", i+1, A[i].nama, A[i].Lokasi)
			for indeksF = 0; indeksF < A[i].jumlahF; indeksF++ {
				fmt.Print(A[i].fasilitas[indeksF])
				if indeksF < A[i].jumlahF-1 {
					fmt.Print(", ")
				}
			}

			fmt.Printf(" | Rp%.0f | Rating: %s | Ulasan: %s\n", A[i].hargaSewa, A[i].rating, A[i].ulasan)
		}
		i++
	}

	if !ditemukan {
		fmt.Println("Tidak ada data dengan fasilitas tersebut.")
	}
}

func SelectionSortHargaDes(A *[NMAX]WorkSpace, N int) {
	var pass, idx, i int
	var temp WorkSpace

	pass = 1
	for pass <= N-1 {
		idx = pass - 1
		i = pass
		for i < N {
			if A[idx].hargaSewa < A[i].hargaSewa {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

func SelectionSortHargaAsc(A *[NMAX]WorkSpace, N int) {
	var pass, idx, i int
	var temp WorkSpace

	pass = 1
	for pass <= N-1 {
		idx = pass - 1
		i = pass
		for i < N {
			if A[idx].hargaSewa > A[i].hargaSewa {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

func InsertionSortRatingAsc(A *[NMAX]WorkSpace, N int) {
	var pass, i int
	var temp WorkSpace

	pass = 1
	for pass <= N-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.rating < A[i-1].rating {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
		pass = pass + 1
	}
}

func InsertionSortRatingDesc(A *[NMAX]WorkSpace, N int) {
	var pass, i int
	var temp WorkSpace

	pass = 1
	for pass <= N-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.rating > A[i-1].rating {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
		pass = pass + 1
	}
}

func InitDummyData(A *[NMAX]WorkSpace, n *int) {
	A[0] = WorkSpace{
		nama:      "StartupHub",
		Lokasi:    "Jakarta",
		fasilitas: [3]string{"WiFi", "Meeting Room", "Snack"},
		jumlahF:   3,
		hargaSewa: 150000,
		rating:    "4.5",
		ulasan:    "Nyaman dan modern",
	}
	A[1] = WorkSpace{
		nama:      "CreativeNest",
		Lokasi:    "Bandung",
		fasilitas: [3]string{"WiFi", "Private Desk", ""},
		jumlahF:   2,
		hargaSewa: 100000,
		rating:    "4.0",
		ulasan:    "Tenang dan cozy",
	}
	A[2] = WorkSpace{
		nama:      "WorkHive",
		Lokasi:    "Surabaya",
		fasilitas: [3]string{"WiFi", "Kopi Gratis", "AC"},
		jumlahF:   3,
		hargaSewa: 120000,
		rating:    "4.8",
		ulasan:    "Fasilitas lengkap dan murah",
	}
	A[3] = WorkSpace{
		nama:      "InovasiSpace",
		Lokasi:    "Jakarta",
		fasilitas: [3]string{"WiFi", "Ruang Presentasi", ""},
		jumlahF:   2,
		hargaSewa: 180000,
		rating:    "4.2",
		ulasan:    "Bagus untuk tim kecil",
	}
	A[4] = WorkSpace{
		nama:      "Remotely",
		Lokasi:    "Yogyakarta",
		fasilitas: [3]string{"WiFi", "Ruang Diskusi", "Kantin"},
		jumlahF:   3,
		hargaSewa: 90000,
		rating:    "4.6",
		ulasan:    "Murah dan asri",
	}
	A[5] = WorkSpace{
		nama:      "ThinkTank",
		Lokasi:    "Bandung",
		fasilitas: [3]string{"WiFi", "Snack", "Ruang Istirahat"},
		jumlahF:   3,
		hargaSewa: 110000,
		rating:    "4.3",
		ulasan:    "Kreatif dan inspiratif",
	}
	A[6] = WorkSpace{
		nama:      "WorkZone",
		Lokasi:    "Depok",
		fasilitas: [3]string{"WiFi", "Kopi Gratis", ""},
		jumlahF:   2,
		hargaSewa: 95000,
		rating:    "4.0",
		ulasan:    "Cocok untuk mahasiswa",
	}
	A[7] = WorkSpace{
		nama:      "CozyCorner",
		Lokasi:    "Bogor",
		fasilitas: [3]string{"WiFi", "Bantal Duduk", "AC"},
		jumlahF:   3,
		hargaSewa: 85000,
		rating:    "3.8",
		ulasan:    "Tenang dan minimalis",
	}
	A[8] = WorkSpace{
		nama:      "FocusRoom",
		Lokasi:    "Bekasi",
		fasilitas: [3]string{"WiFi", "Private Room", ""},
		jumlahF:   2,
		hargaSewa: 130000,
		rating:    "4.7",
		ulasan:    "Fokus kerja sangat nyaman",
	}
	A[9] = WorkSpace{
		nama:      "DigitalSpace",
		Lokasi:    "Tangerang",
		fasilitas: [3]string{"WiFi", "Monitor Ekstra", "Snack"},
		jumlahF:   3,
		hargaSewa: 140000,
		rating:    "4.4",
		ulasan:    "Modern dan profesional",
	}

	*n = 10
}
