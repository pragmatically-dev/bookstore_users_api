package config

import "os"

var (
	DBNAME     = "users_db"
	DBUSER     = "root"
	DBPASSWORD = "password"
	DBHOST     = "127.0.0.1"
	SSLMODE    = "require"
	DBPORT     = "5432"
	PORT       = ""
)

func Load() {
	//DBPORT = os.Getenv("DBPORT")
	PORT = os.Getenv("PORT")
	//DBNAME = os.Getenv("DBNAME")
	//DBPASSWORD = os.Getenv("DBPASSWORD")
}