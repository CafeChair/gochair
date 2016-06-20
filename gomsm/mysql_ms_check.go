package main

import (
	"database/sql"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
)

type MasterInstance struct {
	Role  string
	Addr  string
	Port  int
	Slave *SlaveInstance
}

type SlaveInstance struct {
	Role                string
	Addr                string
	Port                int
	MasterAddr          string
	MasterPort          int
	SlaveIORunning      bool
	SlaveSQLRunning     bool
	SecondsBehindMaster int
}

type Instance struct {
	Username string
	Password string
	Hostname string
	Port     int
}

type InstanceModel struct {
	Role  string
	Addr  string
	Alive bool
}

func GetQueryResult(db *sql.DB, query string) (result map[string]string, err error) {
	result = make(map[string]string)
	var rows *sql.Rows
	glog.V(2).Info("Start query", query)
	if rows, err = db.Query(query); err == nil {
		defer rows.Close()
		glog.V(2).Info("End query", query)
		columns, err := rows.Columns()
		if err != nil {
			glog.V(2).Info("mysql query return error with:", err)
			return result, err
		}
		values := make([]sql.RawBytes, len(columns))
		args := make([]interface{}, len(values))
		for i := range values {
			args[i] = &values[i]
		}

		glog.V(2).Info("Start fetch")
		for rows.Next() {
			if err = rows.Scan(args...); err != nil {
				glog.V(2).Info("Mysql rows scan return error with:", err)
				return result, err
			}
		}

		glog.V(2).Infoln("End fetch")
		var value string
		for i, col := range values {
			if col == nil {
				value = "Null"
			} else {
				value = string(col)
			}
			result[columns[i]] = value
		}
		glog.V(2).Info("Mysql query result:", result)
	} else {
		glog.V(2).Info("Mysql query return error:", err)
	}
	glog.V(2).Info("End query", query)
	return result, err
}

func GetInstanceRole(host, user, pass string) (InstanceModel, error) {
	ins := new(InstanceModel)

	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+")/")
	defer db.Close()
	if err != nil {
		glog.V(1).Info("Connect mysql return error with : ", err)
		return *ins, err
	}

	if result, err := GetQueryResult(db, "show slave status;"); err == nil {
		if value, ok := result["Master_Host"]; ok {
			if value == "Null" {
				ins.Role = "Slave"
				ins.Addr = host
				ins.Alive = true
			} else {
				ins.Role = "Master"
				ins.Addr = host
				ins.Alive = true
			}
		}
	}
	return *ins, nil
}

func main() {
	user := "dbmanager"
	pass := "eyDqNafZ6pq39Rah"
	host := "10.50.17.35:3308"
	instancemodel, err := GetInstanceRole(host, user, pass)
}
