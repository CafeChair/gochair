package main

//后续需要加个redis pool，避免过多的socket连接
import (
	"flag"
	"fmt"
	"gochair/agent/x"
	"gopkg.in/redis.v3"
	"log"
	"os/exec"
	"strconv"
)

func Register() {
	//向redis中注册项目组中的uuid
	redisAddr := x.Config().Redis.Addr + ":" + strconv.Itoa(x.Config().Redis.Port)
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	saddCmd := redisClient.SAdd(x.Config().Tags, x.Config().Uuid)
	err := saddCmd.Err()
	if err != nil {
		log.Fatalln("register agent fail: ", err)
	}
}

func AgentRun() {
	//agent获取任务并执行
	redisAddr := x.Config().Redis.Addr + ":" + strconv.Itoa(x.Config().Redis.Port)
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	request, err := redisClient.HGet(x.Config().Uuid, "taskname").Result()
	if err != nil {
		log.Fatalln("agent get task fail: ", err)
	}
	fmt.Println(request)
}

func ExecTask(taskname string) (string, error) {
	//执行agent获取到的任务并输出日志（数据流形式）到日志文件和缓存到redis中（结果形式）
}

func main() {
	cfg := flag.String("c", "agent.json", "configfile")
	flag.Parse()
	x.ParseConfig(*cfg)
	Register()
	AgentRun()
}
