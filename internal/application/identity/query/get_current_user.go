// Package query 定义 Identity 域查询处理器。
package query

import (
	"context"
	"time"

	"github.com/hcd233/aris-api-tmpl/internal/common/enum"
	"github.com/hcd233/aris-api-tmpl/internal/common/ierr"
	"github.com/hcd233/aris-api-tmpl/internal/domain/identity"
	"github.com/hcd233/aris-api-tmpl/internal/logger"
	"go.uber.org/zap"
)

// UserView 用户详情只读投影。
type UserView struct {
	ID         uint
	Name       string
	Email      string
	Avatar     string
	Permission enum.Permission
	CreatedAt  time.Time
	LastLogin  time.Time
}

// GetCurrentUserQuery 查询当前用户命令。
type GetCurrentUserQuery struct {
	UserID uint
}

// GetCurrentUserHandler 查询处理器。
type GetCurrentUserHandler interface {
	Handle(ctx context.Context, q GetCurrentUserQuery) (*UserView, error)
}

type getCurrentUserHandler struct {
	repo identity.UserRepository
}

// NewGetCurrentUserHandler 构造查询处理器。
func NewGetCurrentUserHandler(repo identity.UserRepository) GetCurrentUserHandler {
	return &getCurrentUserHandler{repo: repo}
}

// Handle 执行当前用户查询。
func (h *getCurrentUserHandler) Handle(ctx context.Context, q GetCurrentUserQuery) (*UserView, error) {
	log := logger.WithCtx(ctx)
	user, err := h.repo.FindByID(ctx, q.UserID)
	if err != nil {
		log.Error("[IdentityQuery] find user failed", zap.Error(err), zap.Uint("userID", q.UserID))
		return nil, err
	}
	if user == nil {
		return nil, ierr.New(ierr.ErrDataNotExists, "user not found")
	}
	return &UserView{
		ID:         user.AggregateID(),
		Name:       user.Name().String(),
		Email:      user.Email().String(),
		Avatar:     user.Avatar().String(),
		Permission: user.Permission(),
		CreatedAt:  user.CreatedAt(),
		LastLogin:  user.LastLogin(),
	}, nil
}
