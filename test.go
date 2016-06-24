package main

import (
	"fmt"
	"math"
)

func sqrt(x,z float64) float64 {
	var y float64 = z
	z = z-(z*z -x)/(z*2)
	if math.Abs(y-z) > 0.000001 {
		return sqrt(x,z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(2,1))
	fmt.Println(math.Sqrt(2))
}
