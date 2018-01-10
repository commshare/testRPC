package main

import (
	"fmt"

	"../../rpc/model"
	"../../../ttrpc"
	"../consts"
	"time"
)

// Handle define handle
type Handle int

// Test  test is example
func (t *Handle) Test(args *model.RequestArg, reply *model.ResponseArg) error {

	fmt.Println("receive:", args)
	reply.RespArgSecond = time.Now().String()
	reply.RespArgOne = "response"
	reply.Code = 200
	reply.Msg = "Success"

	fmt.Println("response:", reply)

	return nil
}

func main() {

	/*0.0.0.0*/
	server := ttrpc.NewServer(consts.MASTER_ADDR+":9999")
	server.Register(new(Handle))

	/*监听请求，接收连接，处理请求*/
	server.Run()
}
