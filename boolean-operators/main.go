package main

import "fmt"

func main() {
	const (
		exam       = 85
		attendance = 75
	)

	var passedExam = exam >= 85
	var passedAtten = attendance >= 80

	var passed = passedExam && passedAtten
	fmt.Println(passed)
	// or
	fmt.Println(exam >= 85 && attendance >= 80)
}
