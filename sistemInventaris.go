package main
import (
	"fmt"
	"os"
    	"os/exec"
    	"runtime"
	//"bufio"
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
	alatLab[0] = inventaris{nama: "Baju", merek: "Lab", lokasi: "Ruang_1", kondisi: "Baik", tanggal: "2021/01/01", stok: 25, harga: 750000}
	alatLab[1] = inventaris{nama: "Penggaris", merek: "Butterfly", lokasi: "Rak_1", kondisi: "Baik", tanggal: "2020/03/22", stok: 100, harga: 15000}
	alatLab[2] = inventaris{nama: "Monitor", merek: "LG", lokasi: "Ruang_1", kondisi: "Baik", tanggal: "2019/03/17", stok: 50, harga: 820000}
	alatLab[3] = inventaris{nama: "Meja", merek: "none", lokasi: "Ruang_2", kondisi: "Sedang", tanggal: "2010/08/07", stok: 3, harga: 2400000}
	alatLab[4] = inventaris{nama: "Gunting", merek: "Kenko", lokasi: "Rak_2", kondisi: "Baik", tanggal: "2024/09/04", stok: 14, harga: 17500}
	alatLab[5] = inventaris{nama: "Kabel_HDMI", merek: "none", lokasi: "Ruang_3", kondisi: "Buruk", tanggal: "2011/12/09", stok: 4, harga: 55000}
	alatLab[6] = inventaris{nama: "NamaPanjangSekali", merek: "MerekPanjangSekali", lokasi: "LokasiPanjangSekali", kondisi: "KondisiPanjangSekali", tanggal: "0001/1/1", stok: 1000, harga: 55000}

	menu(&alatLab, &nData)
}

// Fungsi untuk membersihkan tampilan CMD/Terminal
func clearScreen() {
    var cmd *exec.Cmd //menampung data  dari fungsi di library os/exec
	if runtime.GOOS == "windows" { //mencari os windows
		cmd = exec.Command("cmd", "/c", "cls") // perintah untuk menghapus tampilan sebelumnya
	} else {
		cmd = exec.Command("clear") // perintah untuk menghapus tampilan sebelumnya tetapi untuk os selain windows  
	}
	cmd.Stdout = os.Stdout // untuk menampilkan output dari kondisi yang sudah di buat
	cmd.Run() // untuk mengeksekusi perintah
}

// Fungsi untuk mengubah warna tulisan terminal
func colorPrint(pesan string, color int) string {
	switch color {
	case 1:
		return "\033[31m" + pesan + "\033[0m" 	// Warna merah
	case 2:
		return "\033[32m" + pesan + "\033[0m"	// Warna hijau
	case 3:
		return "\033[33m" + pesan + "\033[0m"	// Warna kuning
	case 10:
		return "\033[43m\033[30m" + pesan + "\033[0m" // Latar kuning teks hitam
	default:
		return pesan
	}
}

// Prosedur untuk menampilkan pilihan menu utama
func menu(alatLab *arrInv, nData *int) { // mempunyai 2 parameter , tipe data arrInv dan int
	var pilihan string
	
	for { //perulangan terus menerus hingga user menginput 0
		clearScreen()
		headerTampilan()
		fmt.Println(padding + colorPrint("║                 Menu Utama                  ║", 3))
		fmt.Println(padding + colorPrint("╠══════════════════════╦══════════════════════╣", 3))
		fmt.Println(padding + colorPrint("║ 1. Data baru         ║ 6. Tampilkan data    ║", 3))
		fmt.Println(padding + colorPrint("║ 2. Ubah data         ║ 7. Urutkan data      ║", 3))
		fmt.Println(padding + colorPrint("║ 3. Hapus data        ║ 8. Perhitungan       ║", 3))
		fmt.Println(padding + colorPrint("║ 4. Cari data         ║ 0. Keluar            ║", 3))
		fmt.Println(padding + colorPrint("║ 5. Cari data (binary)║                      ║", 3))
		fmt.Println(padding + colorPrint("╚══════════════════════╩══════════════════════╝", 3))

		konfirmasiMenu(&pilihan, 1)

		switch pilihan {
		case "1":
			clearScreen()
			dataBaru(alatLab, nData)
		case "2":
			clearScreen()
			ubahData(alatLab, nData)
		case "3":
			clearScreen()
			hapusData(alatLab, nData)
		case "4":
			clearScreen()
			cariData(*alatLab, *nData)
		case "5":
			clearScreen()
			cariDataBinary(alatLab, *nData)
		case "6":
			clearScreen()
			tampilkanData(*alatLab, *nData)
		case "7":
			urutData(alatLab, *nData)
		case "8":
			clearScreen()
			menuHitung(*alatLab, *nData)
		case "0":
			clearScreen()
			fmt.Println()
			byeImage()
			pesanTemplate("        Terima Kasih dan Sampai Jumpa        ", 3)
			return
		default:
			return
		}
	} 
}

