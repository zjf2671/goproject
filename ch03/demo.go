package main

import (
	"fmt"
)

func main() {


	//test()
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	test7()

}
func test()  {
	i:=10
	if i>10{
		fmt.Println("i>10")
	}else if i==10{
		fmt.Println("i==10")
	}else {
		fmt.Println("i<10")
	}
}
func test1()  {
	if i:=8;i>8{
		fmt.Println("i>8")
	}else if i==8{
		fmt.Println("i==8")
	}else {
		fmt.Println("i<8")
	}
}
func test2()  {
	switch i:=11; {
	case i>10:
		fmt.Println("i>10")
	case i>5&&i<10:
		fmt.Println("i>5&&i<10")
	default:
		fmt.Println("i<=5")
	}
}
func test3()  {
	switch i:=1;i {
	case 1:
		fallthrough
		//fmt.Println("匹配到1")
	case 2:
		fmt.Println("匹配到2")
	default:
		fmt.Println("没有匹配到")
	}
}
func test4()  {
	sum:=0
	for i:=1;i<=100;i++{
		sum+=i
	}
	fmt.Println(sum)
}
func test5()  {
	sum:=0
	i:=1
	for i<=100 {
		sum+=i
		i++
	}
	fmt.Println(sum)
}
func test6() {
	sum:=0
	i:=1
	for {
		sum+=i
		i++
		if(i>100){
			break
		}
	}
	fmt.Println(sum)
}
func test7()  {
	sum:=0
	for i:=1; i<=100; i++ {
		if i%2!=0 {
			continue
		}
		sum+=i
	}
	fmt.Println(sum)
}
