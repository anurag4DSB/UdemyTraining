package main

import "fmt"

func greatest(numbers ...int) int {
	var max int
	for i, v := range numbers {
		if i == 0 {
			max = v
		} else {
			if max < v {
				max = v
			}
		}
	}
	return max
}

func main() {
	fmt.Println(greatest(1, 2, 3, 4, 5, 6))
}
