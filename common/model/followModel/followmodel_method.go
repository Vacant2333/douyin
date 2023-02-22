package followModel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

func (m *defaultFollowModel) FindAllByUserId(ctx context.Context, userId int64) ([]*Follow, error) {
	var resp []*Follow
	query := fmt.Sprintf("select `fun_id` from %s where `user_id` = ? and `removed` = 0", m.table)
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

func (m *defaultFollowModel) FindAllByFunId(ctx context.Context, funId int64) ([]*Follow, error) {
	var resp []*Follow
	query := fmt.Sprintf("select `user_id` from %s where `fun_id` = ? and `removed` = 0", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, funId)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFollowModel) CountByFollowRelation(ctx context.Context, id int64, field string) (int64, error) {
	query := fmt.Sprintf("select count(*) from %s where %s = ? and removed = 0", m.table, field)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultFollowModel) CheckIsFollow(ctx context.Context, userId int64, funId int64) (bool, error) {
	query := fmt.Sprintf("select count(*) from %s where user_id = ? and removed = 0 and fun_id = ?", m.table)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId, funId)
	switch err {
	case nil:
		return resp == 1, nil
	case sqlc.ErrNotFound:
		return false, ErrNotFound
	default:
		return false, err
	}
}
