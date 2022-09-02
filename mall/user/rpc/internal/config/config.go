package config

import (
	"github.com/zeromicro/go-zero/core/trace"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Telemetry trace.Config
}
