package logic

import (
	"context"
	"net/http"

	"go-zero-demo/book/search/api/internal/svc"
	"go-zero-demo/book/search/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchReply, err error) {
	// 调用user查询用户信息
	http.Get("http://localhost:8888/api/order/get/1")
	return &types.SearchReply{
		Name:  req.Name,
		Count: 0,
	}, nil
}
