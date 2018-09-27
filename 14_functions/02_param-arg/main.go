package main

import "fmt"

func main() {
	greet("Jane")
	greet("John")
}

func greet(name string) {
	fmt.Println(name)
}


// greet is declared with a parameter "name"
// greet is called by passing an argument(Jane, John).