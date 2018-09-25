package main

import "fmt"

func main() {
	for i := 50000; i <= 50100; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
