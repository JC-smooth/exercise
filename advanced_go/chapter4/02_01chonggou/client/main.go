package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 第一步需要明确服务的名字和接口
// 服务的名字，为了避免名字冲突，我们在RPC服务的名字中增加了包路径前缀(这个是PRC服务抽象的包路径，并非完全等价于Go语言的包路径)
const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, &reply)
}


func main() {
	client, err := DialHelloService("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}