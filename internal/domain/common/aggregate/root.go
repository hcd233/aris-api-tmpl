// Package aggregate 定义聚合根公共契约。
package aggregate

// Root 聚合根接口。
type Root interface {
	AggregateID() uint
	AggregateType() string
}

// Base 聚合根公共字段嵌入体。
type Base struct {
	id uint
}

// AggregateID 返回聚合根 ID。
func (b *Base) AggregateID() uint {
	return b.id
}

// SetID 在聚合首次持久化后回填自增 ID。
func (b *Base) SetID(id uint) {
	b.id = id
}
