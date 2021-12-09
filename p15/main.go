package main

import (
	"fmt"
	"os"
)

//打印输出
func main() {
	fmt.Fprintln(os.Stdout, "内容一", "内容二")
	fmt.Fprintf(os.Stdout, "内容三")
	fmt.Fprintln(os.Stdout, "内容四")

	fmt.Println("内容一")
	fmt.Println("内容二")
	fmt.Print("内容三")
	fmt.Println("内容四")

	a := fmt.Sprintln("1", "2")
	fmt.Println(a)
}
