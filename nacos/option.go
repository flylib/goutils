package nacos

import "fmt"

type Option func(client *Client)

// Default=/nacos,the nacos server contextpath,this is not required in 2.0
func WithContextPath(path string) Option {
	return func(client *Client) {
		client.contextPath = path
	}
}

// the nacos server scheme,default=http,this is not required in 2.0
func WithScheme(path string) Option {
	return func(client *Client) {
		client.scheme = path
	}
}

// Default 30s,the polling time of watch config change
func WithPollingTime(second int) Option {
	return func(client *Client) {
		client.pollingInterval = fmt.Sprintf("%d", second*1000)
	}
}

// Default 30s,the polling time of watch config change
func WithOnConfigChang(f func(change Change, err error)) Option {
	return func(client *Client) {
		client.CallBack = f
	}
}
