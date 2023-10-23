package configure

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

var (
	cfg *Config
)

// run env
const (
	Local = "local"
	Dev   = "dev"
	Test  = "test"
	Prod  = "prod"
)

// file type
const (
	Yaml = "yaml"
	Json = "json"
	Xml  = "xml"
	conf = "conf"
)

type Config struct {
	configChangeCallBack func(origin any, err error)
	env                  string
	path                 string

	configParser *viper.Viper
	origin       reflect.Type
}

func (c *Config) Env() string {
	return c.env
}
func (c *Config) Scan(v any) error {
	err := c.configParser.Unmarshal(v)
	if err != nil {
		return err
	}
	c.origin = reflect.TypeOf(v)
	return nil
}

func Load(file string, options ...Option) *Config {
	cfg = &Config{
		configParser: viper.New(),
		path:         "./",
	}

	for _, f := range options {
		f(cfg)
	}

	splits := strings.Split(file, ".")

	if len(splits) < 2 {
		panic("Wrong file name")
	}
	fileType := strings.ToLower(splits[len(splits)-1])
	cfg.configParser.AutomaticEnv()
	cfg.configParser.SetConfigType(fileType)

	if strings.Contains(cfg.env, "$") {
		environment := strings.Trim(cfg.env, "$")
		cfg.env = cfg.configParser.GetString(environment)
	}

	if cfg.env != "" {
		splits = splits[:len(splits)-1]
		splits = append(splits, cfg.env, fileType)
		file = strings.Join(splits, ".")
	}

	//设置读取文件名
	cfg.configParser.SetConfigName(file)
	//如果没有指定配置文件，则解析默认的配置文件
	cfg.configParser.AddConfigPath(cfg.path)

	if err := cfg.configParser.ReadInConfig(); err != nil {
		panic(err)
	}

	if cfg.configChangeCallBack != nil {
		cfg.configParser.OnConfigChange(func(in fsnotify.Event) {
			v := reflect.New(cfg.origin).Interface()
			err := cfg.configParser.Unmarshal(v)
			cfg.configChangeCallBack(reflect.ValueOf(v).Elem().Interface(), err)
		})
		cfg.configParser.WatchConfig()
	}

	return cfg
}
