package funding_source

import (
	"backend-app/internal/base/models"
)

type FundingSource struct {
	models.BaseModel
	FundingSource string `gorm:"column:funding_source;type:varchar(100);not null" json:"funding_source"`
}

func (FundingSource) TableName() string {
	return "funding_sources_m"
}
