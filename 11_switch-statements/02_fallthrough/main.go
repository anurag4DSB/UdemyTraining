package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Jenny":
		fmt.Println("Wassup Jenny")
		fallthrough
	case "Tim":
		fmt.Println("Wassup Tim")
	case "George":
		fmt.Println("Wassup George")
	default:
		fmt.Println("Have you no friends?")
	}
}
