package main

import "fmt"

func main() {
	if !true {
		fmt.Println("THis did not run")
	}

	if !false {
		fmt.Println("This ran")
	}
}
