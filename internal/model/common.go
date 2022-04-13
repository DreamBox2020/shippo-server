package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Pagination struct {
	Current  int   `json:"current"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
}

type ScopeFunc = func(db *gorm.DB) *gorm.DB

func (t *Pagination) Sql() ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		if t.Current <= 0 {
			t.Current = 1
		}

		switch {
		case t.PageSize > 100:
			t.PageSize = 100
		case t.PageSize <= 0:
			t.PageSize = 20
		}

		return db.Offset((t.Current - 1) * t.PageSize).Limit(t.PageSize)
	}
}

func (t *Pagination) Copy(p Pagination) {
	t.Current = p.Current
	t.PageSize = p.PageSize
	t.Total = p.Total
}
