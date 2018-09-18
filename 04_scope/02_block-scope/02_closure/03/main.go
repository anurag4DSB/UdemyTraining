package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println((increment()))
	fmt.Println((increment()))
}

/*
We have closer outer scope enclosing closure inner scope.
Closure helps us limit the scope of variables used by multiple functions without closure, for two or more funcs to have to access to the same variable, that variable would need to be package scope.
anonymous function: A function without a name

func expression: assigning function to a variable
*/
