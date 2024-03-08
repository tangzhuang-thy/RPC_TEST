package main

import (
	"fmt"
	"github.com/tangzhuang-thy/RPC_TEST/shared"
	"net"
	"net/rpc"
)

type Calculator int

func (c *Calculator) Add(args shared.Args, result *int) error {
	*result = args.A + args.B
	return nil
}

func (c *Calculator) Multiply(args shared.Args, result *int) error {
	*result = args.A * args.B
	return nil
}

func main() {
	calculator := new(Calculator)
	rpc.Register(calculator)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}

	fmt.Println("Server is running on port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
