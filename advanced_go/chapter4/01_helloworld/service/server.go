package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {}

//该方法必须满足Go语言的RPC规则: 方法只能有两个可序列化的参数
//其中第二个参数是指针类型，并且返回一个error类型，同时必须是公开的方法
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

//将HelloService类型的对象注册为一个RPC服务
func main() {
	// rpc.RegisterName()函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数
	// 所有注册的方法会放在HelloService服务的空间之下。然后建立一个唯一的TCP链接
	// 并且通过rpc.ServeConn()函数在该TCP链接上为对方提供RPC服务
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}


