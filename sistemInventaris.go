package main
import (
	"fmt"
	"os"
    "os/exec"
    "runtime"
	/*"bufio"*/
	"strings"
)

// Batas untuk array statis
const maxArr int = 1000

// Mendefinisikan struct untuk simpan data inventaris
type inventaris struct {
	nama 	string	// Nama barang
	merek 	string	// Nama merek barang
	lokasi 	string	// Lokasi barang
	kondisi string	// Kondisi barang
	tanggal string  // Tanggal pembelian barang
	stok 	int		// Stok barang
	harga 	float64	// Harga nilai barang
}

// Mendefinisikan struct untuk simpan array alat
type arrInv [maxArr]inventaris

// Padding untuk menambah estetika antarmuka
var padding = strings.Repeat(" ", 55)

func main() {
	var alatLab arrInv
	var nData int = 7
	
	// Memasukkan beberapa data-data inventaris secara hard-core
	alatLab[0] = inventaris{nama: "Baju", merek: "Lab", lokasi: "Ruang_1", kondisi: "Baik", tanggal: "01/01/2021", stok: 25, harga: 750000}
	alatLab[1] = inventaris{nama: "Penggaris", merek: "Butterfly", lokasi: "Rak_1", kondisi: "Baik", tanggal: "22/03/2020", stok: 100, harga: 15000}
	alatLab[2] = inventaris{nama: "Monitor", merek: "LG", lokasi: "Ruang_1", kondisi: "Baik", tanggal: "17/03/2019", stok: 50, harga: 820000}
	alatLab[3] = inventaris{nama: "Meja", merek: "none", lokasi: "Ruang_2", kondisi: "Sedang", tanggal: "07/08/2010", stok: 3, harga: 2400000}
	alatLab[4] = inventaris{nama: "Gunting", merek: "Kenko", lokasi: "Rak_2", kondisi: "Baik", tanggal: "04/09/2024", stok: 14, harga: 17500}
	alatLab[5] = inventaris{nama: "Kabel_HDMI", merek: "none", lokasi: "Ruang_3", kondisi: "Buruk", tanggal: "09/12/2011", stok: 4, harga: 55000}
	alatLab[6] = inventaris{nama: "NamaPanjangSekali", merek: "MerekPanjangSekali", lokasi: "LokasiPanjangSekali", kondisi: "KondisiPanjangSekali", tanggal: "09/12/2011", stok: 4, harga: 55000}

	menu(&alatLab, &nData)
}

