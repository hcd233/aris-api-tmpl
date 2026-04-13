package dao

import (
	"fmt"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type baseDAOTestModel struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"column:name"`
	GroupName string `gorm:"column:group_name"`
	DeletedAt int64  `gorm:"column:deleted_at;default:0"`
}

func newBaseDAOTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	dsn := fmt.Sprintf("file:%s?mode=memory&cache=private", t.Name())
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm.Open() error = %v", err)
	}
	if err := db.AutoMigrate(&baseDAOTestModel{}); err != nil {
		t.Fatalf("AutoMigrate() error = %v", err)
	}
	return db
}

func TestBaseDAO_Count(t *testing.T) {
	t.Parallel()

	db := newBaseDAOTestDB(t)
	rows := []baseDAOTestModel{
		{Name: "a", GroupName: "g1", DeletedAt: 0},
		{Name: "b", GroupName: "g1", DeletedAt: 1},
		{Name: "c", GroupName: "g2", DeletedAt: 0},
	}
	if err := db.Create(&rows).Error; err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	dao := &baseDAO[baseDAOTestModel]{}
	count, err := dao.Count(db, &baseDAOTestModel{GroupName: "g1"})
	if err != nil {
		t.Fatalf("Count() error = %v", err)
	}
	if count != 1 {
		t.Fatalf("Count() = %d, want 1", count)
	}
}

func TestBaseDAO_BatchDeleteByField(t *testing.T) {
	t.Parallel()

	db := newBaseDAOTestDB(t)
	rows := []baseDAOTestModel{{Name: "a"}, {Name: "b"}, {Name: "c"}}
	if err := db.Create(&rows).Error; err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	dao := &baseDAO[baseDAOTestModel]{}
	if err := dao.BatchDeleteByField(db, "name", []string{"a", "b"}); err != nil {
		t.Fatalf("BatchDeleteByField() error = %v", err)
	}

	var deletedCount int64
	if err := db.Model(&baseDAOTestModel{}).Where("deleted_at != 0").Count(&deletedCount).Error; err != nil {
		t.Fatalf("Count deleted rows error = %v", err)
	}
	if deletedCount != 2 {
		t.Fatalf("deleted row count = %d, want 2", deletedCount)
	}
}

func TestBaseDAO_HardDeleteSoftDeleted(t *testing.T) {
	t.Parallel()

	db := newBaseDAOTestDB(t)
	rows := []baseDAOTestModel{{Name: "active", DeletedAt: 0}, {Name: "deleted", DeletedAt: 1}}
	if err := db.Create(&rows).Error; err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	dao := &baseDAO[baseDAOTestModel]{}
	deletedRows, err := dao.HardDeleteSoftDeleted(db)
	if err != nil {
		t.Fatalf("HardDeleteSoftDeleted() error = %v", err)
	}
	if deletedRows != 1 {
		t.Fatalf("HardDeleteSoftDeleted() deletedRows = %d, want 1", deletedRows)
	}

	var total int64
	if err := db.Model(&baseDAOTestModel{}).Count(&total).Error; err != nil {
		t.Fatalf("Count total rows error = %v", err)
	}
	if total != 1 {
		t.Fatalf("total row count = %d, want 1", total)
	}
}
