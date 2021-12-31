package main

import "fmt"

func main() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	// slice
	for i, v := range l {
		fmt.Println(i, v)
	}

	for _, v := range l {
		fmt.Println(v)
	}

	// map
	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// keyのみ
	for k := range m {
		fmt.Println(k)
	}

	// valueのみ
	for _, v := range m {
		fmt.Println(v)
	}
}