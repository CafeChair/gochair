package main

import (
	"flag"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"gochair/goagent/agent"
	"strconv"
	"time"
)

func main() {
	cfg := flag.String("c", "goagent.json", "configfile")
	flag.Parse()
	agent.ParseConfig(*cfg)

}
