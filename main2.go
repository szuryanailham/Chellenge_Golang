package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// INPUT {Memasukan jumlah uang}
	fmt.Print("Masukan jumlah uang dalam Rupiah :")
	totalInput,_ := reader.ReadString('\n')
	totalInput = strings.TrimSpace(totalInput)
	total,_ := strconv.ParseInt(totalInput,10 , 64)

	fmt.Println("======================================")
	fmt.Println("1. USD")
	fmt.Println("2. EUR")
	fmt.Println("3. JPY")
	fmt.Print("Masukan Menu Konversi Mata Uang Anda: ")

	inputMenu,_ := reader.ReadString('\n')
	inputMenu = strings.TrimSpace(inputMenu)
	menu, err := strconv.ParseInt(inputMenu, 10, 64)
	if err != nil {
		fmt.Println("Input tidak valid. Harus berupa angka.")
		return
	}

	var hasil float64
	var mataUang string
	// SWITCH : logika  switch 
	switch menu {
	case 1:
	hasil = float64(total) / float64(15500)
	mataUang = "USD"
	case 2:
		hasil = float64(total) / float64(17000)
	mataUang = "EUR"

	case 3:
	hasil = float64(total) / float64(110)
	mataUang = "JPY"

	default:
	fmt.Println("Menu tidak tersedia.")
	return
	}
	fmt.Printf("Rp%d = %.2f %s\n", total, hasil, mataUang)
}