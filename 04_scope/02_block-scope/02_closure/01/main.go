package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "This is the inner scope"
		fmt.Println(y)
	}
	//fmt.Println(y)  // Will not work as y i sout of scope now.
}