// Prosedur untuk menampilkan header tampilan
func headerTampilan() {
	fmt.Println()
	fmt.Println(padding + colorPrint("╔═════════════════════════════════════════════╗", 3))
	fmt.Println(padding + colorPrint("║   Aplikasi Sistem Inventaris Laboratorium   ║", 3))
	fmt.Println(padding + colorPrint("║        Created by hafiza and fazli          ║", 3))
	fmt.Println(padding + colorPrint("║        Algoritma Pemrograman 2025           ║", 3))
	fmt.Println(padding + colorPrint("╠═════════════════════════════════════════════╣", 3))
}

func footerTampilan(pesan string) {
	fmt.Printf(padding + colorPrint("║%45s║\n", 3), pesan)
	fmt.Println(padding + colorPrint("╚═════════════════════════════════════════════╝", 3))
}

// Prosedur untuk menampilkan pesan data kosong
func pesanDataKosong() {
	fmt.Println(padding + colorPrint("╔═════════════════════════════════════════════╗", 1))
	fmt.Println(padding + colorPrint("║   [!] Belum ada data, masukkan data baru    ║", 1))
	fmt.Println(padding + colorPrint("║         untuk menggunakan fitur ini         ║", 1))
	fmt.Println(padding + colorPrint("╚═════════════════════════════════════════════╝", 1))
}

// Prosedur untuk menampilkan pesan data berhasil diurut
func pesanDataTerurut(pesan string) {
	fmt.Println()
	fmt.Println(padding + colorPrint("╔═════════════════════════════════════════════╗", 2))
	fmt.Printf(padding + colorPrint("║  [✓] Data sudah diurut berdasarkan %-7s  ║\n", 2), pesan)
	fmt.Println(padding + colorPrint("╚═════════════════════════════════════════════╝", 2))
}

func pesanHitungData(jumlah, warna int, pesan string) {
	fmt.Println()
	fmt.Println(padding + colorPrint("╔═════════════════════════════════════════════╗", warna))
	fmt.Printf(padding + colorPrint("║    Ada %4d barang dengan kondisi %-6s    ║\n", warna), jumlah, pesan)
	fmt.Println(padding + colorPrint("╚═════════════════════════════════════════════╝", warna))
}

func pesanTemplate(pesan string, warna int) {
	fmt.Println()
	fmt.Println(padding + colorPrint("╔═════════════════════════════════════════════╗", warna))
	fmt.Printf(padding + colorPrint("║%45s║\n", warna), pesan)
	fmt.Println(padding + colorPrint("╚═════════════════════════════════════════════╝", warna))
}

func pesanGalat() {
	fmt.Println()
	fmt.Println(padding + "╔═════════════════════════════════════════════╗")
	fmt.Println(padding + "║                  [!] Galat                  ║")
	fmt.Println(padding + "╚═════════════════════════════════════════════╝")
}

// Prosedur untuk menampilkan data menggunakan template tabel
func tabelTampilan(data arrInv, idx int) {
	paddingTable := strings.Repeat(" ", 25)
	
	nama := potongString(data[idx].nama, 15)
    merek := potongString(data[idx].merek, 13)
    lokasi := potongString(data[idx].lokasi, 13)
    kondisi := potongString(data[idx].kondisi, 8)
    tanggal := potongString(data[idx].tanggal, 15)
	
	fmt.Println()
	fmt.Println(paddingTable + colorPrint("┌─────────────────┬───────────────┬───────────────┬──────────┬───────┬───────────────┬─────────────────┐", 3))
	fmt.Println(paddingTable + colorPrint("│ Nama            │ Merek         │ Lokasi        │ Kondisi  │ Stok  │ Harga         │ Tanggal beli    │", 3))
	fmt.Println(paddingTable + colorPrint("├─────────────────┼───────────────┼───────────────┼──────────┼───────┼───────────────┼─────────────────┤", 3))
	line := fmt.Sprintf("│ %-15s │ %-13s │ %-13s │ %-8s │ %5d │ %13.2f │ %-15s │", nama, merek, lokasi, kondisi, data[idx].stok, data[idx].harga, tanggal)
	fmt.Println(paddingTable + colorPrint(line, 3))
	fmt.Println(paddingTable + colorPrint("└─────────────────┴───────────────┴───────────────┴──────────┴───────┴───────────────┴─────────────────┘", 3))
}

// Fungsi untuk memotong string kata sesuai parameter maksKata sehingga tidak merusak antarmuka
func potongString(kata string, maksKata int) string {
	if len(kata) > maksKata {
		return kata[:maksKata]
	}
	return kata
}

