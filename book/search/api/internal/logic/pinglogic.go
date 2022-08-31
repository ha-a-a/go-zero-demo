package logic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/book/search/api/internal/svc"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() error {
	// todo: add your logic here and delete this line
	fmt.Println("ping done")
	return nil
}
