package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	var a int = 0
	var b int = 1
	var i int = 0

	return func() int {
		i++
		if i == 1 {
			return 1
		} else {
			c := a + b
			a = b
			b = c
			return c
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}