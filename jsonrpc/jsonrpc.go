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

type Response struct {
	Code int         `json:"code"` //状态码
	Data interface{} `json:"data"` //数据
	Msg  string      `json:"msg"`  //消息
}

func NewJsonRpClient(addr, version string) JsonRpClient {
	return JsonRpClient{addr: addr, version: version}
}

func (j JsonRpClient) Call(method string, param, result interface{}) error {
	marshal, err := json.Marshal(param)
	if err != nil {
		return err
	}
	resp, err := http.Post(j.addr+"?method="+method, "application/json;charset=UTF-8", bytes.NewReader(marshal))
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respObj := Response{
		Code: 0,
		Data: result,
		Msg:  "",
	}
	if err := json.Unmarshal(body, &respObj); err != nil {
		return err
	}
	if respObj.Code != 200 {
		return errors.New("code:" + strconv.Itoa(respObj.Code) + " " + respObj.Msg)
	}
	return nil
}
