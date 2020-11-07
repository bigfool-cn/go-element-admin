package database

import (
	"bytes"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	config "element-admin-api/configs"
	"log"
	"strconv"
)

var Eloquent *gorm.DB

func init() {

	dbType := config.DatabaseConfig.Dbtype
	host := config.DatabaseConfig.Host
	port := config.DatabaseConfig.Port
	database := config.DatabaseConfig.Database
	username := config.DatabaseConfig.Username
	password := config.DatabaseConfig.Password

	if dbType != "mysql" {
		fmt.Println("db type unknow")
	}
	var err error

	var conn bytes.Buffer
	conn.WriteString(username)
	conn.WriteString(":")
	conn.WriteString(password)
	conn.WriteString("@tcp(")
	conn.WriteString(host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(database)
	//conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")

	log.Println(conn.String())

	var db Database
	if dbType == "mysql" {
		db = new(Mysql)
	} else {
		panic("db type unknow")
	}

	Eloquent, err = db.Open(dbType, conn.String())
	Eloquent.LogMode(true)

	if err != nil {
		log.Fatalln("mysql connect error %v", err)
	} else {
		log.Println("mysql connect success!")
	}

	if Eloquent.Error != nil {
		log.Fatalln("database error %v", Eloquent.Error)
	}

}

type Database interface {
	Open(dbType string, conn string) (db *gorm.DB, err error)
}

type Mysql struct {
}

func (*Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open(dbType, conn)
	return eloquent, err
}

