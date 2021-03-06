package connect

import (
	"bytes"
	"github.com/GPUServerManager/log"
	"golang.org/x/crypto/ssh"
	"time"
)

func GetClient(addr string, config *ssh.ClientConfig) *ssh.Client{
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.ErrorLog("Failed to dial: " + err.Error())
		time.Sleep(time.Second)
		return nil
	}
	return client
}

func GetSession(client *ssh.Client) *ssh.Session {
	session, err := client.NewSession()
	if err != nil {
		log.ErrorLog("Failed to create session: " + err.Error())
		time.Sleep(time.Second)
		return nil
	}
	return session
}

func Run(session *ssh.Session, cmd string) string {
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(cmd); err != nil {
		log.ErrorLog("Failed to run '" + cmd + "': " + err.Error())
		return ""
	}
	return b.String()
}
