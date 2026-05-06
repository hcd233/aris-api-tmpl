// Package vo 定义 OAuth2 域值对象。
package vo

// OAuthUserInfo 第三方 OAuth 平台返回的用户信息值对象。
type OAuthUserInfo struct {
	id     string
	name   string
	email  string
	avatar string
}

// NewOAuthUserInfo 构造 OAuth 用户信息值对象。
func NewOAuthUserInfo(id, name, email, avatar string) OAuthUserInfo {
	return OAuthUserInfo{id: id, name: name, email: email, avatar: avatar}
}

// ID 返回平台方用户唯一 ID。
func (u OAuthUserInfo) ID() string {
	return u.id
}

// Name 返回用户名。
func (u OAuthUserInfo) Name() string {
	return u.name
}

// Email 返回邮箱。
func (u OAuthUserInfo) Email() string {
	return u.email
}

// Avatar 返回头像 URL。
func (u OAuthUserInfo) Avatar() string {
	return u.avatar
}

// IsEmpty 判断是否为空。
func (u OAuthUserInfo) IsEmpty() bool {
	return u.id == ""
}
