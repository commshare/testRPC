package main

import (
	"fmt"

	"../../rpc/model"
	"../../../ttrpc"
	"../consts"
	"sync"
	"net"
	"../utils"
)

func main() {
	/*127.0.0.1*/
	client := ttrpc.NewClient(consts.MASTER_ADDR+":9999", 10)
	clientTest(client)
	clientRegister(client)
}

// Worker holds the state for a server waiting for DoTask or Shutdown RPCs
type Worker struct {
	sync.Mutex
	name   string
	nRPC   int // protected by mutex
	nTasks int // protected by mutex
	l      net.Listener
}


func clientTest(client * ttrpc.Client ){
	req := new(model.RequestArg)
	req.ArgOne = "test"
	var resp model.ResponseArg
	fmt.Printf("send: %+v\n", req)
	err := client.Call("Handle.Test", req, &resp)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Printf("resp: %+v\n", resp)
	}
}

func clientRegister(client * ttrpc.Client){
	registerArgs := new(model.RegisterArgs)
	registerArgs.WorkerName="register_client_1"
	registerArgs.IpAddress = utils.GetFirstIP()
	var resp model.ResponseArg
	fmt.Printf("send: %+v\n", registerArgs)
	err := client.Call("Handle.Register",registerArgs,&resp)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Printf("resp: %+v\n", resp)
	}
}