package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Println(n)

	fmt.Printf("the binary of %d is:", n)

	fmt.Printf("%b\n", n) // 1111011

}
