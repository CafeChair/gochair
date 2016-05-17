package agent

import (
	"github.com/samuel/go-zookeeper/zk"
	"strconv"
	"time"
)

type Zook struct {
	conn *zk.Conn
}

func new(server string) (*Zook, error) {
	conn, _, err := zk.Connect(server, time.Second*5)
	if err != nil {
		// ColorLog("[ERRO] 连接zookeeper %v 失败: %s\n", server, err)
		return nil, err
	}
	return &Zook{conn: conn}, nil
}

func (z *Zook) Close() {
	z.conn.Close()
}

func Watch(key string)
