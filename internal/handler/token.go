package handler

import (
	"context"
	"strings"

	"github.com/hcd233/aris-api-tmpl/internal/application/identity/command"
	"github.com/hcd233/aris-api-tmpl/internal/common/ierr"
	"github.com/hcd233/aris-api-tmpl/internal/dto"
	"github.com/hcd233/aris-api-tmpl/internal/logger"
	"github.com/hcd233/aris-api-tmpl/internal/util"
	"go.uber.org/zap"
)

// TokenHandler 令牌处理器。
type TokenHandler interface {
	HandleRefreshToken(ctx context.Context, req *dto.RefreshTokenReq) (*dto.HTTPResponse[*dto.RefreshTokenRsp], error)
}

// TokenDependencies TokenHandler 依赖项。
type TokenDependencies struct {
	Refresh command.RefreshTokensHandler
}

type tokenHandler struct {
	refresh command.RefreshTokensHandler
}

// NewTokenHandler 创建令牌处理器。
func NewTokenHandler(deps TokenDependencies) TokenHandler {
	return &tokenHandler{refresh: deps.Refresh}
}

// HandleRefreshToken 刷新令牌。
func (h *tokenHandler) HandleRefreshToken(ctx context.Context, req *dto.RefreshTokenReq) (*dto.HTTPResponse[*dto.RefreshTokenRsp], error) {
	rsp := &dto.RefreshTokenRsp{}
	if req == nil || req.Body == nil || strings.TrimSpace(req.Body.RefreshToken) == "" {
		rsp.Error = ierr.ErrValidation.BizError()
		return util.WrapHTTPResponse(rsp, nil)
	}
	pair, err := h.refresh.Handle(ctx, command.RefreshTokensCommand{RefreshToken: req.Body.RefreshToken})
	if err != nil {
		logger.WithCtx(ctx).Warn("[TokenHandler] refresh token failed",
			zap.String("refreshToken", util.MaskSecret(req.Body.RefreshToken)), zap.Error(err))
		rsp.Error = ierr.ToBizError(err, ierr.ErrInternal.BizError())
		return util.WrapHTTPResponse(rsp, nil)
	}
	rsp.AccessToken = pair.AccessToken()
	rsp.RefreshToken = pair.RefreshToken()
	return util.WrapHTTPResponse(rsp, nil)
}
