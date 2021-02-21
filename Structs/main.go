package main

import "fmt"

type contactInfo struct {
	email string
	pinCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	michael := person{
		firstName: "Michael", 
		lastName: "Scott",
		contactInfo: contactInfo{
			email: "michael@dundermifflin.com",
			pinCode: 110069,
		},
	}

	michaelPtr := &michael;
	michaelPtr.updateName("Legend");
	michael.print();

	michael.updateName("Prison Mike");
	michael.print();
}

func (personPtr *person) updateName(newName string) {
	(*personPtr).firstName = newName;
}

func (p person) print() {
	fmt.Printf("%+v \n", p);
}