package main

//后续需要加个redis pool，避免过多的socket连接
import (
	"flag"
	"fmt"
	"github.com/Sirupsen/logrus"
	"gochair/agent/x"
	"gopkg.in/redis.v3"
	"strconv"
	"time"
)

func Register() {
	//向redis中注册项目组中的uuid
	redisAddr := x.Config().Redis.Addr + ":" + strconv.Itoa(x.Config().Redis.Port)
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	saddCmd := redisClient.SAdd(x.Config().Tags, x.Config().Uuid)
	err := saddCmd.Err()
	if err != nil {
		// log.Fatalln("register agent fail: ", err)
		logrus.WithFields(logrus.Fields{"register": "fail"}).Info(err.Error())
	}
}

func AgentRun() {
	//agent获取任务并执行
	redisAddr := x.Config().Redis.Addr + ":" + strconv.Itoa(x.Config().Redis.Port)
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	for {
		request, err := redisClient.HGet(x.Config().Uuid, "taskname").Result()
		if err == redis.Nil {
			time.Sleep(time.Second)
			continue
		}
		if err != nil {
			// log.Fatalln("agent get task fail: ", err)
			logrus.WithFields(logrus.Fields{"get task": "fail"}).Info(err.Error())
			continue
		}
		_, err = redisClient.Del(x.Config().Uuid).Result()
		if err != nil {
			// log.Fatalln("agent delete result fail: ", err)
			logrus.WithFields(logrus.Fields{"delete result": "fail"}).Info(err.Error())
		}
		result, err := x.ExecTask(request)
		if err != nil {
			// log.Fatalln("agent run task fail: ", err)
			logrus.WithFields(logrus.Fields{"run task": "fail"}).Info(err.Error())
		}
		fmt.Println(result)
	}
}

// func LogToFile() bool {
// 	//agent实时日志保存到本地日志文件中

// }

// func LogToRedis() bool {
// 	//任务执行日志保存到redis中
// }

func main() {
	cfg := flag.String("c", "agent.json", "configfile")
	flag.Parse()
	x.ParseConfig(*cfg)
	Register()
	AgentRun()
}
