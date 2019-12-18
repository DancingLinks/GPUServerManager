package log

import (
	"fmt"
	"github.com/GPUServerManager/utils"
	"os"
	"strings"
	"time"
)

const (
	INFO		int = 0
	WARNING		int = 1
	ERROR		int = 2
)

var LogfilePaths []string
var LogtypeDescs []string


func InitLog() {
	LogfilePaths = []string {
		"static/logs/info_logs.txt",
		"static/logs/warning_logs.txt",
		"static/logs/error_logs.txt"}
	LogtypeDescs = []string {
		"INFO",
		"WARNING",
		"ERROR"}
}

func Log(content string, logType int) {

	logTime := time.Now().Format("2006-01-02 15:04:05")
	logContent := strings.Join([]string{"[",logTime,"]","[", LogtypeDescs[logType],"] ",content,"\n"},"")

	if utils.DEBUG {
		fmt.Println(logContent[:len(logContent)-1])
		return
	}

	fd, _ := os.OpenFile(
		utils.GetPath(LogfilePaths[logType]),
		os.O_RDWR | os.O_APPEND | os. O_CREATE,0666)

	buf:=[]byte(logContent)

	if _, err := fd.Write(buf); err != nil {
		fmt.Println(err)
	}

	defer fd.Close()
}

func InfoLog(content string) {
	Log(content, INFO)
}

func WarningLog(content string) {
	Log(content, WARNING)
}

func ErrorLog(content string) {
	Log(content, ERROR)
}