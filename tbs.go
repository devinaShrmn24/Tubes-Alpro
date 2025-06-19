package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const MAX = 100

type Konten struct {
	Judul             string
	Platform          string
	Kategori          string
	TanggalJamPosting string
	JumlahLike        int
	JumlahKomen       int
	JumlahShare       int
	Status            string
	Deadline          string
}

var daftarKonten [MAX]Konten
var jumlahKonten int = 0

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readInt() int {
	var inputStr string
	var num int
	var err error
	for {
		inputStr = readLine()
		num, err = strconv.Atoi(inputStr)
		if err == nil {
			return num
		}
		fmt.Print("Input tidak valid, masukkan angka: ")
	}
}

func tampilkanMenu() {
	fmt.Println("\n=== Aplikasi Manajemen Konten ===")
	fmt.Println("1. Tambah Konten")
	fmt.Println("2. Tampilkan Semua Konten")
	fmt.Println("3. Ubah Status Konten")
	fmt.Println("4. Hapus Konten")
	fmt.Println("5. Urutkan Konten")
	fmt.Println("6. Cari Konten")
	fmt.Println("7. Tampilkan Konten Engagement Tertinggi")
	fmt.Println("8. Keluar")
}

func tambahKonten() {
	if jumlahKonten >= MAX {
		fmt.Println("Kapasitas penuh. Tidak bisa menambah konten lagi.")
		return
	}

	fmt.Print("Judul: ")
	judul := readLine()
	fmt.Print("Platform: ")
	platform := readLine()
	fmt.Print("Kategori: ")
	kategori := readLine()
	fmt.Print("Tanggal dan Jam Posting (YYYY-MM-DD HH:MM): ")
	tanggalJam := readLine()

	fmt.Print("Jumlah Like: ")
	like := readInt()
	fmt.Print("Jumlah Komentar: ")
	komen := readInt()
	fmt.Print("Jumlah Share: ")
	share := readInt()

	fmt.Print("Status (ide/draf/selesai/dipublikasikan): ")
	status := readLine()
	fmt.Print("Deadline (YYYY-MM-DD): ")
	deadline := readLine()

	daftarKonten[jumlahKonten] = Konten{
		Judul:             judul,
		Platform:          platform,
		Kategori:          kategori,
		TanggalJamPosting: tanggalJam,
		JumlahLike:        like,
		JumlahKomen:       komen,
		JumlahShare:       share,
		Status:            status,
		Deadline:          deadline,
	}
	jumlahKonten++
	fmt.Printf("Konten '%s' berhasil ditambahkan.\n", judul)
}

func tampilkanSemua() {
	if jumlahKonten == 0 {
		fmt.Println("Belum ada konten yang terdaftar.")
		return
	}
	fmt.Println("\nDaftar Semua Konten:")
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("%-3s | %-20s | %-10s | %-15s | %-19s | %-10s | %-6s | %-6s | %-6s | %-10s\n",
		"No.", "Judul", "Platform", "Kategori", "Tgl/Jam Posting", "Status", "Like", "Komen", "Share", "Deadline")
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------")
	for i := 0; i < jumlahKonten; i++ {
		k := daftarKonten[i]
		fmt.Printf("%-3d | %-20s | %-10s | %-15s | %-19s | %-10s | %-6d | %-6d | %-6d | %-10s\n",
			i+1, k.Judul, k.Platform, k.Kategori, k.TanggalJamPosting, k.Status, k.JumlahLike, k.JumlahKomen, k.JumlahShare, k.Deadline)
	}
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------")
	fmt.Println()
}

func ubahStatusKonten() {
	fmt.Print("Masukkan judul konten yang ingin diubah statusnya: ")
	judul := readLine()

	index := sequentialSearchJudul(judul)

	if index == -1 {
		fmt.Println("Konten tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan status baru (ide/draf/selesai/dipublikasikan): ")
	statusBaru := readLine()
	daftarKonten[index].Status = statusBaru
	fmt.Printf("Status konten '%s' berhasil diubah menjadi '%s'.\n", judul, statusBaru)
}

func hapusKonten() {
	fmt.Print("Masukkan judul konten yang ingin dihapus: ")
	judul := readLine()

	index := sequentialSearchJudul(judul)

	if index == -1 {
		fmt.Println("Konten tidak ditemukan.")
		return
	}

	for i := index; i < jumlahKonten-1; i++ {
		daftarKonten[i] = daftarKonten[i+1]
	}
	jumlahKonten--
	fmt.Printf("Konten '%s' berhasil dihapus.\n", judul)
}

func hitungEngagement(k Konten) int {
	return k.JumlahLike + k.JumlahKomen + k.JumlahShare
}

func selectionSortJudulAscending() {
	for i := 0; i < jumlahKonten-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahKonten; j++ {
			if daftarKonten[j].Judul < daftarKonten[minIdx].Judul {
				minIdx = j
			}
		}
		daftarKonten[i], daftarKonten[minIdx] = daftarKonten[minIdx], daftarKonten[i]
	}
}

