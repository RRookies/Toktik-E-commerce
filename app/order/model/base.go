package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"column:add_time" json:"-"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"-"`
}


func Paginate(page, pagesize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pagesize > 100:
			pagesize = 100
		case pagesize <= 0:
			pagesize = 10
		}
		offset := (page - 1) * pagesize
		return db.Offset(offset).Limit(pagesize)
	}
}