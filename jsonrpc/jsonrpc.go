package jsonrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type JsonRpClient struct {
	addr    string `json:"addr"`    //地址
	version string `json:"version"` //版本
}

//json rpc
type Request struct {
	Method string      `json:"method"` //方法名
	Param  interface{} `json:"param"`  //参数
}
type Response struct {
	Code int         `json:"code"` //状态码
	Data interface{} `json:"data"` //数据
	Msg  string      `json:"msg"`  //消息
}

func NewJsonRpClient(addr, version string) JsonRpClient {
	return JsonRpClient{addr: addr, version: version}
}

func (j JsonRpClient) Call(method string, param, ret interface{}) error {
	req := Request{
		Method: method,
		Param:  param,
	}
	marshal, err := json.Marshal(req)
	if err != nil {
		return err
	}
	result, err := http.Post(j.addr, "application/json;charset=UTF-8", bytes.NewReader(marshal))
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err
	}
	resp := Response{
		Code: 0,
		Data: ret,
		Msg:  "",
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return err
	}
	if resp.Code != 200 {
		return errors.New("code:" + strconv.Itoa(resp.Code) + " " + resp.Msg)
	}
	return nil
}
