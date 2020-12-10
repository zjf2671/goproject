package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	arr := []string{"a", "b", "c", "d", "e"}
	//arr1:=[5]string{"a","b","c","d","e"}
	//test(arr)
	//test1(arr1)
	//test2(arr)
	test3(arr)
	//test4()
	test5()
}

func test5() {
	s := "harryzhang张"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println(s[0], s[1], s[12])
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i, r := range s {
		fmt.Println(i, r)
	}
}

func test4() {
	nameAgeMap1 := make(map[string]int)
	nameAgeMap1["harry"] = 20
	nameAgeMap1["zhang"] = 22
	fmt.Println(nameAgeMap1["harry"])

	nameAgeMap := map[string]int{"harry": 20}
	age := nameAgeMap["harry"]
	fmt.Println(age)
	fmt.Println(nameAgeMap["test"])

	age, ok := nameAgeMap["test"]
	if ok {
		fmt.Println(age)
	}

	delete(nameAgeMap, "harry")

	nameAgeMap["harry1"] = 21
	nameAgeMap["harry2"] = 22
	for k, v := range nameAgeMap {
		fmt.Println("key is ", k, "value is ", v)
	}

	fmt.Println(len(nameAgeMap))

}

func test3(arr []string) {
	sli := arr[1:4]
	fmt.Println(sli)
}

func test2(arr []string) {
	for _, v := range arr {
		fmt.Printf("对应值:%s\n", v)
	}
}

func test1(arr [5]string) {
	for i, v := range arr {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}
}
func test(arr []string) {
	fmt.Println(arr[2])
}