// Prosedur menjalankan konfirmasi apakah pilihan menu ini tetap dijalankan atau dihentikan
func konfirmasiInput(pesan string, hasil *int) { //prosedur ini mengembsalikan nilai hasil jadi 1 jika user input n dan 2 jika user input y
	var konfirmasi string // variabel baru untuk menyimpan y/n

	for {
		fmt.Printf(padding + colorPrint("%s> ", 3), pesan) //menampilkan pesan
		fmt.Scan(&konfirmasi)
		
		if konfirmasi == "n" || konfirmasi == "N" {
			*hasil = 1 //nilai hasil jadi 1
			return
		} else if konfirmasi == "y" || konfirmasi == "Y" {
			*hasil = 2
			return
		} else {
			pesanTemplate("         [!] Pilihan hanya berupa y/n        ", 1)
		}
	}
}

// Prosedur menjalankan konfirmasi akan pilihan menu yang hanya tersedia, mode dipakai untuk modularitas ketika dipakai pada menu lain
func konfirmasiMenu(pilihan *string, mode int) {
	for {
		if mode == 1 {
			fmt.Print(padding + colorPrint("Pilih menu> ", 3))
		} else {
			fmt.Print(padding + colorPrint("Pilih> ", 3))
		}
		fmt.Scan(pilihan)
		
		if mode == 1 {
			switch *pilihan {
			case "0", "1", "2", "3", "4", "5", "6", "7", "8":
				return
			default:
				pesanTemplate("  [!] Pilihan tidak valid!, mohon coba lagi  ", 1)
			}
		} else if mode == 2 {
			switch *pilihan {
			case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14":
				return
			default:
				pesanTemplate("  [!] Pilihan tidak valid!, mohon coba lagi  ", 1)
			}
		} else if mode == 3 {
			switch *pilihan {
			case "0","1", "2", "3", "4", "5":
				return
			default:
				pesanTemplate("  [!] Pilihan tidak valid!, mohon coba lagi  ", 1)
			}
		}
	}
}

// Prosedur untuk menambahkan data inventaris baru
func dataBaru(data *arrInv, n *int) {
	var hasil int
	
	for {
		if *n >= maxArr {
			fmt.Println()
			fmt.Println(padding + colorPrint("╔═════════════════════════════════════════════╗", 1))
			fmt.Println(padding + colorPrint("║                     [!]                     ║", 1))
			fmt.Println(padding + colorPrint("║ Data sudah penuh! Anda bisa hapus data lain ║", 1))
			fmt.Println(padding + colorPrint("║           untuk menambah data baru          ║", 1))
			fmt.Println(padding + colorPrint("║    (ketik apapun dan Enter untuk keluar)    ║", 1))
			fmt.Println(padding + colorPrint("╚═════════════════════════════════════════════╝", 1))
			
			var dummyInput string
			fmt.Scan(&dummyInput)
			return
		}
		
		headerTampilan()
		footerTampilan("            Menu Utama > Data baru           ")
		pesanTemplate("      Gunakan (_) underscore untuk spasi     ", 3)
		fmt.Print(padding +  "Nama alat\t : ")
		fmt.Scan(&data[*n].nama)	
		fmt.Print(padding + "Merek alat\t : ")
		fmt.Scan(&data[*n].merek)
		fmt.Print(padding + "Lokasi alat\t : ")
		fmt.Scan(&data[*n].lokasi)
		fmt.Print(padding + "Kondisi alat\t : ")
		fmt.Scan(&data[*n].kondisi)
		/*if data[*n].kondisi != "Baik" || data[*n].kondisi != "baik" || data[*n].kondisi != "Sedang" || data[*n].kondisi != "sedang" || data[*n].kondisi != "Buruk" || data[*n].kondisi != "buruk" {
			for {
			
			}
		}*/
		fmt.Print(padding + "Stok alat\t : ")
		fmt.Scan(&data[*n].stok)
		fmt.Print(padding + "Harga alat\t : ")
		fmt.Scan(&data[*n].harga)
		fmt.Print(padding + "Tanggal pembelian (tttt/bb/hh): ")
		fmt.Scan(&data[*n].tanggal)
		
		*n = *n + 1
		
		fmt.Println()
		konfirmasiInput("Masukkan data baru lagi? (y/n)", &hasil)
		
		if hasil == 1 { // hasil menjadi 1 dan keluar dari program
			return
		} else if hasil == 2 { // jika user input y maka hasil jadi 2 , clearscreen dan mengulang program
			clearScreen()
		} else {
			pesanGalat()  // jika hasilnya bukan 1 atau 2 maka akan muncul pesan galat maka akan keluar dari program
			return
		}
	}
}

