package pack

import (
	"context"
	"douyin/pkg/sql/dal/model"
	"douyin/pkg/sql/dal/query"
	"douyin/pkg/sql/pack"
	"testing"
)

func TestTransaction(t *testing.T) {
	db := pack.GetConn()
	q := pack.GetQuery(db)

	err := q.Transaction(func(tx *query.Query) error {
		e := tx.WithContext(context.Background()).Chat.Create(
			&model.Chat{
				Sender:   1,
				Receiver: 2,
				Msg:      "hahahah"},
		)
		if e != nil {
			return e
		}

		e = tx.Video.WithContext(context.Background()).Create(&model.Video{
			Title: "let other all be null, coming error?",
		})
		// no error will be a error
		if e == nil {
			return e
		}

		return nil
	})
	pack.Check(err)

}

func TestWithHandle(t *testing.T) {
	// User
	db := pack.GetConn()
	{
		icd := pack.GetIChatDO(db)
		e := icd.Create(&model.Chat{
			Msg:      "[01_test::TestWithHandle]",
			Sender:   1,
			Receiver: 2,
		})
		check(e, t)

		Chats, err := icd.FindByReceiver(2)
		check(err, t)

		bFound := false
		for _, chat := range Chats {
			if chat.Msg == "[01_test::TestWithHandle]" &&
				chat.Sender == 1 &&
				chat.Receiver == 2 {
				bFound = true
				break
			}
		}
		if !bFound {
			t.Error("Did not find the message that I send!!!!")
		}

	}
}

func check(e error, t *testing.T) {
	if e != nil {
		t.Error(e)
	}
}
