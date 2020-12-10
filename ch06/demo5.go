package main

import "fmt"

func main() {

	p := person{"张飞", 33, address{}}
	fmt.Println(p.age, p.name)
	p1 := person{age: 33, name: "关习习"}
	fmt.Println(p1.name, p1.age)
	p2 := person{name: "刘备"}
	fmt.Println(p2.name, p2.age)

	p3 := person{"曹操", 44, address{
		"江东", "大地",
	}}
	fmt.Println(p3.name, p3.age, p3.addr.province, p3.addr.city)

	printString(p)

}

type person struct {
	name string
	age  int
	addr address
}

type address struct {
	province string
	city     string
}

type Stringer interface {
	String() string
}

func (p person) String() string {
	return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}
