package main

import "fmt"

func main() {
	xs := []int{1,2,3,4,5,9,11}

	for i, value := range xs {
		fmt.Println(i, " - ", value)
	}
}
