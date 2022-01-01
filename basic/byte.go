package main

import "fmt"

func main() {
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b)) // HI

	c := []byte("HI")
	fmt.Println(c) // [72 73]
}
