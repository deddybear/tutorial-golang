package main

import "fmt"

func main() {

	fmt.Println("Hello world")

	// var <nama-variabel> <tipe-data> = <nilai>
	const string1 string = "john"
	const string2 string = "wick"
	const countString int = (len(string1) + len(string2))
	// const string3 string = 'hublah' #WRONG beacuse use single quotes -> ''!

	// fmt.Printf("Hello %s %s! \n", firstName, lastName)
	fmt.Println("hello", string1, string2+"!")
	fmt.Println("Count String : ", countString)
	fmt.Println("Get Frist Char in byte : ", string1[0])
}
