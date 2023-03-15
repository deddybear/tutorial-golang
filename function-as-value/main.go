package main

import "fmt"

func getGoodBye(name string) string {
	return "good bye " + name
}

func main() {

	sayGoodBye := getGoodBye

	result := sayGoodBye("Bear")

	fmt.Println(result)

}
