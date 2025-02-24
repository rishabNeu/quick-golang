package main

import "fmt"

// add takes two integers and returns their sum
// return type is int so we need to specify it
// in the function signature and if no return
// value is specified, it will return 0 by default
func add(x int, y int) int {
	return x + y
}

// sub take same type of arguments and returns their difference
// so at the end we specify the return type as int
// no need to specify at both places if both are same
func sub(x, y int) int {
	return x - y
}

// multiple return values
// we can return multiple values from a function
// by specifying the return types in the function signature
// and then returning the values in the return statement
// the return values are separated by commas
func getLanguages() (string, string, string) {
	return "Go", "Python", "Java"
}

func processIt(fn func(int, int) int) {
	fmt.Println(fn(42, 13))
}

func main() {
	result := add(42, 13)
	sub_result := sub(42, 13)
	fmt.Println("result = ", result, sub_result)

	// we can ignore the return values by using _
	// here we did not want to use the third return value
	// so we ignored it using _
	// so the return values are assigned to lang1 and lang2 & there is no error for the third value
	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)
	// fmt.Println(getLanguages())
}
