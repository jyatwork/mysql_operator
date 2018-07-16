package main

import (
	"conf_yaml/conf"
	"fmt"
	"mysql_operator/db/operator"
)

func main() {
	conf.Load("conf.yaml")
	operator.InitDB(conf.String("mysql.db.host"), conf.String("mysql.db.username"), conf.String("mysql.db.password"), conf.String("mysql.db.database"), conf.Int("mysql.db.MaxOpenConns"), conf.Int("mysql.db.MaxIdleConns"))
	sql := `INSERT INTO user ( create_time,uname, psw) VALUES ( UTC_TIMESTAMP(), ?, ?)`
	if err := operator.DBIns(sql, "username", "password"); err != nil {
		fmt.Println(err.Error())
	}
}
