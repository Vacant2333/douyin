package dal

import (
	"fmt"
	"sync"

	constantx "douyin/pkg/constant"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func init() {
	once.Do(func() {
		DB = ConnectDB()
		_ = DB.AutoMigrate()
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
	return conn
}
