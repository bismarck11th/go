package main

import (
	"fmt"
)

// 関数の外でも宣言可
var (
	i   int     = 1
	f64 float64 = 1.2
	s   string  = "test"
	t   bool    = true
	f   bool    = false
)

func foo() {
	// 関数の外では宣言不可
	xi := 1
	xi = 2
	// 型を指定したい場合にはvarを使う
	var xf32 float32 = 1.2
	xs := "test"
	xt := true
	xf := false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32) // 型を確認
	fmt.Printf("%T\n", xi)   // 型を確認
}

func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}
