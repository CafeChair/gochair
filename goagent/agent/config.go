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
	AgentKey  string
	Tags      string
	Zookeeper *ZookeeperConfig
	App       *AppConfig
	Redis     *RedisConfig
	Task      *TaskConfig
	Script    *ScriptConfig
	Log       *LogConfig
}

type ZookeeperConfig struct {
	Addr string
	Port int
}

type AppConfig struct {
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
		// ColorLog("[ERRO] 请使用参数-c 加上配置文件名称")
		log.Fatalln("请使用参数-c 加上配置文件名称")
	}
	ConfigFile = cfg
	configcontent, err := ToString(cfg)
	if err != nil {
		// ColorLog("[ERRO] 读取配置文件: %v 失败: %s\n", cfg, err)
		log.Fatalln("读取配置文件: ", cfg, "失败: ", err)
	}
	var acfg AgentConfig
	err = json.Unmarshal([]byte(configcontent), &acfg)
	if err != nil {
		// ColorLog("[ERRO] 解析配置文件: %v 失败: %s\n", cfg, err)
		log.Fatalln("解析配置文件: ", cfg, "失败: ", err)
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
