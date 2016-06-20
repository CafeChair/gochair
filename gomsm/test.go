package main

import (
	"database/sql"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	// "time"
)

type InstanceModel struct {
	Role  string
	Addr  string
	Alive bool
}

func main() {
	// user := "dbmanager"
	// pass := "eyDqNafZ6pq39Rah"
	// // host := "10.6.24.40:3308"
	// host := "10.50.17.35:3308"
	// // master_host := "10.6.24.16"
	// //

	// ins := new(InstanceModel)
	// db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+")/")
	// defer db.Close()
	// if err != nil {
	// 	glog.V(1).Info("Connect Mysql return error with:", err)
	// }
	// result, err := GetQueryResult(db, "show slave status;")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if value, ok := result["Master_Host"]; ok {
	// 	if value == "Null" {
	// 		ins.Role = "Slave"
	// 		ins.Addr = host
	// 		ins.Alive = true
	// 	} else {
	// 		ins.Role = "Master"
	// 		ins.Addr = host
	// 		ins.Alive = true
	// 	}
	// }
	// fmt.Println(ins)
	ins := test()
	glog.V(1).Infoln(ins.Role)
}

func test() InstanceModel {
	ins := new(InstanceModel)
	ins.Addr = "127.0.0.1"
	return *ins
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
