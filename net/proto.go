package net

import "net/http"

const (
	StatusOk           = 200 //ok
	BadUserAction      = 204 //用户操作错误
	ServerOrBadRequest = 300 //服务端或请求错误
	BadRequest         = 400 //错误请求
	ServerError        = 500 //服务端错误
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

func ServerOrBadRequestResponse(e error) Response {
	return Response{Code: ServerOrBadRequest, Msg: e.Error()}
}

func ServerOrBadRequestResponseMsg(msg string) Response {
	return Response{Code: ServerOrBadRequest, Msg: msg}
}

func ServerErrorResponse(e error) Response {
	return Response{Code: http.StatusInternalServerError, Msg: e.Error()}
}

func ServerErrorResponseMsg(msg string) Response {
	return Response{Code: http.StatusInternalServerError, Msg: msg}
}

func SuccessResponse(data interface{}) Response {
	return Response{Code: http.StatusOK, Data: data, Msg: ""}
}
func SuccessResponseMsg(msg string) Response {
	return Response{Code: http.StatusOK, Msg: msg}
}
