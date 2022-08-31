package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
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
		return nil, errors.New("参数异常")
	}

	bookUser, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("用户名不存在")
	default:
		return nil, err
	}

	if bookUser.Password != password {
		return nil, errors.New("用户密码不正确")
	}
	fmt.Printf("bookUser=%v", bookUser)

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, bookUser.Id)
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

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	fmt.Println("getJwtToken")
	claims := make(jwt.MapClaims)
	claims["expire"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
