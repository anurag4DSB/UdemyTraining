package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thrusday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)
}
