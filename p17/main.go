package main

import "fmt"

//用户输入
func main() {

	var name string
	var age int
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入年龄：")
	fmt.Scanf("%d", &age)
	fmt.Println("输入的是：", name, age)
}
