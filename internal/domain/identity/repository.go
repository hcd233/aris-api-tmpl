// Package identity 定义 Identity 域仓储接口。
package identity

import (
	"context"

	"github.com/hcd233/aris-api-tmpl/internal/domain/identity/aggregate"
)

// UserRepository User 聚合仓储接口。
type UserRepository interface {
	Save(ctx context.Context, user *aggregate.User) error
	FindByID(ctx context.Context, id uint) (*aggregate.User, error)
	FindByGithubBindID(ctx context.Context, bindID string) (*aggregate.User, error)
	FindByGoogleBindID(ctx context.Context, bindID string) (*aggregate.User, error)
	TouchLastLogin(ctx context.Context, userID uint) error
}
