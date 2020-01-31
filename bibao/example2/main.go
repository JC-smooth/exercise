package main

import "fmt"

func main() {
	// var dummy [3]int
	// for i := 0; i < len(dummy); i++{
	// 	fmt.Println(i)
	// }
	// 输出0， 1， 2

	// var dummy [3]int
	// var f func()
	// for i := 0; i < len(dummy); i++ {
	// 	f = func() {
	// 		fmt.Println(i)
	// 	}
	// }
	// // i自加到3才会跳出循环，所以循环结束后i最后的值为3
	// f()  // 输出3

	// 用for range来实现这个例子就不会这样
	// 因为for range和for底层实现上的不同
	var dummy [3]int
	var f func()
	for i := range dummy {
		f = func(){
			fmt.Println(i)
		}
	}
	f()  // 输出2
}
