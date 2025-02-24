package main

import "fmt"

// changeName takes a pointer to a string and changes the value of the string
// the return type is string so we need to specify it in the function signature
// and if no return value is specified, it will return an empty string by default
func changeName(n *string) string {
	// *n is the value at the address n
	// so we are changing the value at the address n
	*n = "Doe"
	// return the value at the address n
	return *n
}

func main() {
	name := "John"
	// name is a variable of type string
	fmt.Println("name = ", name)
	// &name is the address of the variable name
	// so we are passing the address of the variable name
	// to the function changeName
	// so the value of the variable name will be changed
	// to "Doe" in the function changeName
	// and the value at the address of the variable name will be returned
	fmt.Println(changeName(&name))

}
