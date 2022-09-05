package logic

import (
	"context"
	"fmt"
	"go-zero-demo/book/user/api/common/errorx"
	"go-zero-demo/book/user/api/internal/svc"
	"go-zero-demo/book/user/api/internal/types"
	"go-zero-demo/book/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) GetUsersByName(req *types.GetUserByNameRequest) (resp *types.UsersResponse, err error) {
	username := req.Username

	bookUser, err := l.svcCtx.UserModel.FindAllByName(l.ctx, username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultCodeError("用户名不存在")
	default:
		return nil, err
	}

	fmt.Printf("get bookUser=%v", bookUser)

	users := make([]*types.UserDetailResponse, len(bookUser))

	for i := 0; i < len(bookUser); i++ {
		user := *bookUser[i]
		users[i] = &types.UserDetailResponse{
			Id:     user.Id,
			Name:   user.Name,
			Gender: user.Gender}
	}

	fmt.Printf("get users=%v", users)

	return &types.UsersResponse{
		UserList: users,
	}, nil
}
