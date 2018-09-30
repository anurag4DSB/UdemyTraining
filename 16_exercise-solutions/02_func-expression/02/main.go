package main

import "fmt"

func byTwoEven() func(int) (int, bool) {
	return func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
}

func main() {
	half := byTwoEven()
	fmt.Println(half(66))
}
