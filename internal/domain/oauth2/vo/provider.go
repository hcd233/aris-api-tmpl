package vo

import "github.com/hcd233/aris-api-tmpl/internal/common/constant"

// OAuthProvider OAuth2 平台类型值对象。
type OAuthProvider string

// OAuthProviderGithub Github OAuth2 平台。
var OAuthProviderGithub = OAuthProvider(constant.OAuthProviderGithub)

// OAuthProviderGoogle Google OAuth2 平台。
var OAuthProviderGoogle = OAuthProvider(constant.OAuthProviderGoogle)

// String 返回字符串形态。
func (p OAuthProvider) String() string {
	return string(p)
}

// IsValid 判断是否为支持的平台。
func (p OAuthProvider) IsValid() bool {
	return p == OAuthProviderGithub || p == OAuthProviderGoogle
}
