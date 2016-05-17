/*
1：/usr/local/agent/script/下的备份脚本备份完数据

2：备份格式是（ipaddress_时间格式_数据类型-.tar.gz）

3：agent把备份信息推送到redis队列中

4：服务端从redis队列中取出并通过rsync拉取到异地备份机

*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Unknwon/goconfig"
	"io/ioutil"
	"net"
	"os/exec"
	"path/filepath"
	"time"
)

type BackupInfo struct {
	ServerId        int32
	GameId          int32
	OpId            int32
	ServerIp        string
	Success         bool
	BackupLog       string
	BackupType      string
	LastAllFilename string
	Filename        string
	Instance        int32
	RsyncModel      string
	FileSize        int32
	CreateTime      *time.Time
	LastAllTime     *time.Time
}

func FetchIpaddr() (string, error) {
	var ipaddr string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ipaddr, err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipaddr = ipnet.IP.To4().String()
			}
		}
	}
	return ipaddr, nil
}

func FetchAllScript(scirptPath string) ([]string, error) {
	/*
		获取到备份目录下的所有备份脚本文件，返回文件list
	*/
	scriptList := make([]string, 0)
	scripts, err := ioutil.ReadDir(scirptPath)
	if err != nil {
		return scriptList, err
	}
	for _, script := range scripts {
		if script.IsDir() {
			continue
		}
		scriptName := filepath.Join(scirptPath, script.Name())
		scriptList = append(scriptList, scriptName)
	}
	return scriptList, nil
}

func ExecuteScript(scriptName string) (*BackupInfo, error) {
	/*
		执行备份脚本
	*/
	cmd := exec.Command("sh", "-c", scriptName)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Start()
}

func main() {

}
