package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	// "time"
)

var (
	ConfigFile string
	wconfig    *WorkerConfig
	lock       = new(sync.RWMutex)
)

type EtcdConfig struct {
	Addr    string
	Port    int
	Timeout int
}

type RedisConfig struct {
	Addr    string
	Port    int
	Timeout int
}

type LogConfig struct {
	Path     string
	Filename string
}

type WorkerConfig struct {
	Uuid  string
	Tags  string
	Etcd  *EtcdConfig
	Redis *RedisConfig
	Log   *LogConfig
}

func toString(filename string) (string, error) {
	str, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(str)), nil
}

func ParseWorkerConfigFile(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configfile")
	}
	ConfigFile = cfg
	configcontent, err := toString(cfg)
	if err != nil {
		log.Fatalln("read config file: ", cfg, "fail: ", err)
	}
	var wcfg WorkerConfig
	err = json.Unmarshal([]byte(configcontent), &wcfg)
	if err != nil {
		log.Fatalln("parse config file: ", cfg, "fail: ", err)
	}
	lock.Lock()
	defer lock.Unlock()
	wconfig = &wcfg
}

func Config() *WorkerConfig {
	lock.RLock()
	defer lock.RUnlock()
	return wconfig
}

func main() {
	cfg := flag.String("c", "worker.json", "worker config filename")
	flag.Parse()
	ParseWorkerConfigFile(*cfg)
	fmt.Println(Config().Etcd.Addr)
}