// Fungsi untuk membersihkan tampilan CMD/Terminal
func clearScreen() {
    var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Prosedur untuk menampilkan pilihan menu utama
func menu(alatLab *arrInv, nData *int) {
	var pilihan string
	var clear bool = true
	
	for {
		if clear {
			clearScreen()
		}
		headerTampilan()
		fmt.Println(padding + "║                 Menu Utama                  ║")
		fmt.Println(padding + "╠══════════════════════╦══════════════════════╣")
		fmt.Println(padding + "║ 1. Data baru         ║ 6. Urutkan data      ║")
		fmt.Println(padding + "║ 2. Ubah data         ║ 7. Perhitungan       ║")
		fmt.Println(padding + "║ 3. Hapus data        ║ 0. Keluar            ║")
		fmt.Println(padding + "║ 4. Cari data         ║                      ║")
		fmt.Println(padding + "║ 5. Tampilkan data    ║                      ║")
		fmt.Println(padding + "╚══════════════════════╩══════════════════════╝")
		fmt.Print(padding + "Pilih menu> ")
		
		fmt.Scan(&pilihan)

		switch pilihan {
		case "1":
			clear = true
			clearScreen()
			dataBaru(alatLab, nData)
		case "2":
			clear = true
			clearScreen()
			ubahData(alatLab, nData)
		case "3":
			clear = true
			clearScreen()
			hapusData(alatLab, nData)
		case "4":
			clear = true
			clearScreen()
			cariData(*alatLab, *nData)
		case "5":
			clear = true
			clearScreen()
			tampilkanData(*alatLab, *nData)
		case "6":
			clear = true
			urutData(alatLab, *nData)
		/*case "7":
			clear = true
			clearScreen()
			menuHitung(*alatLab, *nData)*/
		case "0":
			clearScreen()
			fmt.Println()
			byeImage()
			fmt.Println()
			fmt.Println(padding + "╔═════════════════════════════════════════════╗")
			fmt.Println(padding + "║        Terima Kasih dan Sampai Jumpa        ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
			return
		default:
			clear = false
			fmt.Println()
			fmt.Println(padding + "╔═════════════════════════════════════════════╗")
			fmt.Println(padding + "║  [!] Pilihan tidak valid!, mohon coba lagi  ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
		}
	} 
}

// Prosedur untuk menampilkan header tampilan
func headerTampilan() {
	fmt.Println()
	fmt.Println(padding + "╔═════════════════════════════════════════════╗")
	fmt.Println(padding + "║   Aplikasi Sistem Inventaris Laboratorium   ║")
	fmt.Println(padding + "║        Created by hafiza and fazli          ║")
	fmt.Println(padding + "║        Algoritma Pemrograman 2025           ║")
	fmt.Println(padding + "╠═════════════════════════════════════════════╣")
}

// Prosedur untuk menampilkan pesan data kosong
func pesanDataKosong() {
	fmt.Println()
	fmt.Println(padding + "╔═════════════════════════════════════════════╗")
	fmt.Println(padding + "║   [!] Belum ada data, masukkan data baru    ║")
	fmt.Println(padding + "║         untuk menggunakan fitur ini         ║")
	fmt.Println(padding + "╚═════════════════════════════════════════════╝")
}

// Prosedur untuk menampilkan pesan data berhasil diurut
func pesanDataTerurut(pesan string) {
	fmt.Println()
	fmt.Println(padding + "╔═════════════════════════════════════════════╗")
	fmt.Printf(padding + "║  [✓] Data sudah diurut berdasarkan %-7s  ║\n", pesan)
	fmt.Println(padding + "╚═════════════════════════════════════════════╝")
}
// Prosedur untuk menampilkan data menggunakan template tabel
func tabelTampilan(data arrInv, hasilCari int) {
	paddingTable := strings.Repeat(" ", 25)
	
	nama := potongString(data[hasilCari].nama, 15)
    merek := potongString(data[hasilCari].merek, 13)
    lokasi := potongString(data[hasilCari].lokasi, 13)
    kondisi := potongString(data[hasilCari].kondisi, 8)
    tanggal := potongString(data[hasilCari].tanggal, 15)
	
	fmt.Println(paddingTable + "┌─────────────────┬───────────────┬───────────────┬──────────┬───────┬───────────────┬─────────────────┐")
	fmt.Println(paddingTable + "│ Nama            │ Merek         │ Lokasi        │ Kondisi  │ Stok  │ Harga         │ Tanggal beli    │")
	fmt.Println(paddingTable + "├─────────────────┼───────────────┼───────────────┼──────────┼───────┼───────────────┼─────────────────┤")
	line := fmt.Sprintf("│ %-15s │ %-13s │ %-13s │ %-8s │ %5d │ %13.2f │ %-15s │", nama, merek, lokasi, kondisi, data[hasilCari].stok, data[hasilCari].harga, tanggal)
	fmt.Println(paddingTable + line)
	fmt.Println(paddingTable + "└─────────────────┴───────────────┴───────────────┴──────────┴───────┴───────────────┴─────────────────┘")
}

// Fungsi untuk memotong string kata sesuai parameter maksKata sehingga tidak merusak antarmuka
func potongString(kata string, maksKata int) string {
	if len(kata) > maksKata {
		return kata[:maksKata]
	}
	return kata
}

// Prosedur menjalankan konfirmasi apakah pilihan menu ini tetap dijalankan atau dihentikan
func konfirmasiInput(pesan string, hasil *int) {
	var konfirmasi string
	
	for {
		fmt.Printf(padding + "%s> ", pesan)
		fmt.Scan(&konfirmasi)
		
		if konfirmasi == "n" || konfirmasi == "N" {
			*hasil = 1
			return
		} else if konfirmasi == "y" || konfirmasi == "Y" {
			*hasil = 2
			return
		} else {
			fmt.Println()
			fmt.Println(padding + "╔═════════════════════════════════════════════╗")
			fmt.Println(padding + "║         [!] Pilihan hanya berupa y/n        ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
		}
	}
}

// Prosedur untuk menambahkan data inventaris baru
func dataBaru(data *arrInv, n *int) {
	var hasil int
	for {
		if *n >= maxArr {
			fmt.Println(padding + "╔═════════════════════════════════════════════╗")
			fmt.Println(padding + "║                     [!]                     ║")
			fmt.Println(padding + "║ Data sudah penuh! Anda bisa hapus data lain ║")
			fmt.Println(padding + "║           untuk menambah data baru          ║")
			fmt.Println(padding + "║    (tekan apapun lalu Enter untuk keluar)   ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
			
			var dummyInput string
			fmt.Scan(&dummyInput)
			return
		}
		
		headerTampilan()
		fmt.Println(padding + "║            Menu Utama > Data baru           ║")
		fmt.Println(padding + "╚═════════════════════════════════════════════╝")
		fmt.Print(padding +  "Nama alat\t : ")
		fmt.Scan(&data[*n].nama)
		fmt.Print(padding + "Merek alat\t : ")
		fmt.Scan(&data[*n].merek)
		fmt.Print(padding + "Lokasi alat\t : ")
		fmt.Scan(&data[*n].lokasi)
		fmt.Print(padding + "Kondisi alat\t : ")
		fmt.Scan(&data[*n].kondisi)
		fmt.Print(padding + "Stok alat\t : ")
		fmt.Scan(&data[*n].stok)
		fmt.Print(padding + "Harga alat\t : ")
		fmt.Scan(&data[*n].harga)
		fmt.Print(padding + "Tanggal pembelian (hh/bb/tttt): ")
		fmt.Scan(&data[*n].tanggal)
		
		*n = *n + 1
		
		fmt.Println()
		konfirmasiInput("Masukkan data baru lagi? (y/n)", &hasil)
		
		if hasil == 1 {
			return
		} else if hasil == 2 {
			clearScreen()
		} else {
			fmt.Println(padding + "╔═════════════════════════════════════════════╗")
			fmt.Println(padding + "║                  [!] Galat                  ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
			return
		}
	}
}

// Prosedur untuk mengubah data dengan menimpa yang lama dengan yang baru
func ubahData(data *arrInv, n *int) {
	var dataDicari string
	var dataDiubah [1]inventaris
	var hasil, hasilCari int
	
	for {
		headerTampilan()
		fmt.Println(padding + "║            Menu Utama > Ubah data           ║")
		fmt.Println(padding + "╚═════════════════════════════════════════════╝")
		fmt.Print(padding + "Cari nama data: ")
		fmt.Scan(&dataDicari)
		
		hasilCari = sequentialSearch(*data, *n, dataDicari)
		
		if hasilCari >= 0 {
			tabelTampilan(*data, hasilCari)
			fmt.Println()
			fmt.Println(padding + "╔═════════════════════════════════════════════╗")
			fmt.Println(padding + "║     [i] Isi perubahan data di bawah ini     ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
			fmt.Print(padding +  "Nama alat\t : ")
			fmt.Scan(&dataDiubah[0].nama)
			fmt.Print(padding + "Merek alat\t : ")
			fmt.Scan(&dataDiubah[0].merek)
			fmt.Print(padding + "Lokasi alat\t : ")
			fmt.Scan(&dataDiubah[0].lokasi)
			fmt.Print(padding + "Kondisi alat\t : ")
			fmt.Scan(&dataDiubah[0].kondisi)
			fmt.Print(padding + "Stok alat\t : ")
			fmt.Scan(&dataDiubah[0].stok)
			fmt.Print(padding + "Harga alat\t : ")
			fmt.Scan(&dataDiubah[0].harga)
			fmt.Print(padding + "Tanggal pembelian (hh/bb/tttt): ")
			fmt.Scan(&dataDiubah[0].tanggal)
			
			data[hasilCari].nama = dataDiubah[0].nama
			data[hasilCari].merek = dataDiubah[0].merek 
			data[hasilCari].lokasi = dataDiubah[0].lokasi
			data[hasilCari].kondisi = dataDiubah[0].kondisi
			data[hasilCari].tanggal = dataDiubah[0].tanggal
			data[hasilCari].stok = dataDiubah[0].stok
			data[hasilCari].harga = dataDiubah[0].harga
			
			fmt.Println()
			konfirmasiInput("Ubah data lagi? (y/n)", &hasil)
			
		} else {
			fmt.Println()
			fmt.Println(padding + "╔═════════════════════════════════════════════╗")
			fmt.Println(padding + "║          [!] Data tidak ditemukan           ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
			konfirmasiInput("Cari data lagi? (y/n)", &hasil)
		}
		
		if hasil == 1 {
			clearScreen()
			return
		} else if hasil == 2 {
			clearScreen()
		} else {
			fmt.Println(padding + "╔═════════════════════════════════════════════╗")
			fmt.Println(padding + "║                  [!] Galat                  ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
		}
	}

}

// Prosedur untuk menghapus data dengan menimpa data dihapus dan perulangan untuk memindahkannya ke kiri dan mengubah nilai data aktif
func hapusData(data *arrInv, n *int) {
	var dataDicari string
	var idx, hasil int
	
	for {
		headerTampilan()
		fmt.Println(padding + "║           Menu Utama > Hapus data           ║")
		fmt.Println(padding + "╚═════════════════════════════════════════════╝")
		fmt.Print(padding + "Cari nama data: ")
		fmt.Scan(&dataDicari)
		fmt.Println()
		
		idx = sequentialSearch(*data, *n, dataDicari)
		
		if idx >= 0 {
			fmt.Println()
			tabelTampilan(*data, idx)
			konfirmasiInput("[!] Hapus? (y/n)", &hasil)
			fmt.Println()
			
			if hasil == 2 {
				for i := idx; i < *n-1; i++ {
					(*data)[i] = (*data)[i+1]
				}
				(*n)--
				
				fmt.Println(padding + "╔═════════════════════════════════════════════╗")
				fmt.Println(padding + "║         [✓] Data berhasil di hapus          ║")
				fmt.Println(padding + "╚═════════════════════════════════════════════╝")
				konfirmasiInput("Hapus data lagi? (y/n)", &hasil)
			} else {
				fmt.Println(padding + "╔═════════════════════════════════════════════╗")
				fmt.Println(padding + "║       [!] Penghapusan data dibatalkan       ║")
				fmt.Println(padding + "╚═════════════════════════════════════════════╝")
				konfirmasiInput("Cari data lagi? (y/n)", &hasil)	
			}
		} else {
			fmt.Println(padding + "╔═════════════════════════════════════════════╗")
			fmt.Println(padding + "║          [!] Data tidak ditemukan           ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
			konfirmasiInput("Cari data lagi? (y/n)", &hasil)
		}
		
		if hasil == 1 {
			clearScreen()
			return
		} else if hasil == 2 {
			clearScreen()
		} else {
			fmt.Println(padding + "╔═════════════════════════════════════════════╗")
			fmt.Println(padding + "║                  [!] Galat                  ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
			return
		}
	}
}

// Prosedur untuk mencari data dengan menggunakan logika if-else ketika data ditemukan
func cariData(data arrInv, n int) {
	var x string
	var hasil int
	
	for {
		headerTampilan()
		fmt.Println(padding + "║           Menu Utama > Cari data            ║")
		fmt.Println(padding + "╚═════════════════════════════════════════════╝")
		fmt.Print(padding + "Cari nama data: ")
		fmt.Scan(&x)
		fmt.Println()
		
		if sequentialSearch(data, n, x) >= 0 {
			fmt.Println(padding + "╔═════════════════════════════════════════════╗")
			fmt.Println(padding + "║             [✓] Data ditemukan              ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
		} else {
			fmt.Println(padding + "╔═════════════════════════════════════════════╗")
			fmt.Println(padding + "║          [!] Data tidak ditemukan           ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
		}
		
		konfirmasiInput("Cari data lagi? (y/n)", &hasil)
		
		if hasil == 1 {
			clearScreen()
			return
		} else if hasil == 2 {
			clearScreen()
		} else {
			fmt.Println(padding + "╔═════════════════════════════════════════════╗")
			fmt.Println(padding + "║                  [!] Galat                  ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
			return
		}
	}
}


// Fungsi untuk mencari data dengan metode Sequential Search
func sequentialSearch(data arrInv, n int, x string) int {
	var idx int = -1
	var k int = 0
	
	for idx == -1 && k < n {
		if data[k].nama == x {
			idx = k
		}
		k++
	}
	
	return idx
}

// Prosedur untuk melihat seluruh data aktif dengan template tabel
func tampilkanData(data arrInv, n int) {
	/*var idxAwal int = 0
	var idxAkhir int = 10*/
	var numberData = 0
	var dummyInput string
	paddingTable := strings.Repeat(" ", 22)
	
	headerTampilan()
	fmt.Println(padding + "║         Menu Utama > Tampilkan data         ║")
	fmt.Println(padding + "╚═════════════════════════════════════════════╝")
	
	if n == 0 {
		fmt.Println()
		pesanDataKosong()
		fmt.Print(padding + "Keluar? (tekan apapun lalu Enter)> ")
		fmt.Scan(&dummyInput)
		return
	} else {	
		fmt.Println()
		fmt.Println(paddingTable + "┌─────┬─────────────────┬───────────────┬───────────────┬──────────┬───────┬─────────────────┬─────────────────┐")
		fmt.Println(paddingTable + "│ No. │ Nama            │ Merek         │ Lokasi        │ Kondisi  │ Stok  │ Harga           │ Tanggal beli    │")
		fmt.Println(paddingTable + "├─────┼─────────────────┼───────────────┼───────────────┼──────────┼───────┼─────────────────┼─────────────────┤")

		for i := 0; i < n; i++ {
			numberData = numberData + 1
			nama := potongString(data[i].nama, 15)
			merek := potongString(data[i].merek, 13)
			lokasi := potongString(data[i].lokasi, 13)
			kondisi := potongString(data[i].kondisi, 8)
			tanggal := potongString(data[i].tanggal, 15)
			
			line := fmt.Sprintf("│  %2d │ %-15s │ %-13s │ %-13s │ %-8s │ %5d │ %15.2f │ %-15s │", numberData, nama, merek, lokasi, kondisi, data[i].stok, data[i].harga, tanggal)
			fmt.Println(paddingTable + line)
		}

		fmt.Println(paddingTable + "└─────┴─────────────────┴───────────────┴───────────────┴──────────┴───────┴─────────────────┴─────────────────┘")
		fmt.Println()
		fmt.Print(padding + "Keluar? (tekan apapun lalu Enter)> ")
		fmt.Scan(&dummyInput)
		return
	}

}

// Prosedur untuk menampilkan menu untuk mengurutkan data
func urutData(data *arrInv, n int) {
	var pilihan, hasil int
	var dummyInput string
	var clear bool = true
	var input bool = false
	
	for {
		if clear {
			clearScreen()
		}
		headerTampilan()
		fmt.Println(padding + "║          Menu Utama > Urutkan data          ║")
		
		if n == 0 {
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
			pesanDataKosong()
			fmt.Print(padding + "Keluar? (tekan apapun lalu Enter)> ")
			fmt.Scan(&dummyInput)
			return
		} else {
			fmt.Println(padding + "╠═════════════════════════════════════════════╣")
			fmt.Println(padding + "║ 1. Urutkan dari nama (A-Z)                  ║")
			fmt.Println(padding + "║ 2. Urutkan dari nama (Z-A)                  ║")
			fmt.Println(padding + "║ 3. Urutkan dari merek (A-Z)                 ║")
			fmt.Println(padding + "║ 4. Urutkan dari merek (Z-A)                 ║")
			fmt.Println(padding + "║ 5. Urutkan dari lokasi (A-Z)                ║")
			fmt.Println(padding + "║ 6. Urutkan dari lokasi (Z-A)                ║")
			fmt.Println(padding + "║ 7. Urutkan dari kondisi (A-Z)               ║")
			fmt.Println(padding + "║ 8. Urutkan dari kondisi (Z-A)               ║")
			fmt.Println(padding + "║ 9. Urutkan dari stok (menaik)               ║")
			fmt.Println(padding + "║ 10. Urutkan dari stok (menurun)             ║")
			fmt.Println(padding + "║ 11. Urutkan dari harga (menaik)             ║")
			fmt.Println(padding + "║ 12. Urutkan dari harga (menurun)            ║")
			fmt.Println(padding + "║ 0. Kembali                                  ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
		
			
			fmt.Print(padding + "Pilih> ")
			fmt.Scan(&pilihan)
			
			selectionSort(data, n, pilihan)
			
			if pilihan >= 1 && pilihan <= 2 {
				pesanDataTerurut("nama   ")
				input = true
			} else if pilihan >= 3 && pilihan <= 4 {
				pesanDataTerurut("merek  ")
				input = true
			} else if pilihan >= 5 && pilihan <= 6 {
				pesanDataTerurut("lokasi ")
				input = true
			} else if pilihan >= 7 && pilihan <= 8 {
				pesanDataTerurut("kondisi")
				input = true
			} else if pilihan >= 9 && pilihan <= 10 {
				pesanDataTerurut("stok   ")
				input = true
			} else if pilihan >= 11 && pilihan <= 12 {
				pesanDataTerurut("harga  ")
				input = true
			} else if pilihan == 0{
				return
			} else {
				clear = false
				input = false
				fmt.Println()
				fmt.Println(padding + "╔═════════════════════════════════════════════╗")
				fmt.Println(padding + "║  [!] Pilihan tidak valid!, mohon coba lagi  ║")
				fmt.Println(padding + "╚═════════════════════════════════════════════╝")
			}
			
			
			if input {
				konfirmasiInput("Urutkan data lagi? (y/n)", &hasil)
			
				if hasil == 1 {
					clearScreen()
					return
				} else if hasil == 2 {
					clearScreen()
				} else {
					fmt.Println(padding + "╔═════════════════════════════════════════════╗")
					fmt.Println(padding + "║                  [!] Galat                  ║")
					fmt.Println(padding + "╚═════════════════════════════════════════════╝")
					return
				}
			}
		}
	}
}

// Prosedur untuk mengubah urutan data menggunakan Selection Sort
func selectionSort(data *arrInv, n, x int) {
	var i, idx int
	var temp inventaris
	var pass int = 1
	
	for pass < n {
		idx = pass - 1
		i = pass
		
		for i < n {
			switch x {
			case 1:
			if (*data)[i].nama < (*data)[idx].nama{
				idx = i
			}
			case 2:
			if (*data)[i].nama > (*data)[idx].nama{
				idx = i
			}
			case 3:
			if (*data)[i].merek < (*data)[idx].merek{
				idx = i
			}
			case 4:
			if (*data)[i].merek > (*data)[idx].merek{
				idx = i
			}
			case 5:
			if (*data)[i].lokasi < (*data)[idx].lokasi{
				idx = i
			}
			case 6:
			if (*data)[i].lokasi > (*data)[idx].lokasi{
				idx = i
			}
			case 7:
			if (*data)[i].kondisi < (*data)[idx].kondisi{
				idx = i
			}
			case 8:
			if (*data)[i].kondisi > (*data)[idx].kondisi{
				idx = i
			}
			case 9:
			if (*data)[i].stok < (*data)[idx].stok{
				idx = i
			}
			case 10:
			if (*data)[i].stok > (*data)[idx].stok{
				idx = i
			}
			case 11:
			if (*data)[i].harga < (*data)[idx].harga{
				idx = i
			}
			case 12:
			if (*data)[i].harga > (*data)[idx].harga{
				idx = i
			}
			}
			i++ 
		}
		
		temp = (*data)[pass - 1]
		(*data)[pass - 1] = (*data)[idx]
		(*data)[idx] = temp
		pass++
	}
}

// Prosedur untuk menampilkan menu dengan beberapa pilihan perhitungan dasar
func menuHitung(data arrInv, n int) {
	var pilihan, hasil int
	var dummyInput string
	var clear bool = true
	var input bool = false
		
	for {
		if clear {
			clearScreen()
		}
		headerTampilan()
		fmt.Println(padding + "║          Menu Utama > Urutkan data          ║")
		
		if n == 0 {
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
			pesanDataKosong()
			fmt.Print(padding + "Keluar? (tekan apapun lalu Enter)> ")
			fmt.Scan(&dummyInput)
			return
		} else {
			fmt.Println(padding + "╠═════════════════════════════════════════════╣")
			fmt.Println(padding + "║ 1. Hitung kondisi barang 'baik'             ║")
			fmt.Println(padding + "║ 2. Hitung kondisi barang 'buruk'            ║")
			fmt.Println(padding + "║ 0. Kembali                                  ║")
			fmt.Println(padding + "╚═════════════════════════════════════════════╝")
		
			
			fmt.Print(padding + "Pilih> ")
			fmt.Scan(&pilihan)
			
			if
			
			if input {
				konfirmasiInput("Hitung data lagi? (y/n)", &hasil)
			
				if hasil == 1 {
					clearScreen()
					return
				} else if hasil == 2 {
					clearScreen()
				} else {
					fmt.Println(padding + "╔═════════════════════════════════════════════╗")
					fmt.Println(padding + "║                  [!] Galat                  ║")
					fmt.Println(padding + "╚═════════════════════════════════════════════╝")
					return
				}
			}
		}
	}
}

func hitungKondisi(data arrInv, n int) int {
	var jumlahKondisi, N int
	
	jumlahKondisi = 1
	
	return jumlahKondisi
}

// Easter-egg :)
func catImage() {
	fmt.Println(padding + "░░░░░░░░░░░░░░░░░░░░░▄▀░░▌")
	fmt.Println(padding + "░░░░░░░░░░░░░░░░░░░▄▀▐░░░▌")
	fmt.Println(padding + "░░░░░░░░░░░░░░░░▄▀▀▒▐▒░░░▌")
	fmt.Println(padding + "░░░░░▄▀▀▄░░░▄▄▀▀▒▒▒▒▌▒▒░░▌")
	fmt.Println(padding + "░░░░▐▒░░░▀▄▀▒▒▒▒▒▒▒▒▒▒▒▒▒█")
	fmt.Println(padding + "░░░░▌▒░░░░▒▀▄▒▒▒▒▒▒▒▒▒▒▒▒▒▀▄")
	fmt.Println(padding + "░░░░▐▒░░░░░▒▒▒▒▒▒▒▒▒▌▒▐▒▒▒▒▒▀▄")
	fmt.Println(padding + "░░░░▌▀▄░░▒▒▒▒▒▒▒▒▐▒▒▒▌▒▌▒▄▄▒▒▐")
	fmt.Println(padding + "░░░▌▌▒▒▀▒▒▒▒▒▒▒▒▒▒▐▒▒▒▒▒█▄█▌▒▒▌")
	fmt.Println(padding + "░▄▀▒▐▒▒▒▒▒▒▒▒▒▒▒▄▀█▌▒▒▒▒▒▀▀▒▒▐░░░▄")
	fmt.Println(padding + "▀▒▒▒▒▌▒▒▒▒▒▒▒▄▒▐███▌▄▒▒▒▒▒▒▒▄▀▀▀▀")
	fmt.Println(padding + "▒▒▒▒▒▐▒▒▒▒▒▄▀▒▒▒▀▀▀▒▒▒▒▄█▀░░▒▌▀▀▄▄")
	fmt.Println(padding + "▒▒▒▒▒▒█▒▄▄▀▒▒▒▒▒▒▒▒▒▒▒░░▐▒▀▄▀▄░░░░▀")
	fmt.Println(padding + "▒▒▒▒▒▒▒█▒▒▒▒▒▒▒▒▒▄▒▒▒▒▄▀▒▒▒▌░░▀▄")
	fmt.Println(padding + "▒▒▒▒▒▒▒▒▀▄▒▒▒▒▒▒▒▒▀▀▀▀▒▒▒▄▀")
}

func byeImage() {
	fmt.Println(padding + "⠀⠀⠀⠀⠀⠀⠀⠀ ⠀⠀⠀⠀⠀⠀⠀⠀⠀ ⠀⠀⠀⠀⠀⠀⣀⣠⣤⣴⣶⣦⣤⣄⡀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣤⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⠀⠀⠀⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⠀⠀⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡆⠀⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⠓⢾⣿⣿⣹⢹⣿⡟⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣋⣑⡒⠦⣼⢻⣿⡇⢸⣿⠃⢸⣿⡟⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠀⠀⠀⠀⢀⣀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⠈⠉⠻⣷⣼⠚⡇⠃⢸⣿⠀⢸⣿⠉⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠙⠀⠙⠇⠀⢿⠀⢸⣿⣤⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠈⠉⠐⠂⠄⠀⠀⣿⡿⣿⣿⣿⣿⣿⣿⣿⢿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠛⠛⠛⠷⣿⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⣀⣠⡤⠤⠐⠺⣿⡏⢻⣿⣿⣿⣿⣿⣿⠀⢣⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀")
	fmt.Println(padding + "⠄⠂⠈⢀⣀⡠⠤⠒⠚⡟⣿⠙⣿⣿⣿⣿⣿⣿⡇⠀⠳⢤⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⡿⢿⣿⣿⣿⣿⢻⣿⣿⣿⠁⠀⠀⠀⠀")
	fmt.Println(padding + "⠤⠒⠉⠁⠀⠀⠀⠀⣀⣧⠼⡆⠻⣿⣿⢻⣿⣿⣿⣆⠀⠈⠣⠈⠑⠢⢄⠀⠀⠀⠀⠀⣠⣞⡯⢤⣿⣿⣿⣿⡧⣼⣿⣿⢻⡀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠀⢀⡠⠖⠉⠁⢸⡀⢹⡄⠙⢿⡟⣿⣿⣿⠈⢳⡀⠀⠀⠈⠉⠉⠀⢀⣠⡴⠊⢡⠀⠀⣸⣿⣿⣿⣿⠁⢿⣿⣏⡬⣧⠀⠀⠀")
	fmt.Println(padding + "⠀⢀⡠⠖⠉⠀⠀⠀⠀⠀⢳⠋⢈⠃⡌⠙⠻⣿⣿⡀⠈⢽⣦⠤⠤⠒⠒⠊⡹⠀⢱⠀⡏⢀⠜⣿⣿⡟⡹⠃⠀⠘⣿⡇⠀⣸⠀⠀⠀⠀")
	fmt.Println(padding + "⣴⡉⠀⠀⠀⠀⠀⠀⠀⢀⠇⠀⣠⡼⠀⠀⠀⠙⠎⠃⠀⠀⠈⠠⣄⠀⠀⢰⠁⠀⢈⡟⠉⠁⣴⡿⠋⠐⠁⣀⠞⡰⠉⠁⢸⡇⠀⠀⠀⠀")
	fmt.Println(padding + "⢇⢣⡀⠀⠀⠀⢠⠀⠀⠸⠀⢠⣿⠃⠳⣄⠀⠀⢰⠀⠀⠀⠀⠀⠀⠉⡴⠀⠀⠀⣸⠁⠀⠀⠁⢰⠀⡄⢠⠏⡰⠁⠀⠀⣿⠁⠀⠀⠀⠀")
	fmt.Println(padding + "⠈⠦⠳⡀⠀⠀⠀⣇⠀⡇⠀⢸⣿⠀⠀⠈⠑⠒⠚⠀⣦⠀⠀⠀⠀⣰⣁⠧⠔⠉⡏⠀⠀⠀⠀⢸⠀⡇⡜⣴⠁⠀⠀⠀⡟⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⣬⣦⡀⠀⠸⡀⡇⠀⠀⠙⣆⠀⠀⠀⠀⠀⢀⡈⠳⣀⣠⠞⠁⣀⠀⠀⡜⠀⠀⠀⠀⠀⡜⡆⣷⠟⠈⡄⠀⠀⢸⠃⠀⠀⠀⠀⠀")
	fmt.Println(padding + "⠀⠀⠀⠁⠙⠟⠦⡀⢳⡇⠀⠀⡆⠈⢆⠀⠀⠀⠀⠀⣀⣀⣈⣀⠀⠉⠀⠀⢸⡳⢄⠀⠀⠀⢀⠇⢱⡏⠀⠀⡇⠀⠀⣿⠀⠀⠀⠀⠀⠀")
}