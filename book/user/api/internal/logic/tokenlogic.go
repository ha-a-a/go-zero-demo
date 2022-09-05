package logic

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type Tokenlogic struct {
	logx.Logger
	ctx context.Context
}

func NewTokenlogic(ctx context.Context) *Tokenlogic {
	return &Tokenlogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}

func (t *Tokenlogic) GetJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	fmt.Println("getJwtToken")
	claims := make(jwt.MapClaims)
	claims["expire"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
