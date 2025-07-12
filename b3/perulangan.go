// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println("Perulangan ke-", i)
// 	}
// 	fmt.Println("Perulangan selesai.")
// }

// package main

// import "fmt"

// func isEven(number int) bool {
// 	return number%2 == 0
// }

// func isOdd(number int) bool {
// 	return number%2 != 0
// }

// func main() {
// 	number := 10

// 	if isEven(number) {
// 		fmt.Printf("%d adalah bilangan genap\n", number)
// 	} else {
// 		fmt.Printf("%d adalah bilangan ganjil\n", number)
// 	}

// 	number = 7

// 	if isOdd(number) {
// 		fmt.Printf("%d adalah bilangan ganjil\n", number)
// 	} else {
// 		fmt.Printf("%d adalah bilangan genap\n", number)
// 	}
// }

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d adalah bilangan genap\n", i)
		} else {
			fmt.Printf("%d adalah bilangan ganjil\n", i)
		}
	}

}
