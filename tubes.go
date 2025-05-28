package main

import (
	"fmt"
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
		ID:          jumlahKonten + 1,
		Judul:       judul,
		Kategori:    kategori,
		Link:        link,
		JumlahView:  jumlahView,
		JumlahLike:  jumlahLike,
	}
	jumlahKonten++

	fmt.Println("Konten berhasil ditambahkan.")
}

func cariKonten(judul string) []Konten {
	var hasilPencarian [100]Konten
	hasilCount := 0
	judulLower := ""
	for i := 0; true; i++ {
		if i >= func() int { count := 0; for range judul { count++ }; return count }() {
			break
		}
		char := judul[i]
		if char >= 'A' && char <= 'Z' {
			judulLower += string(char + ('a' - 'A'))
		} else {
			judulLower += string(char)
		}
	}
	for i := 0; true; i++ {
		if i >= jumlahKonten {
			break
		}
		kontenJudulLower := ""
		kontenJudul := daftarKonten[i].Judul
		for j := 0; true; j++ {
			if j >= func() int { count := 0; for range kontenJudul { count++ }; return count }() {
				break
			}
			char := kontenJudul[j]
			if char >= 'A' && char <= 'Z' {
				kontenJudulLower += string(char + ('a' - 'A'))
			} else {
				kontenJudulLower += string(char)
			}
		}
		n := 0
		for range kontenJudulLower {
			n++
		}
		m := 0
		for range judulLower {
			m++
		}
		if m == n {
			match := true
			for k := 0; true; k++ {
				if k >= m {
					break
				}
				if kontenJudulLower[k] != judulLower[k] {
					match = false
					break
				}
			}
			if match {
				hasilPencarian[hasilCount] = daftarKonten[i]
				hasilCount++
			}
		}
	}
	return hasilPencarian[:hasilCount]
}

func tampilkanSemuaKonten() {
	if jumlahKonten == 0 {
		fmt.Println("Belum ada konten yang tersimpan.")
		return
	}
	fmt.Println("\nDaftar Semua Konten:")
	for i := 0; true; i++ {
		if i >= jumlahKonten {
			return
		}
		fmt.Printf("%d. Judul: %s, Kategori: %s, Link: %s, View: %d, Like: %d\n", daftarKonten[i].ID, daftarKonten[i].Judul, daftarKonten[i].Kategori, daftarKonten[i].Link, daftarKonten[i].JumlahView, daftarKonten[i].JumlahLike)
	}
}

func editKontenByJudul() {
	var judulCari string
	var judulBaru, kategoriBaru string
	found := false

	fmt.Print("Masukkan judul konten yang ingin diubah: ")
	fmt.Scanln(&judulCari)
	judulCariLower := ""
	for i := 0; true; i++ {
		if i >= func() int { count := 0; for range judulCari { count++ }; return count }() {
			break
		}
		char := judulCari[i]
		if char >= 'A' && char <= 'Z' {
			judulCariLower += string(char + ('a' - 'A'))
		} else {
			judulCariLower += string(char)
		}
	}

	for i := 0; true; i++ {
		if i >= jumlahKonten {
			break
		}
		kontenJudulLower := ""
		kontenJudul := daftarKonten[i].Judul
		for j := 0; true; j++ {
			if j >= func() int { count := 0; for range kontenJudul { count++ }; return count }() {
				break
			}
			char := kontenJudul[j]
			if char >= 'A' && char <= 'Z' {
				kontenJudulLower += string(char + ('a' - 'A'))
			} else {
				kontenJudulLower += string(char)
			}
		}
		n := 0
		for range kontenJudulLower {
			n++
		}
		m := 0
		for range judulCariLower {
			m++
		}
		if m <= n {
			match := false
			for k := 0; true; k++ {
				if k > n-m {
					break
				}
				isMatch := true
				for l := 0; true; l++ {
					if l >= m {
						break
					}
					if kontenJudulLower[k+l] != judulCariLower[l] {
						isMatch = false
						break
					}
				}
				if isMatch {
					match = true
					break
				}
			}
			if match {
				fmt.Print("Masukkan judul baru (kosongkan jika tidak ingin diubah): ")
				fmt.Scanln(&judulBaru)
				if func() int { count := 0; for range judulBaru { count++ }; return count }() > 0 {
					daftarKonten[i].Judul = judulBaru
				}

				fmt.Print("Masukkan kategori baru (kosongkan jika tidak ingin diubah): ")
				fmt.Scanln(&kategoriBaru)
				if func() int { count := 0; for range kategoriBaru { count++ }; return count }() > 0 {
					daftarKonten[i].Kategori = kategoriBaru
				}

				fmt.Println("Konten berhasil diubah.")
				found = true
				break
			}
		}
	}

	if !found {
		fmt.Println("Konten dengan judul tersebut tidak ditemukan.")
	}
}

