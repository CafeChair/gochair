package comm

import (
	"log"
	"time"
	"runtime"
)

const (
	VERSION = "0.1"
	COLLECT_INTERVAL = time.Second
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}