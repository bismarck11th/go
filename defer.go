package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

func main() {
	foo()

	defer fmt.Println("world")
	fmt.Println("hello")

	fmt.Println("hello")

	file, _ := os.Open("./lesson.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read((data))
	fmt.Println(string(data))
}
