package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"reflect"
)

type nacosConfig struct {
	client config_client.IConfigClient
}

func NewNacosConfig(options ...OptionFunc) (*nacosConfig, error) {
	o := &option{
		host:   "localhost",
		port:   8848,
		scheme: "http",
	}
	for _, f := range options {
		f(o)
	}
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			o.host, uint64(o.port),
			constant.WithScheme(o.scheme),
			//说明：指定 Nacos 服务的上下文路径 默认情况下，Nacos 的上下文路径是 /nacos。如果你的 Nacos 服务器配置了自定义上下文路径，则需要相应更改。
			//constant.WithContextPath("/nacos"),
		),
	}

	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(o.namespace), //When namespace is public, fill in the blank string here.
		//constant.WithTimeoutMs(5000),
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

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return nil, err
	}
	client := &nacosConfig{
		configClient,
	}

	return client, nil
}

func getYamlTags(obj interface{}) map[string]string {
	tags := make(map[string]string)

	// 获取传入对象的类型
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// 遍历结构体字段
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		// 获取 "yaml" 标签
		yamlTag := field.Tag.Get("yaml")
		if yamlTag != "" {
			tags[field.Name] = yamlTag
		}
	}
	return tags
}
