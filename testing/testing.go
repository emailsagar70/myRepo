package main

import (
	"fmt"
)

var a int

func main() {

	fmt.Println("Enter number")
	fmt.Scanln(&a)
	numCheck(&a)
}

func numCheck(b *int) {
	for *b < 0 {
		fmt.Println("enter positive number")
		fmt.Scanln(&a)
	}
	if *b > 0 {
		fmt.Println("hello")
	}
}
