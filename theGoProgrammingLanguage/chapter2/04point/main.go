package main

import "fmt"

func main(){
	x := 1
	p := &x
	fmt.Println(&x)
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	// 指针类型的零值是nil, 测试p!= nil, 结果是true说明p指向一个变量。
	// 指针是可以比较的, 两个指针当且仅当指向同一个变量或者两者都是nil的情况下才相等
	var a, y int
	var i, j *int
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(&a == &a, &a == &y, &a == nil)  // true false false
	fmt.Println(&a == &a, a == y, &a == nil)  // true true false
	fmt.Println(&i == &i, i == j, i == nil)  // true true true

	// var b = f()
	// 每次调用f都会返回一个不同的值
	fmt.Println(f() == f())
	fmt.Println(f())
	fmt.Println(f())

	v := 1
	incr(&v)  // v现在等于2
	fmt.Println(incr(&v))  // v现在等于3

}
func f() * int {
	v := 1
	return &v
}

func incr(p *int) int {
	// 递增一个指针参数所指向的变量，然后返回此变量的新值
	*p++  // 递增p所指向的值；p自身保持不变
	return *p
}