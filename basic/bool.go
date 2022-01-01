package main

import "fmt"

func main() {
	// vat t, f bool = true, false
	t, f := true, false

	fmt.Printf("%T %v %t\n", t, t, t)
	fmt.Printf("%T %v %t\n", f, f, f)

	fmt.Println(true && true)   // true
	fmt.Println(true && false)  //false
	fmt.Println(false && false) //false

	fmt.Println(true || true)   // true
	fmt.Println(true || false)  // true
	fmt.Println(false || false) //false

	fmt.Println(!true)  //false
	fmt.Println(!false) //true
}
