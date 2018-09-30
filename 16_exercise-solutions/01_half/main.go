package _1_half

import "fmt"

func byTwoEven(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	fmt.Println(byTwoEven(24))
	h, even := byTwoEven(44)
	fmt.Println(h, even)
}
