package dbconfig

import "fmt"

//DBConfigurations contains the crucial information for connect to a mysql database
type DBConfigurations struct {
DBName   string
User     string
Password string
Host     string
Port     string
}

//GetConnectionURL this methods is in charge of make the connection string
func (dbc *DBConfigurations) GetConnectionURL() string {
connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbc.User, dbc.Password, dbc.Host, dbc.DBName)
return connectionString
}
