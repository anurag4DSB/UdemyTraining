package main

import "fmt"

func main() {
	var largeNum int
	var smallNum int

	fmt.Print("Enter a large number: ")
	fmt.Scan(&largeNum)
	fmt.Print("Enter a large number: ")
	fmt.Scan(&smallNum)
	fmt.Printf("The remainder is: %d\n", largeNum%smallNum)
}