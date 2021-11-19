package config

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

// type DBConfig struct {
// 	Host     string
// 	Port     int
// 	User     string
// 	DBName   string
// 	Password string
// }

type DBconfig struct {
	Host string
	port int
	user string
	// Pass   string
	dbname string
}

func buildconfig() *DBconfig {
	Dbconfig := DBconfig{
		Host:   "localhost",
		port:   3306,
		user:   "root",
		dbname: "golang",
		// Pass:   "",
	}
	return &Dbconfig
}

func DBurl(Dbconfig *DBconfig) string {
	return fmt.Sprintf("%s:tcp(%s:%d)/%s?charset=utf8&parsrTime=true&loc=LOCAL",
		Dbconfig.user,
		Dbconfig.Host,
		Dbconfig.port,
		Dbconfig.dbname)
}
