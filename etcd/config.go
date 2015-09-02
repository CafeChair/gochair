package etcd

import (
	"log"
	"sync"
	"x/comm"
	"encoding/json"
)

type EtcdConfig struct {
	Addr string
	Timeout int
}

type RedisConfig struct {
	Addr string
}

type GlobalConfig struct{
	ServerId string
	Hostname string
	RoleTag string
	Etcd *EtcdConfig
	Redis *RedisConfig
}

var (
	ConfigFile string
	config *GlobalConfig
	lock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func RoleTag() (string,error) {
	roletag := Config().RoleTag
	if roletag != "" {
		return roletag,nil
	}
	return roletag,nil
}

func Serverid() (string, error) {
	serverid := Config().ServerId
	if serverid != "" {
		return serverid,nil
	}
	return serverid,nil
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configure file")
	}
	if !comm.IsExist(cfg) {
		log.Fatalln("config file: ", cfg, "is not exist.")
	}
	ConfigFile = cfg
	configContent,err := comm.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file: ",cfg, "fail: ",err)
	}
	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file: ",cfg, "fail: ",err)
	}
	lock.Lock()
	defer lock.Unlock()
	config = &c
	log.Println("read config file: ",cfg, "successful.")
}