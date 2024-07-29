package main

// import (
// 	"fmt"
// )

// func printPrimes(max int) {
// 	for n := 2; n <= max; n++ {
// 		if n == 2 {
// 			fmt.Print(n, ",")
// 		}
// 		if n%2 == 0 {
// 			continue
// 		}
// 		isPrime := true
// 		for i := 3; i*i <= max; i++ {
// 			if i*i == n {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			fmt.Print(n, ",")
// 		}
// 	}
// }

// // don't edit below this line

// func test(max int) {
// 	fmt.Printf("Primes up to %v:\n", max)
// 	printPrimes(max)
// 	fmt.Println("===============================================================")
// }

// func main() {
// 	test(10)
// 	test(20)
// 	test(30)
// }
