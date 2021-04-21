package usersdb

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pragmatically-dev/bookstore_users_api/datasources/mysql/dbconfig"
	"github.com/pragmatically-dev/bookstore_users_api/utils/config"
)

var (
	Client *sqlx.DB
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
	Client, err = sqlx.Open("mysql", dbc.GetConnectionURL())
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database connected")
}
