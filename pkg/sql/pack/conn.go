package pack

import (
	"douyin/pkg/sql/dal"

	"gorm.io/gorm"
)

func GetConn() *gorm.DB {
	return dal.ConnectDB()
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
