package main

import "fmt"

func main() {
	var value32 int32 = 129
	var value64 int64 = int64(value32)
	var value8 int8 = int8(value32)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value8) // Output will be -127 because max int8 from -128 to 127

	var name = "John"
	var nByte byte = name[0]
	var nString string = string(nByte)

	fmt.Println("Frist char in", name, "is", nString)
}
