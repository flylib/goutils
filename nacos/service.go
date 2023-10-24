package nacos

import (
	"io"
	"net/http"
	"net/url"
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
	NamespaceId string  `json:"namespaceId"`
	GroupName   string  `json:"groupName"`
	Ip          string  `json:"ip"`
	Port        int     `json:"port"`
	Weight      float64 `json:"weight"`
	Enabled     bool    `json:"enabled"`
	Healthy     bool    `json:"healthy"`
	Metadata    string  `json:"metadata"`
	ClusterName string  `json:"clusterName"`
	Ephemeral   bool    `json:"ephemeral"`
}

// curl -X GET 'http://127.0.0.1:8848/nacos/v1/cs/configs?dataId=nacos.example&group=com.alibaba.nacos'
// http://192.168.119.128:8848/nacos/v1/cs/configs?dataId=test&group=test&tenant=8ff90940-a1b1-449c-84e2-ef57fe9f8272
func (c *Client) RegistryService(service Service) ([]byte, string, error) {
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
