package main

import (
	"github.com/hashicorp/consul/api"
	"testing"
)

func TestConsulReadKV(t *testing.T) {
	// 创建 Consul 客户端配置
	config := api.DefaultConfig() // 默认连接到 127.0.0.1:8500
	client, err := api.NewClient(config)
	if err != nil {
		t.Fatal(err)
	}

	// 获取 KV 操作的 Handler
	kv := client.KV()
	// 2. 读取该键值对
	pair, data, err := kv.Get("global", nil)
	if err != nil {
		t.Fatal(err)
	}
	if pair == nil {
		t.Log("Key not found!")
	} else {
		t.Logf("KV pair read: %s = %s\n", pair.Key, pair.Value)
	}
	t.Log(data.CacheHit)
}
