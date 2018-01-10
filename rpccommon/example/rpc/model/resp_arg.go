package model

import (
	"fmt"
)

// ResponseArg response arg example
type ResponseArg struct {
	RespArgOne string `json:"resp_arg_one"`
	RespArgSecond string `json:"resp_arg_second"`

	CommonRPCResponse
}

// String response arg to string
func (resp *ResponseArg) String() string {
	return fmt.Sprintf("<Code: %d, Msg: %s, ResponseArg: %s (and %s) >", resp.Code, resp.Msg, resp.RespArgOne,resp.RespArgSecond)
}
