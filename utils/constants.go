package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

var (
	DEBUG				bool
	REFRESH_INTERVAL	time.Duration
)


type YmlConf struct {
	DEBUG 				bool 	`yaml:"debug"`
	REFRESH_INTERVAL	int 	`yaml:"refresh_interval"`
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
	REFRESH_INTERVAL = time.Duration(conf.REFRESH_INTERVAL)
	return
}
