package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf

	// redis
	RedisCacheConf redis.RedisConf

	// kq : pub sub
	UserCommentOptServiceConf  kq.KqConf
	UserFavoriteOptServiceConf kq.KqConf
	UserFollowOptServiceConf   kq.KqConf

	// rpc
	UserCommentRpc  zrpc.RpcClientConf
	UserFavoriteRpc zrpc.RpcClientConf
	UserFollowRpc   zrpc.RpcClientConf
}
