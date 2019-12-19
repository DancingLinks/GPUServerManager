package main

import (
	"github.com/GPUServerManager/connect"
	"github.com/GPUServerManager/log"
	"github.com/GPUServerManager/socket"
	"github.com/GPUServerManager/utils"
)

func main() {
	utils.InitConstants()
	log.InitLog()
	log.Log("Starting GPU Mangager...", log.INFO)
	ServerSlice, _ := connect.ParseServerConfig()
	ch := make(chan string)
	defer close(ch)
	for i := range ServerSlice.Servers {
		go connect.Work(ServerSlice.Servers[i], ch)
	}
	connect.GPUMap = map[string]connect.GPUList{}
	go socket.StartService()
	for status := range ch {
		connect.UpdateGPUStatus(status)
	}
}
