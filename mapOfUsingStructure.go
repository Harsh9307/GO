package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.6843, -74.39967,
	},
	"Google": Vertex{
		37.42203, -122.08408,
	},
	// "Nagpur": Vertex{
	// 	0, 0,
	// },
}

func main() {
	fmt.Println(m)
	m["Nagpur"] = Vertex{
		0, 0,
	}
	fmt.Println(m)
}
