package app

import "golang-rest-api-template/pkg/e"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(code int, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	}
}
