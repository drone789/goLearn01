package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 键是字符串，值是Vertex类型（结构体）
var m = map[string]Vertex {
	"Bell Labs": Vertex{
		40.45, -88.20,
	},
	"Google": Vertex{
		27.43, -23.20,
	},
}

func main() {
	// m = make(map[string]Vertex)
	// m["Bell Labs"] = Vertex{
	// 	40.68433, -74.39967,
	// }
	fmt.Println(m)
}
 	