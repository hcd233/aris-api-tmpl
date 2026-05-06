// Package service 定义 Identity 域领域服务接口。
package service

// TokenSigner Token 签名器接口。
type TokenSigner interface {
	EncodeToken(userID uint) (token string, err error)
	DecodeToken(tokenString string) (userID uint, err error)
}
