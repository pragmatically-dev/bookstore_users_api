package usersdb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pragmatically-dev/bookstore_users_api/datasources/mysql/dbconfig"
	"github.com/pragmatically-dev/bookstore_users_api/utils/config"
	"log"
)

var(
	Client *sql.DB
)

func init() {
	dbc := dbconfig.DBConfigurations{
		DBName:   config.DBNAME,
		User:     config.DBUSER,
		Password: config.DBPASSWORD,
		Host:     config.DBHOST,
		Port:     config.DBPORT,
	}
	var err error
	Client,err = sql.Open("mysql",dbc.GetConnectionURL())
	if err !=nil{
		panic(err)
	}
	if err= Client.Ping();err!=nil{
		panic(err)
	}
	log.Println("Database connected")
}