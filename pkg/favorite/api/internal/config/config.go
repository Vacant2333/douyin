package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserFavoriteRpc zrpc.RpcClientConf
	UserRpc         zrpc.RpcClientConf

	// kq
	UserFavoriteOptServiceConf kq.KqConf
}
