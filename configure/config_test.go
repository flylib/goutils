package configure

import (
	"fmt"
	"testing"
	"time"
)

type Cfg struct {
	Service Service `json:"service"yaml:"service"`
	Redis   Redis   `json:"redis"yaml:"redis"`
}

type Service struct {
	Name    string `json:"name" yaml:"name"`
	Version string `json:"version" yaml:"version"`
	Address string `json:"address"yaml:"address"`
}

type Redis struct {
	Host     string `json:"host"yaml:"host"`
	DB       int    `json:"db"yaml:"db"`
	User     string `json:"user"yaml:"user"`
	Password string `json:"password"yaml:"password"`
}

func TestReadYaml(t *testing.T) {
	cfg := Cfg{}
	err := Load("config.yaml",
		WithEnvironment("$env"),
		WithFileChangeCallBack(func(config any, err error) {
			if err != nil {
				panic(err)
			}
			c := config.(*Cfg)
			fmt.Println(*c)
		})).Scan(&cfg)
	if err != nil {
		t.Fatal(cfg)
	}

	fmt.Println(cfg)

	for {
		time.Sleep(time.Second)
		fmt.Println("running")
	}
}
