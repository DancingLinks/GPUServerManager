package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)
const (

)

var (
	DEBUG bool
)


type YmlConf struct {
	DEBUG bool `yaml:"debug"`
}

func InitConstants() {
	confYml, err := ioutil.ReadFile(GetPath("static/conf.yaml"))
	if err != nil {
		panic("Fail while read config.yml: " + err.Error())
	}
	conf := YmlConf{}
	err = yaml.Unmarshal(confYml, &conf)
	if err != nil {
		panic("Fail while parse config.yml: " + err.Error())
	}
	DEBUG = conf.DEBUG
	return
}