func selectionSortJudulDescending() {
	for i := 0; i < jumlahKonten-1; i++ {
		maxIdx := i
		for j := i + 1; j < jumlahKonten; j++ {
			if daftarKonten[j].Judul > daftarKonten[maxIdx].Judul {
				maxIdx = j
			}
		}
		daftarKonten[i], daftarKonten[maxIdx] = daftarKonten[maxIdx], daftarKonten[i]
	}
}

func selectionSortEngagement(urutan string) {
	for i := 0; i < jumlahKonten-1; i++ {
		targetIdx := i
		for j := i + 1; j < jumlahKonten; j++ {
			engagementTarget := hitungEngagement(daftarKonten[targetIdx])
			engagementCurrent := hitungEngagement(daftarKonten[j])

			isSwapNeeded := false
			if urutan == "asc" {
				if engagementCurrent < engagementTarget {
					isSwapNeeded = true
				}
			} else if urutan == "desc" {
				if engagementCurrent > engagementTarget {
					isSwapNeeded = true
				}
			}

			if isSwapNeeded {
				targetIdx = j
			}
		}
		daftarKonten[i], daftarKonten[targetIdx] = daftarKonten[targetIdx], daftarKonten[i]
	}
}

func insertionSortTanggalJamPosting(urutan string) {
	const layout = "2006-01-02 15:04"
	for i := 1; i < jumlahKonten; i++ {
		key := daftarKonten[i]
		j := i - 1

		keyTime, _ := time.Parse(layout, key.TanggalJamPosting)

		continueLoop := true
		for j >= 0 && continueLoop {
			currentElementTime, _ := time.Parse(layout, daftarKonten[j].TanggalJamPosting)

			isCompareTrue := false
			if urutan == "asc" {
				isCompareTrue = currentElementTime.After(keyTime)
			} else if urutan == "desc" {
				isCompareTrue = currentElementTime.Before(keyTime)
			}

			if isCompareTrue {
				daftarKonten[j+1] = daftarKonten[j]
				j--
			} else {
				continueLoop = false
			}
		}
		daftarKonten[j+1] = key
	}
}

func menuUrutkanKonten() {
	if jumlahKonten < 2 {
		fmt.Println("Minimal 2 konten diperlukan untuk melakukan pengurutan.")
		return
	}

	var kriteria, urutan string
	fmt.Print("Urutkan berdasarkan (judul/tanggal/engagement): ")
	kriteria = readLine()
	fmt.Print("Urutan (asc/desc): ")
	urutan = readLine()

	if urutan != "asc" && urutan != "desc" {
		fmt.Println("Urutan tidak valid. Gunakan 'asc' atau 'desc'.")
		return
	}

	switch kriteria {
	case "judul":
		if urutan == "asc" {
			selectionSortJudulAscending()
			fmt.Println("Konten diurutkan berdasarkan Judul (Ascending) dengan Selection Sort.")
		} else {
			selectionSortJudulDescending()
			fmt.Println("Konten diurutkan berdasarkan Judul (Descending) dengan Selection Sort.")
		}
	case "tanggal":
		insertionSortTanggalJamPosting(urutan)
		fmt.Printf("Konten diurutkan berdasarkan Tanggal/Jam Posting (%s) dengan Insertion Sort.\n", urutan)
	case "engagement":
		selectionSortEngagement(urutan)
		fmt.Printf("Konten diurutkan berdasarkan Engagement (%s) dengan Selection Sort.\n", urutan)
	default:
		fmt.Println("Kriteria pengurutan tidak valid. Pilih 'judul', 'tanggal', atau 'engagement'.")
	}
}

func sequentialSearchJudul(judul string) int {
	foundIndex := -1
	i := 0
	for i < jumlahKonten {
		if daftarKonten[i].Judul == judul {
			foundIndex = i
			i = jumlahKonten
		}
		i++
	}
	return foundIndex
}

func sequentialSearchByKategori(kategori string) []int {
	var foundIndices []int
	for i := 0; i < jumlahKonten; i++ {
		if daftarKonten[i].Kategori == kategori {
			foundIndices = append(foundIndices, i)
		}
	}
	return foundIndices
}

