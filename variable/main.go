package main

import "fmt"

func main() {

	//declare variabel
	var varName string
	varName = "Test var"
	fmt.Println("Declare Variable : ", varName)
	// or without data type
	var varName1 = "test var 2"
	// varName = true #ERROR because varName has been declared typed string
	fmt.Println("Declare Variable Without Data Type : ", varName1)

	// or without var and data type
	varName3 := "Test Var 3"
	fmt.Println("Declare Variable Without Data Type and keyword var/const : ", varName3)
	// varName3 := "Test Var 4" #ERROR because ":=" just for frist initialization variable

	// for re-assign value variable use "="
	varName3 = "Test Var 4"
	fmt.Println("Re-Assign Variable Without Data Type and keyword var/const : ", varName3)

	// or with const
	const varConst string = "test const"
	fmt.Println("Re-Assign Variable Wit Data Type and keyword const : ", varConst)

	// declarate multiple variable
	var (
		name   string = "number "
		number int8   = 9
	)

	fmt.Println("Declarate multiple Variable : ", name, number)

	name = "angka"
	number = 10

	fmt.Println("Re-assign multiple Variable : ", name, number)

}
