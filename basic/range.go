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

	// sliceから最小値取得
	slice := []int{100, 300, 23, 11, 2, 4, 6, 4}

	var min int
	for i := 0; i < len(slice); i++ {

		for i, num := range slice {
			if i == 0 {
				min = num
				continue
			}

			if min >= num {
				min = num
			}
		}
	}

	fmt.Println("最小値:", min)

	// mapのvalueの合計値取得
	m2 := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	sum := 0
	for _, v := range m2 {
		sum += v
	}

	fmt.Println(sum)
}
