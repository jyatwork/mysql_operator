/**
 * database operator
 * Created by pigsatwork 2017.12.10
 */
package operator

import (
	"conf_yaml/conf"
	"mysql_operator/db/adapter"

	"github.com/alecthomas/log4go"
	_ "github.com/go-sql-driver/mysql"
)

//declear an object grih Database Connection
var dbConn *adapter.Mysql

func InitDB(host string, user string, psw string, dbName string, MaxOpenConns int, MaxIdleConns int) error {
	dbConn = new(adapter.Mysql)
	dbConn.InitMYSQL(host, user, psw, dbName, MaxOpenConns, MaxIdleConns)
	//断开重连
	if err := dbConn.Db.Ping(); err != nil {
		//noinspection ALL
		log4go.Error("Database: SQL connection failed %s", err.Error())
		return err
	} else {
		log4go.Info("Database: " + conf.String("mysql.db.database") + " Connection Established! ")
		return nil
	}
}

//Insert to grih databases
func DBIns(sql string, args ...interface{}) error {
	stmtIns, err := dbConn.Db.Prepare(sql)
	if err != nil {
		//noinspection ALL
		log4go.Error("Database: DBIns() Prepare %s", err.Error())
		return err
	}
	defer stmtIns.Close()

	if _, err := stmtIns.Exec(args...); err != nil {
		//noinspection ALL
		log4go.Error("Database: DBIns() Exec: %s", err.Error())
		return err
	}

	return nil
}
