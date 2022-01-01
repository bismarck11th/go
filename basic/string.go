package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello" + "World")
	fmt.Println(string("Hello World"[0])) // H

	var s string = "Hello World"
	// s[0] = "x"  これは不可

	fmt.Println(strings.Replace(s, "H", "x", 1)) // xello World (copy)
	fmt.Println(s)                               // Hello World

	fmt.Println(strings.Contains(s, "World")) // true
}
