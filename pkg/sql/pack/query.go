/*
@brief: Get the `Query` or Table's method handle
*/
package pack

import (
	"douyin/pkg/sql/dal/query"

	"gorm.io/gorm"
)

func GetQuery(db *gorm.DB) *query.Query {
	return query.Use(db)
}

func GetIUserDo(db *gorm.DB) query.IUserDo {
	q := GetQuery(db)
	ctx := db.Statement.Context
	return q.User.WithContext(ctx)
}

func GetIChatDO(db *gorm.DB) query.IChatDo {
	q := GetQuery(db)
	ctx := db.Statement.Context
	return q.Chat.WithContext(ctx)
}

func GetICommentDO(db *gorm.DB) query.ICommentDo {
	q := GetQuery(db)
	ctx := db.Statement.Context
	return q.Comment.WithContext(ctx)
}

func GetIFavoriteDO(db *gorm.DB) query.IFavoriteDo {
	q := GetQuery(db)
	ctx := db.Statement.Context
	return q.Favorite.WithContext(ctx)
}

func GetIFollowDO(db *gorm.DB) query.IFollowDo {
	q := GetQuery(db)
	ctx := db.Statement.Context
	return q.Follow.WithContext(ctx)
}

func GetIVideoDO(db *gorm.DB) query.IVideoDo {
	q := GetQuery(db)
	ctx := db.Statement.Context
	return q.Video.WithContext(ctx)
}
