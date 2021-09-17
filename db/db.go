package db

import (
	"database/sql"
	"fmt"
	"github.com/sugiantodenny01/apibook/config"
	_"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Initial()  {
	conf:=config.GetConfig()
	connectionString := conf.DB_USERNAME +":"+ conf.DB_PASSWORD +"@tcp("+ conf.DB_HOST +":"+ conf.DB_PORT + ")/" + conf.DB_NAME+"?charset=utf8&parseTime=True"
	fmt.Println(connectionString)
	db,err=sql.Open("mysql", connectionString)

	if err != nil {
		panic("connectionString Error")
	}

	err = db.Ping()
	if err != nil {
		panic("DSB Invalid")
	}

}

func CreateConn() *sql.DB  {
	return db
}
