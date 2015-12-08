package main

import (
	"flag"
	// "fmt"
	"gochair/agent/x"
	"gopkg.in/redis.v3"
	"log"
	"strconv"
)

func Register() {
	redisAddr := x.Config().Redis.Addr + ":" + strconv.Itoa(x.Config().Redis.Port)
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	saddCmd := redisClient.SAdd(x.Config().Tags, x.Config().Uuid)
	err := saddCmd.Err()
	if err != nil {
		log.Fatalln("register agent fail: ", err)
	}
}

func main() {
	cfg := flag.String("c", "agent.json", "configfile")
	flag.Parse()
	x.ParseConfig(*cfg)
	Register()
}
