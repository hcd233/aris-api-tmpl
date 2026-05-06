package util

import (
	"context"

	"github.com/hcd233/aris-api-tmpl/internal/common/enum"
)

// CtxValueUint 从 context 中读取 uint 值。
func CtxValueUint(ctx context.Context, key string) uint {
	value, ok := ctx.Value(key).(uint)
	if !ok {
		return 0
	}
	return value
}

// CtxValuePermission 从 context 中读取用户权限。
func CtxValuePermission(ctx context.Context, key string) enum.Permission {
	permission, ok := ctx.Value(key).(enum.Permission)
	if !ok {
		return ""
	}
	return permission
}
