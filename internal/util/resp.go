// Package util 工具包
package util

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/hcd233/go-backend-tmpl/internal/protocol"
)

// SendHTTPResponse 发送HTTP响应
//
//	param c *fiber.Ctx
//	param data interface{}
//	param err error
//	author centonhuang
//	update 2025-01-04 17:34:06
func SendHTTPResponse(c *fiber.Ctx, data interface{}, err error) error {
	status := http.StatusOK
	rsp := protocol.HTTPResponse{}

	if err == nil {
		rsp.Data = data
		return c.Status(status).JSON(rsp)
	}
	rsp.Error = err.Error()

	switch err {
	case protocol.ErrDataNotExists, protocol.ErrDataExists: // 200
	case protocol.ErrBadRequest: // 400
		status = http.StatusBadRequest
	case protocol.ErrBadRequest: // 400
		status = http.StatusBadRequest
	case protocol.ErrUnauthorized: // 401
		status = http.StatusUnauthorized
	case protocol.ErrNoPermission, protocol.ErrInsufficientQuota: // 403
		status = http.StatusForbidden
	case protocol.ErrTooManyRequests: // 429
		status = http.StatusTooManyRequests
	case protocol.ErrInternalError: // 500
		status = http.StatusInternalServerError
	case protocol.ErrNoImplement: // 501
		status = http.StatusNotImplemented
	}

	return c.Status(status).JSON(rsp)
}

// WrapHTTPResponse 包装HTTP响应错误
//
//	@param rsp rspT
//	@param err error
//	@return *protocol.HumaHTTPResponse[rspT]
//	@return error
//	@author centonhuang
//	@update 2025-10-31 01:47:14
func WrapHTTPResponse[rspT any](rsp rspT, err error) (*protocol.HumaHTTPResponse[rspT], error) {
	switch err {
	case protocol.ErrDataNotExists: // 404
		return nil, huma.Error404NotFound(err.Error())
	case protocol.ErrDataExists, protocol.ErrBadRequest, protocol.ErrInsufficientQuota: // 400
		return nil, huma.Error400BadRequest(err.Error())
	case protocol.ErrUnauthorized: // 401
		return nil, huma.Error401Unauthorized(err.Error())
	case protocol.ErrNoPermission: // 403
		return nil, huma.Error403Forbidden(err.Error())
	case protocol.ErrTooManyRequests: // 429
		return nil, huma.Error429TooManyRequests(err.Error())
	case protocol.ErrInternalError: // 500
		return nil, huma.Error500InternalServerError(err.Error())
	case protocol.ErrNoImplement: // 501
		return nil, huma.Error501NotImplemented(err.Error())
	case nil:
		return &protocol.HumaHTTPResponse[rspT]{
			Body: rsp,
		}, nil
	default:
		return nil, huma.Error500InternalServerError("Unknown error: " + err.Error())
	}
}
