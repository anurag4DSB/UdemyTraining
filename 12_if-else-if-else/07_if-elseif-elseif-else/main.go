package main

import "fmt"

func main() {
	if false {
		fmt.Println("First print statement")
	} else if false {
		fmt.Println("second print statement")
	} else if true {
		fmt.Println("third print statement")
	} else {
		fmt.Println("forth print statement")
	}
}
