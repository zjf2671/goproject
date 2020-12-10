package main

import "fmt"

func main() {
	age := Age(2)
	age.String()
	age.Modify()
	age.String()
	(&age).Modify()
	age.String()
}

type Age uint

func (age Age) String() {
	fmt.Println("this age is", age)
}

func (age *Age) Modify() {
	*age = Age(40)
}
