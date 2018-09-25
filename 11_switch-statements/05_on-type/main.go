/*
	switch on types
	-- normally we switch on value of variable
	-- go allows you to switch on the type of variable
*/
package main

import "fmt"

type Contact struct {
	greetings string
	name      string
}

//we'll learn more about interfaces later
func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting x of this type
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("McLeod")
}
