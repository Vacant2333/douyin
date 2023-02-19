package logic

import (
	"context"
	myToken "douyin/common/help/token"
	"douyin/common/xerr"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"

	"douyin/pkg/user/rpc/internal/svc"
	"douyin/pkg/user/rpc/userInfoPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *userInfoPb.LoginReq) (*userInfoPb.LoginResp, error) {
	fmt.Printf("loginRpc:::::::::::::::::::")
	user, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, in.UserName)
	if err != nil {
		logx.Errorf("find user failed, err: %s", err.Error())
		return nil, errors.Wrap(err, "find user failed")
	}

	// 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password))
	if err != nil {
		logx.Errorf("password not match, err: %s", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "password not match")
	}

	// 通过userId查 redis 是否有此token
	token, err := l.svcCtx.RedisCache.GetCtx(l.ctx, "token:"+strconv.FormatInt(user.Id, 10))
	if err != nil {
		logx.Errorf("get token from redis failed, err: %s", err.Error())
		return nil, errors.Wrap(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "get token from redis failed")
	}
	// 如果存在，则直接返回
	if token != "" {
		return &userInfoPb.LoginResp{
			UserId: user.Id,
			Token:  token,
		}, nil
	}

	//如果不存在，则生成token，并存入redis
	var genToken myToken.GenToken
	now := time.Now()
	token, err = genToken.GenToken(now, user.Id, nil)
	_, err = l.svcCtx.RedisCache.SetnxExCtx(l.ctx, "token:"+strconv.FormatInt(user.Id, 10), token, myToken.AccessExpire)
	if err != nil {
		logx.Errorf("set token to redis failed, err: %s", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "set token to redis error")
	}

	return &userInfoPb.LoginResp{
		UserId: user.Id,
		Token:  token,
	}, nil
}
