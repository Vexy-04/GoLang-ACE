// program to find factorial in go lang
package main

import "fmt"

func main() {
	var i, n, fact int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fact = 1
	for i = 1; i <= n; i++ {
		fact = fact * i
	}
	fmt.Println("Factorial of %d is %d", n, fact)
}
