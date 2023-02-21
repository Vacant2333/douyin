package followModel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

func (m *defaultFollowModel) FindAllByUserId(ctx context.Context, userId string) ([]*Follow, error) {
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

func (m *defaultFollowModel) FindAllByFunId(ctx context.Context, funId string) ([]*Follow, error) {
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
	// todo: 与向交流后完善该函数

	return 0, nil
}
