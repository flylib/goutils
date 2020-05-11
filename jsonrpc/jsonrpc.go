package jsonrpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type JsonRpClient struct {
	addr    string `json:"addr"`    //地址
	version string `json:"version"` //版本
}

func NewJsonRpClient(addr, version string) JsonRpClient {
	return JsonRpClient{addr: addr, version: version}
}

func (j JsonRpClient) Call(param, result interface{}) error {
	marshal, err := json.Marshal(param)
	if err != nil {
		return err
	}
	resp, err := http.Post(j.addr, "application/json;charset=UTF-8", bytes.NewReader(marshal))
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, result); err != nil {
		return err
	}
	return nil
}
