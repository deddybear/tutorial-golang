package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "dedi"
	}

	registerUser("dedi", blacklist)
	registerUser("admin", blacklist)
	// or
	registerUser("budi", func(name string) bool {
		return name == "budi"
	})
}
