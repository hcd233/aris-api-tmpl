package vo

// TokenPair 访问令牌对值对象。
type TokenPair struct {
	accessToken  string
	refreshToken string
}

// NewTokenPair 构造令牌对值对象。
func NewTokenPair(accessToken, refreshToken string) TokenPair {
	return TokenPair{accessToken: accessToken, refreshToken: refreshToken}
}

// AccessToken 返回访问令牌。
func (p TokenPair) AccessToken() string {
	return p.accessToken
}

// RefreshToken 返回刷新令牌。
func (p TokenPair) RefreshToken() string {
	return p.refreshToken
}

// IsEmpty 判断令牌对是否为空。
func (p TokenPair) IsEmpty() bool {
	return p.accessToken == "" && p.refreshToken == ""
}
