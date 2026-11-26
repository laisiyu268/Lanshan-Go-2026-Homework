package main

import (
	"fmt"
)

func main() {
	arr1 := [3]string{"a", "b", "c"}
	var v int = 8
	fmt.Printf("%v\n", v)
	fmt.Printf("%+v\n", v)
	fmt.Printf("%#v\n", v)
	fmt.Printf("The arr1: %v\n", arr1)
	fmt.Printf("The arr1: %+v\n", arr1)
	fmt.Printf("The arr1: %#v\n\n", arr1)

	arr2 := modifyArray(arr1)
	fmt.Printf("The arr2: %v\n", arr2)
	fmt.Printf("The arr1: %v\n", arr1)

	s := []int{1, 2, 5}
	modifySlice(s)
	fmt.Println(s)
}
func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
	/*
		所有传给函数的参数值都会被复制，
		函数内部操作的只是副本。数组是值类型，
		复制时会连同所有元素一起拷贝，
		所以函数里改的只是副本。
	*/
}
func modifySlice(s []int) {
	s[0] = 100
	/*
		切片、map、channel 这类"引用语义"的类型，
		复制时只拷贝它们本身（对切片来说就是那个小结构体）
		底层数据是共享的——这叫浅表复制，不是深拷贝。

	*/
}
