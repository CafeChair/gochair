package main

import (
	"flag"
	"fmt"
	"gochair/goagent/agent"
)

func main() {
	cfg := flag.String("c", "goagent.json", "configfile")
	flag.Parse()
	agent.ParseConfig(*cfg)
	fmt.Println(agent.Config().Zookeeper.Addr)
}
