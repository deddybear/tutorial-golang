package main

import "fmt"

func main() {
	name := "testaz"

	switch name {
	case "test":
		fmt.Println("is dummy")
	case "joko":
		fmt.Println("is human")
	default:
		fmt.Println("None")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Too long")
	case false:
		fmt.Println("Ok")
	}

	length := len(name)

	switch {
	case length > 5:
		fmt.Println("Too long")
	case length < 4:
		fmt.Println("Too Short")
	default:
		fmt.Println("ok")
	}
}
