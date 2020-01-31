package main

import "fmt"

// 闭包是由函数和与其相关的引用环境组合而成的实体
func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main(){
	// 通过把这个函数变量赋值给i,i就成为了一个闭包
	// 所以i保存着对x的引用，可以想象i中有着一个指针指向x或i中有x的地址
	// 由于i有着指向x的指针，所以可以修改x，且保持着状态
	// i := incr()
	// fmt.Println(i())  // 1
	// fmt.Println(i())  // 2
	// fmt.Println(i())  // 3
	fmt.Println(incr()())
	fmt.Println(incr()())
	fmt.Println(incr()())
}
