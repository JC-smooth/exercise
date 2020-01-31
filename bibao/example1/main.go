package main

import "fmt"

func main() {
	// 每次迭代后都对i进行了解引用并使用得到的值不再使用，所以这段代码会正常输出
	for i := 0; i < 3; i++ {
		func() {
			fmt.Println(i)
		}()
	}
}
