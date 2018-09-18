package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) // max is not the result, not the function.
}


/*
This is bad coding practice to shadow variables. Don't do this.

Error:  cannot call non-function max (type int)
 */

