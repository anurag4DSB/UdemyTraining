package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names starts with M")
	case "Julien", "Sishant":
		fmt.Println("Wassup Julien / Sushant")
	}
}
