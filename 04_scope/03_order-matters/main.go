package main

import "fmt"

func main() {
	fmt.Println(y)
	fmt.Println(x)
	x := 42
}

var y = 42

/*
This wont run as we are trying to print x before its scope starts.
Will through an error: # command-line-arguments
./main.go:6:14: undefined: x
But the order does not matter for outer scope.
*/
