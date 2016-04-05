/*
1：/usr/local/agent/script/下的备份脚本备份完数据库
2：备份格式是（ipaddress-timestamp-数据类型-.tar.gz）
3：agent把备份信息推送到redis队列中
4：服务端从redis队列中取出并通过rsync拉取
*/
package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"io/ioutil"
	"path/filepath"
)

type AgentBackupInfo struct {
	IPaddr     string
	Module     string
	BackupFile string
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

func ExecuteScript(script string) (*AgentBackupInfo, error) {
	/*
		执行备份脚本
	*/
}

func main() {

}
