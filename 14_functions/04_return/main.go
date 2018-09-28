package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}


// Sprint: string print. It will do printing but to a string.
// Not to a std out but to a string, basically concatinating them with a space between them.