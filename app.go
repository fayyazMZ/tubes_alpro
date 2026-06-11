package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const NMAX int = 100

type subscription struct {
	id, no   int
	name     string
	category string
	cost     int
	dueDate  dueDate
	method   string
	status   string
}

type dueDate struct {
	day, month, year int
	converted        int
}

type subData [NMAX]subscription

func cls() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func header() {
	fmt.Println()
	fmt.Println("                ╔══════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("                ║                        ▶▶  S u b s c r i b e W i s e  ◀◀                  ║")
	fmt.Println("                ║                      💡 Kelola langganan, hemat pengeluaran               ║")
	fmt.Println("                ╚══════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func mainMenu(choice *int, counter *int) {
	cls()
	fmt.Println()
	fmt.Println("                ┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("                │         ▶▶ Selamat Datang di SubscribeWise! ◀◀              │")
	fmt.Println("                │              💡 Kelola langganan, hemat pengeluaran           │")
	fmt.Println("                ├─────────────────────────────────────────────────────────────┤")
	fmt.Println("                │                  °❀⋆.ೃ࿔*:･   MAIN MENU   °❀⋆.ೃ࿔*:･          │")
	fmt.Println("                └─────────────────────────────────────────────────────────────┘")
	if (*choice < 1 || *choice > 2) && *counter != 0 {
		fmt.Println("                Pilihan tidak valid.")
	}
	fmt.Println()
	fmt.Println("                [1] 🔧 Kelola Langganan")
	fmt.Println("                [2] ↩️  Keluar")
	fmt.Println()
	fmt.Print("                ╰┈➤ ")
	fmt.Scan(choice)
	*counter++
}

func displaySubs(A subData, n int, status *string) {
	cls()
	header()
	if n != 0 {
		fmt.Println("┌────┬────────────────────────────┬──────────────┬────────────┬──────────────┬─────────────────────┬──────────────┐")
		fmt.Println("│ No │ Nama Layanan               │ Kategori     │ Biaya      │ Metode       │ Jatuh Tempo         │ Status       │")
		fmt.Println("├────┼────────────────────────────┼──────────────┼────────────┼──────────────┼─────────────────────┼──────────────┤")
		for i := 0; i < n; i++ {
			A[i].no = i + 1
			fmt.Printf("│ %-2d │ %-26s │ %-12s │ Rp%-8d│ %-12s │ %02d/%02d/%04d           │ %-12s │\n",
				A[i].no, A[i].name, A[i].category, A[i].cost, A[i].method,
				A[i].dueDate.day, A[i].dueDate.month, A[i].dueDate.year, A[i].status)
		}
		fmt.Println("└────┴────────────────────────────┴──────────────┴────────────┴──────────────┴─────────────────────┴──────────────┘")
	} else {
		fmt.Println("┌──────────────────────────────────────────────────────────────────────────────────────────────────────────┐")
		fmt.Println("│                               ⌕  Belum ada langganan                                                     │")
		fmt.Println("└──────────────────────────────────────────────────────────────────────────────────────────────────────────┘")
	}
	fmt.Println()
	if *status != "" {
		fmt.Printf("                %v\n", *status)
		*status = ""
	}
}

func subMenu(choice *int, A subData, n *int, status *string, budget *int) {
	displaySubs(A, *n, status)
	fmt.Println("                [1] ➕ Tambah Langganan")
	fmt.Println("                [2] 📝 Edit Langganan")
	fmt.Println("                [3] 🗑️  Hapus Langganan")
	fmt.Println("                [4] 🔎 Cari Langganan")
	fmt.Println("                [5] 📶 Urutkan Langganan")
	fmt.Println("                [6] ⏰ Pengingat Jatuh Tempo")
	fmt.Println("                [7] 💰 Total Pengeluaran & Rekomendasi")
	fmt.Println("                [8] ⬅️ Kembali ke Menu Utama")
	fmt.Println()
	fmt.Print("                ╰┈➤ ")
	fmt.Scan(choice)

	if *n == 0 && (*choice >= 2 && *choice <= 7) {
		*status = "Belum ada data untuk melakukan perintah."
		*choice = 0
	} else if *choice < 1 || *choice > 8 {
		*status = "Pilihan tidak valid."
		*choice = 0
	}
}

func dateToDays(d dueDate) int {
	return d.year*365 + d.month*30 + d.day
}

func addSubscription(A *subData, n *int, status *string) {
	i := *n
	var name string
	var dummy string
	displaySubs(*A, *n, status)
	fmt.Println("                Format: Nama Kategori Biaya Metode Hari/Bulan/Tahun")
	fmt.Println("                Contoh: Netflix Hiburan 89000 KartuKredit 15/06/2026")
	fmt.Println("                Ketik '-' untuk berhenti.")
	fmt.Println()

	for {
		if i >= NMAX {
			*status = "Kapasitas penuh!"
			break
		}
		fmt.Print("                Masukkan data: ")
		fmt.Scan(&name)
		if name == "-" {
			break
		}
		A[i].name = name
		fmt.Scanf("%s %d %s %d/%d/%d", &A[i].category, &A[i].cost, &A[i].method,
			&A[i].dueDate.day, &A[i].dueDate.month, &A[i].dueDate.year)
		A[i].status = "Active"
		A[i].dueDate.converted = dateToDays(A[i].dueDate)
		A[i].id = i + 1
		A[i].no = A[i].id
		i++
		*n = i
		*status = "Langganan berhasil ditambahkan!"
	}
	// Bersihkan buffer enter sisa input data terakhir
	fmt.Scanln(&dummy)
}

func findIndexByNo(A subData, n int, target int) int {
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if target < A[mid].no {
			right = mid - 1
		} else if target > A[mid].no {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func editSubscription(A *subData, n int, status *string) {
	var no, idx int
	var field string
	displaySubs(*A, n, status)
	fmt.Println("                Format: No_Langganan Kolom")
	fmt.Println("                Kolom: name, category, cost, method, duedate, status")
	fmt.Print("                ╰┈➤ ")
	fmt.Scan(&no, &field)
	field = strings.ToLower(field)
	idx = findIndexByNo(*A, n, no)
	if idx != -1 {
		fmt.Printf("                Nilai baru untuk %s: ", field)
		switch field {
		case "name":
			fmt.Scan(&A[idx].name)
		case "category":
			fmt.Scan(&A[idx].category)
		case "cost":
			fmt.Scan(&A[idx].cost)
		case "method":
			fmt.Scan(&A[idx].method)
		case "duedate":
			fmt.Scanf("%d/%d/%d", &A[idx].dueDate.day, &A[idx].dueDate.month, &A[idx].dueDate.year)
			A[idx].dueDate.converted = dateToDays(A[idx].dueDate)
		case "status":
			fmt.Scan(&A[idx].status)
		default:
			*status = "Kolom tidak valid."
			return
		}
		*status = "Data berhasil diubah."
	} else {
		*status = "Data tidak ditemukan."
	}
}

func deleteSubscription(A *subData, n *int, status *string) {
	var no, idx int
	displaySubs(*A, *n, status)
	fmt.Print("                Nomor langganan (0 = batal): ")
	fmt.Scan(&no)
	if no == 0 {
		return
	}
	idx = findIndexByNo(*A, *n, no)
	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			A[i] = A[i+1]
			A[i].no--
			A[i].id--
		}
		*n--
		*status = "Langganan berhasil dihapus."
	} else {
		*status = "Data tidak ditemukan."
	}
}

func sequentialSearchByName(A subData, n int, keyword string) subData {
	var result subData
	count := 0
	keyword = strings.ToLower(keyword)
	for i := 0; i < n; i++ {
		if strings.Contains(strings.ToLower(A[i].name), keyword) {
			result[count] = A[i]
			result[count].no = count + 1
			count++
		}
	}
	return result
}

func insertionSortByName(A *subData, n int) {
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && strings.ToLower(A[j].name) > strings.ToLower(key.name) {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
	}
}

func binarySearchByName(A subData, n int, target string) int {
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if strings.ToLower(target) < strings.ToLower(A[mid].name) {
			right = mid - 1
		} else if strings.ToLower(target) > strings.ToLower(A[mid].name) {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func searchMenu(A *subData, n *int, status *string) {
	var method, keyword string
	var wait string
	displaySubs(*A, *n, status)
	fmt.Println("                [1] Sequential Search (nama mengandung kata kunci)")
	fmt.Println("                [2] Binary Search (nama persis)")
	fmt.Println("                [3] Kembali")
	fmt.Print("                ╰┈➤ ")
	fmt.Scan(&method)

	if method == "1" {
		fmt.Print("                Kata kunci: ")
		fmt.Scan(&keyword)
		result := sequentialSearchByName(*A, *n, keyword)
		count := 0
		for count < NMAX && result[count].name != "" {
			count++
		}
		if count > 0 {
			displaySubs(result, count, status)
			fmt.Println()
			fmt.Print("                Ketik angka/huruf bebas + Enter untuk kembali ke menu: ")
			fmt.Scan(&wait)
		} else {
			*status = "Tidak ditemukan."
		}
	} else if method == "2" {
		var sorted subData
		copy(sorted[:], (*A)[:*n])
		insertionSortByName(&sorted, *n)
		fmt.Print("                Nama persis: ")
		fmt.Scan(&keyword)
		idx := binarySearchByName(sorted, *n, keyword)
		if idx != -1 {
			cls()
			header()
			fmt.Println("┌──────────────────────────────────────────────────────────────────────────────────────────────────────────┐")
			fmt.Printf("│               ✅ Ditemukan: %-26s (Rp%-8d/bulan)                                      │\n", sorted[idx].name, sorted[idx].cost)
			fmt.Println("└──────────────────────────────────────────────────────────────────────────────────────────────────────────┘")
			fmt.Println()
			fmt.Print("                Ketik angka/huruf bebas + Enter untuk kembali ke menu: ")
			fmt.Scan(&wait)
		} else {
			*status = "Tidak ditemukan."
		}
	}
}

func selectionSortByCost(A *subData, n int, ascending bool) {
	for i := 0; i < n-1; i++ {
		extremeIdx := i
		for j := i + 1; j < n; j++ {
			if ascending {
				if A[j].cost < A[extremeIdx].cost {
					extremeIdx = j
				}
			} else {
				if A[j].cost > A[extremeIdx].cost {
					extremeIdx = j
				}
			}
		}
		A[i], A[extremeIdx] = A[extremeIdx], A[i]
	}
}

func insertionSortByDueDate(A *subData, n int, ascending bool) {
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		if ascending {
			for j >= 0 && A[j].dueDate.converted > key.dueDate.converted {
				A[j+1] = A[j]
				j--
			}
		} else {
			for j >= 0 && A[j].dueDate.converted < key.dueDate.converted {
				A[j+1] = A[j]
				j--
			}
		}
		A[j+1] = key
	}
}

func sortMenu(A *subData, n *int, status *string) {
	var sortBy, order int
	displaySubs(*A, *n, status)
	fmt.Println("                Urutkan berdasarkan:")
	fmt.Println("                [1] Biaya (Selection Sort)")
	fmt.Println("                [2] Tanggal Jatuh Tempo (Insertion Sort)")
	fmt.Print("                ╰┈➤ ")
	fmt.Scan(&sortBy)
	fmt.Print("                [1] Naik / [2] Turun: ")
	fmt.Scan(&order)
	asc := order == 1

	if sortBy == 1 {
		selectionSortByCost(A, *n, asc)
		*status = "Diurutkan berdasarkan biaya (Selection Sort)."
	} else if sortBy == 2 {
		insertionSortByDueDate(A, *n, asc)
		*status = "Diurutkan berdasarkan jatuh tempo (Insertion Sort)."
	}
}

func remindDue(A subData, n int) {
	now := time.Now()
	todayDays := now.Year()*365 + int(now.Month())*30 + now.Day()
	var wait string
	cls()
	fmt.Println()
	fmt.Println("                ╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("                ║                ⏰ PENGINGAT JATUH TEMPO (7 HARI)              ║")
	fmt.Println("                ╚══════════════════════════════════════════════════════════════╝")
	fmt.Println()
	found := false
	for i := 0; i < n; i++ {
		if A[i].status != "Active" {
			continue
		}
		diff := A[i].dueDate.converted - todayDays
		if diff >= 0 && diff <= 7 {
			fmt.Printf("                ⚠️  %s jatuh tempo %02d/%02d/%04d (dalam %d hari)\n",
				A[i].name, A[i].dueDate.day, A[i].dueDate.month, A[i].dueDate.year, diff)
			found = true
		}
	}
	if !found {
		fmt.Println("                ✅ Tidak ada yang jatuh tempo dalam 7 hari.")
	}
	fmt.Println()
	fmt.Print("                Ketik angka/huruf bebas + Enter untuk kembali ke menu: ")
	fmt.Scan(&wait)
}

func totalAndRecommend(A subData, n int, budget *int) {
	var total int
	var wait string
	for i := 0; i < n; i++ {
		if A[i].status == "Active" {
			total += A[i].cost
		}
	}
	cls()
	fmt.Println()
	fmt.Println("                ╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("                ║                💰 TOTAL PENGELUARAN & REKOMENDASI             ║")
	fmt.Println("                ╚══════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("                Total pengeluaran aktif: Rp%d/bulan\n", total)

	if *budget == 0 {
		fmt.Print("                Masukkan anggaran bulanan Anda: Rp")
		fmt.Scan(budget)
	}

	if *budget > 0 {
		fmt.Printf("                Anggaran Anda: Rp%d/bulan\n", *budget)
		if total <= *budget {
			fmt.Println("                ✅ Pengeluaran masih dalam batas anggaran.")
		} else {
			fmt.Println("                ⚠️  Melebihi anggaran! Pertimbangkan hentikan:")
			var actives subData
			idx := 0
			for i := 0; i < n; i++ {
				if A[i].status == "Active" {
					actives[idx] = A[i]
					idx++
				}
			}
			selectionSortByCost(&actives, idx, false)
			sisa := total
			for i := 0; i < idx; i++ {
				if sisa <= *budget {
					break
				}
				fmt.Printf("                   - %s (Rp%d/bulan)\n", actives[i].name, actives[i].cost)
				sisa -= actives[i].cost
			}
			fmt.Printf("                Jika dihemat: ~Rp%d/bulan\n", sisa)
		}
	}
	fmt.Println()
	fmt.Print("                Ketik angka/huruf bebas + Enter untuk kembali ke menu: ")
	fmt.Scan(&wait)
}

func initDummy(A *subData, n *int) {
	i := 0
	A[i] = subscription{id: 1, no: 1, name: "Netflix", category: "Hiburan", cost: 89000, method: "Kartu Kredit", status: "Active",
		dueDate: dueDate{day: 10, month: 6, year: 2026, converted: dateToDays(dueDate{day: 10, month: 6, year: 2026})}}
	i++
	A[i] = subscription{id: 2, no: 2, name: "Spotify", category: "Musik", cost: 54000, method: "GoPay", status: "Active",
		dueDate: dueDate{day: 5, month: 6, year: 2026, converted: dateToDays(dueDate{day: 5, month: 6, year: 2026})}}
	i++
	A[i] = subscription{id: 3, no: 3, name: "Disney+", category: "Hiburan", cost: 45000, method: "Kartu Debit", status: "Active",
		dueDate: dueDate{day: 16, month: 6, year: 2026, converted: dateToDays(dueDate{day: 16, month: 6, year: 2026})}}
	i++
	A[i] = subscription{id: 4, no: 4, name: "YouTube Premium", category: "Hiburan", cost: 69000, method: "Dana", status: "Active",
		dueDate: dueDate{day: 7, month: 6, year: 2026, converted: dateToDays(dueDate{day: 7, month: 6, year: 2026})}}
	i++
	A[i] = subscription{id: 5, no: 5, name: "AWS Cloud", category: "Produktivitas", cost: 150000, method: "Transfer", status: "Active",
		dueDate: dueDate{day: 1, month: 6, year: 2026, converted: dateToDays(dueDate{day: 1, month: 6, year: 2026})}}
	*n = i + 1
}

func main() {
	var data subData
	var n int = 0
	var choice int = 1
	var status string
	var counter int = 0
	var budget int = 0

	initDummy(&data, &n)

	for choice != 2 {
		mainMenu(&choice, &counter)
		if choice == 1 {
			counter = 0
			for choice != 8 {
				subMenu(&choice, data, &n, &status, &budget)
				switch choice {
				case 1:
					addSubscription(&data, &n, &status)
				case 2:
					editSubscription(&data, n, &status)
				case 3:
					deleteSubscription(&data, &n, &status)
				case 4:
					searchMenu(&data, &n, &status)
				case 5:
					sortMenu(&data, &n, &status)
				case 6:
					remindDue(data, n)
				case 7:
					totalAndRecommend(data, n, &budget)
				}
			}
		}
	}
}
