package main

import (
	"fmt"
	"strings"
)

type Konten struct {
	ID          int
	Judul       string
	Kategori    string
	Link        string
	JumlahView  int
	JumlahLike  int
}

var daftarKonten [100]Konten
var jumlahKonten int = 0
var nextAvailableID int = 1 

func tambahKonten() {
	if jumlahKonten >= 100 {
		fmt.Println("Kapasitas penyimpanan konten penuh.")
		return
	}

	var judul, kategori, link string
	var jumlahView, jumlahLike int

	fmt.Print("Masukkan Judul Konten: ")
	fmt.Scanln(&judul)

	fmt.Print("Masukkan Kategori Konten: ")
	fmt.Scanln(&kategori)

	fmt.Print("Masukkan Link Konten: ")
	fmt.Scanln(&link)

	fmt.Print("Masukkan Jumlah View: ")
	fmt.Scanln(&jumlahView)

	fmt.Print("Masukkan Jumlah Like: ")
	fmt.Scanln(&jumlahLike)

	daftarKonten[jumlahKonten] = Konten{
		ID:          nextAvailableID, 
		Judul:       judul,
		Kategori:    kategori,
		Link:        link,
		JumlahView:  jumlahView,
		JumlahLike:  jumlahLike,
	}
	jumlahKonten++
	nextAvailableID++ 

	fmt.Println("Konten berhasil ditambahkan.")
}

func cariKonten(judul string) []Konten {
	var hasilPencarian []Konten
	judulLower := strings.ToLower(judul)

	for i := 0; i < jumlahKonten; i++ {
		kontenJudulLower := strings.ToLower(daftarKonten[i].Judul)
		if strings.Contains(kontenJudulLower, judulLower) {
			hasilPencarian = append(hasilPencarian, daftarKonten[i])
		}
	}
	return hasilPencarian
}

func tampilkanSemuaKonten() {
	if jumlahKonten == 0 {
		fmt.Println("Belum ada konten yang tersimpan.")
		return
	}
	fmt.Println("\nDaftar Semua Konten:")
	for i := 0; i < jumlahKonten; i++ {
		fmt.Printf("%d. Judul: %s, Kategori: %s, Link: %s, View: %d, Like: %d\n", daftarKonten[i].ID, daftarKonten[i].Judul, daftarKonten[i].Kategori, daftarKonten[i].Link, daftarKonten[i].JumlahView, daftarKonten[i].JumlahLike)
	}
}

func editKontenByJudul() {
	var judulCari string
	var judulBaru, kategoriBaru string
	found := false

	fmt.Print("Masukkan judul konten yang ingin diubah: ")
	fmt.Scanln(&judulCari)
	judulCariLower := strings.ToLower(judulCari)

	for i := 0; i < jumlahKonten; i++ {
		kontenJudulLower := strings.ToLower(daftarKonten[i].Judul)
		if strings.Contains(kontenJudulLower, judulCariLower) { 
			fmt.Print("Masukkan judul baru (kosongkan jika tidak ingin diubah): ")
			fmt.Scanln(&judulBaru)
			if len(judulBaru) > 0 { 
				daftarKonten[i].Judul = judulBaru
			}

			fmt.Print("Masukkan kategori baru (kosongkan jika tidak ingin diubah): ")
			fmt.Scanln(&kategoriBaru)
			if len(kategoriBaru) > 0 { 
				daftarKonten[i].Kategori = kategoriBaru
			}

			fmt.Println("Konten berhasil diubah.")
			found = true
			break 
		}
	}

	if !found {
		fmt.Println("Konten dengan judul tersebut tidak ditemukan.")
	}
}

func hapusKontenByJudul() {
	var judulHapus string
	foundIndex := -1

	fmt.Print("Masukkan judul konten yang ingin dihapus: ")
	fmt.Scanln(&judulHapus)
	judulHapusLower := strings.ToLower(judulHapus)

	for i := 0; i < jumlahKonten; i++ {
		kontenJudulLower := strings.ToLower(daftarKonten[i].Judul)
		if strings.Contains(kontenJudulLower, judulHapusLower) { 
			foundIndex = i
			break 
		}
	}

	if foundIndex != -1 {
		for i := foundIndex; i < jumlahKonten-1; i++ {
			daftarKonten[i] = daftarKonten[i+1]
		}
		jumlahKonten--


		for i := 0; i < jumlahKonten; i++ {
			daftarKonten[i].ID = i + 1
		}

		fmt.Println("Konten berhasil dihapus.")
	} else {
		fmt.Println("Konten dengan judul tersebut tidak ditemukan.")
	}
}

func selectionSortByLike(data *[100]Konten, n int, ascending bool) {
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if ascending {
				if data[j].JumlahLike < data[minIndex].JumlahLike {
					minIndex = j
				}
			} else {
				if data[j].JumlahLike > data[minIndex].JumlahLike {
					minIndex = j
				}
			}
		}
		temp := data[i]
		data[i] = data[minIndex]
		data[minIndex] = temp
	}
}