// Prosedur untuk mengubah data dengan menimpa yang lama dengan yang baru
func ubahData(data *arrInv, n *int) {
	var dataDicari string
	var dataDiubah [1]inventaris
	var hasil, idx int
	
	for {
		headerTampilan()
		footerTampilan("            Menu Utama > Ubah data           ")
		fmt.Print(padding + "Cari nama data: ")
		fmt.Scan(&dataDicari)
		
		idx = sequentialSearch(*data, *n, dataDicari)
		
		if idx >= 0 {
			tabelTampilan(*data, idx)
			fmt.Println(padding + colorPrint("╔═════════════════════════════════════════════╗", 3))
			fmt.Println(padding + colorPrint("║     [i] Isi perubahan data di bawah ini     ║", 3))
			fmt.Println(padding + colorPrint("║      Gunakan (_) underscore untuk spasi     ║", 3))
			fmt.Println(padding + colorPrint("╚═════════════════════════════════════════════╝", 3))
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
			fmt.Print(padding + "Tanggal pembelian (tttt/bb/hh): ")
			fmt.Scan(&dataDiubah[0].tanggal)
			
			data[idx].nama = dataDiubah[0].nama
			data[idx].merek = dataDiubah[0].merek 
			data[idx].lokasi = dataDiubah[0].lokasi
			data[idx].kondisi = dataDiubah[0].kondisi
			data[idx].tanggal = dataDiubah[0].tanggal
			data[idx].stok = dataDiubah[0].stok
			data[idx].harga = dataDiubah[0].harga
			
			fmt.Println()
			konfirmasiInput("Ubah data lagi? (y/n)", &hasil)
			
		} else {
			pesanTemplate("          [!] Data tidak ditemukan           ", 1)
			konfirmasiInput("Cari data lagi? (y/n)", &hasil)
		}
		
		if hasil == 1 {
			clearScreen()
			return
		} else if hasil == 2 {
			clearScreen()
		} else {
			pesanGalat()
			return
		}
	}
}

// Prosedur untuk menghapus data dengan menimpa data dihapus dan perulangan untuk memindahkannya ke kiri dan mengubah nilai data aktif
func hapusData(data *arrInv, n *int) {
	var dataDicari string
	var idx, hasil int
	
	for {
		headerTampilan()
		footerTampilan("           Menu Utama > Hapus data           ")
		fmt.Print(padding + "Cari nama data: ")
		fmt.Scan(&dataDicari)
		
		idx = sequentialSearch(*data, *n, dataDicari)
		
		if idx >= 0 {
			tabelTampilan(*data, idx)
			fmt.Println()
			konfirmasiInput("[!] Hapus? (y/n)", &hasil)
			
			if hasil == 2 {
				for i := idx; i < *n-1; i++ {
					(*data)[i] = (*data)[i+1]
				}
				(*n)--
				
				pesanTemplate("         [✓] Data berhasil di hapus          ", 2)
				konfirmasiInput("Hapus data lagi? (y/n)", &hasil)
			} else {
				pesanTemplate("       [!] Penghapusan data dibatalkan       ", 1)
				konfirmasiInput("Cari data lagi? (y/n)", &hasil)	
			}
		} else {
			pesanTemplate("          [!] Data tidak ditemukan           ", 1)
			konfirmasiInput("Cari data lagi? (y/n)", &hasil)
		}
		
		if hasil == 1 {
			clearScreen()
			return
		} else if hasil == 2 {
			clearScreen()
		} else {
			pesanGalat()
			return
		}
	}
}

// Prosedur untuk mencari data dengan menggunakan logika if-else dan sequentialSearch ketika data ditemukan
func cariData(data arrInv, n int) {
	var x string
	var idx, hasil int
	
	for {
		headerTampilan()
		footerTampilan("           Menu Utama > Cari data            ")
		fmt.Print(padding + "Cari nama data: ")
		fmt.Scan(&x)
		
		idx = sequentialSearch(data, n, x)
		
		if idx >= 0 {
			tabelTampilan(data, idx)
			pesanTemplate("             [✓] Data ditemukan              ", 2)
		} else {
			pesanTemplate("          [!] Data tidak ditemukan           ", 1)
		}
		
		konfirmasiInput("Cari data lagi? (y/n)", &hasil)
		
		if hasil == 1 {
			clearScreen()
			return
		} else if hasil == 2 {
			clearScreen()
		} else {
			pesanGalat()
			return
		}
	}
}

