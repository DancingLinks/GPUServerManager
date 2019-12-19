package socket

import (
	"encoding/json"
	"github.com/GPUServerManager/connect"
	"github.com/GPUServerManager/log"
)

const (
	OP_QUERY int = 1
)

type Message struct {
	Op       int
}

func ParseMessage(jsonString string) Message {
	var msg Message
	if err := json.Unmarshal([]byte(jsonString), &msg); err != nil {
		log.ErrorLog("Parse json err: " + jsonString)
	}
	return msg
}

type SenderMsgReply struct {
	Data 	 string
	Status 	 int
}

func GetJsonString(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}

func FormatReplyStatus(GPUMap map[string]connect.GPUList) SenderMsgReply {
	GPUList := []connect.GPUList{}
	for _, v := range GPUMap {
		GPUList = append(GPUList, v)
	}
	return SenderMsgReply{Data: GetJsonString(GPUList), Status: 0}
}
