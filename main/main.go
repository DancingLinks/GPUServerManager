package main

import (
	"fmt"
	"github.com/GPUServerManager/connect"
	"github.com/GPUServerManager/log"
	"github.com/GPUServerManager/utils"
)

var (
	ServerSlice connect.Server
	GPUMap map[string]connect.GPUList
)

func main() {
	utils.InitConstants()
	log.InitLog()
	log.Log("Hello world!", log.INFO)
	ServerSlice, _ := connect.ParseServerConfig()
	ch := make(chan string)
	defer close(ch)
	for i := range ServerSlice.Servers {
		go connect.Work(ServerSlice.Servers[i], ch)
	}
	GPUMap = map[string]connect.GPUList{}
	for status := range ch {
		gpuList := connect.Parse(status)
		GPUMap[gpuList.ID] = gpuList
		fmt.Println(GPUMap)
	}
}

