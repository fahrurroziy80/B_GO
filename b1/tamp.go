// package main

// import "fmt"

// func main() {
// 	var nama string = "fahrurrozi"
// 	umur := 20
// 	var alamat string = "lombok"

// 	fmt.Println("nama saya ", nama)
// 	fmt.Println("Umur saya", umur, "tahun.")
// 	fmt.Println("Alamat saya di", alamat)
// }

package main

import "fmt"

func main() {
	var nama_depan string = "Fahrur "
	var nama_belakang string = "Rozi"
	usia := 20
	status_mahasiswa := true

	fmt.Println("nama depan saya adalah", nama_depan)
	fmt.Println("nama belakang saya adalah", nama_belakang)
	fmt.Println("usia saya adalah", usia, "tahun")
	fmt.Println("Apakah saya mahasiswa?", status_mahasiswa)
	fmt.Println("Nama lengkap saya adalah", nama_depan+nama_belakang)
}
