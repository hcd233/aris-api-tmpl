package handler

import (
	"context"
	"time"

	identitycommand "github.com/hcd233/aris-api-tmpl/internal/application/identity/command"
	identityquery "github.com/hcd233/aris-api-tmpl/internal/application/identity/query"
	"github.com/hcd233/aris-api-tmpl/internal/common/constant"
	"github.com/hcd233/aris-api-tmpl/internal/common/ierr"
	"github.com/hcd233/aris-api-tmpl/internal/dto"
	"github.com/hcd233/aris-api-tmpl/internal/logger"
	"github.com/hcd233/aris-api-tmpl/internal/util"
	"go.uber.org/zap"
)

// UserHandler 用户处理器。
type UserHandler interface {
	HandleGetCurUser(ctx context.Context, req *dto.EmptyReq) (*dto.HTTPResponse[*dto.GetCurUserRsp], error)
	HandleUpdateUser(ctx context.Context, req *dto.UpdateUserReq) (*dto.HTTPResponse[*dto.EmptyRsp], error)
}

// UserDependencies UserHandler 依赖项。
type UserDependencies struct {
	GetCurrentUser identityquery.GetCurrentUserHandler
	UpdateProfile  identitycommand.UpdateProfileHandler
}

type userHandler struct {
	getCurrentUser identityquery.GetCurrentUserHandler
	updateProfile  identitycommand.UpdateProfileHandler
}

// NewUserHandler 创建用户处理器。
func NewUserHandler(deps UserDependencies) UserHandler {
	return &userHandler{getCurrentUser: deps.GetCurrentUser, updateProfile: deps.UpdateProfile}
}

// HandleGetCurUser 获取当前用户信息。
func (h *userHandler) HandleGetCurUser(ctx context.Context, _ *dto.EmptyReq) (*dto.HTTPResponse[*dto.GetCurUserRsp], error) {
	rsp := &dto.GetCurUserRsp{}
	userID := util.CtxValueUint(ctx, constant.CtxKeyUserID)
	if userID == 0 {
		rsp.Error = ierr.ErrUnauthorized.BizError()
		return util.WrapHTTPResponse(rsp, nil)
	}
	view, err := h.getCurrentUser.Handle(ctx, identityquery.GetCurrentUserQuery{UserID: userID})
	if err != nil {
		logger.WithCtx(ctx).Error("[UserHandler] get current user failed", zap.Error(err))
		rsp.Error = ierr.ToBizError(err, ierr.ErrInternal.BizError())
		return util.WrapHTTPResponse(rsp, nil)
	}
	rsp.User = &dto.DetailedUser{
		ID:         view.ID,
		CreatedAt:  view.CreatedAt.Format(time.DateTime),
		LastLogin:  view.LastLogin.Format(time.DateTime),
		Permission: string(view.Permission),
		User: dto.User{
			Name:   view.Name,
			Email:  view.Email,
			Avatar: view.Avatar,
		},
	}
	return util.WrapHTTPResponse(rsp, nil)
}

// HandleUpdateUser 更新当前用户资料。
func (h *userHandler) HandleUpdateUser(ctx context.Context, req *dto.UpdateUserReq) (*dto.HTTPResponse[*dto.EmptyRsp], error) {
	rsp := &dto.EmptyRsp{}
	userID := util.CtxValueUint(ctx, constant.CtxKeyUserID)
	if userID == 0 || req == nil || req.Body == nil || req.Body.User == nil {
		rsp.Error = ierr.ErrBadRequest.BizError()
		return util.WrapHTTPResponse(rsp, nil)
	}
	if err := h.updateProfile.Handle(ctx, identitycommand.UpdateProfileCommand{
		UserID: userID,
		Name:   req.Body.User.Name,
		Email:  req.Body.User.Email,
		Avatar: req.Body.User.Avatar,
	}); err != nil {
		logger.WithCtx(ctx).Error("[UserHandler] update user failed", zap.Error(err))
		rsp.Error = ierr.ToBizError(err, ierr.ErrInternal.BizError())
		return util.WrapHTTPResponse(rsp, nil)
	}
	return util.WrapHTTPResponse(rsp, nil)
}
