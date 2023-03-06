package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	// init statement				//post statement
	for counter := 0; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Dedi", "Dedu", "Dede"}

	person := map[string]int{
		"age":  1,
		"anak": 3,
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index ke", i, "=", value)
	}

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
