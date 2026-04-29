package pagination

import (
	"math"

	"gorm.io/gorm"
)

type Request struct {
	Page     int `form:"page,default=1"`
	Paginate int `form:"paginate,default=10"`
	Search   int `form:"search,default=''"`
}

type Meta struct {
	Total       int64 `json:"total"`
	Paginate    int   `json:"paginate"`
	CurrentPage int   `json:"current_page"`
	From        int   `json:"from"`
	To          int   `json:"to"`
	LastPage    int   `json:"last_page"`
}

func BuildMeta(total int64, page int, paginate int, dataLength int) *Meta {
	lastPage := int(math.Ceil(float64(total) / float64(paginate)))
	from := 0
	to := 0

	if total > 0 {
		from = ((page - 1) * paginate) + 1
		to = from + dataLength - 1
	}

	return &Meta{
		Total:       total,
		Paginate:    paginate,
		CurrentPage: page,
		From:        from,
		To:          to,
		LastPage:    lastPage,
	}
}

func PaginateScope(req Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if req.Page <= 0 {
			req.Page = 1
		}
		if req.Paginate <= 0 {
			req.Paginate = 10
		}

		offset := (req.Page - 1) * req.Paginate
		return db.Select("*, COUNT(*) OVER() AS total_count").
			Offset(offset).
			Limit(req.Paginate)
	}
}
