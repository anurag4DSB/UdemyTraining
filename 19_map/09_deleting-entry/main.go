package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good Morning",
		1: "Bonjour",
	}
	myGreeting[2] = "Howdy"
	fmt.Println(myGreeting)
	delete(myGreeting, 0)
	fmt.Println(myGreeting)
}
