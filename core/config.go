package core

import (
	"fmt"
	"gin_blog/config"
	"gin_blog/global"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const configFile = "settings.yaml"

// ConfInit 读取yaml配置
func ConfInit() {

	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		logrus.Fatalf("config Init Unmarshal: %v", err)
	}
	logrus.Println("config yamlFile load Init success.")
	fmt.Println(c)
}

func SetYaml(conf config.Server) {
	data, err := yaml.Marshal(conf)
	if err != nil {
		global.LOG.Error(err)
		return
	}
	err = ioutil.WriteFile(configFile, data, 0777)
	if err != nil {
		global.LOG.Error(err)
		return
	}
}
