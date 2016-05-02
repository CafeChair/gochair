package agent

import (
	"github.com/samuel/go-zookeeper/zk"
	"strconv"
	"time"
)

func Connection() *zk.Conn {
	conn, _, err := zk.Connect([]string{Config().Zookeeper.Addr + ":" + strconv.Itoa(Config().Zookeeper.Port)}, time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	return conn
}
