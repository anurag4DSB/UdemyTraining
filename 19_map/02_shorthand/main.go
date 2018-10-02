package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good Morning",
		"Jenny": "Bonjour",
	}

	fmt.Println(myGreeting)
}

/*
	This is called composite literal
	basically its composed of literal values.
	Even if there are no values in there we call it composite literal because it can have values.
*/
