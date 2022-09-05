package logic

import (
	"context"
	"fmt"
	"go-zero-demo/book/user/api/common/errorx"
	"go-zero-demo/book/user/model"
	"strings"
	"time"

	"go-zero-demo/book/user/api/internal/svc"
	"go-zero-demo/book/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	username := req.Username
	password := req.Password
	if len(strings.TrimSpace(username)) == 0 || len(strings.TrimSpace(password)) == 0 {
		return nil, errorx.NewDefaultCodeError("参数异常")
	}

	bookUser, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultCodeError("用户名不存在")
	default:
		return nil, err
	}

	if bookUser.Password != password {
		return nil, errorx.NewDefaultCodeError("用户密码不正确")
	}
	fmt.Printf("bookUser=%v", bookUser)

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
