package main

import (
	"fmt"
)

func main() {
	arr:=[]string{"a","b","c","d","e"}
	//arr1:=[5]string{"a","b","c","d","e"}
	//test(arr)
	//test1(arr1)
	//test2(arr)
	test3(arr)
}

func test3(arr []string) {
	sli:=arr[1:4]
	fmt.Println(sli)
}

func test2(arr []string) {
	for _,v:=range arr{
		fmt.Printf("对应值:%s\n", v)
	}
}

func test1(arr [5]string) {
	for i,v:=range arr{
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}
}
func test(arr []string) {
	fmt.Println(arr[2])
}


