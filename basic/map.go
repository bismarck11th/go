package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)

	fmt.Println(m["apple"])

	m["banana"] = 300
	fmt.Println(m)

	m["new"] = 500
	fmt.Println(m)

	fmt.Println(m["nothing"]) // 0

	v, ok := m["apple"]
	fmt.Println(v, ok) // 100 true

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2) // 0 false

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)
}
