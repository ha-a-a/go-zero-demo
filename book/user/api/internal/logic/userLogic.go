package logic

import (
	"context"
	"fmt"
	"go-zero-demo/book/user/api/common/errorx"
	"go-zero-demo/book/user/api/internal/svc"
	"go-zero-demo/book/user/api/internal/types"
	"go-zero-demo/book/user/model"
	"time"

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

func (l *UserLogic) AddUser(req *types.AddUserRequest) (resp *types.LoginResponse, err error) {

	result, err := l.svcCtx.UserModel.Insert(l.ctx, &model.BookUser{Number: req.Number,
		Name:     req.Username,
		Password: req.Password,
		Gender:   req.Gender,
	})
	if result == nil {
		return nil, errorx.NewDefaultCodeError("添加失败")
	}
	id, err := result.LastInsertId()
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultCodeError("用户名不存在")
	default:
		return nil, err
	}
	bookUser, err := l.svcCtx.UserModel.FindOne(l.ctx, id)
	if bookUser == nil {
		return nil, errorx.NewDefaultCodeError("用户不存在")
	}
	fmt.Printf("get bookUser=%v", bookUser)
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	token, err := NewTokenlogic(l.ctx).GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, bookUser.Id)
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		Id:           bookUser.Id,
		Name:         bookUser.Name,
		Gender:       bookUser.Gender,
		AccessToken:  token,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
