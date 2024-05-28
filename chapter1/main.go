package main // package declaration

import "fmt"

// this is a comment!

func main() {
	var input string
	fmt.Println("Input your name: ")
	fmt.Scan(&input)
	fmt.Println("Hello", input)
}
