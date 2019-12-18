package connect

import (
	"encoding/json"
	"github.com/GPUServerManager/log"
	"github.com/GPUServerManager/utils"
	"io/ioutil"
	"os"
)

type Server struct {
	Servers []ServerConfig
}

type ServerConfig struct {
	ID 			string
	Host 		string
	Port 		string
	User 		string
	Pwd 		string
	GpuCount 	int
}

func readServerConfig() (string, error) {
	fi, err := os.Open(utils.GetRoot()+"/static/server.json")
	if err != nil {
		log.ErrorLog("Fail while read server config: " + err.Error())
		return "", err
	}
	defer fi.Close()
	fd, _ := ioutil.ReadAll(fi)
	return string(fd), nil
}

func ParseServerConfig() (Server, error) {
	var server Server
	config, err := readServerConfig()
	if err != nil {
		return Server{}, err
	}
	log.InfoLog("Read server config successfully: " + config)
	err = json.Unmarshal([]byte(config), &server)
	if err != nil {
		log.ErrorLog("Fail while parse server config: " + err.Error())
		return Server{}, err
	}
	return server, nil
}
