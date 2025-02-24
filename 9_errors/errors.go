package main

import "fmt"

func checkByZero(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by 0 not allowed")
	}
	return (a / b), nil
}

func main() {

	result, err := checkByZero(2, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result:", result)
	}

}
