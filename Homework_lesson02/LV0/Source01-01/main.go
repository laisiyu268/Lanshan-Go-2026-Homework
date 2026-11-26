package main

import (
	"fmt"
)

func main() {
	var a int32 = 20
	var ip *int32
	ip = &a
	fmt.Printf("a 变量的地址: %x\n", &a)
	fmt.Printf("ip 变量储存的变量地址: %x\n", ip)
	fmt.Printf("ip 变量的值: %d\n\n", *ip)
	a = 30
	fmt.Printf("a 变量的地址: %x\n", &a)
	fmt.Printf("ip 变量储存的变量地址: %x\n", ip)
	fmt.Printf("ip 变量的值: %d\n\n", *ip)
	*ip = 40
	fmt.Printf("a 变量的地址: %x\n", &a)
	fmt.Printf("ip 变量储存的变量地址: %x\n", ip)
	fmt.Printf("ip 变量的值: %d\n\n", *ip)

	var ptr *int32
	fmt.Printf("ptr变量的地址: %x\n", ptr)
	// fmt.Printf("ptr变量的值:%d\n\n", *ptr) runtime error
	fmt.Println((ptr == nil))
}

/*
a 变量的地址: c000188008
ip 变量储存的变量地址: c000188008
ip 变量的值: 20

a 变量的地址: c000188008
ip 变量储存的变量地址: c000188008
ip 变量的值: 30

a 变量的地址: c000188008
ip 变量储存的变量地址: c000188008
ip 变量的值: 40

成功: 进程退出代码 0.
*/
