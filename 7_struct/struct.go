package main

type person struct {
	name string
	age  int
}

// getDetails returns the name of the person
// in order to bind the function to the person struct
// we need to pass the person struct as the receiver
// so the function signature will be func (p person) getDetails() string
// here p is the receiver and it is of type person
// so we can access the fields of the person struct using p
// and we can return the name of the person
func (p person) getDetails() string {
	return p.name
}

// changeName takes a pointer to a person and changes the name of the person
// in order to bind the function to the person struct
// we need to pass the pointer to the person struct as the receiver
// so the function signature will be func (p *person) changeName(newName string)
// here p is the receiver and it is of type *person
// so we can access the fields of the person struct using p
// and we can change the name of the person
func (p *person) changeName(newName string) {
	p.name = newName
}

func main() {

	// create a new person
	p := person{name: "John", age: 27}

	// call the getDetails function on the person struct
	// the function is bound to the person struct
	// so we can call the function on the person struct
	// and it will return the name of the person
	name := p.getDetails()
	println(name)

	// call the changeName function on the person struct
	// the function is bound to the person struct
	// so we can call the function on the person struct
	// and it will change the name of the person
	p.changeName("Doe")
	println(p.getDetails())

}
