package constant

import "time"

const (
	// HTTPClientTimeout HTTP 客户端总超时时间。
	HTTPClientTimeout = 5 * time.Minute

	// HTTPDialTimeout HTTP 建连超时时间。
	HTTPDialTimeout = 10 * time.Second

	// HTTPKeepAlive HTTP keepalive 周期。
	HTTPKeepAlive = 30 * time.Second

	// HTTPTLSHandshakeTimeout TLS 握手超时时间。
	HTTPTLSHandshakeTimeout = 10 * time.Second

	// HTTPResponseHeaderTimeout 等待响应头超时时间。
	HTTPResponseHeaderTimeout = 30 * time.Second

	// HTTPIdleConnTimeout 空闲连接回收时间。
	HTTPIdleConnTimeout = 90 * time.Second
)

const (
	// HTTPMaxIdleConns 全局空闲连接上限。
	HTTPMaxIdleConns = 100

	// HTTPMaxIdleConnsPerHost 单 Host 空闲连接上限。
	HTTPMaxIdleConnsPerHost = 20
)
