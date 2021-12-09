package main

import "fmt"

//转义字符
func main() {
	fmt.Printf("%d %x %X %o %b", 18, 11, 12, 14, 44)

	a := fmt.Sprintf("%x", 18)
	fmt.Println("\n", a)

	fmt.Printf("%t", false)

	fmt.Printf("%c", 65)

	fmt.Printf("%s %q", "你好", "你好")

	fmt.Printf("花费了%d%%总钱数", 15)

}
