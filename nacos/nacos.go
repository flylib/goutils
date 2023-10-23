package nacos

import "fmt"

type Client struct {
	host            string
	baseURI         string
	namespaceId     string
	contextPath     string
	scheme          string
	pollingInterval string
	CallBack        func(change Change, err error)
}

type Change struct {
	NamespaceId string
	Group       string
	DataId      string
	Version     string
	Md5         string
	Body        []byte
}

func NewClient(host, namespaceId string, options ...Option) *Client {
	cli := &Client{
		scheme:          "http",
		host:            host,
		namespaceId:     namespaceId,
		baseURI:         fmt.Sprintf("http://%s", host),
		contextPath:     "/nacos",
		pollingInterval: "30000",
	}
	for _, f := range options {
		f(cli)
	}
	cli.baseURI = fmt.Sprintf("%s://%s%s", cli.scheme, host, cli.contextPath)
	return cli
}
