package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

type nacosConfig struct {
	client       config_client.IConfigClient
	configParser *viper.Viper
	option       *option
}

func (n *nacosConfig) Env() string {
	return "test"
}

func (n *nacosConfig) Scan(conf any) error {
	n.configParser.SetConfigType("yaml") //默认读取yaml格式

	//tags := getTags(conf, "nacos")
	// 获取传入对象的值和类型
	v := reflect.ValueOf(conf)
	//t := reflect.TypeOf(conf)
	fmt.Println(v.Kind())
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("conf must be a pointer")
	}
	v = v.Elem()
	t := v.Type()

	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// 获取 "nacos" 标签
		tag := field.Tag.Get("nacos")
		if tag == "" {
			continue
		}
		tagParts := strings.Split(tag, ",")
		dataId := tagParts[0]

		if len(dataId) > 0 {
			content, err := n.client.GetConfig(vo.ConfigParam{
				DataId: dataId,
				Group:  n.option.group,
			})
			if err != nil {
				return fmt.Errorf("failed to get config for DataId %s: %w", dataId, err)
			}

			// 解析配置并赋值
			if err := n.configParser.ReadConfig(strings.NewReader(content)); err != nil {
				return fmt.Errorf("failed to parse config for DataId %s: %w", dataId, err)
			}

			fieldValue := v.Field(i)
			if !fieldValue.CanSet() {
				return fmt.Errorf("cannot set field %s", field.Name)
			}

			fieldInterface := fieldValue.Addr().Interface()
			if err := n.configParser.Unmarshal(fieldInterface); err != nil {
				return fmt.Errorf("failed to unmarshal field %s: %w", field.Name, err)
			}

			// 配置监听
			if len(tagParts) == 2 && tagParts[1] == "listen" {
				err = n.client.ListenConfig(vo.ConfigParam{
					DataId: dataId,
					Group:  n.option.group,
					OnChange: func(namespace, group, dataId, data string) {
						if err = n.configParser.ReadConfig(strings.NewReader(data)); err != nil {
							n.option.onChange(dataId, err)
						} else {
							err = n.configParser.Unmarshal(fieldInterface)
							n.option.onChange(dataId, err)
						}
					},
				})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewNacosConfig(options ...OptionFunc) (*nacosConfig, error) {
	o := &option{
		host:   "localhost",
		port:   8848,
		scheme: "http",
		group:  "DEFAULT_GROUP",
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
		viper.New(),
		o,
	}

	return client, nil
}

func getTags(obj interface{}, tag string) map[string]string {
	tags := make(map[string]string)

	// 获取传入对象的类型

	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// 获取 "yaml" 标签
		yamlTag := field.Tag.Get(tag)
		if yamlTag != "" {
			tags[field.Name] = yamlTag
		}
	}
	return tags
}

func setField(obj any, fieldName string, value any) error {
	v := reflect.ValueOf(obj)

	// 确保传入的是指针，且指针指向的是结构体
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct")
	}

	// 获取结构体的实际值
	v = v.Elem()

	// 获取字段
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("no such field: %s in obj", fieldName)
	}

	// 确保字段是可设置的
	if !field.CanSet() {
		return fmt.Errorf("cannot set field %s", fieldName)
	}

	// 设置字段值
	val := reflect.ValueOf(value)
	if field.Type() != val.Type() {
		return fmt.Errorf("provided value type doesn't match field type")
	}

	field.Set(val)
	return nil
}
