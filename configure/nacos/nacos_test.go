package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"testing"
)

func TestReadConfig(t *testing.T) {
	//create clientConfig
	//clientConfig := constant.ClientConfig{
	//	NamespaceId:         "e525eafa-f7d7-4029-83d9-008937f9d468", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
	//	TimeoutMs:           5000,
	//	NotLoadCacheAtStart: true,
	//	LogDir:              "/tmp/nacos/log",
	//	CacheDir:            "/tmp/nacos/cache",
	//	LogLevel:            "debug",
	//}
	//Another way of create clientConfig
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId("c7b54532-63a1-4267-8e12-ee445fef9ecb"), //When namespace is public, fill in the blank string here.
		constant.WithTimeoutMs(5000),
		/*说明：设置启动时是否从本地缓存加载配置：
			true: 启动时不加载缓存，直接从 Nacos 服务器获取配置。
			false: 启动时优先从本地缓存加载配置（如果服务器不可用，这种方式可能更可靠）。
		建议：
			如果你正在开发和调试，使用 true 是合理的，因为它可以避免读取陈旧配置。
			在生产环境中，建议设置为 false，以提高容灾能力。*/
		constant.WithNotLoadCacheAtStart(true),
		//说明：日志文件的存储路径。
		constant.WithLogDir("/tmp/nacos/log"),
		//说明：本地缓存文件的存储路径
		constant.WithCacheDir("/tmp/nacos/cache"),
		//说明：设置日志级别为 debug，适合开发和调试阶段。在生产环境中，建议将日志级别降为 info 或 error，以减少日志量
		constant.WithLogLevel("debug"),
	)
	// At least one ServerConfig
	//serverConfigs := []constant.ServerConfig{
	//	{
	//		IpAddr:      "localhost",
	//		ContextPath: "/nacos",
	//		Port:        8848,
	//		Scheme:      "http",
	//	},
	//}
	//Another way of create serverConfigs
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			"localhost",
			8848,
			constant.WithScheme("grpc"),
			//说明：指定 Nacos 服务的上下文路径 默认情况下，Nacos 的上下文路径是 /nacos。如果你的 Nacos 服务器配置了自定义上下文路径，则需要相应更改。
			//constant.WithContextPath("/nacos"),
		),
	}

	// Another way of create config client for dynamic configuration (recommend)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "global.yaml",
		Group:  "DEFAULT_GROUP",
	})
	if err != nil {
		t.Fatalf("获取配置失败: %v", err)
	}
	t.Log(content)
}

func TestGetYamlFlag(t *testing.T) {
	// 定义 Redis 结构体
	type Redis struct {
		Host     string `json:"host" yaml:"host,listen"`
		DB       int    `json:"db" yaml:"db,listen"`
		User     string `json:"user" yaml:"user"`
		Password string `json:"password" yaml:"password"`
	}
	redis := Redis{
		Host:     "127.0.0.1",
		DB:       0,
		User:     "admin",
		Password: "secret",
	}

	// 获取 Redis 结构体的 yaml 标签
	yamlTags := getYamlTags(redis)

	// 打印 yaml 标签

	for fieldName, yamlTag := range yamlTags {
		t.Logf("Field: %s, YAML Tag: %s\n", fieldName, yamlTag)
	}
}