func binarySearchJudul(judul string) int {
	left := 0
	right := jumlahKonten - 1
	resultIndex := -1

	for left <= right {
		mid := (left + right) / 2
		if daftarKonten[mid].Judul == judul {
			resultIndex = mid
			left = right + 1
		} else if daftarKonten[mid].Judul < judul {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return resultIndex
}

func menuCariKonten() {
	var kriteria string
	fmt.Print("Cari berdasarkan (judul/kategori): ")
	kriteria = readLine()
	fmt.Print("Masukkan nilai yang dicari: ")
	nilai := readLine()

	var found bool = false
	var tempIndices []int

	if kriteria == "judul" {
		selectionSortJudulAscending()
		idx := binarySearchJudul(nilai)
		if idx != -1 {
			k := daftarKonten[idx]
			fmt.Println("\nKonten ditemukan:")
			fmt.Printf("Judul: %s | Platform: %s | Kategori: %s | Tgl/Jam Posting: %s | Status: %s | Deadline: %s | Likes: %d | Komen: %d | Shares: %d\n",
				k.Judul, k.Platform, k.Kategori, k.TanggalJamPosting, k.Status, k.Deadline, k.JumlahLike, k.JumlahKomen, k.JumlahShare)
			found = true
		}
	} else if kriteria == "kategori" {
		tempIndices = sequentialSearchByKategori(nilai)
		if len(tempIndices) > 0 {
			fmt.Println("\nKonten ditemukan dengan kategori tersebut:")
			for _, idx := range tempIndices {
				k := daftarKonten[idx]
				fmt.Printf("Judul: %s | Platform: %s | Kategori: %s | Tgl/Jam Posting: %s | Status: %s | Deadline: %s | Likes: %d | Komen: %d | Shares: %d\n",
					k.Judul, k.Platform, k.Kategori, k.TanggalJamPosting, k.Status, k.Deadline, k.JumlahLike, k.JumlahKomen, k.JumlahShare)
			}
			found = true
		}
	} else {
		fmt.Println("Kriteria pencarian tidak valid. Pilih 'judul' atau 'kategori'.")
		return
	}

	if !found {
		fmt.Println("Konten tidak ditemukan.")
	}
}

func tampilkanKontenEngagementTertinggi() {
	if jumlahKonten == 0 {
		fmt.Println("Belum ada konten untuk dianalisis engagement-nya.")
		return
	}

	fmt.Print("Masukkan tanggal awal (YYYY-MM-DD): ")
	startDateStr := readLine()
	fmt.Print("Masukkan tanggal akhir (YYYY-MM-DD): ")
	endDateStr := readLine()

	const dateFormat = "2006-01-02"
	parsedStartDate, err1 := time.Parse(dateFormat, startDateStr)
	parsedEndDate, err2 := time.Parse(dateFormat, endDateStr)

	if err1 != nil || err2 != nil {
		fmt.Println("Format tanggal tidak valid. Gunakan YYYY-MM-DD.")
		return
	}

	var highestEngagementKonten Konten
	maxEngagement := -1
	foundAny := false

	for i := 0; i < jumlahKonten; i++ {
		k := daftarKonten[i]
		tanggalPostingPart := strings.Split(k.TanggalJamPosting, " ")[0]

		parsedPostingDate, err3 := time.Parse(dateFormat, tanggalPostingPart)

		if err3 == nil {
			isInRange := (parsedPostingDate.Equal(parsedStartDate) || parsedPostingDate.After(parsedStartDate)) &&
				(parsedPostingDate.Equal(parsedEndDate) || parsedPostingDate.Before(parsedEndDate))

			if isInRange {
				currentEngagement := hitungEngagement(k)
				if currentEngagement > maxEngagement {
					maxEngagement = currentEngagement
					highestEngagementKonten = k
					foundAny = true
				}
			}
		}
	}

	if foundAny {
		fmt.Println("\nKonten dengan Engagement Tertinggi dalam rentang tersebut:")
		fmt.Printf("Judul: %s | Platform: %s | Kategori: %s | Tgl/Jam Posting: %s | Status: %s | Deadline: %s | Likes: %d | Komen: %d | Shares: %d | Total Engagement: %d\n",
			highestEngagementKonten.Judul, highestEngagementKonten.Platform, highestEngagementKonten.Kategori,
			highestEngagementKonten.TanggalJamPosting, highestEngagementKonten.Status, highestEngagementKonten.Deadline,
			highestEngagementKonten.JumlahLike, highestEngagementKonten.JumlahKomen, highestEngagementKonten.JumlahShare, maxEngagement)
	} else {
		fmt.Println("Tidak ada konten yang ditemukan dalam rentang tanggal tersebut, atau semua konten memiliki engagement 0.")
	}
}

func main() {
	var pilihan int

	for {
		tampilkanMenu()
		fmt.Print("Pilih menu: ")
		_, err := fmt.Scanln(&pilihan)
		if err != nil {
			fmt.Println("Input tidak valid. Harap masukkan angka menu.")
			continue
		}

		switch pilihan {
		case 1:
			tambahKonten()
		case 2:
			tampilkanSemua()
		case 3:
			ubahStatusKonten()
		case 4:
			hapusKonten()
		case 5:
			menuUrutkanKonten()
		case 6:
			menuCariKonten()
		case 7:
			tampilkanKontenEngagementTertinggi()
		case 8:
			fmt.Println("Terima kasih telah menggunakan Aplikasi Manajemen Konten!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
