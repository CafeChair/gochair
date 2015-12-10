package main

//后续需要加个redis pool，避免过多的socket连接
import (
	"flag"
	"github.com/Sirupsen/logrus"
	"gochair/agent/x"
	"gopkg.in/redis.v3"
	"strconv"
	"time"
)

func AgentRun() {
	redisAddr := x.Config().Redis.Addr + ":" + strconv.Itoa(x.Config().Redis.Port)
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	logfile := x.Config().Log.Path + "/" + x.Config().Log.File
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
			logrus.WithFields(logrus.Fields{"register": "fail"}).Info(err.Error())
		}
		//agent获取任务并执行
		request, err := redisClient.HGet(x.Config().Uuid, "taskname").Result()
		if err == redis.Nil {
			time.Sleep(time.Second)
			continue
		}
		if err != nil {
			logrus.WithFields(logrus.Fields{"get task": "fail"}).Info(err.Error())
			continue
		}
		_, err = redisClient.Del(x.Config().Uuid).Result()
		if err != nil {
			logrus.WithFields(logrus.Fields{"delete result": "fail"}).Info(err.Error())
		}
		result, err := x.ExecTask(request)
		if err != nil {
			logrus.WithFields(logrus.Fields{"run task": "fail"}).Info(err.Error())
		}
		//记录日志到本地文件和远端redis中
		ok := x.WriteLogToFile(logfile, request, result)
		if ok {
			continue
		}
	}
}

func main() {
	cfg := flag.String("c", "agent.json", "configfile")
	flag.Parse()
	x.ParseConfig(*cfg)
	AgentRun()
}
