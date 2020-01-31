package main

import "fmt"

type address struct {
	phone int64
	addr string
}

type company struct {
	name string
	address
}

func main(){
	var com company = company{
		name:    "xxxx",
		address: address{31231231, "dfasdf"},
	}
	fmt.Println(com.addr)
}

