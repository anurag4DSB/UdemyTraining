package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"zero":		"Good Morning",
		"one":		"Bonjour",
		"Two":		"Buenos dias",
		"Three":	"Bongiorno",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, "Two")
	fmt.Println(myGreeting)
}
