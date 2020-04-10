package main

import (
	"fmt"
	"github.com/theGoProgrammingLanguage/chapter2/tempconv"
)

func main() {
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.FreezingC))
	// fmt.Printf("%g\n", boilingF-tempconv.FreezingC)  // 编译错误， 类型不匹配

	// var c tempconv.Celsius
	// var f tempconv.Fahrenheit
	// fmt.Println(c == 0) // true
	// fmt.Println(f >= 0) // true
	// // fmt.Println(c == f)  // 编译错误， 类型不匹配
	// fmt.Println(c == tempconv.Celsius(f)) // true

	c := tempconv.FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}
