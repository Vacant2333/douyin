package followModel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

func (m *defaultFollowModel) FindAllByUserId(ctx context.Context, userId int64) ([]*Follow, error) {
	var resp []*Follow
	query := fmt.Sprintf("select * from %s where `user_id` = ? and `removed` = 0", m.table)
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
	query := fmt.Sprintf("select * from %s where `fun_id` = ? and `removed` = 0", m.table)
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

func (m *defaultFollowModel) FindIfExist(ctx context.Context, userId int64, funId int64) (int64, error) {
	query := fmt.Sprintf("select id from %s where user_id = ? and removed = 0 and fun_id = ?", m.table)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId, funId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultFollowModel) UpdateMsg(ctx context.Context, sender int64, receiver int64, msg string) error {
	query := fmt.Sprintf("update %s set msg = ?,sender = 0 where user_id = ? and fun_id = ? ", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, msg, sender, receiver)
	query = fmt.Sprintf("update %s set msg = ?,sender = 1 where user_id = ? and fun_id = ? ", m.table)
	_, err = m.ExecNoCacheCtx(ctx, query, msg, receiver, sender)
	return err
}

func (m *defaultFollowModel) FindMsg(ctx context.Context, userId int64, funId int64) (*Follow, error) {
	query := fmt.Sprintf("select msg,sender from %s where user_id = ? and removed = 0 and fun_id = ?", m.table)
	var resp Follow
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId, funId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
