// Package httpclient 通用 HTTP 客户端模块。
package httpclient

import (
	"crypto/tls"
	"net"
	"net/http"

	"github.com/hcd233/aris-api-tmpl/internal/common/constant"
	"github.com/hcd233/aris-api-tmpl/internal/logger"
	"go.uber.org/zap"
)

var client *http.Client

// GetHTTPClient 获取通用 HTTP 客户端单例。
func GetHTTPClient() *http.Client {
	return client
}

// InitHTTPClient 初始化通用 HTTP 客户端。
func InitHTTPClient() {
	client = &http.Client{
		Timeout: constant.HTTPClientTimeout,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   constant.HTTPDialTimeout,
				KeepAlive: constant.HTTPKeepAlive,
			}).DialContext,
			TLSClientConfig:       &tls.Config{MinVersion: tls.VersionTLS12},
			TLSHandshakeTimeout:   constant.HTTPTLSHandshakeTimeout,
			ResponseHeaderTimeout: constant.HTTPResponseHeaderTimeout,
			MaxIdleConns:          constant.HTTPMaxIdleConns,
			MaxIdleConnsPerHost:   constant.HTTPMaxIdleConnsPerHost,
			IdleConnTimeout:       constant.HTTPIdleConnTimeout,
			ForceAttemptHTTP2:     true,
		},
	}

	logger.Logger().Info("[HTTPClient] initialized upstream HTTP client",
		zap.Duration("timeout", constant.HTTPClientTimeout),
		zap.Int("maxIdleConns", constant.HTTPMaxIdleConns),
		zap.Int("maxIdleConnsPerHost", constant.HTTPMaxIdleConnsPerHost),
	)
}
