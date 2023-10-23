package nacos

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	SPLIT        = "%02"
	HeaderMD5Key = "Content-MD5"
)

// curl -X GET 'http://127.0.0.1:8848/nacos/v1/cs/configs?dataId=nacos.example&group=com.alibaba.nacos'
// http://192.168.119.128:8848/nacos/v1/cs/configs?dataId=test&group=test&tenant=8ff90940-a1b1-449c-84e2-ef57fe9f8272
func (c *Client) GetConfig(group, dataID string) ([]byte, string, error) {
	args := url.Values{}
	args.Set("group", group)
	args.Set("dataId", dataID)
	args.Set("tenant", c.namespaceId)

	resp, err := http.Get(c.baseURI + "/v1/cs/configs?" + args.Encode())
	if err != nil {
		return nil, "", err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, "", errors.New("Response code " + resp.Status)
	}

	md5 := resp.Header.Get(HeaderMD5Key)
	body, err := io.ReadAll(resp.Body)
	return body, md5, err
}

// http://serverIp:8848/nacos/v1/cs/configs/listener
//
// POST 请求体数据内容：
//
// Listening-Configs=dataId%02group%02contentMD5%02tenant%01
// dataId^2Group^2contentMD5^2tenant^1

func (c *Client) WatchConfig(group, dataID, md5 string) error {
LOOP:
	args := []string{"Listening-Configs=", dataID, SPLIT, group, SPLIT, md5, SPLIT, c.namespaceId, "%01"}
	req, err := http.NewRequest(http.MethodPost, c.baseURI+"/v1/cs/configs/listener", strings.NewReader(strings.Join(args, "")))
	if err != nil {
		return err
	}
	//req.Body = io.NopCloser())
	req.Header.Set("Long-Pulling-Timeout", c.pollingInterval)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var resp *http.Response
	for {
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return formatError(resp)
		}
		body, _ := io.ReadAll(resp.Body)
		if len(body) != 0 {
			body, md5, err = c.GetConfig(group, dataID)
			e := Change{
				NamespaceId: c.namespaceId,
				Group:       group,
				DataId:      dataID,
				Md5:         md5,
				Body:        body,
			}
			c.CallBack(e, err)
			if err != nil {
				return err
			}
			goto LOOP
		}

	}

	return err
}

func formatError(resp *http.Response) error {
	body, _ := io.ReadAll(resp.Body)
	return fmt.Errorf(`
		{
			"code": %d,
			"desc": %s
		}
		`, resp.StatusCode, string(body))
}
