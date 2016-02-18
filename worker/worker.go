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
	"gopkg.in/redis.v3"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	// "time"
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

//--------------------------define worker config----------------------------
//
//--------------------------define worker agent-----------------------------
type Worker struct {
	workerCfg   *WorkerConfig
	etcdAddr    string
	redisClient *redis.Client
}

func NewWorker(cfg *WorkerConfig) (*Worker, error) {
	var err error
	w := new(Worker)
	w.workerCfg = cfg
}

func (w *Worker) Run() error {

}

func (w *Worker) Close() {

}

func (w *Worker) RegisterWorker() error {

}

func (w *Worker) RunTask() (string, error) {

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
