package main

import "fmt"

func main() {
	myFriendsName := "Al"
	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Hello my friend with name length of 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Jen":
		fmt.Println("Wassup Jen")
	case myFriendsName == "George", myFriendsName == "Tim":
		fmt.Println("Wassup GeorgeTim")
	default:
		fmt.Println("nothing matched, this is default")
	}
}
