package handler

import (
	"context"

	oauth2command "github.com/hcd233/aris-api-tmpl/internal/application/oauth2/command"
	"github.com/hcd233/aris-api-tmpl/internal/common/ierr"
	"github.com/hcd233/aris-api-tmpl/internal/dto"
	"github.com/hcd233/aris-api-tmpl/internal/logger"
	"github.com/hcd233/aris-api-tmpl/internal/util"
	"go.uber.org/zap"
)

// Oauth2Handler OAuth2处理器接口。
type Oauth2Handler interface {
	HandleLogin(ctx context.Context, req *dto.LoginReq) (*dto.HTTPResponse[*dto.LoginResp], error)
	HandleCallback(ctx context.Context, req *dto.CallbackReq) (*dto.HTTPResponse[*dto.CallbackRsp], error)
}

// Oauth2Dependencies OAuth2Handler 依赖项。
type Oauth2Dependencies struct {
	Initiate oauth2command.InitiateLoginHandler
	Callback oauth2command.HandleCallbackHandler
}

type oauth2Handler struct {
	initiate oauth2command.InitiateLoginHandler
	callback oauth2command.HandleCallbackHandler
}

// NewOauth2Handler 创建 OAuth2 处理器。
func NewOauth2Handler(deps Oauth2Dependencies) Oauth2Handler {
	return &oauth2Handler{initiate: deps.Initiate, callback: deps.Callback}
}

// HandleLogin OAuth2 登录。
func (h *oauth2Handler) HandleLogin(ctx context.Context, req *dto.LoginReq) (*dto.HTTPResponse[*dto.LoginResp], error) {
	rsp := &dto.LoginResp{}
	if req == nil || req.Platform == "" {
		rsp.Error = ierr.ErrBadRequest.BizError()
		return util.WrapHTTPResponse(rsp, nil)
	}
	result, err := h.initiate.Handle(ctx, oauth2command.InitiateLoginCommand{Platform: req.Platform})
	if err != nil {
		logger.WithCtx(ctx).Error("[OAuth2Handler] initiate login failed", zap.String("platform", req.Platform), zap.Error(err))
		rsp.Error = ierr.ToBizError(err, ierr.ErrInternal.BizError())
		return util.WrapHTTPResponse(rsp, nil)
	}
	rsp.RedirectURL = result.RedirectURL
	return util.WrapHTTPResponse(rsp, nil)
}

// HandleCallback OAuth2 回调。
func (h *oauth2Handler) HandleCallback(ctx context.Context, req *dto.CallbackReq) (*dto.HTTPResponse[*dto.CallbackRsp], error) {
	rsp := &dto.CallbackRsp{}
	if req == nil || req.Body == nil {
		rsp.Error = ierr.ErrBadRequest.BizError()
		return util.WrapHTTPResponse(rsp, nil)
	}
	result, err := h.callback.Handle(ctx, oauth2command.HandleCallbackCommand{
		Platform: req.Body.Platform,
		Code:     req.Body.Code,
		State:    req.Body.State,
	})
	if err != nil {
		logger.WithCtx(ctx).Error("[OAuth2Handler] callback failed", zap.String("platform", req.Body.Platform), zap.Error(err))
		rsp.Error = ierr.ToBizError(err, ierr.ErrInternal.BizError())
		return util.WrapHTTPResponse(rsp, nil)
	}
	rsp.AccessToken = result.TokenPair.AccessToken()
	rsp.RefreshToken = result.TokenPair.RefreshToken()
	return util.WrapHTTPResponse(rsp, nil)
}