// Prosedur untuk mencari data dengan menggunakan logika if-else dan binarySearchTanggal ketika data ditemukan
func cariDataBinary(data *arrInv, n int) {headerTampilan()
	var x string
	var idx, hasil int
	
	for {
		footerTampilan("Menu Utama > Cari data dengan tanggal(binary)")
		fmt.Print(padding + "Masukkan tanggal yang dicari (format tttt/bb/hh): ")
		fmt.Scan(&x)
		
		for i := 0; i < n; i++ {
			(*data)[i].tanggal = padTanggal((*data)[i].tanggal)
		}
		x = padTanggal(x)
		
		insertionSortForBinary(data, n)
		idx = binarySearchTanggal(data, n, x)

		
		if idx >= 0 {
			tabelTampilan(*data, idx)
			pesanTemplate("             [✓] Data ditemukan              ", 2)
		} else {
			pesanTemplate("          [!] Data tidak ditemukan           ", 1)
		}
		
		konfirmasiInput("Cari data lagi? (y/n)", &hasil)
		
		if hasil == 1 {
			clearScreen()
			return
		} else if hasil == 2 {
			clearScreen()
		} else {
			pesanGalat()
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

// Fungsi untuk normalisasi tanggal menjadi format "YYYY/MM/DD"
func padTanggal(t string) string {
	var y, m, d int
	fmt.Sscanf(t, "%d/%d/%d", &y, &m, &d)
	return fmt.Sprintf("%04d/%02d/%02d", y, m, d)
}

// Fungsi binary search dengan insert posisi (juga dipakai di insertion sort)
func binarySearchTanggalInsertPos(data *arrInv, item inventaris, low, high int) int {
	for low <= high {
		mid := low + (high-low)/2
		if (*data)[mid].tanggal < item.tanggal {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

// Fungsi pencarian Binary berdasarkan tanggal
func binarySearchTanggal(data *arrInv, n int, target string) int {
	var low int = 0
	var high int = n - 1
	var mid int
    for low <= high {
        mid = low + (high-low)/2
        if (*data)[mid].tanggal == target {
            return mid 
        } else if (*data)[mid].tanggal < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1
}

// Fungsi untuk pengurutan khusus binary
func insertionSortForBinary(data *arrInv, n int) {
	var i, j, loc int
	var temp inventaris

	for i = 1; i < n; i++ {
		temp = (*data)[i]
		loc = binarySearchTanggalInsertPos(data, temp, 0, i-1)
		j = i - 1
		for j >= loc {
			(*data)[j+1] = (*data)[j]
			j--
		}
		(*data)[loc] = temp
	}
}

// Prosedur untuk melihat seluruh data aktif dengan template tabel
func tampilkanData(data arrInv, n int) {
	/*var idxAwal int = 0
	var idxAkhir int = 10*/
	var numberData = 0
	var dummyInput string
	paddingTable := strings.Repeat(" ", 22)
	
	headerTampilan()
	footerTampilan("         Menu Utama > Tampilkan data         ")
	fmt.Println()
	
	if n == 0 {
		pesanDataKosong()
		fmt.Print(padding + colorPrint("Keluar? (tekan apapun lalu Enter)> ", 3))
		fmt.Scan(&dummyInput)
		return
	} else {
		fmt.Println(paddingTable + colorPrint("┌─────┬─────────────────┬───────────────┬───────────────┬──────────┬───────┬─────────────────┬─────────────────┐", 3))
		fmt.Println(paddingTable + colorPrint("│ No. │ Nama            │ Merek         │ Lokasi        │ Kondisi  │ Stok  │ Harga           │ Tanggal beli    │", 3))
		fmt.Println(paddingTable + colorPrint("├─────┼─────────────────┼───────────────┼───────────────┼──────────┼───────┼─────────────────┼─────────────────┤", 3))

		for i := 0; i < n; i++ {
			numberData = numberData + 1
			nama := potongString(data[i].nama, 15)
			merek := potongString(data[i].merek, 13)
			lokasi := potongString(data[i].lokasi, 13)
			kondisi := potongString(data[i].kondisi, 8)
			tanggal := potongString(data[i].tanggal, 15)
			
			line := fmt.Sprintf("│  %2d │ %-15s │ %-13s │ %-13s │ %-8s │ %5d │ %15.2f │ %-15s │", numberData, nama, merek, lokasi, kondisi, data[i].stok, data[i].harga, tanggal)
			fmt.Println(paddingTable + colorPrint(line, 3))
		}

		fmt.Println(paddingTable + colorPrint("└─────┴─────────────────┴───────────────┴───────────────┴──────────┴───────┴─────────────────┴─────────────────┘", 3))
		fmt.Println()
		fmt.Print(padding + colorPrint("Keluar? (tekan apapun lalu Enter)> ", 3))		
		fmt.Scan(&dummyInput)
		return
	}

}

// Prosedur untuk menampilkan menu untuk mengurutkan data
func urutData(data *arrInv, n int) {
	var hasil int
	var pilihan, dummyInput string
	
	for {
		clearScreen()
		headerTampilan()
		fmt.Println(padding + colorPrint("║          Menu Utama > Urutkan data          ║", 3))
		
		if n == 0 {
			fmt.Println(padding + colorPrint("╚═════════════════════════════════════════════╝", 3))
			fmt.Println()
			pesanDataKosong()
			fmt.Print(padding + colorPrint("Keluar? (tekan apapun lalu Enter)> ", 3))
			fmt.Scan(&dummyInput)
			return
		} else {
			fmt.Println(padding + colorPrint("╠═════════════════════════════════════════════╣", 3))
			fmt.Println(padding + colorPrint("║               Selection Sort                ║", 3))
			fmt.Println(padding + colorPrint("║ 1. Urutkan dari nama (A-Z)                  ║", 3))
			fmt.Println(padding + colorPrint("║ 2. Urutkan dari nama (Z-A)                  ║", 3))
			fmt.Println(padding + colorPrint("║ 3. Urutkan dari merek (A-Z)                 ║", 3))
			fmt.Println(padding + colorPrint("║ 4. Urutkan dari merek (Z-A)                 ║", 3))
			fmt.Println(padding + colorPrint("║ 5. Urutkan dari lokasi (A-Z)                ║", 3))
			fmt.Println(padding + colorPrint("║ 6. Urutkan dari lokasi (Z-A)                ║", 3))
			fmt.Println(padding + colorPrint("║ 7. Urutkan dari kondisi (A-Z)               ║", 3))
			fmt.Println(padding + colorPrint("║ 8. Urutkan dari kondisi (Z-A)               ║", 3))
			fmt.Println(padding + colorPrint("╠═════════════════════════════════════════════╣", 3))
			fmt.Println(padding + colorPrint("║               Insertion Sort                ║", 3))
			fmt.Println(padding + colorPrint("║ 9. Urutkan dari stok (menaik)               ║", 3))
			fmt.Println(padding + colorPrint("║ 10. Urutkan dari stok (menurun)             ║", 3))
			fmt.Println(padding + colorPrint("║ 11. Urutkan dari harga (menaik)             ║", 3))
			fmt.Println(padding + colorPrint("║ 12. Urutkan dari harga (menurun)            ║", 3))
			fmt.Println(padding + colorPrint("║ 13. Urutkan dari tanggal (menaik)           ║", 3))
			fmt.Println(padding + colorPrint("║ 14. Urutkan dari tanggal (menurun)          ║", 3))
			fmt.Println(padding + colorPrint("║ 0. Kembali                                  ║", 3))
			fmt.Println(padding + colorPrint("╚═════════════════════════════════════════════╝", 3))
		
			konfirmasiMenu(&pilihan, 2)
						
			switch pilihan {
			case "1":
				selectionSort(data, n, "1")
				pesanDataTerurut("nama   ")
			case "2":
				selectionSort(data, n, "2")
				pesanDataTerurut("nama   ")
			case "3":
				selectionSort(data, n, "3")
				pesanDataTerurut("merek  ")			
			case "4":
				selectionSort(data, n, "4")
				pesanDataTerurut("merek  ")
			case "5":
				selectionSort(data, n, "5")
				pesanDataTerurut("lokasi ")
			case "6":
				selectionSort(data, n, "6")
				pesanDataTerurut("lokasi ")
			case "7":
				selectionSort(data, n, "7")
				pesanDataTerurut("kondisi")
			case "8":
				selectionSort(data, n, "8")
				pesanDataTerurut("kondisi")
			case "9":
				insertionSort(data, n, "1")
				pesanDataTerurut("stok   ")
			case "10":
				insertionSort(data, n, "2")
				pesanDataTerurut("stok   ")
			case "11":
				insertionSort(data, n, "3")
				pesanDataTerurut("harga  ")
			case "12":
				insertionSort(data, n, "4")
				pesanDataTerurut("harga  ")
			case "13":
				insertionSort(data, n, "5")
				pesanDataTerurut("tanggal")
			case "14":
				insertionSort(data, n, "6")
				pesanDataTerurut("tanggal")
			case "0":
				return
			default:
				return
			}
			
			konfirmasiInput("Urutkan data lagi? (y/n)", &hasil)
		
			if hasil == 1 {
				clearScreen()
				return
			} else if hasil == 2 {
				clearScreen()
			} else {
				pesanGalat()
				return
			}
		}
	}
}

// Prosedur untuk mengubah urutan data menggunakan Selection Sort
func selectionSort(data *arrInv, n int, x string) {
	var i, idx int
	var temp inventaris
	var pass int = 1
	
	for pass < n {
		idx = pass - 1
		i = pass
		
		for i < n {
			switch x {
			case "1":
				if (*data)[i].nama < (*data)[idx].nama{
					idx = i
				}
			case "2":
				if (*data)[i].nama > (*data)[idx].nama{
					idx = i
				}
			case "3":
				if (*data)[i].merek < (*data)[idx].merek{
					idx = i
				}
			case "4":
				if (*data)[i].merek > (*data)[idx].merek{
					idx = i
				}
			case "5":
				if (*data)[i].lokasi < (*data)[idx].lokasi{
					idx = i
				}
			case "6":
				if (*data)[i].lokasi > (*data)[idx].lokasi{
					idx = i
				}
			case "7":
				if (*data)[i].kondisi < (*data)[idx].kondisi{
					idx = i
				}
			case "8":
				if (*data)[i].kondisi > (*data)[idx].kondisi{
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

// Prosedur untuk mengubah urutan data menggunakan Insertion Sort
func insertionSort(data *arrInv, n int, x string) {
	var i, j int
	var temp inventaris
	
	for i = 1; i < n; i++ {
		temp = data[i]
		j = i - 1
		
		switch x {
		case "1":
			for j >= 0 && data[j].stok > temp.stok {
				data[j+1] = data[j]
				j--
			}
		case "2":
			for j >= 0 && data[j].stok < temp.stok {
				data[j+1] = data[j]
				j--
			}
		case "3":
			for j >= 0 && data[j].harga > temp.harga {
				data[j+1] = data[j]
				j--
			}
		case "4":
			for j >= 0 && data[j].harga < temp.harga {
				data[j+1] = data[j]
				j--
			}
		case "5":
			for j >= 0 && data[j].tanggal > temp.tanggal {
				data[j+1] = data[j]
				j--
			}
		case "6":
			for j >= 0 && data[j].tanggal < temp.tanggal {
				data[j+1] = data[j]
				j--
			}
		}
		data[j+1] = temp	
	}
}

// Prosedur untuk menampilkan menu dengan beberapa pilihan perhitungan dasar
func menuHitung(data arrInv, n int) {
	var hasil int
	var pilihan, dummyInput string
		
	for {
		clearScreen()
		headerTampilan()
		fmt.Println(padding + colorPrint("║          Menu Utama > Urutkan data          ║", 3))
		
		if n == 0 {
			fmt.Println(padding + colorPrint("╚═════════════════════════════════════════════╝", 3))
			pesanDataKosong()
			fmt.Print(padding + colorPrint("Keluar? (tekan apapun lalu Enter)> ", 3))
			fmt.Scan(&dummyInput)
			return
		} else {
			fmt.Println(padding + colorPrint("╠═════════════════════════════════════════════╣", 3))
			fmt.Println(padding + colorPrint("║ 1. Hitung kondisi barang 'baik'             ║", 3))
			fmt.Println(padding + colorPrint("║ 2. Hitung kondisi barang 'sedang'           ║", 3))
			fmt.Println(padding + colorPrint("║ 3. Hitung kondisi barang 'buruk'            ║", 3))
			fmt.Println(padding + colorPrint("║ 4. Hitung nilai kerugian barang 'buruk'     ║", 3))
			fmt.Println(padding + colorPrint("║ 0. Kembali                                  ║", 3))
			fmt.Println(padding + colorPrint("╚═════════════════════════════════════════════╝", 3))

			konfirmasiMenu(&pilihan, 3)
			
			switch pilihan {
			case "1":
				pesanHitungData(hitungKondisi(data, n, "Baik"), 2, "baik  ")
			case "2":
				pesanHitungData(hitungKondisi(data, n, "Sedang"), 3, "sedang")
			case "3":
				pesanHitungData(hitungKondisi(data, n, "Buruk"), 1, "buruk ")
			case "4":
				fmt.Println()
				fmt.Println(padding + colorPrint("╔═════════════════════════════════════════════╗", 1))
				fmt.Println(padding + colorPrint("║      Kerugian barang 'buruk' mencapai       ║", 1))
				fmt.Printf(padding + colorPrint("║               %15.2f               ║\n", 1), hitungKerugian(data, n))
				fmt.Println(padding + colorPrint("╚═════════════════════════════════════════════╝", 1))
			case "0":
				return
			}
			
			konfirmasiInput("Hitung data lagi? (y/n)", &hasil)
			
			if hasil == 1 {
				clearScreen()
				return
			} else if hasil == 2 {
				clearScreen()
			} else {
				pesanGalat()
				return
			}
		}
	}
}

func hitungKondisi(data arrInv, n int, kondisi string) int {
	var jumlahKondisi int
	
	for i := 0; i < n; i++ {
		if data[i].kondisi == kondisi {
			jumlahKondisi = jumlahKondisi + data[i].stok
		}
	}
	
	return jumlahKondisi
}

func hitungKerugian(data arrInv, n int) float64 {
	var jumlahRugi float64
	
	for i := 0; i < n; i++ {
		if data[i].kondisi == "Buruk" {
			jumlahRugi = jumlahRugi + (data[i].harga * float64(hitungKondisi(data, n, "Buruk")))
		}
	}
	
	return jumlahRugi
}

// Easter-egg :)
func byeImage() {
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⠀⠀⠀⠀ ⠀⠀⠀⠀⠀⠀⠀⠀⠀ ⠀⠀⠀⠀⠀⠀⣀⣠⣤⣴⣶⣦⣤⣄⡀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣤⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⠀⠀⠀⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⠀⠀⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡆⠀⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⠓⢾⣿⣿⣹⢹⣿⡟⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣋⣑⡒⠦⣼⢻⣿⡇⢸⣿⠃⢸⣿⡟⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⠀⠀⠀⢀⣀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⠈⠉⠻⣷⣼⠚⡇⠃⢸⣿⠀⢸⣿⠉⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠙⠀⠙⠇⠀⢿⠀⢸⣿⣤⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠈⠉⠐⠂⠄⠀⠀⣿⡿⣿⣿⣿⣿⣿⣿⣿⢿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠛⠛⠛⠷⣿⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⣀⣠⡤⠤⠐⠺⣿⡏⢻⣿⣿⣿⣿⣿⣿⠀⢣⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠄⠂⠈⢀⣀⡠⠤⠒⠚⡟⣿⠙⣿⣿⣿⣿⣿⣿⡇⠀⠳⢤⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⡿⢿⣿⣿⣿⣿⢻⣿⣿⣿⠁⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠤⠒⠉⠁⠀⠀⠀⠀⣀⣧⠼⡆⠻⣿⣿⢻⣿⣿⣿⣆⠀⠈⠣⠈⠑⠢⢄⠀⠀⠀⠀⠀⣠⣞⡯⢤⣿⣿⣿⣿⡧⣼⣿⣿⢻⡀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠀⢀⡠⠖⠉⠁⢸⡀⢹⡄⠙⢿⡟⣿⣿⣿⠈⢳⡀⠀⠀⠈⠉⠉⠀⢀⣠⡴⠊⢡⠀⠀⣸⣿⣿⣿⣿⠁⢿⣿⣏⡬⣧⠀⠀⠀ ", 10))
	fmt.Println(padding + colorPrint("⠀⢀⡠⠖⠉⠀⠀⠀⠀⠀⢳⠋⢈⠃⡌⠙⠻⣿⣿⡀⠈⢽⣦⠤⠤⠒⠒⠊⡹⠀⢱⠀⡏⢀⠜⣿⣿⡟⡹⠃⠀⠘⣿⡇⠀⣸⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⣴⡉⠀⠀⠀⠀⠀⠀⠀⢀⠇⠀⣠⡼⠀⠀⠀⠙⠎⠃⠀⠀⠈⠠⣄⠀⠀⢰⠁⠀⢈⡟⠉⠁⣴⡿⠋⠐⠁⣀⠞⡰⠉⠁⢸⡇⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⢇⢣⡀⠀⠀⠀⢠⠀⠀⠸⠀⢠⣿⠃⠳⣄⠀⠀⢰⠀⠀⠀⠀⠀⠀⠉⡴⠀⠀⠀⣸⠁⠀⠀⠁⢰⠀⡄⢠⠏⡰⠁⠀⠀⣿⠁⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠈⠦⠳⡀⠀⠀⠀⣇⠀⡇⠀⢸⣿⠀⠀⠈⠑⠒⠚⠀⣦⠀⠀⠀⠀⣰⣁⠧⠔⠉⡏⠀⠀⠀⠀⢸⠀⡇⡜⣴⠁⠀⠀⠀⡟⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⣬⣦⡀⠀⠸⡀⡇⠀⠀⠙⣆⠀⠀⠀⠀⠀⢀⡈⠳⣀⣠⠞⠁⣀⠀⠀⡜⠀⠀⠀⠀⠀⡜⡆⣷⠟⠈⡄⠀⠀⢸⠃⠀⠀⠀⠀⠀", 10))
	fmt.Println(padding + colorPrint("⠀⠀⠀⠁⠙⠟⠦⡀⢳⡇⠀⠀⡆⠈⢆⠀⠀⠀⠀⠀⣀⣀⣈⣀⠀⠉⠀⠀⢸⡳⢄⠀⠀⠀⢀⠇⢱⡏⠀⠀⡇⠀⠀⣿⠀⠀⠀⠀⠀⠀", 10))
}
