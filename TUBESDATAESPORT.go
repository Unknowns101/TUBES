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

func AddTeam(teams *Teams, nTim *int) {
//Berfungsi untuk menambahkan Nama tim kedalam Array serta pemain dalam tim tersebut//
	if *nTim >= NMAX {
		fmt.Println("Kapasitas tim penuh.")
		return nTim = NMAX
	}

	var t Tim
	fmt.Print("Nama tim: ")
	fmt.Scan(&t.Nama)
	var i int

	for i = 0; i < 5; i++ {
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

func UpdateTeam(teams *Teams, nTim int) {
//Berfungsi untuk mengubah atau mengupdate kemenagan, kekalahan, total pertandingan, poin tim dan poin tiap pemain dalam tim yang akan di update/
	var nama string
	fmt.Print("Masukkan nama tim : ")
	fmt.Scan(&nama)
	var index, i int
	var t Teams

	index = CariTim(teams, nTim, nama)
	if index == -1 {
		fmt.Println("Tim tidak ditemukan.")
		return
	}

	t = &teams[index]
	fmt.Print("Jumlah menang: ")
	fmt.Scan(&t.Menang)
	fmt.Print("Jumlah kalah: ")
	fmt.Scan(&t.Kalah)

	if t.Menang == t.Kalah {
		t.Poin += 1
	} else {
		t.Poin = t.Poin + (t.Menang - t.Kalah)
	}
	t.TotalPertandingan = t.Menang + t.Kalah

	for i = 0; i < 5; i++ {
		fmt.Printf("Masukkan poin terbaru untuk %s: ", t.NamaPemain[i].Nama)
		fmt.Scan(&t.NamaPemain[i].Poin)
	}

	fmt.Println("Tim berhasil diupdate.")
}

func DeleteTeam(teams *Teams, nTim *int) {
//Berfungsi untuk mengapus Tim dari Array Tim//
	var nama string
	fmt.Print("Nama tim yang akan dihapus: ")
	fmt.Scan(&nama)
	var i, j int

	for i = 0; i < *nTim; i++ {
		if teams[i].Nama == nama {
			for j = i; j < *nTim-1; j++ {
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
//Berfungsi untuk menampilkan Seluruh Tim yang ikut berkompetisi dan menampilkan tim terbaik dan pemain terbaik//
	var i, j, maxPoin, maxPoinPemain, maxPoinTim int
	fmt.Println("Daftar Tim: ")
	fmt.Println("-------------")
	for i = 0; i < nTim; i++ {
		fmt.Println("Nama Tim :", t[i].Nama)
		fmt.Println("Menang :", t[i].Menang)
		fmt.Println("Kalah :", t[i].Kalah)
		fmt.Println("Total :", t[i].TotalPertandingan)
		fmt.Println("Poin :", t[i].Poin)
		fmt.Println("Daftar Pemain :")

		maxPoin = -1
		var pemainTerbaik string
		for j = 0; j < 5; j++ {
			fmt.Println(t[i].NamaPemain[j])
			if t[i].NamaPemain[j].Poin > maxPoin {
				maxPoin = t[i].NamaPemain[j].Poin
				pemainTerbaik = t[i].NamaPemain[j].Nama
			}
		}
		fmt.Printf("Pemain terbaik di dalam tim: %s (Poin: %d)\n", pemainTerbaik, maxPoin)
		fmt.Println("----------------")
	}

	maxPoinTim = -1
	maxPoinPemain = -1
	var namaTimTerbaik, namaPemainTerbaik, namaTimPemainTerbaik string

	for i = 0; i < nTim; i++ {
		if t[i].Poin > maxPoinTim {
			maxPoinTim = t[i].Poin
			namaTimTerbaik = t[i].Nama
		}
		for j = 0; j < 5; j++ {
			if t[i].NamaPemain[j].Poin > maxPoinPemain {
				maxPoinPemain = t[i].NamaPemain[j].Poin
				namaPemainTerbaik = t[i].NamaPemain[j].Nama
				namaTimPemainTerbaik = t[i].Nama
			}
		}
	}

	fmt.Printf("Tim dengan poin terbanyak: %s (Poin: %d)\n", namaTimTerbaik, maxPoinTim)
	fmt.Printf("Pemain terbaik dari semua tim: %s (Poin: %d) dari tim %s\n", namaPemainTerbaik, maxPoinPemain, namaTimPemainTerbaik)
}

func UrutkanTimByNama_SelectionSort(teams *[NMAX]Tim, nTim int) {
//Untuk mengurutkan Tim bedasarkan Alfabetis menggunakan Selection Sort//
	var i, minIdx, j int
	for i = 0; i < nTim-1; i++ {
		minIdx = i
		for j = i + 1; j < nTim; j++ {
			if teams[j].Nama < teams[minIdx].Nama {
				minIdx = j
			}
		}
		if minIdx != i {
			teams[i], teams[minIdx] = teams[minIdx], teams[i]
		}
	}
}

func UrutkanTimByPoin_InsertionSort(teams *[NMAX]Tim, nTim int) {
//Untuk mengurutkan Tim bedasarkan Poin yang dimiliki menggunakan insertion sort//
	var i, j  int
	for i = 1; i < nTim; i++ {
		temp := teams[i]
		j = i - 1
		for j >= 0 && teams[j].Poin < temp.Poin {
			teams[j+1] = teams[j]
			j--
		}
		teams[j+1] = temp
	}
}

func CariTim(teams *[NMAX]Tim, nTim int, nama string) int {
//Berfungsi untuk mencari Tim menggunakan Nama tim yang ingin dicari dengan Sequential search//
	var i, idx int

	idx = -1
	i = 0
	for i < nTim && idx == -1 {
		if teams[i].Nama == nama {
			idx = i
		}
	}
	return idx
}

func CariTimByPoin(teams *[NMAX]Tim, nTim int, poin int) int {
//Berfungsi untuk mencari tim bedasarkan poin yang dimiliki tim saat itu menggunakan binary search//
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
//Menampilkan Fitur fitur yang dapat dipakai dalam aplikasi//
	var choice int
	var nTim int
	var Team Teams
	for {
		fmt.Println("===== MENU =====")
		fmt.Println("1. Tambah Tim")
		fmt.Println("2. Update Tim")
		fmt.Println("3. Hapus Tim")
		fmt.Println("4. Tampilkan Semua Tim")
		fmt.Println("5. Cari Tim (Binary Search berdasarkan Nama)")
		fmt.Println("6. Cari Tim (Binary Search berdasarkan Poin)")
		fmt.Println("7. Urutkan Tim Berdasarkan Nama")
		fmt.Println("8. Urutkan Tim Berdasarkan Poin")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			AddTeam(&Team, &nTim)
		case 2:
			UpdateTeam(&Team, nTim)
		case 3:
			DeleteTeam(&Team, &nTim)
		case 4:
			DisplayTeams(&Team, nTim)
		case 5:
			var nama string
			fmt.Print("Masukkan nama tim yang dicari: ")
			fmt.Scan(&nama)
			UrutkanTimByNama_SelectionSort(&Team, nTim)
			idx := CariTim(&Team, nTim, nama)
			if idx != -1 {
				fmt.Printf("Tim ditemukan: %s | Poin: %d\n", Team[idx].Nama, Team[idx].Poin)
			} else {
				fmt.Println("Tim tidak ditemukan.")
			}
		case 6:
			var poin int
			fmt.Print("Masukkan poin tim yang dicari: ")
			fmt.Scan(&poin)
			UrutkanTimByPoin_InsertionSort(&Team, nTim)
			idx := CariTimByPoin(&Team, nTim, poin)
			if idx != -1 {
				fmt.Printf("Tim ditemukan: %s | Poin: %d\n", Team[idx].Nama, Team[idx].Poin)
			} else {
				fmt.Println("Tim dengan poin tersebut tidak ditemukan.")
			}
		case 7:
			UrutkanTimByNama_SelectionSort(&Team, nTim)
			fmt.Println("Tim berhasil diurutkan berdasarkan nama (Selection Sort).")
		case 8:
			UrutkanTimByPoin_InsertionSort(&Team, nTim)
			fmt.Println("Tim berhasil diurutkan berdasarkan poin (Insertion Sort).")
		case 9:
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
