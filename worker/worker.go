/*
@auth:Jay
@Version:1.0
@Profile:
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"github.com/flike/golog"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"
)

//--------------------------define worker config----------------------------
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

type LogConfig struct {
	Path     string
	Filename string
}

type WorkerConfig struct {
	Uuid string
	Tags string
	Etcd *EtcdConfig
	Log  *LogConfig
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

//--------------------------define worker config----------------------------
//
//--------------------------define worker agent-----------------------------

func pingEtcd(cfg *WorkerConfig) bool {
	etcdAddr := []string{"http://" + cfg.Etcd.Addr + cfg.Etcd.Port}
	etcdClient := etcd.NewClient(etcdAddr)
	if _, err := etcdClient.Set("/workerping", "pong", 0); err != nil {
		golog.Error("RegisterWroker", "ping", "ping etcd fail", 0, "err", err.Error())
		return false
	}
	return true
}
func RegisterWorker(cfg *WorkerConfig) bool {
	if ok := pingEtcd(cfg); ok != false {

	}
}

func FetchJobs(cfg *WorkerConfig) bool {

}

func RunJobs(cfg *WorkerConfig) bool {

}

//--------------------------define worker agent-----------------------------
//
//--------------------------define result storage---------------------------
type TaskResult struct {
	Uuid       string
	TaskName   string
	MaxRunTime int64
	IsSuccess  bool
	Result     string
}

//--------------------------define result storage---------------------------
func main() {
	cfg := flag.String("c", "worker.json", "worker config filename")
	flag.Parse()
	ParseWorkerConfigFile(*cfg)
}
