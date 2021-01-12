package database

import (
  "bytes"
  _ "github.com/go-sql-driver/mysql" //加载mysql
  "github.com/jinzhu/gorm"
  "go-element-admin/configs"
  "go-element-admin/utils"
  "strconv"
)

var Eloquent *gorm.DB

var lgr = utils.DefaultLogger(false)

func init() {

	dbType := configs.DatabaseConfig.Dbtype
	host := configs.DatabaseConfig.Host
	port := configs.DatabaseConfig.Port
	database := configs.DatabaseConfig.Database
	username := configs.DatabaseConfig.Username
	password := configs.DatabaseConfig.Password

	if dbType != "mysql" {
    lgr.Error("db type unknow")
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
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")

  lgr.Println(conn.String())

	var db Database
	if dbType == "mysql" {
		db = new(Mysql)
	} else {
		panic("db type unknow")
	}

	Eloquent, err = db.Open(dbType, conn.String())

	Eloquent.LogMode(configs.ApplicationConfig.Debug)

	if err != nil {
    lgr.Errorf("mysql connect error %v", err)
	} else {
    lgr.Println("mysql connect success!")
	}

	if Eloquent.Error != nil {
    lgr.Errorf("database error %v", Eloquent.Error)
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

