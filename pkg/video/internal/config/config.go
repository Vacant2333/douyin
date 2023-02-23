package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	CacheRedis cache.CacheConf

	FavoriteRpc    zrpc.RpcClientConf
	UserRpc        zrpc.RpcClientConf
	MinIOClientRpc zrpc.RpcClientConf
	FollowRPC      zrpc.RpcClientConf
}
