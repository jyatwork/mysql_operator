/**
 * database connector
 * Created by pigsatwork 2017.12.10
 */
package adapter

import (
	"database/sql"
	"fmt"

	"github.com/alecthomas/log4go"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	Host   string
	User   string
	Psw    string
	DbName string

	MaxOpenConns int
	MaxIdleConns int
	Db           *sql.DB
}

func (s *Mysql) InitMYSQL(_host string, _user string, _psw string, _dbName string,
	_MaxOpenConns int, _MaxIdleConns int) {
	//sqlConn = &Mysql{}
	s.Host = _host
	s.User = _user
	s.Psw = _psw
	s.DbName = _dbName
	s.MaxOpenConns = _MaxOpenConns
	s.MaxIdleConns = _MaxIdleConns

	s.Db, _ = s.newMysqlPool()
}

func (s *Mysql) newMysqlPool() (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		s.User, s.Psw, s.Host, s.DbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log4go.Error("mysqlConnectingUnexpected exception be thrown: %s" + err.Error())
		return db, err
	}
	db.SetMaxOpenConns(s.MaxOpenConns)
	db.SetMaxIdleConns(s.MaxIdleConns)

	db.Ping()

	return db, nil
}
