package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func main() {
	// 创建 Consul 客户端配置
	config := api.DefaultConfig() // 默认连接到 127.0.0.1:8500
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	// 获取 KV 操作的 Handler
	kv := client.KV()
	// 2. 读取该键值对
	pair, data, err := kv.Get("global", nil)
	if err != nil {
		panic(err)
	}
	if pair == nil {
		fmt.Println("Key not found!")
	} else {
		fmt.Printf("KV pair read: %s = %s\n", pair.Key, pair.Value)
	}
	fmt.Println(data)
}
