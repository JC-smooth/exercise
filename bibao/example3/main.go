package main

import "fmt"

// func main() {
// 	// var funcSlice []func()
// 	// for i := 0; i < 3; i++ {
// 	// 	funcSlice = append(funcSlice, func() {
// 	// 		println(i)
// 	// 	})
// 	// }
// 	// // 这三个函数引用的都是同一个变量(i)的地址，所以之后i递增
// 	// // 解引用得到的值也会递增，所以这三个函数都会输出3
// 	// for j := 0; j < 3; j++ {
// 	// 	funcSlice[j]()
// 	// }
//
// 	// 声明新匿名函数并传参
// 	var funcSlice []func()
// 	for i := 0; i < 3; i++ {
// 		func(i int) {
// 			funcSlice = append(funcSlice, func() {
// 				fmt.Println(i)
// 			})
// 		}(i)
// 	}
//
// 	for j := 0; j < 3; j++ {
// 		funcSlice[j]()
// 	}
// }

// type student struct {
// 	name string
// 	age  int
// }
//
// func main() {
// 	m := make(map[string]*student)
// 	stus := []student{
// 		{name: "小王子", age: 18},
// 		{name: "娜扎", age: 23},
// 		{name: "大王八", age: 9000},
// 	}
//
// 	for _, stu := range stus {
// 		m[stu.name] = &stu
// 	}
// 	for k, v := range m {
// 		fmt.Println(k, "=>", v.name)
// 	}
// }

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}