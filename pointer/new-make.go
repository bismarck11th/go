package main

import "fmt"

func main() {
	// makeとnewの違い -> ポインタを返すのがnew
	s := make([]int, 0)   // slice
	fmt.Printf("%T\n", s) // int[]

	m := make(map[string]int)
	fmt.Printf("%T\n", m) // map[string]int

	ch := make(chan int)
	fmt.Printf("%T\n", ch) // chan int

	var p *int = new(int)
	fmt.Printf("%T\n", p) // *int

	var st = new(struct{})
	fmt.Printf("%T\n", st) // *struct {}

	var p2 *int = new(int)
	fmt.Println(p2)
	fmt.Println(*p2) // 0

	var p3 *int     // メモリ領域を確保しない
	fmt.Println(p3) // nil
}
