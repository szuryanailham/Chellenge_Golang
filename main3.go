package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type  Book struct {
	id int
	title string
}

func removeById(books []Book, id int)[]Book{
results := []Book{}
for _, book := range books {
	if book.id != id {
		results = append(results,book)
	}
}
return results
}

func main() {
	books := []Book{
		{id :1 , title : " Hallo Bgas"},
		{id :2 , title : " Hallo world"},
		{id :3 , title : " Hallo kucay"},
	}
	reader := bufio.NewReader(os.Stdin)
	for {
	fmt.Println("====================================")
	fmt.Println("1 .Menambah Buku")
	fmt.Println("2.Melihat Daftar Buku")
	fmt.Println("3.Mencari Buku berdasarkan ID")
	fmt.Println("4.Menghapus Buku")
	fmt.Println("5.Selesai")
	fmt.Println("====================================")
		fmt.Println("Masukan Menu yang kamu inginkan : ")
	
	inputMenu ,_ := reader.ReadString('\n')
	inputMenu = strings.TrimSpace(inputMenu);
	menu, err := strconv.ParseInt(inputMenu,10 , 64)

	if err != nil {
		fmt.Println("Input tidak valid. Harus berupa angka.")
		return
	}

	switch (menu) {
	case 1 :
	fmt.Print("masukan jumlah book baru :")
	inputJumlah ,_ := reader.ReadString('\n')
	inputJumlah = strings.TrimSpace(inputJumlah);
	jumlah, _ := strconv.ParseInt(inputJumlah,10 , 64)
	for i := 0; i <int(jumlah); i++ {
		// INPUT :: input data book id
	fmt.Print("Masukkan ID: ")
	inputID, _ := reader.ReadString('\n')
	inputID = strings.TrimSpace(inputID)
	idBook, err := strconv.Atoi(inputID)

if err != nil {
	fmt.Println("ID harus berupa angka!")
	return
}

	fmt.Println(idBook);
		// INPUT :: input data book name
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
		
		// input
	newbook := Book{id:idBook , title: title}
	books = append(books, newbook);

	
	}
		fmt.Println("selamat data berhasil dimasukan ")
	case 2 :
		fmt.Println("===========LIST BOOK =====================")
		for _, book := range books {
		fmt.Printf("id : %d , judul : %s\n", book.id, book.title)
		}
	case 3 :
		fmt.Printf("Masukan id buku yang anda cari :")
		inputSearch,_ := reader.ReadString('\n')
		inputSearch = strings.TrimSpace(inputSearch)
		search,_ := strconv.Atoi(inputSearch)
			for _, book := range books {
				if book.id == search {
					fmt.Println("============= Buku ditemukan ===============")
					fmt.Printf("id : %d , judul : %s\n", book.id, book.title)
						fmt.Println("=======================================")
				}
		}
    case 4 :
			fmt.Printf("Masukan id buku yang ingin anda hapus:")
			inputDelete,_ := reader.ReadString('\n')
			inputDelete = strings.TrimSpace(inputDelete)
			Delete,_ := strconv.Atoi(inputDelete)
			
			fmt.Println("sebelum dihapus :")
			for _, book := range books {
		fmt.Printf("id : %d , judul : %s\n", book.id, book.title)
		}
		fmt.Println("sesudah dihapus :")
		books = removeById(books,Delete)
for _, book := range books {
		fmt.Printf("id : %d , judul : %s\n", book.id, book.title)
		}

	
	case 5 :
		fmt.Println("=========Terimakasih =======")
		break
	}

	}

}