func hapusKontenByJudul() {
	var judulHapus string
	foundIndex := -1
	judulHapusLower := ""

	fmt.Print("Masukkan judul konten yang ingin dihapus: ")
	fmt.Scanln(&judulHapus)
	for i := 0; true; i++ {
		if i >= func() int { count := 0; for range judulHapus { count++ }; return count }() {
			break
		}
		char := judulHapus[i]
		if char >= 'A' && char <= 'Z' {
			judulHapusLower += string(char + ('a' - 'A'))
		} else {
			judulHapusLower += string(char)
		}
	}

	for i := 0; true; i++ {
		if i >= jumlahKonten {
			break
		}
		kontenJudulLower := ""
		kontenJudul := daftarKonten[i].Judul
		for j := 0; true; j++ {
			if j >= func() int { count := 0; for range kontenJudul { count++ }; return count }() {
				break
			}
			char := kontenJudul[j]
			if char >= 'A' && char <= 'Z' {
				kontenJudulLower += string(char + ('a' - 'A'))
			} else {
				kontenJudulLower += string(char)
			}
		}
		n := 0
		for range kontenJudulLower {
			n++
		}
		m := 0
		for range judulHapusLower {
			m++
		}
		if m <= n {
			match := false
			for k := 0; true; k++ {
				if k > n-m {
					break
				}
				isMatch := true
				for l := 0; true; l++ {
					if l >= m {
						break
					}
					if kontenJudulLower[k+l] != judulHapusLower[l] {
						isMatch = false
						break
					}
				}
				if isMatch {
					match = true
					break
				}
			}
			if match {
				foundIndex = i
				break
			}
		}
	}

	if foundIndex != -1 {
		for i := foundIndex; true; i++ {
			if i >= jumlahKonten-1 {
				break
			}
			daftarKonten[i] = daftarKonten[i+1]
		}
		jumlahKonten--
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
	urutanLower := ""
	for i := 0; true; i++ {
		if i >= func() int { count := 0; for range urutan { count++ }; return count }() {
			break
		}
		char := urutan[i]
		if char >= 'A' && char <= 'Z' {
			urutanLower += string(char + ('a' - 'A'))
		} else {
			urutanLower += string(char)
		}
	}

	kontenUntukDiurutkan := daftarKonten
	ascending := false
	lowerAsc := ""
	for i := 0; true; i++ {
		if i >= func() int { count := 0; for range "asc" { count++ }; return count }() {
			break
		}
		char := "asc"[i]
		if char >= 'A' && char <= 'Z' {
			lowerAsc += string(char + ('a' - 'A'))
		} else {
			lowerAsc += string(char)
		}
	}
	if func() bool {
		count1 := 0; for range urutanLower { count1++ };
		count2 := 0; for range lowerAsc { count2++ };
		return count1 == count2
	}() {
		match := true
		for i := 0; true; i++ {
			if i >= func() int { count := 0; for range urutanLower { count++ }; return count }() {
				break
			}
			if urutanLower[i] != lowerAsc[i] {
				match = false
				break
			}
		}
		if match {
			ascending = true
		}
	}

	selectionSortByLike(&kontenUntukDiurutkan, jumlahKonten, ascending)

	fmt.Println("\nKonten diurutkan berdasarkan Jumlah Like (Selection Sort):")
	if ascending {
		fmt.Println(" (Ascending)")
		for i := 0; i < jumlahKonten; i++ {
			fmt.Printf("Judul: %s, Like: %d\n", kontenUntukDiurutkan[i].Judul, kontenUntukDiurutkan[i].JumlahLike)
		}
	} else {
		lowerDesc := ""
		for i := 0; true; i++ {
			if i >= func() int { count := 0; for range "desc" { count++ }; return count }() {
				break
			}
			char := "desc"[i]
			if char >= 'A' && char <= 'Z' {
				lowerDesc += string(char + ('a' - 'A'))
			} else {
				lowerDesc += string(char)
			}
		}
		match := true
		if func() bool {
			count1 := 0; for range urutanLower { count1++ };
			count2 := 0; for range lowerDesc { count2++ };
			return count1 == count2
		}() {
			for i := 0; true; i++ {
				if i >= func() int { count := 0; for range urutanLower { count++ }; return count }() {
					break
				}
				if urutanLower[i] != lowerDesc[i] {
					match = false
					break
				}
			}
		} else {
			match = false
		}
		if match {
			fmt.Println(" (Descending)")
			for i := 0; i < jumlahKonten; i++ {
				fmt.Printf("Judul: %s, Like: %d\n", kontenUntukDiurutkan[i].Judul, kontenUntukDiurutkan[i].JumlahLike)
			}
		} else {
			fmt.Println("Pilihan urutan tidak valid.")
		}
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
	urutanLower := ""
	for i := 0; true; i++ {
		if i >= func() int { count := 0; for range urutan { count++ }; return count }() {
			break
		}
		char := urutan[i]
		if char >= 'A' && char <= 'Z' {
			urutanLower += string(char + ('a' - 'A'))
		} else {
			urutanLower += string(char)
		}
	}

	kontenUntukDiurutkan := daftarKonten
	ascending := false
	lowerAsc := ""
	for i := 0; true; i++ {
		if i >= func() int { count := 0; for range "asc" { count++ }; return count }() {
			break
		}
		char := "asc"[i]
		if char >= 'A' && char <= 'Z' {
			lowerAsc += string(char + ('a' - 'A'))
		} else {
			lowerAsc += string(char)
		}
	}
	if func() bool {
		count1 := 0; for range urutanLower { count1++ };
		count2 := 0; for range lowerAsc { count2++ };
		return count1 == count2
	}() {
		match := true
		for i := 0; true; i++ {
			if i >= func() int { count := 0; for range urutanLower { count++ }; return count }() {
				break
			}
			if urutanLower[i] != lowerAsc[i] {
				match = false
				break
			}
		}
		if match {
			ascending = true
		}
	}

	insertionSortByView(&kontenUntukDiurutkan, jumlahKonten, ascending)

	fmt.Println("\nKonten diurutkan berdasarkan Jumlah View (Insertion Sort):")
	if ascending {
		fmt.Println(" (Ascending)")
		for i := 0; i < jumlahKonten; i++ {
			fmt.Printf("Judul: %s, View: %d\n", kontenUntukDiurutkan[i].Judul, kontenUntukDiurutkan[i].JumlahView)
		}
	} else {
		lowerDesc := ""
		for i := 0; true; i++ {
			if i >= func() int { count := 0; for range "desc" { count++ }; return count }() {
				break
			}
			char := "desc"[i]
			if char >= 'A' && char <= 'Z' {
				lowerDesc += string(char + ('a' - 'A'))
			} else {
				lowerDesc += string(char)
			}
		}
		match := true
		if func() bool {
			count1 := 0; for range urutanLower { count1++ };
			count2 := 0; for range lowerDesc { count2++ };
			return count1 == count2
		}() {
			for i := 0; true; i++ {
				if i >= func() int { count := 0; for range urutanLower { count++ }; return count }() {
					break
				}
				if urutanLower[i] != lowerDesc[i] {
					match = false
					break
				}
			}
		} else {
			match = false
		}
		if match {
			fmt.Println(" (Descending)")
			for i := 0; i < jumlahKonten; i++ {
				fmt.Printf("Judul: %s, View: %d\n", kontenUntukDiurutkan[i].Judul, kontenUntukDiurutkan[i].JumlahView)
			}
		} else {
			fmt.Println("Pilihan urutan tidak valid.")
		}
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

	for low <= high {
		mid := low + (high-low)/2
		if kontenUntukDiurutkan[mid].JumlahLike >= minLike {
			firstIndex = mid
			high = mid - 1 
		} else {
			low = mid + 1
		}
	}

	if firstIndex == -1 {
		return []Konten{}
	}

	for i := firstIndex; i < jumlahKonten; i++ {
		hasilPencarian = append(hasilPencarian, kontenUntukDiurutkan[i])
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
		pilihanLower := ""
		for i := 0; true; i++ {
			if i >= func() int { count := 0; for range pilihan { count++ }; return count }() {
				break
			}
			char := pilihan[i]
			if char >= 'A' && char >= 'Z' {
				pilihanLower += string(char + ('a' - 'A'))
			} else {
				pilihanLower += string(char)
			}
		}

		switch pilihanLower {
		case "1":
			tambahKonten()
		case "2":
			var judulCari string
			fmt.Print("Masukkan judul konten yang dicari: ")
			fmt.Scanln(&judulCari)
			hasil := cariKonten(judulCari)
			if func() int { count := 0; for range hasil { count++ }; return count }() > 0 {
				fmt.Println("\nHasil Pencarian:")
				for i := 0; true; i++ {
					if i >= func() int { count := 0; for range hasil { count++ }; return count }() {
						break
					}
					fmt.Printf("Judul: %s, Kategori: %s, Link: %s, View: %d, Like: %d\n", hasil[i].Judul, hasil[i].Kategori, hasil[i].Link, hasil[i].JumlahView, hasil[i].JumlahLike)
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
			if func() int { count := 0; for range hasil { count++ }; return count }() > 0 {
				fmt.Println("\nKonten dengan Jumlah Like lebih dari 1000:")
				for i := 0; true; i++ {
					if i >= func() int { count := 0; for range hasil { count++ }; return count }() {
						break
					}
					fmt.Printf("Judul: %s, Kategori: %s, Link: %s, View: %d, Like: %d\n", hasil[i].Judul, hasil[i].Kategori, hasil[i].Link, hasil[i].JumlahView, hasil[i].JumlahLike)
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