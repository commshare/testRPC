package main

import (
	"fmt"

	"../../rpc/model"
	"../../../ttrpc"
)

// Handle define handle
type Handle int

// Test  test is example
func (t *Handle) Test(args *model.RequestArg, reply *model.ResponseArg) error {

	fmt.Println("receive:", args)

	reply.RespArgOne = "response"
	reply.Code = 200
	reply.Msg = "Success"

	fmt.Println("response:", reply)

	return nil
}

func main() {

	server := ttrpc.NewServer("0.0.0.0:9999")
	server.Register(new(Handle))

	/*监听请求，接收连接，处理请求*/
	server.Run()
}
