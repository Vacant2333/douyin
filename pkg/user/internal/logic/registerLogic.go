package logic

import (
	"context"
	"douyin/common/help/token"
	model "douyin/common/model/userModel"
	"douyin/common/xerr"
	"douyin/pkg/user/internal/svc"
	"douyin/pkg/user/userInfoPb"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------user-----------------------
func (l *RegisterLogic) Register(in *userInfoPb.RegisterReq) (*userInfoPb.RegisterResp, error) {
	_, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, in.UserName)
	if err == nil {
		fmt.Printf("exist::::::::::::::::::")
		fmt.Printf("re:::::::::::::::::::::::")
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR), "User %s already exists ", in.UserName)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(in.Password), 12)
	if err != nil {
		logx.Errorf("generate password failed, err:%s", err.Error())
		return nil, err
	}
	res, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username:        in.UserName,
		Password:        string(bytes),
		Avatar:          "https://www.shunvzhi.com/uploads/allimg/180731/1TF952E-3.jpg",
		BackgroundImage: "https://inews.gtimg.com/newsapp_bt/0/13250363674/1000.jpg",
		Signature:       "打工魂",
	})
	if err != nil {
		logx.Errorf("insert user failed, err: %s", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insert user failed, user_name: %s", in.UserName)
	}
	userId, _ := res.LastInsertId()

	var genToken *token.GenToken
	now := time.Now()
	tokenString, err := genToken.GenToken(now, userId, nil)
	if err != nil {
		logx.Errorf("gen token error: %s", err.Error())
		return nil, errors.Wrapf(err, "genToken error")
	}

	_, err = l.svcCtx.RedisCache.SetnxExCtx(l.ctx, "token:"+strconv.FormatInt(userId, 10), tokenString, token.AccessExpire)
	if err != nil {
		logx.Errorf("set token to redis error: %s", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "genToken error")
	}

	return &userInfoPb.RegisterResp{
		UserId: userId,
		Token:  tokenString,
	}, nil
}
