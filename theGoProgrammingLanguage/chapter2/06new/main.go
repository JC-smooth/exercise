package main

import "fmt"

// newInt1与newInt2有相同的行为
func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}


func main() {
	// 使用内置new函数创建变量
	// 表达式new(T)创建一个未命名的T类型变量，初始化为T类型的零值，并返回其地址（地址类型为*T）
	p := new(int)   // *int类型的P,指向其未命名的int变量
	fmt.Println(*p) // 输出“0”
	*p = 2          // 把未命名的int设置为2
	fmt.Println(*p) // 输出“2”

	// 每次调用new返回一个具有唯一地址的不同变量
	i := new(int)
	j := new(int)
	fmt.Println(i == j)  // false
	// 该规则有一个例外，两个变量的类型不携带任何信息且是零值，例如struct{}或[0]int, 当前的实现里面，它们有相同的地址。
	// new是一个预声明的函数，不是一个关键字，所以它可以重定义为另外的其他类型，例如
	// func delta(old, new int) int {return new - old} 函数内，内置的new函数是不可用的

}
