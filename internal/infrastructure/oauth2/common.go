// Package oauth2 实现 OAuth2 平台策略。
package oauth2

import (
	"context"

	"github.com/hcd233/aris-api-tmpl/internal/domain/oauth2/vo"
	"golang.org/x/oauth2"
)

// Platform OAuth2 提供商接口。
type Platform interface {
	GetAuthURL() string
	ExchangeToken(ctx context.Context, code string) (*oauth2.Token, error)
	GetUserInfo(ctx context.Context, token *oauth2.Token) (vo.OAuthUserInfo, error)
}
