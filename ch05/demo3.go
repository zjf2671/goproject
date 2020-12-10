package main

import (
	"errors"
	"fmt"
)

func main() {

	test()
	result := sum(1, 3)
	fmt.Println(result)

	result, err := sum1(1, -8)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	result1, _ := sum1(2, 4)
	fmt.Println(result1)

	result2, _ := sum2(3, 4)
	fmt.Println(result2)

	myPrintln := MyPrintln(2, 3, 4)
	fmt.Println(myPrintln)

	test1()

	c1 := test2()
	fmt.Println(c1(2))
	fmt.Println(c1(2))

}

func test2() func(a int) int {
	i := 0
	return func(a int) int {
		i++
		return i + a
	}
}

func test1() {
	sum := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(2, 3))
}

func MyPrintln(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}
	return sum
}

func sum2(a int, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a or b is not Negative")
	}
	sum = a + b
	err = nil
	return
}

func sum1(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a or b is not Negative")
	}
	return a + b, nil
}

func sum(a int, b int) int {
	return a + b
}
func test() {
	aa := [3][3]int{}
	aa[0][0] = 1
	aa[0][1] = 2
	aa[0][2] = 3
	aa[1][0] = 4
	aa[1][1] = 5
	aa[1][2] = 6
	aa[2][0] = 7
	aa[2][1] = 8
	aa[2][2] = 9
	fmt.Println(aa)
}
