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

	p4 := NewPerson("xttt")
	fmt.Println(p4.name)

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
	fmt.Println(s)
}

func NewPerson(name string) *person {
	return &person{name: name}
}

//结构体，内部一个字段s，存储错误信息
type errorString struct {
	s    string
	code int
}

//用于实现error接口
func (e errorString) Error() string {
	return e.s
}

//工厂函数，返回一个error接口，其实具体实现是*errorString
func New(text string) error {
	return &errorString{text, 1002}
}

type WalkRun interface {
	Walk()
	Run()
}

func (p *person) Walk() {
	fmt.Printf("%s能走\n", p.name)
}
func (p *person) Run() {
	fmt.Printf("%s能跑\n", p.name)
}
