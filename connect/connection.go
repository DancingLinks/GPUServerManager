package connect

import (
	"github.com/GPUServerManager/log"
	"golang.org/x/crypto/ssh"
	"net"
	"strings"
	"time"
)

type Connection struct {
	conf *ssh.ClientConfig
	session *ssh.Session
}

func Work(config ServerConfig, ch chan string) {
	addr := config.Host+":"+config.Port
	sshClientConf := &ssh.ClientConfig{
		User: config.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Pwd),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	for {
		var session *ssh.Session
		for session == nil {
			session = GetSession(addr, sshClientConf)
		}
		response := Run(session, "nvidia-smi dmon -c 1 -s pum")
		statusList := strings.Split(response, "\n")
		statusList = statusList[2:len(statusList)-1]
		status := config.ID+"|"
		for i := 0; i < config.GpuCount; i ++ {
			if i < len(statusList) {
				p := strings.Fields(statusList[i])
				// gpu pwr gtemp mtemp sm mem enc dec fb bar1
				status += "1,"+p[1]+","+p[2]+","+p[4]+","+p[8]
			} else {
				status += "0"
			}
			status += "|"
		}
		log.InfoLog(status)
		ch <- status
		time.Sleep(time.Second * 10)
	}
}
