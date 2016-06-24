package main

import "fmt"

func main() {
	i, j := 4, 50
	// 创建一个指针p, 指向i的地址
	p := &i
	fmt.Println(*p)

	// 给p指向的内存地址赋值
	*p = 44
	fmt.Println(i)

	p = &j
	*p = *p /25
	fmt.Println(j)
}