func urutkanKontenBerdasarkanLike() {
	if jumlahKonten == 0 {
		fmt.Println("Belum ada konten yang tersimpan.")
		return
	}

	var urutan string
	fmt.Print("Urutkan berdasarkan jumlah like (asc/desc): ")
	fmt.Scanln(&urutan)
	urutanLower := strings.ToLower(urutan)

	kontenUntukDiurutkan := daftarKonten
	
	ascending := (urutanLower == "asc")

	selectionSortByLike(&kontenUntukDiurutkan, jumlahKonten, ascending)

	fmt.Println("\nKonten diurutkan berdasarkan Jumlah Like (Selection Sort):")
	if urutanLower == "asc" || urutanLower == "desc" {
		if ascending {
			fmt.Println(" (Ascending)")
		} else {
			fmt.Println(" (Descending)")
		}
		for i := 0; i < jumlahKonten; i++ {
			fmt.Printf("Judul: %s, Like: %d\n", kontenUntukDiurutkan[i].Judul, kontenUntukDiurutkan[i].JumlahLike)
		}
	} else {
		fmt.Println("Pilihan urutan tidak valid.")
	}
}

func insertionSortByView(data *[100]Konten, n int, ascending bool) {
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		if ascending {
			for j >= 0 && data[j].JumlahView > key.JumlahView {
				data[j+1] = data[j]
				j--
			}
		} else {
			for j >= 0 && data[j].JumlahView < key.JumlahView {
				data[j+1] = data[j]
				j--
			}
		}
		data[j+1] = key
	}
}

func urutkanKontenBerdasarkanView() {
	if jumlahKonten == 0 {
		fmt.Println("Belum ada konten yang tersimpan.")
		return
	}

	var urutan string
	fmt.Print("Urutkan berdasarkan jumlah view (asc/desc): ")
	fmt.Scanln(&urutan)
	urutanLower := strings.ToLower(urutan)

	kontenUntukDiurutkan := daftarKonten

	ascending := (urutanLower == "asc")

	insertionSortByView(&kontenUntukDiurutkan, jumlahKonten, ascending)

	fmt.Println("\nKonten diurutkan berdasarkan Jumlah View (Insertion Sort):")
	if urutanLower == "asc" || urutanLower == "desc" {
		if ascending {
			fmt.Println(" (Ascending)")
		} else {
			fmt.Println(" (Descending)")
		}
		for i := 0; i < jumlahKonten; i++ {
			fmt.Printf("Judul: %s, View: %d\n", kontenUntukDiurutkan[i].Judul, kontenUntukDiurutkan[i].JumlahView)
		}
	} else {
		fmt.Println("Pilihan urutan tidak valid.")
	}
}

func binarySearch(minLike int) []Konten {
	if jumlahKonten == 0 {
		return []Konten{}
	}

	kontenUntukDiurutkan := daftarKonten
	selectionSortByLike(&kontenUntukDiurutkan, jumlahKonten, true) 

	var hasilPencarian []Konten
	low := 0
	high := jumlahKonten - 1
	firstIndex := -1 

	for low <= high {
		mid := low + (high-low)/2
		if kontenUntukDiurutkan[mid].JumlahLike >= minLike {
			firstIndex = mid
			high = mid - 1 
		} else {
			low = mid + 1
		}
	}

	if firstIndex != -1 {
		for i := firstIndex; i < jumlahKonten; i++ {
			hasilPencarian = append(hasilPencarian, kontenUntukDiurutkan[i])
		}
	}

	return hasilPencarian
}

func main() {
	var pilihan string

	for {
		fmt.Println("\n--- Manajemen Konten Kreator ---")
		fmt.Println("1. Tambah Konten")
		fmt.Println("2. Cari Konten berdasarkan Judul")
		fmt.Println("3. Tampilkan Semua Konten")
		fmt.Println("4. Urutkan Konten berdasarkan Like (Selection Sort)")
		fmt.Println("5. Urutkan Konten berdasarkan Jumlah View (Insertion Sort)")
		fmt.Println("6. Edit Konten berdasarkan Judul")
		fmt.Println("7. Hapus Konten berdasarkan Judul")
		fmt.Println("8. Cari Konten dengan Like Lebih dari 1000 (Binary Search)")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih opsi: ")

		fmt.Scanln(&pilihan)
		pilihanLower := strings.ToLower(pilihan)

		switch pilihanLower {
		case "1":
			tambahKonten()
		case "2":
			var judulCari string
			fmt.Print("Masukkan judul konten yang dicari: ")
			fmt.Scanln(&judulCari)
			hasil := cariKonten(judulCari)
			if len(hasil) > 0 { 
				fmt.Println("\nHasil Pencarian:")
				for i := 0; i < len(hasil); i++ { 
					fmt.Printf("ID: %d, Judul: %s, Kategori: %s, Link: %s, View: %d, Like: %d\n", hasil[i].ID, hasil[i].Judul, hasil[i].Kategori, hasil[i].Link, hasil[i].JumlahView, hasil[i].JumlahLike)
				}
			} else {
				fmt.Println("Konten dengan judul tersebut tidak ditemukan.")
			}
		case "3":
			tampilkanSemuaKonten()
		case "4":
			urutkanKontenBerdasarkanLike()
		case "5":
			urutkanKontenBerdasarkanView()
		case "6":
			editKontenByJudul()
		case "7":
			hapusKontenByJudul()
		case "8":
			hasil := binarySearch(1000)
			if len(hasil) > 0 { 
				fmt.Println("\nKonten dengan Jumlah Like lebih dari 1000:")
				for i := 0; i < len(hasil); i++ { 
					fmt.Printf("ID: %d, Judul: %s, Kategori: %s, Link: %s, View: %d, Like: %d\n", hasil[i].ID, hasil[i].Judul, hasil[i].Kategori, hasil[i].Link, hasil[i].JumlahView, hasil[i].JumlahLike)
				}
			} else {
				fmt.Println("Tidak ada konten dengan Jumlah Like lebih dari 1000.")
			}
		case "9":
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Opsi tidak valid. Silakan coba lagi.")
		}
	}
}
