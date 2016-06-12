/*
1：/usr/local/agent/script/下的备份脚本备份完数据
2：备份格式是（ipaddress_时间格式_数据类型-.tar.gz）
3：agent把备份信息推送到redis队列中
4：服务端从redis队列中取出并通过rsync拉取到异地备份机
*/

package main

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	"github.com/Unknwon/goconfig"
	"io/ioutil"
	"net"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type BackupInfo struct {
	ServerId        int64 //唯一ID
	Interval        int64 //上报时间间隔
	GameId          int64
	OpId            int64
	ServerIp        string //本机的内网IP地址
	Success         bool   //备份成功与否
	BackupLog       string
	BackupType      string //备份的类型(MySQL or redis)
	LastAllFilename string //上次全备的文件名称
	Filename        string //备份的文件名称
	Instance        int64  //数据库的端口
	RsyncModel      string //rsync模块名称
	FileSize        int64  //备份文件的大小
	CreateTime      int64  //备份文件创建时间
	LastAllTime     int64  //上次备份文件的创建时间
}

func FetchIpaddr() (string, error) {
	/*
		获取本机的内网IP地址
	*/
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

func FetchConfig(filename, section, key string) (string, error) {
	/*
		读取配置文件,并获取value值
	*/
	config, err := goconfig.LoadConfigFile(fileName)
	if err != nil {
		return "Error", err
	}
	value, err := config.GetValue(section, key)
	if err != nil {
		return "Error", err
	}
	return value, nil

}
func ExecuteScript(script string) ([]byte, error) {
	/*
		执行备份脚本
	*/
	command := strings.Split(script, " ")
	cmd := exec.Command(command[0], command[1], command[2])
	stdout, err := cmd.Output()
	if err != nil {
		return stdout, err
	}
	return stdout, nil
}

func main() {

}
