// Package repository 实现领域仓储接口。
package repository

import (
	"context"
	"errors"
	"time"

	"github.com/hcd233/aris-api-tmpl/internal/common/constant"
	"github.com/hcd233/aris-api-tmpl/internal/common/ierr"
	"github.com/hcd233/aris-api-tmpl/internal/domain/identity"
	"github.com/hcd233/aris-api-tmpl/internal/domain/identity/aggregate"
	"github.com/hcd233/aris-api-tmpl/internal/domain/identity/vo"
	"github.com/hcd233/aris-api-tmpl/internal/infrastructure/database"
	"github.com/hcd233/aris-api-tmpl/internal/infrastructure/database/dao"
	dbmodel "github.com/hcd233/aris-api-tmpl/internal/infrastructure/database/model"
	"gorm.io/gorm"
)

// userRepository UserRepository 的 GORM 实现。
type userRepository struct {
	dao *dao.UserDAO
}

// NewUserRepository 构造用户仓储。
func NewUserRepository() identity.UserRepository {
	return &userRepository{dao: dao.GetUserDAO()}
}

// Save 持久化用户聚合。
func (r *userRepository) Save(ctx context.Context, user *aggregate.User) error {
	db := database.GetDBInstance(ctx)
	if user.AggregateID() == 0 {
		record := &dbmodel.User{
			Name:         user.Name().String(),
			Email:        user.Email().String(),
			Avatar:       user.Avatar().String(),
			Permission:   user.Permission(),
			LastLogin:    user.LastLogin(),
			GithubBindID: user.GithubBindID(),
			GoogleBindID: user.GoogleBindID(),
		}
		if err := r.dao.Create(db, record); err != nil {
			return ierr.Wrap(ierr.ErrDBCreate, err, "create user")
		}
		user.SetID(record.ID)
		return nil
	}
	updates := map[string]any{
		constant.FieldName:       user.Name().String(),
		constant.FieldEmail:      user.Email().String(),
		constant.FieldAvatar:     user.Avatar().String(),
		constant.FieldPermission: user.Permission(),
		constant.FieldLastLogin:  user.LastLogin(),
	}
	if err := r.dao.Update(db, &dbmodel.User{ID: user.AggregateID()}, updates); err != nil {
		return ierr.Wrap(ierr.ErrDBUpdate, err, "update user")
	}
	return nil
}

// TouchLastLogin 仅更新 last_login 字段。
func (r *userRepository) TouchLastLogin(ctx context.Context, userID uint) error {
	db := database.GetDBInstance(ctx)
	if err := r.dao.Update(db, &dbmodel.User{ID: userID}, map[string]any{
		constant.FieldLastLogin: time.Now().UTC(),
	}); err != nil {
		return ierr.Wrap(ierr.ErrDBUpdate, err, "touch last login")
	}
	return nil
}

// FindByID 按 ID 查询用户聚合。
func (r *userRepository) FindByID(ctx context.Context, id uint) (*aggregate.User, error) {
	db := database.GetDBInstance(ctx)
	record, err := r.dao.Get(db, &dbmodel.User{ID: id}, constant.UserRepoFieldsFull())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, ierr.Wrap(ierr.ErrDBQuery, err, "get user by id")
	}
	return toUserAggregate(record), nil
}

// FindByGithubBindID 按 Github 绑定 ID 查询用户聚合。
func (r *userRepository) FindByGithubBindID(ctx context.Context, bindID string) (*aggregate.User, error) {
	db := database.GetDBInstance(ctx)
	record, err := r.dao.Get(db, &dbmodel.User{GithubBindID: bindID}, constant.UserRepoFieldsFull())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, ierr.Wrap(ierr.ErrDBQuery, err, "get user by github bind id")
	}
	return toUserAggregate(record), nil
}

// FindByGoogleBindID 按 Google 绑定 ID 查询用户聚合。
func (r *userRepository) FindByGoogleBindID(ctx context.Context, bindID string) (*aggregate.User, error) {
	db := database.GetDBInstance(ctx)
	record, err := r.dao.Get(db, &dbmodel.User{GoogleBindID: bindID}, constant.UserRepoFieldsFull())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, ierr.Wrap(ierr.ErrDBQuery, err, "get user by google bind id")
	}
	return toUserAggregate(record), nil
}

func toUserAggregate(record *dbmodel.User) *aggregate.User {
	return aggregate.RestoreUser(
		record.ID,
		vo.UserName(record.Name),
		vo.Email(record.Email),
		vo.Avatar(record.Avatar),
		record.Permission,
		record.LastLogin,
		record.CreatedAt,
		record.GithubBindID,
		record.GoogleBindID,
	)
}
