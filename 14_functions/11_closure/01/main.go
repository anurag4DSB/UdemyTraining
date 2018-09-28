package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "This is only in this block"
		fmt.Println(y)
	}
	// fmt.Println(y) //outside of scope
}
