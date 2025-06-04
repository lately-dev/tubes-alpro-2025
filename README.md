# tubes-alpro-2025

# Aplikasi Sistem Inventaris Laboratorium
Program ini merupakan aplikasi berbasis CLI (Command Line Interface) untuk mengelola data inventaris laboratorium yang diimplementasikan menggunakan bahasa pemograman Go (Golang).

## Deskripsi
Program ini dibangun untuk memenuhi kebutuhan pengelolaan inventaris yang secara tradisional kurang efisien dan kesalahan manusia dalam pencatatan data tersebut. Program ini termasuk pencatatan data inventaris, perhitungan kondisi dan nilai barang, pencarian dan pengurutan data inventaris berdasarkan kriteria yang dapat dipilih pengguna.

## Fitur-fitur
1. Manajemen data inventaris
    - Pencatatan inventaris dengan nama, merek, lokasi, kondisi, stok, harga dan tanggal pembelian barang
    - Mengubah atau memperbarui data tersimpan
    - Menghapus data tersimpan dari sistem
	- Menampilkan daftar data inventaris dengan antarmuka tabel yang rapi dan jelas
2. Pencarian data
    - Pencarian berdasarkan nama barang inventaris
3. Pengurutan data
    - Pengurutan menggunakan algoritma Selection Sort
	- Opsi pengurutan berdasarkan huruf (A-Z) dan angka melalui tipe inventaris yang tersedia
	- Dapat diurut secara menaik atau menurun
4. Perhitungan data
    - Mengkalkulasikan kondisi barang inventaris
	- Mengkalkulasikan nilai barang inventaris

## Struktur program
Program terdiri dari total 33 fungsi utama yang saling bekerja sama untuk menjalankan seluruh fitur aplikasi. Fungsi-fungsi utama tersebut antara lain:
- main — Fungsi utama aplikasi
- dataBaru — Menambah data inventaris baru
- ubahData — Mengubah data inventaris yang sudah ada
- hapusData — Menghapus data inventaris
- cariData — Mencari data inventaris berdasarkan kriteria tertentu
- tampilkanData — Menampilkan daftar data inventaris dalam format tabel
- urutData — Mengurutkan data inventaris berdasarkan pilihan pengguna
- menuHitung — Menu untuk melakukan perhitungan nilai dan kondisi inventaris

## Cara menggunakan program
1. Jalankan aplikasi melalui terminal/command prompt dengan perintah go run sistemInventaris.go
2. Pilih menu yang diinginkan untuk menambah, mengedit, menghapus, mencari, atau mengurutkan data inventaris
3. Masukkan data sesuai instruksi yang muncul di layar (contoh: format tanggal harus YYYY/MM/DD)
4. Gunakan fitur pencarian untuk mencari data berdasarkan nama atau tanggal pembelian barang
5. Gunakan fitur pengurutan untuk menampilkan data berdasarkan kriteria yang dipilih
6. Ikuti instruksi dan konfirmasi pada layar untuk melanjutkan atau keluar dari program

## Batasan Program
- Program menggunakan array statis dengan kapasitas hingga 1000 data inventaris
- Format tanggal wajib mengikuti YYYY/MM/DD untuk keakuratan pencarian dan pengurutan
- Antarmuka program berbasis teks (CLI), belum mendukung GUI

# Pembagian Tugas
## Hafiza
1. Dokumentasi kode 
2. Laporan perkembangan kode
3. Testing

## Fazli
1. Pengembangan program pusat
2. Pengujian dan debugging
