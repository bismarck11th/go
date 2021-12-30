package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx) // float64 1 1.000000

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy) // int 1 1

	var s string = "14"
	// z = int(s)  NG
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i) // int 1 1
}
