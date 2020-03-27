package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main(){
	// rpc.Dial拨号RPC服务
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	// 通过client.Call()调用具体的RPC方法。在调用client.Call()时
	// 第一个参数是用点号链接的RPC服务名字和方法名字
	// 第二个和第三个参数分别是定义RPC方法的两个参数
	err = client.Call("HelloService.Hello", "xxxx", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
