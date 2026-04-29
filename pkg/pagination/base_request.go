package pagination

import (
	"math"
	"strings"

	"gorm.io/gorm"
)

type BaseRequest struct {
	Page      int    `form:"page,default=1"`
	Paginate  int    `form:"paginate,default=10"`
	Search    string `form:"search,default=''"`
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
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

func PaginateScope(req BaseRequest) func(db *gorm.DB) *gorm.DB {
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

func (r BaseRequest) SearchScope(columns ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		cleanSearch := strings.Trim(r.Search, " '\"")

		if cleanSearch == "" || len(columns) == 0 {
			return db
		}

		var query string
		var args []interface{}
		for i, col := range columns {
			if i > 0 {
				query += " OR "
			}
			query += col + " ILIKE ?"
			args = append(args, "%"+cleanSearch+"%")
		}

		return db.Where("("+query+")", args...)
	}
}

func (r BaseRequest) DateRangeScope(column string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		cleanStart := strings.TrimSpace(r.StartDate)
		cleanEnd := strings.TrimSpace(r.EndDate)

		if cleanStart == "" && cleanEnd == "" {
			return db
		}

		if cleanStart != "" && cleanEnd != "" {
			start := cleanStart + " 00:00:00"
			end := cleanEnd + " 23:59:59"
			return db.Where(column+" BETWEEN ? AND ?", start, end)
		}

		if cleanStart != "" {
			start := cleanStart + " 00:00:00"
			return db.Where(column+" >= ?", start)
		}

		if cleanEnd != "" {
			end := cleanEnd + " 23:59:59"
			return db.Where(column+" <= ?", end)
		}

		return db
	}
}
