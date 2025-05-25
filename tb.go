package main

import (
	"fmt"
	"strings"
)

type Tanggal struct {
	Hari  int
	Bulan int
	Tahun int
	Jam   int
	Menit int
}

type Konten struct {
	Judul       string
	Platform    string
	Kategori    string
	TanggalPost Tanggal
	Like        int
	Komentar    int
	Share       int
}

const MAX = 100

var daftarKonten [MAX]Konten
var jumlahKonten int = 0

func TambahKonten(k Konten) {
	if jumlahKonten < MAX {
		daftarKonten[jumlahKonten] = k
		jumlahKonten++
		fmt.Println("Konten berhasil ditambahkan!")
	} else {
		fmt.Println("Data penuh, tidak bisa menambahkan konten lagi.")
	}
}

func UbahKonten(judul string, baru Konten) bool {
	for i := 0; i < jumlahKonten; i++ {
		if strings.EqualFold(daftarKonten[i].Judul, judul) {
			daftarKonten[i] = baru
			return true
		}
	}
	return false
}

func HapusKonten(judul string) bool {
	index := -1
	for i := 0; i < jumlahKonten; i++ {
		if strings.EqualFold(daftarKonten[i].Judul, judul) {
			index = i
			break
		}
	}
	if index == -1 {
		return false
	}

	for i := index; i < jumlahKonten-1; i++ {
		daftarKonten[i] = daftarKonten[i+1]
	}
	jumlahKonten--
	return true
}

func CariKategoriSequential(kategori string) []Konten {
	var hasil []Konten
	for i := 0; i < jumlahKonten; i++ {
		if strings.EqualFold(daftarKonten[i].Kategori, kategori) {
			hasil = append(hasil, daftarKonten[i])
		}
	}
	return hasil
}

func CariJudulBinary(target string) int {
	kiri := 0
	kanan := jumlahKonten - 1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if daftarKonten[tengah].Judul == target {
			return tengah
		} else if daftarKonten[tengah].Judul < target {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func SelectionSortEngagement(ascending bool) {
	for i := 0; i < jumlahKonten-1; i++ {
		idx := i
		for j := i + 1; j < jumlahKonten; j++ {
			e1 := daftarKonten[idx].Like + daftarKonten[idx].Komentar + daftarKonten[idx].Share
			e2 := daftarKonten[j].Like + daftarKonten[j].Komentar + daftarKonten[j].Share
			if (ascending && e2 < e1) || (!ascending && e2 > e1) {
				idx = j
			}
		}
		daftarKonten[i], daftarKonten[idx] = daftarKonten[idx], daftarKonten[i]
	}
}

func tanggalToInt(t Tanggal) int {
	return t.Tahun*10000 + t.Bulan*100 + t.Hari
}

func InsertionSortTanggal(ascending bool) {
	for i := 1; i < jumlahKonten; i++ {
		temp := daftarKonten[i]
		j := i - 1
		for j >= 0 {
			t1 := tanggalToInt(daftarKonten[j].TanggalPost)
			t2 := tanggalToInt(temp.TanggalPost)

			if (ascending && t1 > t2) || (!ascending && t1 < t2) {
				daftarKonten[j+1] = daftarKonten[j]
				j--
			} else {
				break
			}
		}
		daftarKonten[j+1] = temp
	}
}

func InRange(t, dari, sampai Tanggal) bool {
	return tanggalToInt(t) >= tanggalToInt(dari) && tanggalToInt(t) <= tanggalToInt(sampai)
}

func KontenEngagementTertinggi(dari, sampai Tanggal) Konten {
	maxIndex := -1
	maxEngagement := -1
	for i := 0; i < jumlahKonten; i++ {
		t := daftarKonten[i].TanggalPost
		if InRange(t, dari, sampai) {
			total := daftarKonten[i].Like + daftarKonten[i].Komentar + daftarKonten[i].Share
			if total > maxEngagement {
				maxEngagement = total
				maxIndex = i
			}
		}
	}
	if maxIndex != -1 {
		return daftarKonten[maxIndex]
	}
	return Konten{}
}

func TampilKonten() {
	fmt.Println("Daftar Konten:")
	for i := 0; i < jumlahKonten; i++ {
		k := daftarKonten[i]
		fmt.Printf("%d. %s [%s] di %s - %d/%d/%d %02d:%02d | Like: %d, Komentar: %d, Share: %d",
			i+1, k.Judul, k.Kategori, k.Platform,
			k.TanggalPost.Hari, k.TanggalPost.Bulan, k.TanggalPost.Tahun, k.TanggalPost.Jam, k.TanggalPost.Menit,
			k.Like, k.Komentar, k.Share)
	}
}

func main() {
	for {
		fmt.Println("\n===== MENU APLIKASI KONTEN =====")
		fmt.Println("1. Tambah Konten")
		fmt.Println("2. Tampil Semua Konten")
		fmt.Println("3. Urutkan Berdasarkan Tanggal (Insertion Sort)")
		fmt.Println("4. Urutkan Berdasarkan Engagement (Selection Sort)")
		fmt.Println("5. Keluar")

		var pilihan int
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			var k Konten
			fmt.Print("Judul: ")
			fmt.Scanln(&k.Judul)
			fmt.Print("Platform: ")
			fmt.Scanln(&k.Platform)
			fmt.Print("Kategori: ")
			fmt.Scanln(&k.Kategori)
			fmt.Print("Tanggal (dd mm yyyy hh mm): ")
			fmt.Scan(&k.TanggalPost.Hari, &k.TanggalPost.Bulan, &k.TanggalPost.Tahun, &k.TanggalPost.Jam, &k.TanggalPost.Menit)
			fmt.Print("Like: ")
			fmt.Scanln(&k.Like)
			fmt.Print("Komentar: ")
			fmt.Scanln(&k.Komentar)
			fmt.Print("Share: ")
			fmt.Scanln(&k.Share)
			TambahKonten(k)
		} else if pilihan == 2 {
			TampilKonten()
		} else if pilihan == 3 {
			var asc string
			fmt.Print("Ascending (y/n)? ")
			fmt.Scanln(&asc)
			InsertionSortTanggal(strings.ToLower(asc) == "y")
			fmt.Println("Konten berhasil diurutkan berdasarkan tanggal.")
		} else if pilihan == 4 {
			var asc string
			fmt.Print("Ascending (y/n)? ")
			fmt.Scanln(&asc)
			SelectionSortEngagement(strings.ToLower(asc) == "y")
			fmt.Println("Konten berhasil diurutkan berdasarkan engagement.")
		} else if pilihan == 5 {
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			break
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
