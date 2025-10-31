// Package dto 用户DTO
package dto

// GetCurUserInfoReq represents a request to get the current authenticated user's information
//
//	author centonhuang
//	update 2025-01-04 21:00:54
type GetCurUserInfoReq struct {
	UserID uint `json:"userID" doc:"User ID extracted from the JWT token"`
}

// GetCurUserInfoResp represents the response containing the current user's detailed information
//
//	author centonhuang
//	update 2025-01-04 21:00:59
type GetCurUserInfoResp struct {
	User *User `json:"user" doc:"Complete user information including permissions"`
}

// GetUserInfoReq represents a request to get a specific user's public information
//
//	author centonhuang
//	update 2025-01-04 21:19:41
type GetUserInfoReq struct {
	UserID uint `json:"userID" path:"userID" doc:"Unique identifier of the user to retrieve"`
}

// GetUserInfoResp represents the response containing a user's public information
//
//	author centonhuang
//	update 2025-01-04 21:19:44
type GetUserInfoResp struct {
	User *User `json:"user" doc:"Public user information"`
}

// UpdateUserInfoReq represents a request to update the current user's information
//
//	author centonhuang
//	update 2025-01-04 21:19:47
type UpdateUserInfoReq struct {
	Body *UpdateUserInfoReqBody `json:"body" doc:"Request body containing fields to update"`
}

// UpdateUserInfoReqBody contains the fields that can be updated for a user
//
//	author centonhuang
//	update 2025-10-31 02:33:48
type UpdateUserInfoReqBody struct {
	UserName string `json:"userName" doc:"New display name for the user"`
}
