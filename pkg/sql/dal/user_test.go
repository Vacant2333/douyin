package dal

import (
	"context"
	"douyin/pkg/sql/dal/model"
	"douyin/pkg/sql/dal/query"
	"testing"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB
var q *query.Query
var ctx context.Context
var ud query.IUserDo

func init() {
	db = ConnectDB()
	q = query.Use(db)
	ctx = db.Statement.Context
	ud = q.User.WithContext(ctx)

}

func TestCreate(t *testing.T) {
	var err error

	var ID uint

	{
		userdata := &model.User{
			Username:   "123",
			Password:   "123",
			CreateTime: time.Now(),
		}

		err = ud.Create(userdata)
		check(err, t)
		ID = userdata.ID
	}
	{
		user, err := ud.FindById(ID)
		check(err, t)

		if user.Username != "123" {
			t.Error("User name not the same")
		}
	}

}

func TestUpdate(t *testing.T) {
	{
		u, e := ud.FindById(12)
		check(e, t)

		u.Username = "I'm 12"
		u.Password = "I'm 12 passwd"

		e = ud.Save(&u)
		check(e, t)
	}
	{
		u, e := ud.FindById(12)
		check(e, t)

		if u.Username != "I'm 12" || u.Password != "I'm 12 passwd" {
			t.Error("Update model failed: username or password not the modified revision")
		}
	}
}

func check(e error, t *testing.T) {
	if e != nil {
		t.Error(e)
	}
}
