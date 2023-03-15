package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// struct method

func (cust Customer) Greetings(age int) {
	fmt.Println("Greetings", cust.Name, "My age is", age)
}

func main() {
	var cus Customer
	cus.Name = "Abi"
	cus.Address = "Indonesia"
	cus.Age = 13

	//init stuct
	cus1 := Customer{
		Name:    "Joko",
		Address: "Cirebon",
		Age:     0,
	}

	cus.Greetings(19)
	//fmt.Println(cus)
	//fmt.Println(cus.Name)

	fmt.Println(cus1)
}
