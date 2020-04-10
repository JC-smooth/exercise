package main

import (
	"flag"
	"fmt"
	"strings"
)

// echo输出其命令行参数
// flag.Bool
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main(){
	// 在使用标识前， 必须调用flag.Parse来更新标识变量的默认值
	flag.Parse()
	fmt.Println(n)
	fmt.Println(*n)
	// 非标识参数也可以从flag.Args()返回的字符串slice来访问。
	// 如果flag.Parse遇到错误，它输出一条帮助消息，然后调用os.Exit(2)来结束程序
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
