package main

import "fmt"

//without params
func sayHello() {
	fmt.Println("Hello World")
}

//with params
func sayHelloTo(fristname string, lastname string, age int) {
	fmt.Println(fristname, lastname, "selamat ulang tahun ke", age)
}

//with params + return
func sum(val, val2 int) int {
	return val + val2
}

//with params + multiple return
func getInformation(name string, age int) (string, int) {
	return name, age
}

//with params + named return values

func listPet(fname, lname string, ages, sexx int, life bool) (firstName, lastName string, age, sex int, die bool) {

	//must init if not will null
	firstName = fname
	lastName = lname
	age = ages
	sex = sexx
	die = life
	return
}

//variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	sayHello()
	sayHelloTo("Dudu", "Hermawan", 25)
	result := sum(1, 5)
	fmt.Println(result)
	// _, == ignore return
	_, age := getInformation("test", 10)
	fmt.Println(age)

	fristname, _, age, _, die := listPet("bobo", "lumbo", 10, 1, true)

	fmt.Println(fristname, age, die)

	total := sumAll(10, 20, 30, 40)
	fmt.Println(total)

	numbers := []int{10, 10, 10, 10}
	total = sumAll(numbers...)
	fmt.Println(total)
}
