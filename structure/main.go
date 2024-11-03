package main

import "fmt"

//Defines a new struct
type Human struct {
	name  string
	age   int
	email string
}

func main() {
	//An instance
	var man Human
	man.name = "Bob"
	man.age = 50
	man.email = "bob@internet.net"
	fmt.Println(man)

	//Another instance
	woman := Human{"Emma", 32, "emma@web.com"}
	fmt.Println(woman)

	//Pointers
	var i int = 10
	var pointer_of_i *int = &i
	fmt.Println(i, pointer_of_i)
	//Change the value of the memory
	change(pointer_of_i)
	fmt.Println(i, pointer_of_i)
	h := Human{"Justin", 20, "justin@aol.com"}
	h.introduction()
}

func change(i *int) {
	*i = 20
}

//Struct method. The infomation on the first () is a receiver that differentiates it from a regular function
func (h *Human) introduction() {
	fmt.Printf("My name is %s, I am %v years old and my e-mail is %s", h.name, h.age, h.email)
}
