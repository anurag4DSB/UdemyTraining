package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world()
	hello()
}


// defer basically deferred the execution of this function
// Where did it defer it to? right before main exits