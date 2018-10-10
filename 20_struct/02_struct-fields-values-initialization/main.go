package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"james", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 19}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}


// struct is an aggregate type and this is type of encapsulation as we are encapsulating data types together