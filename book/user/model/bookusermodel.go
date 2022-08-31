package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BookUserModel = (*customBookUserModel)(nil)

type (
	// BookUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBookUserModel.
	BookUserModel interface {
		bookUserModel
	}

	customBookUserModel struct {
		*defaultBookUserModel
	}
)

// NewBookUserModel returns a model for the database table.
func NewBookUserModel(conn sqlx.SqlConn, c cache.CacheConf) BookUserModel {
	return &customBookUserModel{
		defaultBookUserModel: newBookUserModel(conn, c),
	}
}
