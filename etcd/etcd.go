package etcd
import (
	"github.com/coreos/go-etcd/etcd"
)

func Get(key string) (string,error) {
	etcdserver := Config().Etcd.Addr
	machine := []string{"http://"+etcdserver}
	client := etcd.NewClient(machine)
	res,err := client.Get(key,true,true)
	if err != nil {
    	return "",err
    }
    return res.Node.Value,nil
}

func Set(key string) (bool, error) {
	etcdserver := Config().Etcd.Addr
	machine := []string{"http://"+etcdserver}
	client := etcd.NewClient(machine)
	_,err := client.Set(key,"",60)
	if err != nil {
		return false,err
	}
	return true,nil
}

