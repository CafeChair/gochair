package agent

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

var (
	ConfigFile string
	aconfig    *AgentConfig
	lock       = new(sync.RWMutex)
)

type AgentConfig struct {
	Uuid  string
	Tags  string
	Zook  *ZookeeperConfig
	Redis *RedisConfig
	Task  *TaskConfig
	Log   *LogConfig
}

type ZookeeperConfig struct {
	Addr string
	Port int
}

type RedisConfig struct {
	Addr string
	Port int
}

type TaskConfig struct {
	Addr string
}

type ScriptConfig struct {
	Addr string
}

type LogConfig struct {
	Addr string
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
