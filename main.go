package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mahasiswa struct {
	NIM   string
	Name  string
	Nilai float64
}

func main() {
	var mahasiswaList []Mahasiswa
	reader := bufio.NewReader(os.Stdin)

	// INPUT: mengambil jumlah mahasiswa
	fmt.Print("Masukkan jumlah mahasiswa: ")
	inputJumlah, _ := reader.ReadString('\n')
	inputJumlah = strings.TrimSpace(inputJumlah)
	jumlah, _ := strconv.Atoi(inputJumlah)

	// Inisialisasi wadah
	var totalScore float64
	var scoreTertinggi float64
	var namaTertinggi string

	for i := 0; i < jumlah; i++ {
		fmt.Printf("\nData mahasiswa ke-%d\n", i+1)

		// INPUT: NIM
		fmt.Print("NIM: ")
		nim, _ := reader.ReadString('\n')
		nim = strings.TrimSpace(nim)

		// INPUT: Nama
		fmt.Print("Name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		// INPUT: Nilai
		fmt.Print("Score: ")
		scoreInput, _ := reader.ReadString('\n')
		scoreInput = strings.TrimSpace(scoreInput)
		score, _ := strconv.ParseFloat(scoreInput, 64)

		// Buat struct Mahasiswa
		mahasiswa := Mahasiswa{NIM: nim, Name: name, Nilai: score}
		mahasiswaList = append(mahasiswaList, mahasiswa)

		// Hitung total nilai
		totalScore += score

		// Cek nilai tertinggi
		if score > scoreTertinggi {
			scoreTertinggi = score
			namaTertinggi = name
		}
	}

	// HITUNG & TAMPILKAN HASIL
	rataRata := totalScore / float64(jumlah)
	fmt.Printf("\nRata-rata nilai: %.2f\n", rataRata)
	fmt.Printf("Nilai tertinggi: %.2f oleh %s\n", scoreTertinggi, namaTertinggi)
}
