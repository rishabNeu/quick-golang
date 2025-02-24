package main

import "fmt"

// for only loop in Go
func main() {
	// while loop implementation using for
	i := 1
	for i < 5 {
		fmt.Println(i)
		i = i + 1
	}

	// classic for loop
	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

	// range loop
	for i := range 10 {
		fmt.Println(i)
	}

}
