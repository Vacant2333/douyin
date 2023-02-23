package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	UserRPC zrpc.RpcClientConf
	Mysql   struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
}
