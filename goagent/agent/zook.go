package agent

import (
	"github.com/samuel/go-zookeeper/zk"
	"strconv"
	"sync"
	"time"
)

type ZKConn struct {
	conn *zk.Conn
	lock sync.Mutex
	init bool
}

func (zkc *ZKConn) Connection() *zk.Conn {
	conn, _, err := zk.Connect([]string{Config().Zookeeper.Addr + ":" + strconv.Itoa(Config().Zookeeper.Port)}, time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	return conn
}

func (zkc *ZKConn) initial() {
	zkc.conn = zkc.Connection()
	zkc.lock = sync.Mutex{}
}

var zkconn ZKConn

func GetConn() ZKConn {
	if zkconn.init {
		return zkconn
	} else {
		zkconn = ZKConn{init: true}
		zkconn.initial()
	}
	return zkconn
}
