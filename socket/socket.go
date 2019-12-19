package socket

import (
	"github.com/GPUServerManager/connect"
	"github.com/GPUServerManager/log"
	"golang.org/x/net/websocket"
	"net/http"
)

func Handle(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		//websocket接受信息
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			log.ErrorLog("Receive failed:" + err.Error())
			break
		}
		msg := ParseMessage(reply)
		if msg.Op == OP_QUERY {
			//这里是发送消息
			msgReply := SenderMsgReply{Data: GetJsonString(connect.GPUMap), Status: 0}
			log.InfoLog(GetJsonString(msgReply))
			if err = websocket.Message.Send(ws, GetJsonString(msgReply)); err != nil {
				log.ErrorLog("Send failed: " + err.Error())
				break
			}
		}
	}
}

func StartService() {
	http.Handle("/websocket", websocket.Handler(Handle))
	if errWs := http.ListenAndServe(":88", nil); errWs != nil {
		log.ErrorLog("Start socket service failed: " + errWs.Error())
	}
}