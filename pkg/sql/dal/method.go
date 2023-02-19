/*
*

	@Desc: Define your sql methods here

*
*/
package dal

import (
	"time"

	"gorm.io/gen"
)

type DefaultMethod interface {

	// sql(select * from @@table)
	FindAll() ([]gen.T, error)

	// Where("id=@id")
	FindById(id uint) (gen.T, error)

	// Where("removed=@removed")
	FindByRemoved(removed uint) ([]gen.T, error)
}

type UserMethod interface {

	// Where("username=@username")
	FindByUsernmae(username string) ([]gen.T, error)

	// Where("type=@theType")
	FindByType(theType uint) ([]gen.T, error)

	// Where("enable=@enable")
	FindByEnable(enable uint) ([]gen.T, error)
}

type ChatMethod interface {

	// Where("sender=@sender")
	FindBySender(sender uint) ([]gen.T, error)

	// Where("receiver=@receiver")
	FindByReceiver(receiver uint) ([]gen.T, error)

	// Where("sender=@userId or receiver=@userId")
	FindMessageByUserId(userId uint) ([]gen.T, error)
}

type FollowMethod interface {

	// Where("user_id=@userId")
	FindFolloweesByUserId(userId uint) ([]gen.T, error)

	// Where("fun_id=@userId")
	FindFollowersByUserId(userId uint) ([]gen.T, error)
}

type VideoMethod interface {
	// Where("user_id=@userId")
	FindByAuthor(userId uint) ([]gen.T, error)

	// Where("play_url=@playUrl")
	FindByPlayUrl(playUrl string) (gen.T, error)

	// Where("time=@time")
	FindByTime(time uint) ([]gen.T, error)
	// Where("time>=@time")
	FindByTimeLongerThan(time uint) ([]gen.T, error)

	// where("title=@title or title like @title")
	FindByTitle(title string) ([]gen.T, error)
}

type CommentMethod interface {

	// where("user_id=@userId")
	FindByUserId(userId uint) ([]gen.T, error)

	// Where("video_id=@videoId")
	FindByVideoId(videoId uint) ([]gen.T, error)

	// Where("create_time >= @createTime")
	FindNewerThanCreateTime(createTime *time.Time)

	// Where("deleted=@deleted")
	FindByDeleted(deleted uint) ([]gen.T, error)
}

type FavoriteMethod interface {
	// where("user_id=@userId")
	FindByUserId(userId uint) ([]gen.T, error)

	// Where("video_id=@videoId")
	FindByVideoId(videoId uint) ([]gen.T, error)
}
