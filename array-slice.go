package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// 配列は要素数を指定する必要があるが、スライスは要素数を指定する必要がない
	// 配列
	var b [2]int = [2]int{100, 200}
	// b = append(b, 300) NG
	fmt.Println(b)

	/* スライス
		スライスの作成方法は３つ
		1. リテラル
		2. 簡易スライス式
		3. 組み込み関数make
	*/

	// リテラル
	c := []int{1, 2, 3, 4}
	c = append(c, 5)
	fmt.Println(c)
	fmt.Println(c[1])   // 2
	fmt.Println(c[2:4]) // [3 4]
	fmt.Println(c[:2])  // [1 2]
	fmt.Println(c[2:])  // [3 4 5]
	fmt.Println(c[:])   // [1 2 3 4 5]

	var board = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	fmt.Println(board)

	// make (make([]要素の型, 要素の数, 容量))
	d := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(d), cap(d), d)
	d = append(d, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(d), cap(d), d)
	d = append(d, 0, 0, 0, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(d), cap(d), d) // ok

	e := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(e), cap(e), e)

	// f := make([]int, 5)
	f := make([]int, 0, 5)
	for i := 1; i <= 5; i++ {
		f = append(f, i)
		fmt.Println(f)
	}
}
