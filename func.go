package main

import "fmt"

func add(x, y int) (int, int) {
	return x + y, x - y
}

// int型の変数resultを返すよう指定
func cal(price, item int) (result int) {
	result = price * item
	return // result 書かなくてもいい
}

func convert(price int) float64 {
	return float64(price)
}

// クロージャー
func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

// 可変長引数
func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
	r1, r2 := add(1, 2)
	fmt.Println(r1, r2)

	r3 := cal(100, 3)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(1)

	// クロージャ
	counter := incrementGenerator()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	c1 := circleArea(3.14)
	fmt.Println(c1(2))

	c2 := circleArea(3)
	fmt.Println(c2(2))

	// 可変長引数
	foo(1, 2, 3)
	foo(1, 2, 3, 4)
	
	s := []int{1, 2, 3}
	foo(s...)
}
