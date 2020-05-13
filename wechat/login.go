package wechat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type WxAuthCode2Session struct {
	Openid     string `json:"openid"`      //用户唯一标识
	SessionKey string `json:"session_key"` //会话密钥
	Uinionid   string `json:"uinionid"`    //用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
	Errcode    int32  `json:"errcode"`     //错误码
	ErrMsg     string `json:"errmsg"`      //错误信息
}

func WeChatLogin(wxAppId, wxAppSecret, jsCode string) (wxToken *WxAuthCode2Session, err error) {
	var res *http.Response
	res, err = http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + wxAppId + "&secret=" + wxAppSecret + "&js_code=" + jsCode + "&grant_type=authorization_code")
	if err != nil {
		return wxToken, err
	}
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	//将序列化的byte[]重写反序列化为对象。
	err = json.Unmarshal(body, &res)
	if err != nil {
		return wxToken, err
	}
	return wxToken, nil
}
