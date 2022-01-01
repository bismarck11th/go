package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
}

func main() {
	a := Vertex{1, 2, "test"}
	changeVertex(a)
	fmt.Println(a) // {1 2 test}

	a2 := &Vertex{1, 2, "test"}
	changeVertex2(a2)
	fmt.Println(a2) // {1000 2 test}

	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)

	var v5 Vertex
	fmt.Println(v5)

	v6 := new(Vertex)
	fmt.Println(v6)

	v7 := &Vertex{}
	fmt.Println(v7)
}
