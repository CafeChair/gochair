package main

//后续需要加个redis pool，避免过多的socket连接
import (
	"flag"
	"fmt"
	"gochair/agent/x"
	"gopkg.in/redis.v3"
	"log"
	"bytes"
	"time"
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
	result,err := ExecTask(request)
	if err !=nil {
		log.Fatalln("agent run task fail: ", err)
	}
	
}

func ExecTask(taskname string) (string,error) {
	//执行任务命令
	cmd := exec.Command("sh","-c", taskname)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Start()
	cmdDone := make(chan error, 1)
	go func(){
		cmdDone <- cmd.Wait()
	}()

	select{
	case <- time.After(time.Duration(60) * time.Second):
		KillCmd(cmd)
		<- cmdDone
		return "", errors.New("Command timeout")
	case err := <- cmdDone:
		if err != nil {
			log.Println(stderr.String())
		}
		return out.String(), err
	}
}

func KillCmd(cmd *exec.Cmd) {
	if err := cmd.Process.Kill(); err != nil {
		log.Printf("Failed to kill command: %v", err)
	}
}

func LogToFile() bool {
//任务执行日志保存到本地日志文件中	
}

func LogToRedis() bool {
//任务执行日志保存到redis中
}

func main() {
	cfg := flag.String("c", "agent.json", "configfile")
	flag.Parse()
	x.ParseConfig(*cfg)
	Register()
	AgentRun()
}
