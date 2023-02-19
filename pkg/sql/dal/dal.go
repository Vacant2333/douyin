package dal

import (
	"fmt"
	"sync"

	constantx "douyin/pkg/constant"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBI *gorm.DB
var once sync.Once

func init() {
	once.Do(func() {
		DBI = ConnectDB()
		err := DBI.AutoMigrate()
		if err != nil {
			panic(err)
		}
	})
}

func ConnectDB() (conn *gorm.DB) {
	conn, err := gorm.Open(mysql.Open(constantx.MYSQL_Dsn), &gorm.Config{
		SkipDefaultTransaction: constantx.MYSQL_SkipDefaultTransaction, // close default tx
		PrepareStmt:            constantx.MYSQL_PrepareStmt,            // cache precompile sentence
	})
	if err != nil {
		panic(fmt.Errorf("cannot setup db conn: %v", err))
	}
	if conn == nil {
		panic("conn is null")
	}
	return conn
}
