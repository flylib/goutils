package net

import "net/http"

const (
	StatusError        = 0   //处理失败
	StatusOk           = 200 //处理成功
	BadUserAction      = 204 //用户操作错误
	BadRequest         = 400 //请求错误
	StatusUnauthorized = 401 //无访问权限或者权限已经过期
	ServerError        = 500 //服务内部错误
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func BadRequestResponse(e error) Response {
	return Response{Code: BadRequest, Msg: e.Error()}
}
func BadRequestResponseMsg(msg string) Response {
	return Response{Code: BadRequest, Msg: msg}
}

func UserBadActionResponse(e error) Response {
	return Response{Code: BadUserAction, Msg: e.Error()}
}

func UserBadActionResponseMsg(msg string) Response {
	return Response{Code: BadUserAction, Msg: msg}
}

func ServerErrorResponse(e error) Response {
	return Response{Code: http.StatusInternalServerError, Msg: e.Error()}
}

func ServerErrorResponseMsg(msg string) Response {
	return Response{Code: http.StatusInternalServerError, Msg: msg}
}

func SuccessResponse(data interface{}) Response {
	return Response{Code: http.StatusOK, Data: data}
}
func SuccessResponseMsg(msg string) Response {
	return Response{Code: http.StatusOK, Msg: msg}
}
