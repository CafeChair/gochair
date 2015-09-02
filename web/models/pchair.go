package models
import (
	"github.com/astaxie/beego"
	"github.com/coreos/go-etcd/etcd"
)

type Projects struct {
    ProjectName string
}

func AddProject(projectname string) error {
	machines := []string{"http://127.0.0.1:2379"}
    client := etcd.NewClient(machines)
    _,err := client.SetDir(projectname,0)
    if err != nil {
    	return err
    }
    return nil
}

func GetAllProject(key string) ([]Projects,error) {
    machines := []string{"http://127.0.0.1:2379"}
    client := etcd.NewClient(machines)
    res, err := client.Get(key,true,true)
    if err != nil {
        return nil,err
    }
    var a []string
    projectnames := a[:] 
    for _,v := range res.Node.Nodes {
    projectnames = append(projectnames, v.Key)
    }
    return projectnames ,nil
}
