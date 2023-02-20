package videoModel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

func (m *defaultVideoModel) FindManyByTime(ctx context.Context, time int64, num int64) ([]*Video, error) {
	var resp []*Video
	query := fmt.Sprintf("select %s from %s where `removed` = 0 and `time` > ? limit ?", videoRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, time, num)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoModel) FindAllByUserId(ctx context.Context, userId int64) ([]*Video, error) {
	var resp []*Video
	query := fmt.Sprintf("select %s from %s where `removed` = 0 and `author_id` > ?", videoRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
