package main

import "fmt"

func main() {
	m := make([ ]string, 2, 25) // this is making a slice
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
}

func changeMe(z []string) {
	z[0] = "Todd"
	z[1] = "Anurag"
	fmt.Println(z)
}


// slice is a reference type Its an address pointing to an underlying array.