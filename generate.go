package oauth2

import (
	"context"
	"net/http"
	"time"
)

type (
	// GenerateBasic 提供生成的令牌数据的基础
	GenerateBasic struct {
		Client    ClientInfo
		UserID    string
		CreateAt  time.Time
		TokenInfo TokenInfo
		Request   *http.Request
	}

	// AuthorizeGenerate 不带刷新
	AuthorizeGenerate interface {
		Token(ctx context.Context, data *GenerateBasic) (code string, err error)
	}

	// AccessGenerate 生成token的接口，带刷新
	AccessGenerate interface {
		Token(ctx context.Context, data *GenerateBasic, isGenRefresh bool) (access, refresh string, err error)
	}
)
