package main

import "fmt"

func main() {
	var a,b []int
	printSlice("a", a)

	a = append(a,0)
	printSlice("a", a)

	b = append(a,1,2,3,4,5,6,7,8)
	printSlice("a", b)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}