// package main

// import "fmt"

// func main() {

// 	umur := 20

//		if umur >= 18 {
//			fmt.Println("Anda sudah dewasa.")
//		} else {
//			fmt.Println("Anda masih anak-anak.")
//		}
//	}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Membuat scanner untuk membaca input dari pengguna
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan umur Anda: ")
	input, _ := reader.ReadString('\n')

	// Membersihkan input dari karakter newline
	input = strings.TrimSpace(input)

	// Mengonversi input string ke integer
	umur, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}

	// Menampilkan umur yang diinputkan

	if umur >= 18 {
		fmt.Println("Anda sudah dewasa.")
	} else {
		fmt.Println("Anda masih anak-anak.")
	}
}
