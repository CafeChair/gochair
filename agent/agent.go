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

func AgentRun() {
	redisAddr := x.Config().Redis.Addr + ":" + strconv.Itoa(x.Config().Redis.Port)
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	for {
		//注册project
		existCmd := redisClient.Exists(x.Config().Tags)
		if err := existCmd.Err(); err != nil {
			logrus.WithFields(logrus.Fields{"exist key": "fail"}).Info(err.Error())
			continue
		}
		saddCmd := redisClient.SAdd(x.Config().Tags, x.Config().Uuid)
		err := saddCmd.Err()
		if err != nil {
			// log.Fatalln("register agent fail: ", err)
			logrus.WithFields(logrus.Fields{"register": "fail"}).Info(err.Error())
		}
		//agent获取任务并执行
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

func main() {
	cfg := flag.String("c", "agent.json", "configfile")
	flag.Parse()
	x.ParseConfig(*cfg)
	// Register()
	AgentRun()
}
