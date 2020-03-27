package main

import (
	"log"
	"net"
	"net/rpc"
)

// 第一步需要明确服务的名字和接口
// 服务的名字，为了避免名字冲突，我们在RPC服务的名字中增加了包路径前缀(这个是PRC服务抽象的包路径，并非完全等价于Go语言的包路径)
const HelloServiceName = "path/to/pkg.HelloService"

type HelloService struct{}

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}
// RegisterHelloService注册服务时， 编译器会要求传入的对象满足HelloServiceInterface接口
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main(){
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		// 支持多个TCP链接
		go rpc.ServeConn(conn)
	}

}
