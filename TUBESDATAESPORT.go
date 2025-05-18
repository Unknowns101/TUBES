package main

import "fmt"

const NMAX = 10

type Pemain struct {
	Nama string
	Poin int
}

type Tim struct {
	Nama              string
	Menang            int
	Kalah             int
	TotalPertandingan int
	Poin              int
	NamaPemain        [5]Pemain
}

var Teams [NMAX]Tim
var nTim int

func AddTeam(teams *[NMAX]Tim, nTim *int) {
	if *nTim >= NMAX {
		fmt.Println("Kapasitas tim penuh.")
		return
	}

	var t Tim
	fmt.Print("Nama tim: ")
	fmt.Scan(&t.Nama)

	for i := 0; i < 5; i++ {
		fmt.Print("Nama pemain : ")
		fmt.Scan(&t.NamaPemain[i].Nama)
		t.NamaPemain[i].Poin = 0
	}

	t.Menang = 0
	t.Kalah = 0
	t.TotalPertandingan = 0
	t.Poin = 0

	teams[*nTim] = t
	(*nTim)++

	fmt.Println("Tim ditambahkan.")
}

func UpdateTeam(teams *[NMAX]Tim, nTim int) {
	var nama string
	fmt.Print("Masukkan nama tim : ")
	fmt.Scan(&nama)

	UrutkanTim(teams, nTim)
	index := CariTim(teams, nTim, nama)
	if index == -1 {
		fmt.Println("Tim tidak ditemukan.")
		return
	}

	t := &teams[index]
	fmt.Print("Jumlah menang: ")
	fmt.Scan(&t.Menang)
	fmt.Print("Jumlah kalah: ")
	fmt.Scan(&t.Kalah)

	fmt.Print("Jumlah pertandingan seri: ")
	var seri int
	fmt.Scan(&seri)
	t.Poin = t.Menang*3 + seri*2 + t.Kalah*1
	t.TotalPertandingan += 1
	for i := 0; i < 5; i++ {
		fmt.Printf("Masukkan poin terbaru untuk %s: ", t.NamaPemain[i].Nama)
		fmt.Scan(&t.NamaPemain[i].Poin)
	}
	fmt.Println("Tim berhasil diupdate.")
}

func DeleteTeam(teams *[NMAX]Tim, nTim *int) {
	var nama string
	fmt.Print("Nama tim yang akan dihapus: ")
	fmt.Scan(&nama)

	for i := 0; i < *nTim; i++ {
		if teams[i].Nama == nama {
			for j := i; j < *nTim-1; j++ {
				teams[j] = teams[j+1]
			}
			(*nTim)--
			fmt.Println("Tim dihapus.")
			return
		}
	}
	fmt.Println("Tim tidak ditemukan.")
}

func DisplayTeams(t *[NMAX]Tim, nTim int) {
	fmt.Println("Daftar Tim: ")
	for i := 0; i < nTim; i++ {
		fmt.Println("Nama Tim :", t[i].Nama)
		fmt.Println("Menang :", t[i].Menang)
		fmt.Println("Kalah :", t[i].Kalah)
		fmt.Println("Total :", t[i].TotalPertandingan)
		fmt.Println("Poin :", t[i].Poin)
		fmt.Println("Daftar Pemain :")

		maxPoin := -1
		var pemainTerbaik string
		for j := 0; j < 5; j++ {
			fmt.Println(t[i].NamaPemain[j])
			if t[i].NamaPemain[j].Poin > maxPoin {
				maxPoin = t[i].NamaPemain[j].Poin
				pemainTerbaik = t[i].NamaPemain[j].Nama
			}
		}
		fmt.Printf("Pemain terbaik: %s (Poin: %d)\n", pemainTerbaik, maxPoin)
		fmt.Println("----------------")
	}
}

func UrutkanTim(teams *[NMAX]Tim, nTim int) {
	for i := 0; i < nTim-1; i++ {
		minIdx := i
		for j := i + 1; j < nTim; j++ {
			if teams[j].Nama < teams[minIdx].Nama {
				minIdx = j
			}
		}
		if minIdx != i {
			teams[i], teams[minIdx] = teams[minIdx], teams[i]
		}
	}
}

func UrutkanTimByPoin(teams *[NMAX]Tim, nTim int) {
	for i := 0; i < nTim-1; i++ {
		maxIdx := i
		for j := i + 1; j < nTim; j++ {
			if teams[j].Poin > teams[maxIdx].Poin {
				maxIdx = j
			}
		}
		if maxIdx != i {
			teams[i], teams[maxIdx] = teams[maxIdx], teams[i]
		}
	}
}

func CariTim(teams *[NMAX]Tim, nTim int, nama string) int {
	var left, right, mid int
	left = 0
	right = nTim - 1

	for left <= right {
		mid = (left + right) / 2
		if teams[mid].Nama == nama {
			return mid
		} else if teams[mid].Nama < nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func CariTimByPoin(teams *[NMAX]Tim, nTim int, poin int) int {
	var left, right, mid int
	left = 0
	right = nTim - 1

	for left <= right {
		mid = (left + right) / 2
		if teams[mid].Poin == poin {
			return mid
		} else if teams[mid].Poin < poin {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func Menu() {
	var choice int
	for {
		fmt.Println(" ")
		fmt.Println("===== MENU =====")
		fmt.Println("1. Tambah Tim")
		fmt.Println("2. Update Tim")
		fmt.Println("3. Hapus Tim")
		fmt.Println("4. Tampilkan Semua Tim")
		fmt.Println("5. Cari Tim (Binary Search berdasarkan Nama)")
		fmt.Println("6. Cari Tim (Binary Search berdasarkan Poin)")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			AddTeam(&Teams, &nTim)
		case 2:
			UpdateTeam(&Teams, nTim)
		case 3:
			DeleteTeam(&Teams, &nTim)
		case 4:
			DisplayTeams(&Teams, nTim)
		case 5:
			var nama string
			fmt.Print("Masukkan nama tim yang dicari: ")
			fmt.Scan(&nama)
			UrutkanTim(&Teams, nTim)
			idx := CariTim(&Teams, nTim, nama)
			if idx != -1 {
				fmt.Printf("Tim ditemukan: %s | Poin: %d\n", Teams[idx].Nama, Teams[idx].Poin)
			} else {
				fmt.Println("Tim tidak ditemukan.")
			}
		case 6:
			var poin int
			fmt.Print("Masukkan poin tim yang dicari: ")
			fmt.Scan(&poin)
			UrutkanTimByPoin(&Teams, nTim)
			idx := CariTimByPoin(&Teams, nTim, poin)
			if idx != -1 {
				fmt.Printf("Tim ditemukan: %s | Poin: %d\n", Teams[idx].Nama, Teams[idx].Poin)
			} else {
				fmt.Println("Tim dengan poin tersebut tidak ditemukan.")
			}
		case 7:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main() {
	Menu()
}
