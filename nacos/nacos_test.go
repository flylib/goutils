package nacos

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	//client := NewClient("192.168.119.128:8848", "8ff90940-a1b1-449c-84e2-ef57fe9f8272")
	client := NewClient("192.168.119.128:8848",
		"c2a0851f-f466-47ea-a509-e12901757e5a",
		WithChangeCallBack(func(change Change, err error) {
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println("--- change ---", change.DataId)
			fmt.Println(string(change.Body))
		}))

	body, md5, err := client.GetConfig("test", "test")
	fmt.Println("get data ", string(body), "\n md5", md5)
	err = client.WatchConfig("test", "test", md5)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(body))
}
