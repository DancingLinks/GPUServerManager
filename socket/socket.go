package socket

import (
	"github.com/GPUServerManager/connect"
	"github.com/GPUServerManager/log"
	"github.com/GPUServerManager/utils"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"net/http"
	"os"
)

func WebsocketHandle(ws *websocket.Conn) {
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
			msgReply := FormatReplyStatus(connect.GPUMap)
			log.InfoLog(GetJsonString(msgReply))
			if err = websocket.Message.Send(ws, GetJsonString(msgReply)); err != nil {
				log.ErrorLog("Send failed: " + err.Error())
				break
			}
		}
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fd, _ := os.OpenFile(
		utils.GetPath("/web/html/index.html"),
		os.O_RDONLY,0666)

	buf, _ := ioutil.ReadAll(fd)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := w.Write(buf)
	if err != nil {
		log.ErrorLog("Failed while writing response: " + err.Error())
	}
}


func StartService() {
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir(utils.GetPath("/web")))))
	http.Handle("/index", http.HandlerFunc(IndexHandler))
	http.Handle("/websocket", websocket.Handler(WebsocketHandle))
	if errWs := http.ListenAndServe(":88", nil); errWs != nil {
		log.ErrorLog("Start service failed: " + errWs.Error())
	}
}
