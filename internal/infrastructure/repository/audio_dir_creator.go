package repository

import (
	"context"

	"github.com/hcd233/aris-api-tmpl/internal/common/ierr"
	objdao "github.com/hcd233/aris-api-tmpl/internal/infrastructure/storage/obj_dao"
)

// AudioDirCreator 将对象存储 DAO 适配为创建目录能力。
type AudioDirCreator struct {
	dao objdao.ObjDAO
}

// NewAudioDirCreator 构造音频目录创建器。
func NewAudioDirCreator() *AudioDirCreator {
	return &AudioDirCreator{dao: objdao.GetAudioObjDAO()}
}

// CreateDir 为指定用户创建独立音频目录。
func (a *AudioDirCreator) CreateDir(ctx context.Context, userID uint) error {
	if _, err := a.dao.CreateDir(ctx, userID); err != nil {
		return ierr.Wrap(ierr.ErrObjStorage, err, "create audio dir")
	}
	return nil
}
