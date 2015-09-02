package main
import (
	"flag"
	"fmt"
	"os"
	"log"
	"x/etcd"
	"x/comm"
)


func main(){
	cfg := flag.String("c", "cfg.json","config file")
	version := flag.Bool("v",false,"show version")
	flag.Parse()

	if *version {
		fmt.Println(comm.VERSION)
		os.Exit(0)
	}

	etcd.ParseConfig(*cfg)
	roletag,err := etcd.RoleTag()
	if err != nil {
		log.Println(err)
	}
	serverid,err := etcd.Serverid()
	if err !=nil {
		log.Println(err)
	}
	key := roletag + "/command/" + serverid

	res,err := etcd.Get(key)
	if err != nil {
		log.Println(err)
	}
	str,err := etcd.RunCmd(res)
	if err != nil {
		log.Println(err)
	}
	ok := etcd.LogTo(key,str)
	if !ok {
		log.Println("Write log faild")
	}
}