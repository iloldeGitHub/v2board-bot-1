package service

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Conf struct {
	Bot      BotConf      `yaml:"bot"`
	Database DatabaseConf `yaml:"database"`
}

type BotConf struct {
	Token string `yaml:"token"`
	Name  string `yaml:"name"`
	Byte  int64  `yaml:"byte"`
}
type DatabaseConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (c *Conf) GetConfig() *Conf {
	yamlFile, err := ioutil.ReadFile("uuBot.yaml")
	if err != nil {
		fmt.Printf("打开配置文件错误...\n错误信息:%s", err)
		os.Exit(1)
	}
	if err = yaml.Unmarshal(yamlFile, c); err != nil {
		fmt.Printf("配置文件解析错误... \n错误信息:%s", err)
		os.Exit(1)
	}
	return c
}
