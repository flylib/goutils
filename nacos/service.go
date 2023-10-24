package nacos

import (
	"encoding/json"
	"fmt"
	"github.com/flylib/goutils/structs"
	"io"
	"net/http"
	url2 "net/url"
	"time"
)

// ip	字符串	是	服务实例IP
// port	int	是	服务实例port
// namespaceId	字符串	否	命名空间ID
// weight	double	否	权重
// enabled	boolean	否	是否上线
// healthy	boolean	否	是否健康
// metadata	字符串	否	扩展信息
// clusterName	字符串	否	集群名
// serviceName	字符串	是	服务名
// groupName	字符串	否	分组名
// ephemeral	boolean	否	是否临时实例
type Service struct {
	NamespaceId string  `json:"namespaceId" structs:"namespaceId"`
	GroupName   string  `json:"groupName,omitempty" structs:"groupName,omitempty"`
	Ip          string  `json:"ip" structs:"ip"`
	Port        int     `json:"port" structs:"port"`
	Weight      float64 `json:"weight" structs:"weight"`
	Enabled     bool    `json:"enabled" structs:"enabled"`
	Healthy     bool    `json:"healthy" structs:"healthy"`
	Metadata    string  `json:"metadata,omitempty" structs:"metadata,omitempty"`
	ClusterName string  `json:"clusterName,omitempty" structs:"clusterName,omitempty"`
	ServiceName string  `json:"serviceName" structs:"serviceName"`
	Ephemeral   bool    `json:"ephemeral" structs:"ephemeral"`
}

// curl -X POST 'http://127.0.0.1:8848/nacos/v1/ns/instance?port=8848&healthy=true&ip=11.11.11.11&weight=1.0&serviceName=nacos.test.3&encoding=GBK&namespaceId=n1'
func (c *Client) RegistryService(service Service) error {
	service.NamespaceId = c.namespaceId
	args := structs.Map(service)
	url := c.baseURI + "/v1/ns/instance?" + UrlEncode(args)
	fmt.Println(url)
	resp, err := http.Post(url, "text/plain", nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return formatError(resp)
	}
	//md5 := resp.Header.Get(HeaderMD5Key)
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	return err
}

// curl -X PUT '127.0.0.1:8848/nacos/v1/ns/instance/beat?serviceName=nacos.test.2&beat=%7b%22cluster%22%3a%22c1%22%2c%22ip%22%3a%22127.0.0.1%22%2c%22metadata%22%3a%7b%7d%2c%22port%22%3a8080%2c%22scheduled%22%3atrue%2c%22serviceName%22%3a%22jinhan0Fx4s.173TL.net%22%2c%22weight%22%3a1%7d'
func (c *Client) HeartBeat(service Service, duration time.Duration) error {
	service.NamespaceId = c.namespaceId
	marshal, _ := json.Marshal(service)
	args := structs.Map(service)
	args["beat"] = url2.QueryEscape(string(marshal))
	url := c.baseURI + "/v1/ns/instance/beat?" + UrlEncode(args)
	fmt.Println(url)
	request, _ := http.NewRequest(http.MethodPut, url, nil)
	//request.Header.Add("Content-Type", "text/plain")

	//resp, err := http.Post(url, "text/plain", nil)
	ticker := time.NewTicker(duration)

	for range ticker.C {
		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return formatError(resp)
		}
	}

	return nil
}
