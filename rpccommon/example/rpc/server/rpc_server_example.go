package main

import (
	"fmt"

	"../../rpc/model"
	"../../../ttrpc"
	"../consts"
	"time"
	"sync"
)

// Handle define handle
type Handle struct{
	sync.Mutex
	registerChannel chan *model.RegisterArgs

}

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

func (t *Handle) Register(args *model.RegisterArgs, reply *model.ResponseArg) error {

	fmt.Println("Register :receive:", args)
	reply.RespArgSecond = time.Now().String()
	reply.RespArgOne = "response"
	reply.Code = 200
	reply.Msg = "Success"

	t.Lock()
	defer t.Unlock()
	/*加入注册*/
	t.registerChannel <- args
	fmt.Println("response:", reply)

	return nil
}

func (t *Handle)showRpcAgent(){
	for {
		select {
		case agent,ok := <- t.registerChannel :
			if ok == true {
				fmt.Printf("agent  name %v AND  ip : %v \n",agent.WorkerName ,agent.IpAddress)
			}else{
				fmt.Printf("agent channel closed !~! \n")
			}

		}
	}

}
func main() {

	/*0.0.0.0*/
	server := ttrpc.NewServer(consts.MASTER_ADDR+":9999")
	h := new(Handle)
	h.registerChannel = make(chan *model.RegisterArgs)

	server.Register(h)

	go h.showRpcAgent()
	/*监听请求，接收连接，处理请求*/
	server.Run()
}
