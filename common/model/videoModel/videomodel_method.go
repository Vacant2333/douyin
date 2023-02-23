package videoModel

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	query := fmt.Sprintf("select %s from %s where `removed` = 0 and `author_id` = ?", videoRows, m.table)
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

func (m *defaultVideoModel) UpdateCount(ctx context.Context, videoId int64, filed string, actionType int64) error {
	var query string
	if actionType == 1 {
		query = fmt.Sprintf("update %s set %s = %s + 1 where `id` = ?", m.table, filed, filed)
	} else {
		query = fmt.Sprintf("update %s set %s = %s - 1 where `id` = ?", m.table, filed, filed)
	}
	tiktokVideoIdKey := fmt.Sprintf("%s%v", cacheTiktokVideoIdPrefix, videoId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		return conn.ExecCtx(ctx, query, videoId)
	}, tiktokVideoIdKey)
	if err != nil {
		return err
	}
	return nil
}
