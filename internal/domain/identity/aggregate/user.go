// Package aggregate 定义 Identity 域聚合根。
package aggregate

import (
	"time"

	"github.com/hcd233/aris-api-tmpl/internal/common/constant"
	"github.com/hcd233/aris-api-tmpl/internal/common/enum"
	"github.com/hcd233/aris-api-tmpl/internal/common/ierr"
	commonaggregate "github.com/hcd233/aris-api-tmpl/internal/domain/common/aggregate"
	"github.com/hcd233/aris-api-tmpl/internal/domain/identity/vo"
)

// User 用户聚合根。
type User struct {
	commonaggregate.Base

	name         vo.UserName
	email        vo.Email
	avatar       vo.Avatar
	permission   enum.Permission
	lastLogin    time.Time
	githubBindID string
	googleBindID string
	createdAt    time.Time
}

// RegisterUser 创建新用户聚合。
func RegisterUser(name vo.UserName, email vo.Email, avatar vo.Avatar, provider, bindID string, now time.Time) (*User, error) {
	if name.IsEmpty() {
		return nil, ierr.New(ierr.ErrValidation, "user name is empty")
	}
	user := &User{
		name:       name,
		email:      email,
		avatar:     avatar,
		permission: enum.PermissionPending,
		lastLogin:  now,
		createdAt:  now,
	}
	switch provider {
	case constant.OAuthProviderGithub:
		user.githubBindID = bindID
	case constant.OAuthProviderGoogle:
		user.googleBindID = bindID
	}
	return user, nil
}

// RestoreUser 从仓储重建用户聚合。
func RestoreUser(id uint, name vo.UserName, email vo.Email, avatar vo.Avatar,
	permission enum.Permission, lastLogin, createdAt time.Time, githubBindID, googleBindID string) *User {
	user := &User{
		name:         name,
		email:        email,
		avatar:       avatar,
		permission:   permission,
		lastLogin:    lastLogin,
		createdAt:    createdAt,
		githubBindID: githubBindID,
		googleBindID: googleBindID,
	}
	user.SetID(id)
	return user
}

// UpdateProfile 更新用户资料。
func (u *User) UpdateProfile(name vo.UserName, email vo.Email, avatar vo.Avatar) error {
	if name.IsEmpty() {
		return ierr.New(ierr.ErrValidation, "user name is empty")
	}
	u.name = name
	u.email = email
	u.avatar = avatar
	return nil
}

// RecordLogin 记录最新登录时间。
func (u *User) RecordLogin(now time.Time) {
	u.lastLogin = now
}

// ChangePermission 变更权限。
func (u *User) ChangePermission(permission enum.Permission) {
	if u.permission == permission {
		return
	}
	u.permission = permission
}

// AggregateType 返回聚合类型。
func (*User) AggregateType() string {
	return constant.AggregateTypeUser
}

// Name 返回用户名。
func (u *User) Name() vo.UserName {
	return u.name
}

// Email 返回邮箱。
func (u *User) Email() vo.Email {
	return u.email
}

// Avatar 返回头像。
func (u *User) Avatar() vo.Avatar {
	return u.avatar
}

// Permission 返回权限。
func (u *User) Permission() enum.Permission {
	return u.permission
}

// LastLogin 返回最近登录时间。
func (u *User) LastLogin() time.Time {
	return u.lastLogin
}

// CreatedAt 返回创建时间。
func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

// GithubBindID 返回 Github 绑定 ID。
func (u *User) GithubBindID() string {
	return u.githubBindID
}

// GoogleBindID 返回 Google 绑定 ID。
func (u *User) GoogleBindID() string {
	return u.googleBindID
}
