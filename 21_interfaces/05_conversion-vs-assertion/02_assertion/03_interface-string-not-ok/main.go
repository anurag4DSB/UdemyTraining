package main

import "fmt"

func main() {
	var name interface{} = 2
	str, ok := name.(string)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
