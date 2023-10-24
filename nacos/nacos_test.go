package nacos

import (
	"fmt"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	//client := NewClient("192.168.119.128:8848", "8ff90940-a1b1-449c-84e2-ef57fe9f8272")
	client := NewClient("192.168.119.128:8848",
		"c2a0851f-f466-47ea-a509-e12901757e5a",
		WithPollingTime(10),
		WithOnConfigChang(func(change Change, err error) {
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

func TestService(t *testing.T) {
	cli := NewClient("192.168.119.128:8848", "c2a0851f-f466-47ea-a509-e12901757e5a")

	service := Service{
		GroupName:   "test",
		ServiceName: "wallet-service",
		Ip:          "192.168.1.7",
		Port:        8080,
		Weight:      1.0,
		Enabled:     true,
		Healthy:     true,
		Ephemeral:   true,
	}

	err := cli.RegistryService(service)
	if err != nil {
		t.Fatal(err)
	}

	err = cli.HeartBeat(service, time.Second*4)
}
