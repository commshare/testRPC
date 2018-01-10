package main

import (
	"fmt"

	"../../rpc/model"
	"../../../ttrpc"
	"../consts"
)

func main() {
	/*127.0.0.1*/
	client := ttrpc.NewClient(consts.MASTER_ADDR+":9999", 10)
	clientTest(client)

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