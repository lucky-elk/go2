package main

import "fmt"

//变量声明和赋值
func main() {
	var v1 string
	v1 = "a"
	fmt.Println(v1)

	v2 := 1
	fmt.Println(v2)

	var v3, v4 = 1, "a"
	fmt.Println(v3, v4)

	var (
		v5 = 6
		v6 = "b"
	)
	fmt.Println(v5, v6)
}
