package main

func main() {
	age := 18

	if age >= 18 {
		println("You are an adult")
	} else {
		println("You are a minor")
	}

	// we can declare a variable in the if statement
	if admin := true; admin == true {
		println("You are an admin")
	} else {
		println("You are not an admin")
	}

	// go does not support ternary operator

}
