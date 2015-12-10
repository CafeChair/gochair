package x

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var (
	ConfigFile string
	aconfig    *AgentConfig
	lock       = new(sync.RWMutex)
)

type AgentConfig struct {
	Uuid  string
	Tags  string
	Redis *RedisConfig
	Task  *TaskConfig
	Log   *LogConfig
}

type RedisConfig struct {
	Addr string
	Port int
}

type TaskConfig struct {
	Path    string
	TimeOut int
}

type LogConfig struct {
	Path string
	File string
}

func ToString(filename string) (string, error) {
	str, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(str)), nil
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configfile")
	}
	ConfigFile = cfg
	configcontent, err := ToString(cfg)
	if err != nil {
		log.Fatalln("read config file: ", cfg, "fail: ", err)
	}
	var acfg AgentConfig
	err = json.Unmarshal([]byte(configcontent), &acfg)
	if err != nil {
		log.Fatalln("parse config file: ", cfg, "fail: ", err)
	}
	lock.Lock()
	defer lock.Unlock()
	aconfig = &acfg
}

func Config() *AgentConfig {
	lock.RLock()
	defer lock.RUnlock()
	return aconfig
}

// func WriteLog(filename, logstring string) {

// }
func ExecTask(taskname string) (string, error) {
	// cmd := exec.Command("sh", "-c", taskname)
	cmd := exec.Command(taskname)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	cmd.Start()
	err, isTimeout := execTaskTimeOut(cmd)
	if isTimeout {
		if err == nil {
			return "run task timeout and kill process", err
		}
		if err != nil {
			return "kill process fail", err
		}
	}
	if err != nil {
		return "run task fail", err
	}
	return stdout.String(), nil
}

func execTaskTimeOut(cmd *exec.Cmd) (error, bool) {
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(time.Duration(10) * time.Second):
		go func() {
			<-done
		}()
		if err = cmd.Process.Kill(); err != nil {
			log.Printf("Failed to kill command: %v", err)
		}
		return err, true
	case err = <-done:
		return err, false
	}
}
